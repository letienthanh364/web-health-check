package bizwebsite

import (
	"github.com/teddlethal/web-health-check/appCommon"
	bizecontact "github.com/teddlethal/web-health-check/modules/contact/biz"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
	storagecontact "github.com/teddlethal/web-health-check/modules/contact/storage"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// FetchWebsites fetches the list of websites from the database
func FetchWebsites(db *gorm.DB) []modelwebsite.WebConfig {
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

	var configs []modelwebsite.WebConfig
	for _, website := range websites {
		checkTimeList, err := FetchCheckTimesForWebsite(db, website.Id)
		if err != nil {
			log.Printf("Failed to fetch check time for website: %s, error: %v\n", website.Name, err)
			continue
		}

		contactList, err := FetchContactsForWebsite(db, website.Id)
		if err != nil {
			log.Printf("Failed to fetch contact list for website: %s, error: %v\n", website.Name, err)
			continue
		}

		config := modelwebsite.WebConfig{
			WebId:        website.Id,
			Name:         website.Name,
			Path:         website.Path,
			TimeInterval: website.TimeInterval,
			Retry:        website.Retry,
			DefaultEmail: website.DefaultEmail,
			Status:       website.Status,
			CheckTimes:   checkTimeList,
			Contacts:     contactList,
			TimeZone:     "Asia/Ho_Chi_Minh",
		}
		configs = append(configs, config)
	}

	return configs
}

func FetchWebsite(db *gorm.DB, websiteId int) *modelwebsite.WebConfig {
	websiteStorage := storagewebsite.NewSqlStore(db)

	getWebsiteBiz := NewGetWebsiteBiz(websiteStorage)

	// Fetch the website
	website, err := getWebsiteBiz.GetWebsiteById(nil, websiteId)
	if err != nil {
		log.Printf("Error refetching website: %v", appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err))
		return nil
	}

	contactList, err := FetchContactsForWebsite(db, websiteId)
	if err != nil {
		log.Printf("Error refetching website: %v", appCommon.ErrCannotListEntity(modelcontact.EntityName, err))
		return nil
	}

	checkTimeList, err := FetchCheckTimesForWebsite(db, websiteId)
	if err != nil {
		log.Printf("Error refetching website: %v", appCommon.ErrCannotListEntity("website check time", err))
		return nil
	}

	// Convert newData to Config
	config := &modelwebsite.WebConfig{
		WebId:        website.Id,
		Name:         website.Name,
		Path:         website.Path,
		TimeInterval: website.TimeInterval,
		Retry:        website.Retry,
		DefaultEmail: website.DefaultEmail,
		TimeZone:     "Asia/Ho_Chi_Minh",
		Contacts:     contactList,
		CheckTimes:   checkTimeList,
	}

	return config
}

func FetchCheckTimesForWebsite(db *gorm.DB, websiteId int) ([]string, error) {
	websiteStorage := storagewebsite.NewSqlStore(db)
	listCheckTimesForWebsiteBiz := NewListCheckTimesForWebsiteBiz(websiteStorage)

	paging := &appCommon.Paging{
		Page:  1,
		Limit: 50,
	}

	checkTimeFilter := &modelwebsite.WebsiteCheckTimeFilter{WebsiteId: strconv.Itoa(websiteId)}
	checkTimeRes, err := listCheckTimesForWebsiteBiz.ListCheckTimesForWebsite(nil, websiteId, checkTimeFilter, paging)
	if err != nil {

		return nil, err
	}
	var checkTimeList []string
	for _, checkTime := range checkTimeRes {
		checkTimeList = append(checkTimeList, checkTime.CheckTime)
	}

	return checkTimeList, nil
}

func FetchContactsForWebsite(db *gorm.DB, websiteId int) ([]modelwebsite.WebsiteContact, error) {
	contactStorage := storagecontact.NewSqlStore(db)
	contactBiz := bizecontact.NewListContactBiz(contactStorage)
	contactFilter := &modelcontact.Filter{WebsiteId: strconv.Itoa(websiteId)}

	paging := &appCommon.Paging{
		Page:  1,
		Limit: 50,
	}

	contactRes, err := contactBiz.ListContacts(nil, contactFilter, paging)
	if err != nil {
		return nil, err
	}
	var contactList []modelwebsite.WebsiteContact
	for _, contact := range contactRes {
		c := modelwebsite.WebsiteContact{
			Id:             contact.Id,
			ContactAddress: contact.Address,
			ContactMethod:  contact.ContactMethod,
		}
		contactList = append(contactList, c)
	}

	return contactList, nil
}
