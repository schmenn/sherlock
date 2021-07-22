package sites

import "net/http"

// TikTok site
func TikTok(username string) bool {
	res, err := http.Head("https://tiktok.com/@" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
