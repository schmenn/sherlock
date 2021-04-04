package sites

import "net/http"

// Osu site
func Osu(username string) bool {
	res, err := http.Get("https://osu.ppy.sh/users/" + username)
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false

}
