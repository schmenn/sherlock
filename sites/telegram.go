package sites

import "net/http"

func Telegram(username string) bool {
	res, err := http.Get("https://t.me/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
