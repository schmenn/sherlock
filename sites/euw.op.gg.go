package sites

import (
	"bytes"
	"io"
	"net/http"
)

func OPGG(username string) bool {
	res, err := http.Get("https://euw.op.gg/summoner/userName=" + username)
	if err != nil {
		return false
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return false
	}

	return !bytes.Contains(b, []byte("This summoner is not registered at OP.GG."))

}
