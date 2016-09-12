package respoke

import (
	"encoding/json"
	"log"
	"runtime"

	"github.com/ruffrey/go-socketio09"
)

const goRespokeVersion = "0.1.0"

// SDK is the user agent type used as an informational header to the Respoke backend
var SDK = "GoRespoke%2f" + goRespokeVersion + "%20(" + runtime.GOOS + "-" +
	runtime.GOARCH + ")%20" + runtime.Version()

type respokeHeaders struct {
	RespokeSDK string `json:"Respoke-SDK,omitempty"`
	AppToken   string `json:"App-Token,omitempty"`
	AppSecret  string `json:"App-Secret,omitempty"`
	AdminToken string `json:"Admin-Token,omitempty"`
	EndpointID string `json:"endpointId, omitempty"`
}

/*
Presence is the respoke ws presence format
*/
type Presence struct {
	Type string `json:"type"`
}

type wsData struct {
	EndpointID string   `json:"endpointId,omitempty"`
	AppSecret  string   `json:"appSecret,omitempty"`
	AppID      string   `json:"appId,omitempty"`
	Groups     []string `json:"groups,omitempty"`
	Presence   Presence `json:"presence,omitempty"`
}

type wsBody struct {
	URL     string         `json:"url"`
	Headers respokeHeaders `json:"headers"`
	Data    interface{}    `json:"data"`
}

func (b wsBody) EncodeAsString() string {
	bodyBytes, _ := json.Marshal(b)
	return string(bodyBytes[:len(bodyBytes)])
}

/*
OneToOneMessage is the parsed JSON received on respoke pubsub event
*/
type OneToOneMessage struct {
	Message interface{} `json:"message"`
	From    string      `json:"from"`
}

/*
GroupMessage is the parsed JSON received on respoke pubsub event
*/
type GroupMessage struct {
	Message interface{} `json:"message"`
	GroupID string      `json:"groupId"`
	From    string      `json:"from"`
}

/*
Client is main exported
*/
type Client struct {
	Headers respokeHeaders
	Socket  *socketio09.SocketIOClient
}

/*
ConnectAsEndpoint is for a non-admin end user to make a web socket connection to Respoke.
*/
func (client *Client) ConnectAsEndpoint(endpointID string, appToken string) (err error) {
	client.Headers.RespokeSDK = SDK
	client.Headers.AppToken = appToken
	client.Headers.EndpointID = endpointID
	urlWithToken := "https://api.respoke.io/socket.io/1/" +
		"?__sails_io_sdk_version=0.10.0&Respoke-SDK=" + SDK + "&&app-token=" + appToken

	log.Println("respoke: connecting via websocket", urlWithToken)
	wst := socketio09.NewConnection()
	// initial connection
	client.Socket, err = wst.Connect(urlWithToken)
	if err != nil {
		log.Println("respoke: dial failed")
		return err
	}
	log.Println("respoke: web socket is alive")
	connectData := wsData{EndpointID: endpointID}
	connectBody := wsBody{
		URL:     "/v1/connections",
		Headers: client.Headers,
		Data:    connectData,
	}
	log.Println("respoke: posting to connections", connectBody)
	err = client.Socket.Emit("post", connectBody.EncodeAsString())
	if err != nil {
		log.Println("respoke: connections post after after successful handshake failed", err)
		return err
	}
	log.Println("respoke: connections post OK")

	return nil
}

/*
Join a group
*/
func (client *Client) Join(groupID string) (err error) {
	_, err = client.Socket.EmitWithAck("post", wsBody{
		URL:     "/v1/channels/" + groupID + "/subscribers",
		Headers: client.Headers,
		Data: wsData{
			EndpointID: client.Headers.EndpointID,
		},
	})
	return err
}

/*
SendMsg sends a message to a respoke user endpoint
*/
func (client *Client) SendMsg(endpointID string, message interface{}) (err error) {
	_, err = client.Socket.EmitWithAck("post", wsBody{
		URL:     "/v1/messages",
		Headers: client.Headers,
		Data:    message,
	})
	return err
}

/*
SendGroupMsg sends a message to a respoke group
*/
func (client *Client) SendGroupMsg(groupID string, message interface{}) (err error) {
	_, err = client.Socket.EmitWithAck("post", wsBody{
		URL:     "/v1/channels/" + groupID + "/publish",
		Headers: client.Headers,
		Data:    message,
	})
	return err
}
