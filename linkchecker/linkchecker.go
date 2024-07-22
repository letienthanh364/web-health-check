// linkchecker/linkchecker.go
package linkchecker

import (
	"fmt"
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
	configs []WebConfig
	cron    *cron.Cron
}

// NewLinkChecker initializes and returns a new LinkChecker with given configurations
func NewLinkChecker(configs []WebConfig) *LinkChecker {
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
			} else {
				break
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
			res, err := business.ListContactsForWebsite(nil, config.WebId, &modelcontact.Filter{}, &appCommon.Paging{
				Page:  1,
				Limit: 100,
			})
			if err != nil {
				log.Printf("Failed to get contact list for website %s: %v", config.Name, err)
				return
			}

			for _, contact := range res {
				switch contact.ContactMethod {
				case "email":
					subject := "Link Down Notification"
					msg := fmt.Sprintf("The link %s is down. Website: %s", config.Path, config.Name)
					if err := appCommon.SendEmail(contact.ContactAddress, subject, msg); err != nil {
						log.Printf("Failed to sended contact to %s: %v", config.Email, err)
					}
					log.Printf("Sucessfully sended notifacation to address: %s, method: %s.", contact.ContactAddress, contact.ContactMethod)

				case "discord":
					log.Printf("Sucessfully send notifacation to address: %s, method: %s.", contact.ContactAddress, contact.ContactMethod)

				default:
					log.Printf("Invalid method contact: %s.", contact.ContactMethod)
				}
			}

			//subject := "Link Down Notification"
			//msg := fmt.Sprintf("The link %s is down. Website: %s", config.Path, config.Name)
			//
			//if err := appCommon.SendEmail(config.Email, subject, msg); err != nil {
			//	log.Printf("Failed to send contact to %s: %v", config.Email, err)
			//}
		}
	}
}
