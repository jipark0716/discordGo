package main

import (
    "os"
    "discord"
)

var seq interface{}

func main() {
    discord.Connect(os.Getenv("DISCORD_BOT_TOKEN"))
    select{}
}
