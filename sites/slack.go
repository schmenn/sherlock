package sites

import "net/http"

func Slack(username string) bool {
	res, err := http.Head("https://" + username + ".slack.com")
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
