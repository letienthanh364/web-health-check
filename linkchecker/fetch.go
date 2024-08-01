package linkchecker

import (
	"github.com/teddlethal/web-health-check/appCommon"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"log"
)

// FetchWebsites fetches the list of websites from the database
func FetchWebsites(db *gorm.DB) []WebConfig {
	websiteStore := storagewebsite.NewSqlStore(db)

	var queryString struct {
		appCommon.Paging
		modelwebsite.Filter
	}

	queryString.Paging.Process()

	websites, err := websiteStore.ListWebsite(nil, &queryString.Filter, &queryString.Paging)
	if err != nil {
		log.Fatalf("Failed to fetch websites: %v", err)
	}

	websiteBiz := bizwebsite.NewListCheckTimesForWebsiteBiz(websiteStore)

	var configs []WebConfig
	for _, website := range websites {
		checkTimeRes, err := websiteBiz.ListCheckTimesForWebsite(nil, website.Id, &modelwebsite.WebsiteCheckTimeFilter{}, &appCommon.Paging{})
		if err != nil {
			log.Fatalf("Failed to fetch check time for website: %s, error: %v", website.Name, err)
		}

		var checkTimeList []string
		for _, checkTime := range checkTimeRes {
			checkTimeList = append(checkTimeList, checkTime.CheckTime)
		}
		config := WebConfig{
			WebId:        website.Id,
			Name:         website.Name,
			Path:         website.Path,
			TimeInterval: website.TimeInterval,
			Retry:        website.Retry,
			DefaultEmail: website.DefaultEmail,
			Status:       website.Status,
			CheckTimes:   checkTimeList,
			TimeZone:     "Asia/Ho_Chi_Minh",
		}
		configs = append(configs, config)
	}

	return configs
}
