package sites

import "net/http"

// Kik site
func Kik(username string) bool {
	res, err := http.Head("https://kik.me/" + username)
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false
}
