package sites

import "net/http"

func SoundCloud(username string) bool {
	res, err := http.Head("https://soundcloud.com/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
