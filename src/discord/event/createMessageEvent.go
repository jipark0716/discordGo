package event

type MessageCreateEvent struct {
    Event
    data map[string]interface{}
}

func NewMessageCreateEvent(data map[string]interface{}) MessageCreateEvent {
    data = data["d"].(map[string]interface{})
    return MessageCreateEvent{
        data: data,
    }
}
