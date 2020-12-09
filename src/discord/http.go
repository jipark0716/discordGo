package discord

import (
    "log"
    "fmt"
    "bytes"
    "net/http"
    "discord/obj"
    "encoding/json"
)

type Http struct {
    token string
}

const baseUrl = "https://discord.com/api"

func NewHttp(token string) Http {
    return Http{
        token:token,
    }
}

func (this *Http) SendMessage(message obj.Message, channelId string) (*http.Response) {
    return this.Post(fmt.Sprintf("/channels/%s/messages", channelId), message)
}

func (this *Http) Post(url string, body interface{}) (*http.Response) {
    url = baseUrl + url
    bodyB, _ := json.Marshal(body)
    buff := bytes.NewBuffer(bodyB)
    req, _ := http.NewRequest("POST", url, buff)
    this.SetHeader(req)
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        log.Println(err)
        return nil
    }
    return res
}

func (this *Http) SetHeader(req *http.Request) {
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authorization", this.token)
}
