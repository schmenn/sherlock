package sites

import (
	"io"
	"net/http"
)

// GitLab site
func GitLab(username string) bool {
	res, err := http.Get("https://gitlab.com/api/v4/users?username=" + username)
	if err != nil {
		return false
	}
	body, err := io.ReadAll(res.Body)
	if string(body) != "[]" {
		return true
	}
	return false
}
