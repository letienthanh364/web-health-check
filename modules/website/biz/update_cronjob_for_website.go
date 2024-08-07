package bizwebsite

import (
	"github.com/teddlethal/web-health-check/linkchecker"
	"gorm.io/gorm"
	"log"
)

type updateCronJobForWebsiteBiz struct {
	linkChecker *linkchecker.LinkChecker
}

func NewUpdateCronJobForWebsiteBiz(lc *linkchecker.LinkChecker) *updateCronJobForWebsiteBiz {
	return &updateCronJobForWebsiteBiz{linkChecker: lc}
}

func (biz *updateCronJobForWebsiteBiz) UpdateCronJobForWebsite(db *gorm.DB, websiteId int) {
	newConfig := FetchWebsite(db, websiteId)
	log.Printf("New Website config: %v", *newConfig)

	log.Printf("Restarting the cron job for the website: %s", newConfig.Name)
	// Stop the existing cron job
	biz.linkChecker.StopCronJob(websiteId)

	// Add the new cron job
	biz.linkChecker.AddCronJob(*newConfig)

	log.Printf("Successfully restarted the cron job for the website: %s", newConfig.Name)
}
