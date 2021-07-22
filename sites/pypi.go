package sites

import (
	"net/http"
)

func PyPi(username string) bool {
	res, err := http.Head("https://pypi.org/user/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
