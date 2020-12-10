package main

import (
    "os"
    "discord"
    "discord/obj"
    "discord/event"
)

var seq interface{}

func main() {
    client := discord.Connect(os.Getenv("DISCORD_BOT_TOKEN"))
    client.EventAdapter.OnMessageCreate = func(event event.MessageCreateEvent) {
        event.Reply(obj.NewMessage("test"))
    }
    select{}
}
