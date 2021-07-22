package sites

import "net/http"

// GitHub site
func GitHub(username string) bool {
	res, err := http.Head("https://www.github.com/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
