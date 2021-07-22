package sites

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
)

// Reddit site
func Reddit(username string) bool {
	req, _ := http.NewRequest("GET", "https://www.reddit.com/user/"+username, nil)

	req.Header.Set("user-agent", gofakeit.UserAgent())
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	var client http.Client

	res, err := client.Do(req)
	if res.StatusCode == http.StatusOK && err == nil {
		return true
	}
	return false
}
