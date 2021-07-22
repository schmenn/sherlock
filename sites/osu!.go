package sites

import "net/http"

// Osu site
func Osu(username string) bool {
	res, err := http.Head("https://osu.ppy.sh/users/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
