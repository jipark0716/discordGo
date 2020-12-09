package main

import (
    // "fmt"
    "log"
    "encoding/json"
    "github.com/sacOO7/gowebsocket"
)

var seq interface{}

func main() {
    socket := gowebsocket.New("wss://gateway.discord.gg/")

    socket.OnConnected = func(socket gowebsocket.Socket) {
        log.Println("Connected to server");
    };
    socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
        log.Println("Recieved connect error ", err)
    };
    socket.OnTextMessage = onTextMessageEvent;
    socket.OnBinaryMessage = func(data [] byte, socket gowebsocket.Socket) {
        log.Println("Recieved binary data ", data)
    };
    socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
        log.Println("Recieved ping " + data)
    };
    socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
        log.Println("Recieved pong " + data)
    };
    socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
        log.Println(socket)
        log.Println("Disconnected from server ")
        return
    };
    socket.Connect()
    select{}
}

func onTextMessageEvent (message string, socket gowebsocket.Socket) {
    var data map[string]interface{}
    err := json.Unmarshal([]byte(message), &data);
    if err != nil {
        panic(err)
    }
    if data["t"] == nil {
        req := map[string]interface{}{
            "op" : 2,
            "d" : map[string]interface{}{
                "token" : "NTcwNTMxMTk4MzQ3MjQ3NjE2.XMAilQ.UaKV3WD77ynJu8GQ015KeR9sSuY",
                "intents" : 513,
                "properties" : map[string]interface{}{
                    "$os" : "window",
                    "$browser" : "lib",
                    "$device" : "lib",
                },
            },
        }
        seq = data["d"].(map[string]interface{})["heartbeat_interval"]
        reqT, _ := json.Marshal(req)
        println("op2")
        socket.SendText(string(reqT))
    } else if data["op"].(int) == 10 {
        req := map[string]interface{}{
            "op" : 6,
            "d" : map[string]interface{}{
                "token" : "NTcwNTMxMTk4MzQ3MjQ3NjE2.XMAilQ.UaKV3WD77ynJu8GQ015KeR9sSuY",
                "session_id" : data["d"].(map[string]interface{})["session_id"],
                "seq" : seq,
            },
        }
        reqT, _ := json.Marshal(req)
        println("op6")
        socket.SendText(string(reqT))
    }
    log.Println("Recieved text message", message)
    println()
}
