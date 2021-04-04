package sites

import "net/http"

// Reddit site
func Reddit(username string) bool {
	res, err := http.Get("https://www.reddit.com/user/" + username)
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false
}
