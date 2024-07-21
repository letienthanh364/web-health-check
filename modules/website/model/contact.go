package modelwebsite

import "strings"

type WebsiteContact struct {
	ContactAddress string `json:"contact_address"`
	ContactMethod  string `json:"contact_method"`
}

type WebsiteContactCreation struct {
	ContactAddress string `json:"contact_address"`
	ContactMethod  string `json:"contact_method"`
}

func (data *WebsiteContactCreation) Validate() error {
	data.ContactAddress = strings.TrimSpace(data.ContactAddress)
	if data.ContactAddress == "" {
		return ErrContactLinkCannotBeEmpty
	}

	data.ContactMethod = strings.TrimSpace(data.ContactMethod)
	if data.ContactMethod == "" {
		return ErrContactMethodCannotBeEmpty
	}

	return nil
}
