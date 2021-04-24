package sites

import (
	"net/http"
	"strings"
)

// NewGrounds site
func NewGrounds(username string) bool {
	if strings.Contains(username, ".") {
		return false
	}
	res, err := http.Head("https://" + username + ".newgrounds.com")
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false
}
