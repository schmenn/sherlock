package sites

import (
	"io"
	"net/http"
	"strings"
)

// ArchiveOrg site
func ArchiveOrg(username string) bool {
	res, err := http.Get("https://archive.org/details/@"+username)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "cannot find account") {
		return true
	}
	return false
}
