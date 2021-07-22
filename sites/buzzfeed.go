package sites

import "net/http"

func BuzzFeed(username string) bool {
	res, err := http.Head("https://www.buzzfeed.com/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
