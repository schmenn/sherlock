package sites

import "net/http"

func SourceForge(username string) bool {
	res, err := http.Head("https://sourceforge.net/u/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
