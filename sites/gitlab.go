package sites

import (
	"net/http"
)

// GitLab site
func GitLab(username string) bool {
	res, err := http.Head("https://gitlab.com/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
