// linkchecker/linkchecker.go
package linkchecker

import (
	"github.com/robfig/cron/v3"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/checker"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
	storageemail "github.com/teddlethal/web-health-check/modules/contact/storage"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"log"
	"time"
)

type LinkChecker struct {
	configs        []WebConfig
	cron           *cron.Cron
	lastCheckTime  time.Time
	alertEmail     string
	checkInterval  time.Duration
	alertThreshold time.Duration
}

// NewLinkChecker initializes and returns a new LinkChecker with given configurations
func NewLinkChecker(configs []WebConfig, alertEmail string, checkInterval, alertThreshold time.Duration) *LinkChecker {
	return &LinkChecker{
		configs: configs,
		cron:    cron.New(),
	}
}

// Start begins the cron job to check links at regular intervals
func (lc *LinkChecker) Start(db *gorm.DB) {
	log.Println(lc.configs)

	for _, config := range lc.configs {
		interval := "@every " + time.Duration((24*3600*1e9)/config.Limit).String()
		lc.cron.AddFunc(interval, lc.checkLink(config, db))
	}
	lc.cron.Start()
}

// Stop stops the cron job gracefully
func (lc *LinkChecker) Stop() {
	lc.cron.Stop()
}

// checkLink checks the link for a given configuration
func (lc *LinkChecker) checkLink(config WebConfig, db *gorm.DB) func() {
	return func() {
		status := "alive"
		for i := 0; i < config.Retry; i++ {
			isDead := checker.CheckLink(config.Path)
			status = "alive"
			if isDead {
				status = "dead"
			}
			log.Printf("Website: %s, URL: %s, Status: %s\n", config.Name, config.Path, status)
			if status == "alive" {
				break
			}
		}
		if status == "dead" {
			websiteStorage := storagewebsite.NewSqlStore(db)
			emailStorage := storageemail.NewSqlStore(db)
			business := bizwebsite.NewListContactsForWebsiteBiz(websiteStorage, emailStorage)
			contacts, err := business.ListContactsForWebsite(nil, config.WebId, &modelcontact.Filter{}, &appCommon.Paging{
				Page:  1,
				Limit: 100,
			})
			if err != nil {
				log.Printf("Failed to get contact list for website %s: %v", config.Name, err)
			}

			SendNotifications(contacts, config)

		}
		lc.lastCheckTime = time.Now()
	}
}

func (lc *LinkChecker) selfCheck() {
	if time.Since(lc.lastCheckTime) > lc.alertThreshold {
		subject := "Health-Check Service Alert"
		body := "The health-check service has not performed checks for the configured threshold period."
		if err := appCommon.SendEmail(lc.alertEmail, subject, body); err != nil {
			log.Printf("Failed to send self-check alert email to %s: %v", lc.alertEmail, err)
		}
	}
}
