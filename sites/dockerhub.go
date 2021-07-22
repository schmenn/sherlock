package sites

import "net/http"

// DockerHub site
func DockerHub(username string) bool {
	res, err := http.Head("https://hub.docker.com/v2/users/" + username)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
