package sites

import "net/http"

// Instagram site
func Instagram(username string) bool {
	req, _ := http.NewRequest("GET", "https://instagram.com/"+username, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36")

	var client http.Client
	res, err := client.Do(req)
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false
}
