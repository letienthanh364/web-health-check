package linkchecker

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/checker"
	"github.com/teddlethal/web-health-check/modules/website/model"
	"log"
	"strconv"
	"strings"
	"time"
)

type LinkChecker struct {
	configs        []modelwebsite.WebConfig
	cron           *cron.Cron
	lastCheckTime  time.Time
	alertEmail     string
	checkInterval  time.Duration
	alertThreshold time.Duration
	cronEntries    map[int][]cron.EntryID
}

// NewLinkChecker initializes and returns a new LinkChecker with given configurations
func NewLinkChecker(configs []modelwebsite.WebConfig, alertEmail string, checkInterval, alertThreshold time.Duration) *LinkChecker {
	return &LinkChecker{
		configs:        configs,
		cron:           cron.New(),
		alertEmail:     alertEmail,
		checkInterval:  checkInterval,
		alertThreshold: alertThreshold,
		cronEntries:    make(map[int][]cron.EntryID), // Initialize the map to hold slices of EntryIDs
	}
}

// Start begins the cron job to check links at regular intervals
func (lc *LinkChecker) Start() {
	for _, config := range lc.configs {
		log.Println(config)
		lc.AddCronJob(config)
	}
	lc.cron.Start()
}

// Stop stops the cron job gracefully
func (lc *LinkChecker) Stop() {
	lc.cron.Stop()
}

// checkLink checks the link for a given configuration
func (lc *LinkChecker) checkLink(config *modelwebsite.WebConfig) func() {
	return func() {
		status := "alive"
		for i := 0; i < config.Retry; i++ {
			isDead := checker.CheckLink(config.Path)
			if isDead {
				status = "dead"
			}
			log.Printf("Website: %s, URL: %s, Status: %s\n", config.Name, config.Path, status)
			if status == "alive" {
				// If the link is alive, reset the notification flag
				config.NotificationSent = false
				config.LastNotificationDate = time.Time{} // Reset the last notification date
				break
			}
		}
		if status == "dead" {
			today := time.Now().Format("2006-01-02")
			lastNotified := config.LastNotificationDate.Format("2006-01-02")

			if !config.NotificationSent || lastNotified != today {
				SendNotifications(*config)
				config.NotificationSent = true
				config.LastNotificationDate = time.Now()
			}
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

func (lc *LinkChecker) StopCronJob(websiteId int) {
	if entryIDs, exists := lc.cronEntries[websiteId]; exists {
		for _, entryID := range entryIDs {
			lc.cron.Remove(entryID)
		}
		delete(lc.cronEntries, websiteId)
	}
}

func (lc *LinkChecker) AddCronJob(config modelwebsite.WebConfig) {
	//_, err := time.LoadLocation(config.TimeZone)
	//if err != nil {
	//	log.Printf("Invalid time zone %s for website %s: %v", config.TimeZone, config.Name, err)
	//	return
	//}

	// Add cron jobs for specific check times
	for _, checkTime := range config.CheckTimes {
		//adjustedCheckTime, err := lc.adjustCronExpression(checkTime, loc)
		entryID, err := lc.cron.AddFunc(checkTime, lc.checkLink(&config))
		if err != nil {
			log.Printf("Failed to add cron job for website %s: %v", config.Name, err)
			continue
		}
		lc.cronEntries[config.WebId] = append(lc.cronEntries[config.WebId], entryID)
	}

	// Add cron job for the interval specified in TimeInterval (seconds)
	if config.TimeInterval > 0 {
		interval := fmt.Sprintf("@every %ds", config.TimeInterval)
		entryID, err := lc.cron.AddFunc(interval, lc.checkLink(&config))
		if err != nil {
			log.Printf("Failed to add cron job for website %s: %v", config.Name, err)
			return
		}
		lc.cronEntries[config.WebId] = append(lc.cronEntries[config.WebId], entryID)
	}
}

// adjusts the cron expression based on the specified timezone
func (lc *LinkChecker) adjustCronExpression(cronExpr string, loc *time.Location) (string, error) {
	parts := strings.Fields(cronExpr)
	if len(parts) != 5 {
		return "", fmt.Errorf("invalid cron expression: %s", cronExpr)
	}

	// Parse the minute and hour
	minute, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid minute in cron expression: %s", parts[0])
	}
	hour, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid hour in cron expression: %s", parts[1])
	}

	// Create a time.Time with the parsed hour and minute in UTC
	t := time.Date(0, 1, 1, hour, minute, 0, 0, time.UTC)

	// Convert the time to the specified timezone
	t = t.In(loc)

	// Adjust the cron expression with the new hour and minute
	adjustedCronExpr := fmt.Sprintf("%d %d %s %s %s", t.Minute(), t.Hour(), parts[2], parts[3], parts[4])

	return adjustedCronExpr, nil
}
