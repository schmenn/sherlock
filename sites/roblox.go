package sites

import "net/http"

func Roblox(username string) bool {
	res, err := http.Head("https://www.roblox.com/user.aspx?username=" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
