package sites

import (
	"io"
	"net/http"
	"strings"

	json "github.com/json-iterator/go"
)

type twitchStruct struct {
	OperationName string `json:"operationName"`
	Query         string `json:"query"`
	Variables     *vars  `json:"variables"`
}

type vars struct {
	IsLive     bool   `json:"isLive"`
	Login      string `json:"login"`
	IsVod      bool   `json:"isVod"`
	VodID      string `json:"vodID"`
	PlayerType string `json:"playerType"`
}

type twitchResponse struct {
	Data struct {
		StreamPlaybackAccessToken interface{} `json:"streamPlaybackAccessToken"`
	} `json:"data"`
	Extensions struct {
		DurationMilliseconds int    `json:"durationMilliseconds"`
		OperationName        string `json:"operationName"`
		RequestID            string `json:"requestID"`
	} `json:"extensions"`
}

// Twitch site
func Twitch(username string) bool {
	body, err := json.MarshalIndent(&twitchStruct{
		OperationName: "PlaybackAccessToken_Template",
		Query:         "query PlaybackAccessToken_Template($login: String!, $isLive: Boolean!, $vodID: ID!, $isVod: Boolean!, $playerType: String!) {  streamPlaybackAccessToken(channelName: $login, params: {platform: \"web\", playerBackend: \"mediaplayer\", playerType: $playerType}) @include(if: $isLive) {    value    signature    __typename  }  videoPlaybackAccessToken(id: $vodID, params: {platform: \"web\", playerBackend: \"mediaplayer\", playerType: $playerType}) @include(if: $isVod) {    value    signature    __typename  }}",
		Variables: &vars{
			IsLive:     true,
			Login:      username,
			IsVod:      false,
			VodID:      "",
			PlayerType: "site",
		},
	}, "", "  ")
	if err != nil {
		return false
	}
	req, err := http.NewRequest("POST", "https://gql.twitch.tv/gql", strings.NewReader(string(body)))
	if err != nil {
		return false
	}
	req.Header.Set("Client-ID", "kimne78kx3ncx6brgo4mv6wki5h1ko")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var r twitchResponse

	err = json.Unmarshal(b, &r)
	if err != nil {
		return false
	}
	return r.Data.StreamPlaybackAccessToken != nil
}
