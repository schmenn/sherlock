package sites

import "net/http"

// DockerHub site
func DockerHub(username string) bool {
	res, err := http.Get("https://hub.docker.com/v2/users/"+username)
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false
}
