package sites

import "net/http"

// Kik site
func Kik(username string) bool {
	res, err := http.Head("http://ws2.kik.com/user/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
