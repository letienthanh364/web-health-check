package linkchecker

import (
	"fmt"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/website/model"
	"log"
)

func SendNotifications(config modelwebsite.WebConfig) {
	contacts := config.Contacts

	subject := "Link Down Notification"
	msg := fmt.Sprintf("The link %s is down. Website: %s", config.Path, config.Name)
	for _, contact := range contacts {
		contactAddress := contact.Address
		switch contact.ContactMethod {
		case "email":
			if err := appCommon.SendEmail(contactAddress, subject, msg); err != nil {
				log.Printf("Failed to sended contact to %s: %v", config.DefaultEmail, err)
			}
			log.Printf("Sucessfully sended notifacation to address: %s, method: %s.", contact.Address, contact.ContactMethod)

		case "discord":
			if contactAddress != "" {
				if err := appCommon.SendDiscordNotification(contactAddress, msg); err != nil {
					log.Printf("Failed to send Discord notification: %v", err)
				}
			}
			log.Printf("Sucessfully send notifacation to address: %s, method: %s.", contact.Address, contact.ContactMethod)

		default:
			log.Printf("Invalid method contact: %s.", contact.ContactMethod)
		}
	}

	if err := appCommon.SendEmail(config.DefaultEmail, subject, msg); err != nil {
		log.Printf("Failed to sended contact to %s: %v", config.DefaultEmail, err)
	}
	log.Printf("Sucessfully sended notifacation to address: %s, method: email.", config.DefaultEmail)
}
