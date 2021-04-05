package sites

import "net/http"

// Reddit site
func Reddit(username string) bool {
	req, _ := http.NewRequest("GET", "https://www.reddit.com/user/" + username, nil)

	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	var client http.Client

	res, err := client.Do(req)
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false
}
