package linkchecker

type WebConfig struct {
	WebId        int
	Name         string
	Path         string
	TimeInterval int // The time interval in seconds
	Retry        int
	DefaultEmail string
	Status       string
	CheckTimes   []string // List of cron format strings for specific times
	TimeZone     string   // Time zone for the checks
}
