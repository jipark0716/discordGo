package obj

type Message struct {
    content string
    tts bool
    embed map[string]string
}

func NewMessage(content string) Message {
    return Message{
        content: content,
        tts: false,
        embed: nil,
    }
}
