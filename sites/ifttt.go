package sites

import "net/http"

// IFTTT site
func IFTTT(username string) bool {
	res, err := http.Head("https://ifttt.com/p/"+username)
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false
}
