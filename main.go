package main

import (
    "log"
    "github.com/sacOO7/gowebsocket"
    "os"
    "os/signal"
)

func main() {
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)

    socket := gowebsocket.New("wss://gateway.discord.gg/")

    socket.OnConnected = func(socket gowebsocket.Socket) {
        log.Println("Connected to server");
    };
    socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
        log.Println("Recieved connect error ", err)
    };
    socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
        socket.SendText("{\"op\": 2,\"d\": {\"token\": \"NTcwNTMxMTk4MzQ3MjQ3NjE2.XMAilQ.EwooL9k3rFh63P7nL1CJ_e2JxXw\",\"intents\": 513,\"properties\": {\"$os\": \"linux\",\"$browser\": \"my_library\",\"$device\": \"my_library\"}}}")
        log.Println("Recieved message " + message)
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
    for {
        select {
        case <-interrupt:
            log.Println("interrupt")
            socket.Close()
            return
        }
    }
}
