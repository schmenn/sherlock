package sites

import (
	"io"
	"net/http"
	"strings"
)

// Steam site
func Steam(username string) bool {
	res, err := http.Get("https://steamcommunity.com/id/" + username)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return false
	}
	return !strings.Contains(string(b), "<title>Steam Community :: Error</title>")
}
