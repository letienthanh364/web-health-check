package linkchecker

import (
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"log"
)

// FetchWebsites fetches the list of websites from the database
func FetchWebsites(db *gorm.DB) []Config {
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

	var configs []Config
	for _, website := range websites {
		config := Config{
			Name:   website.Name,
			Path:   website.Path,
			Limit:  website.Limit,
			Retry:  website.Retry,
			Emails: website.DefaultEmail,
			Status: website.Status,
		}
		configs = append(configs, config)
	}

	return configs
}
