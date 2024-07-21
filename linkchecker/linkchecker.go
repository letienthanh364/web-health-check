// linkchecker/linkchecker.go
package linkchecker

import (
	"fmt"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/checker"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// Config holds the configuration for the LinkChecker
type Config struct {
	Name   string
	Path   string
	Limit  int
	Retry  int
	Emails string
	Status string
}

// LinkChecker holds the cron job instance
type LinkChecker struct {
	configs []Config
	cron    *cron.Cron
}

// NewLinkChecker initializes and returns a new LinkChecker with given configurations
func NewLinkChecker(configs []Config) *LinkChecker {
	return &LinkChecker{
		configs: configs,
		cron:    cron.New(),
	}
}

// Start begins the cron job to check links at regular intervals
func (lc *LinkChecker) Start() {
	for _, config := range lc.configs {
		log.Printf("%d", config.Limit)
		interval := "@every " + time.Duration((24*3600*1e9)/config.Limit).String()
		log.Print(interval)
		lc.cron.AddFunc(interval, lc.checkLink(config))
	}
	lc.cron.Start()
}

// Stop stops the cron job gracefully
func (lc *LinkChecker) Stop() {
	lc.cron.Stop()
}

// checkLink checks the link for a given configuration
func (lc *LinkChecker) checkLink(config Config) func() {
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
			subject := "Link Down Notification"
			msg := fmt.Sprintf("The link %s is down. Website: %s", config.Path, config.Name)
			if err := appCommon.SendEmail(config.Emails, subject, msg); err != nil {
				log.Printf("Failed to send email to %s: %v", config.Emails, err)
			}
		}
	}
}
