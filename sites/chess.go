package sites

import "net/http"

// Chess site
func Chess(username string) bool {
	res, err := http.Head("https://www.chess.com/member/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
