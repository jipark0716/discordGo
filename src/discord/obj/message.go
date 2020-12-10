package obj

type Message struct {
    Content string `json:"content"`
    Tts     bool   `json:"tts"`
    Embed  *Embed  `json:"embed"`
}

type Embed struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

func NewMessage(content string) Message {
    return Message{
        Content: content,
        Tts: false,
        Embed: nil,
    }
}
