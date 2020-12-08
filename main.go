package main

import (
    "log"
    "github.com/sacOO7/gowebsocket"
)

func main() {
    socket := gowebsocket.New("wss://gateway.discord.gg/")

    first := true
    socket.OnConnected = func(socket gowebsocket.Socket) {
        log.Println("Connected to server");
    };
    socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
        log.Println("Recieved connect error ", err)
    };
    socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
        if first {
            socket.SendText("{\"op\": 2,\"d\": {\"token\": \"NTcwNTMxMTk4MzQ3MjQ3NjE2.XMAilQ.SzqJgFxGCv7Pjd3Eeox7RPqg9O4\",\"intents\": 513,\"properties\": {\"$os\": \"linux\",\"$browser\": \"my_library\",\"$device\": \"my_library\"}}}")
            first = false
        }
        log.Println("Recieved message " + message)
        println()
    };
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
        log.Println("Disconnected from server ")
        return
    };
    socket.Connect()
    select{}
}
