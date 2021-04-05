package sites

import "net/http"

// Instagram site
func Instagram(username string) bool {
	req, _ := http.NewRequest("GET", "https://instagram.com/"+username, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0.1; SM-G935T Build/MMB29M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.81 Mobile Safari/537.36 Instagram 8.4.0 Android (23/6.0.1; 560dpi; 1440x2560; samsung; SM-G935T; hero2qltetmo; qcom; en_US")

	var client http.Client
	res, err := client.Do(req)
	if res.StatusCode == 200 && err == nil {
		return true
	}
	return false
}
