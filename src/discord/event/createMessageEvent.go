package event

import (
    "discord/obj"
    "discord/http"
)

type MessageCreateEvent struct {
    Event
    data map[string]interface{}
}

func NewMessageCreateEvent(data map[string]interface{}, http http.Http) MessageCreateEvent {
    data = data["d"].(map[string]interface{})
    return MessageCreateEvent{
        data: data,
        Event: Event{
            http : http,
        },
    }
}

func (this *MessageCreateEvent) Reply(msg obj.Message) {
    this.Event.http.SendMessage(msg, this.data["channel_id"].(string))
}
