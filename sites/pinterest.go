package sites

import "net/http"

func Pinterest(username string) bool {
	res, err := http.Head("https://www.pinterest.com/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
