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

	// Stop the existing cron job
	log.Printf("Removing the cron job of the website: %s", newConfig.Name)
	biz.linkChecker.StopCronJob(websiteId)

	// Add the new cron job
	log.Printf("Adding the cron job for the website: %s", newConfig.Name)
	biz.linkChecker.AddCronJob(*newConfig)

	log.Printf("Successfully added the cron job for the website: %s", newConfig.Name)
}
