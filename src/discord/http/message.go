package http

import (
    "fmt"
    "net/http"
    "discord/obj"
)

func (this *Http) SendMessage(message obj.Message, channelId string) (*http.Response) {
    return this.Post(fmt.Sprintf("/channels/%s/messages", channelId), message)
}
