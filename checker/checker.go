package checker

import (
	"net/http"
	"time"
)

func CheckLink(url string) bool {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return true
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return true
	}
	return false
}
