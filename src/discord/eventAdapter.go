package discord

import (
    "fmt"
    "discord/event"
    "encoding/json"
    "github.com/sacOO7/gowebsocket"
)

type EventAdapter struct {
    OnMessage          func(data map[string]interface{}, soket gowebsocket.Socket)

    OnDispatch         func(data map[string]interface{}, soket gowebsocket.Socket)
    OnHeartbeat        func(data map[string]interface{}, soket gowebsocket.Socket)
    OnIdentify         func(data map[string]interface{}, soket gowebsocket.Socket)
    OnPresenceUpdate   func(data map[string]interface{}, soket gowebsocket.Socket)
    OnVoiceStateUpdate func(data map[string]interface{}, soket gowebsocket.Socket)
    OnVoiceServerPing  func(data map[string]interface{}, soket gowebsocket.Socket)
    OnResume           func(data map[string]interface{}, soket gowebsocket.Socket)
    OnReconnect        func(data map[string]interface{}, soket gowebsocket.Socket)
    OnGuildMemberChunk func(data map[string]interface{}, soket gowebsocket.Socket)
    OnInvalidSession   func(data map[string]interface{}, soket gowebsocket.Socket)
    OnHello            func(data map[string]interface{}, soket gowebsocket.Socket)
    OnHeartbeatAck     func(data map[string]interface{}, soket gowebsocket.Socket)

    OnReady            func(data map[string]interface{}, soket gowebsocket.Socket)
    OnMessageCreate    func(event event.MessageCreateEvent)
}

func NewEventAdapter() EventAdapter {
    return EventAdapter{}
}

func (this *EventAdapter) onTextMessageEvent(message string, socket gowebsocket.Socket) {
    var data map[string]interface{}
    err := json.Unmarshal([]byte(message), &data)
    if err != nil {
        panic(err)
    }
    println()
    println(message)
    println()
    go func(data map[string]interface{}, socket gowebsocket.Socket) {
        if this.OnMessage != nil {
            this.OnMessage(data, socket)
        }
    }(data, socket)
    go func(data map[string]interface{}, op string, socket gowebsocket.Socket) {
        switch op {
            case "0":
                if this.OnDispatch != nil {
                    this.OnDispatch(data, socket)
                }
            case "1":
                if this.OnHeartbeat != nil {
                    this.OnHeartbeat(data, socket)
                }
            case "2":
                if this.OnIdentify != nil {
                    this.OnIdentify(data, socket)
                }
            case "3":
                if this.OnPresenceUpdate != nil {
                    this.OnPresenceUpdate(data, socket)
                }
            case "4":
                if this.OnVoiceStateUpdate != nil {
                    this.OnVoiceStateUpdate(data, socket)
                }
            case "5":
                if this.OnVoiceServerPing != nil {
                    this.OnVoiceServerPing(data, socket)
                }
            case "6":
                if this.OnResume != nil {
                    this.OnResume(data, socket)
                }
            case "7":
                if this.OnReconnect != nil {
                    this.OnReconnect(data, socket)
                }
            case "8":
                if this.OnGuildMemberChunk != nil {
                    this.OnGuildMemberChunk(data, socket)
                }
            case "9":
                if this.OnInvalidSession != nil {
                    this.OnInvalidSession(data, socket)
                }
            case "10":
                if this.OnHello != nil {
                    this.OnHello(data, socket)
                }
            case "11":
                if this.OnHeartbeatAck != nil {
                    this.OnHeartbeatAck(data, socket)
                }
        }
    }(data, fmt.Sprintf("%#v", data["op"]), socket)

    go func(data map[string]interface{}, t string, socket gowebsocket.Socket){
        switch t {
            case "READY":
                if this.OnReady != nil {
                    this.OnReady(data, socket)
                }
            case "MESSAGE_CREATE":
                if this.OnMessageCreate != nil {
                    this.OnMessageCreate(event.NewMessageCreateEvent(data))
                }
        }
    }(data, fmt.Sprintf("%v", data["t"]), socket)
}
