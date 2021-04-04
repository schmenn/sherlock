package sites

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Steam site
func Steam(username string) bool {
	//<title>Steam Community :: Error</title>
	resp, err := http.Get(fmt.Sprintf("https://steamcommunity.com/profiles/%v/", username))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	return !strings.Contains(string(b), "<title>Steam Community :: Error</title>")
}
