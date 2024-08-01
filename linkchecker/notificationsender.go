package linkchecker

import (
	"fmt"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	"log"
)

func SendNotifications(contacts []modelwebsite.WebsiteContact, config WebConfig) {
	for _, contact := range contacts {
		contactAddress := contact.ContactAddress
		switch contact.ContactMethod {
		case "email":
			subject := "Link Down Notification"
			msg := fmt.Sprintf("The link %s is down. Website: %s", config.Path, config.Name)
			if err := appCommon.SendEmail(contactAddress, subject, msg); err != nil {
				log.Printf("Failed to sended contact to %s: %v", config.DefaultEmail, err)
			}
			log.Printf("Sucessfully sended notifacation to address: %s, method: %s.", contact.ContactAddress, contact.ContactMethod)

		case "discord":
			if contactAddress != "" {
				discordMessage := fmt.Sprintf("The link %s is down.", config.Path)
				if err := appCommon.SendDiscordNotification(contactAddress, discordMessage); err != nil {
					log.Printf("Failed to send Discord notification: %v", err)
				}
			}
			log.Printf("Sucessfully send notifacation to address: %s, method: %s.", contact.ContactAddress, contact.ContactMethod)

		default:
			log.Printf("Invalid method contact: %s.", contact.ContactMethod)
		}
	}
}
