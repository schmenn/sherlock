package sites

import "net/http"

func FaceBook(username string) bool {
	res, err := http.Head("https://www.facebook.com/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
