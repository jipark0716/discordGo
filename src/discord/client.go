package discord

import (
    "log"
    "fmt"
    "time"
    "strconv"
    "discord/http"
    "encoding/json"
    "github.com/sacOO7/gowebsocket"
)

type Client struct {
    socket gowebsocket.Socket
    token string
    EventAdapter *EventAdapter
    heartbeat int
    sessionId string
    seq int
    Http http.Http
}

func Connect(token string) Client {
    socket := gowebsocket.New("wss://gateway.discord.gg/")
    this := Client{
        socket: socket,
        token: token,
        Http:http.NewHttp(token),
    }
    this.EventAdapter = NewEventAdapter(this.Http)
    this.EventAdapter.OnHello = this.OnHello
    this.EventAdapter.OnReady = this.OnReady
    this.EventAdapter.OnMessage = this.OnMessage
    socket.OnTextMessage = this.EventAdapter.onTextMessageEvent
    socket.OnConnected = func(socket gowebsocket.Socket) {
        log.Println("Connected to server");
    }
    socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
        log.Println("Recieved connect error ", err)
    }
    socket.OnBinaryMessage = func(data [] byte, socket gowebsocket.Socket) {
        log.Println("Recieved binary data ", data)
    }
    socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
        log.Println("Recieved ping " + data)
    }
    socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
        log.Println("Recieved pong " + data)
    }
    socket.OnDisconnected = this.reconnect
    socket.Connect()
    return this
}

func (this *Client) reconnect(err error, socket gowebsocket.Socket) {
    log.Println("Disconnected from server ")
    socket.Connect()
}

func SendJson(data map[string]interface{}, socket gowebsocket.Socket) {
    dataB, _ := json.Marshal(data)
    socket.SendText(string(dataB))
}

func (this *Client) identify() (map[string]interface{}) {
    return map[string]interface{}{
        "op" : 2,
        "d" : map[string]interface{}{
            "token" : this.token,
            "intents" : 513,
            "properties" : map[string]interface{}{
                "$os" : "window",
                "$browser" : "lib",
                "$device" : "lib",
            },
        },
    }
}

func (this *Client) setHeartbeat(intval int, socket gowebsocket.Socket) {
    if intval == 0 {
        return
    }
    for {
        time.Sleep(time.Millisecond * time.Duration(intval))
        payload := map[string]interface{}{
            "op" : 1,
            "d" : this.seq,
        }
        SendJson(payload, socket)
    }
}

func (this *Client) OnHello(data map[string]interface{}, socket gowebsocket.Socket) {
    intval, _ := strconv.Atoi(fmt.Sprintf("%v", data["d"].(map[string]interface{})["heartbeat_interval"]))
    go this.setHeartbeat(intval, socket)
    SendJson(this.identify(), socket)
}

func (this *Client) OnReady(data map[string]interface{}, socket gowebsocket.Socket) {
    sessionId, _ := data["d"].(map[string]interface{})["session_id"].(string)
    this.sessionId = sessionId
}

func (this *Client) OnMessage(data map[string]interface{}, socket gowebsocket.Socket) {
    if seq, ok := data["s"].(int); ok {
        this.seq = seq
    }
}
