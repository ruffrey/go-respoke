package respoke

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

/*
SwapTokenIDForSessionToken contacts Respoke to activate a session token by tokenID and receive
back the session token, which may then be used to initiate a web socket.
*/
func (client *Client) SwapTokenIDForSessionToken(tokenID string) (sessionToken string, err error) {
	postData := bytes.NewBufferString(`{"tokenId":"` + tokenID + `"}`)
	var sessionData struct {
		Token string `json:"token"`
	}
	resp, err := http.Post("https://api.respoke.io/v1/session-tokens", "application/json", postData)
	if err != nil {
		return sessionToken, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sessionToken, err
	}
	if resp.StatusCode != 200 {
		return sessionToken, errors.New(string(body))
	}
	err = json.Unmarshal(body, &sessionData)
	if err != nil {
		return sessionToken, err
	}
	sessionToken = sessionData.Token
	return sessionToken, err
}
