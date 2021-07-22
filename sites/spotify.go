package sites

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
)

func Spotify(username string) bool {
	req, _ := http.NewRequest("GET", "https://open.spotify.com/user/"+username, nil)
	req.Header.Set("User-Agent", gofakeit.UserAgent())

	res, err := http.DefaultClient.Do(req)

	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
