package main

import (
	"os"
)

func main() {
	slack := NewSlack(os.Getenv("SLACK_API_TOKEN"))
	MessageBus.Subscribe(slack)
	httpd := NewHttpd()
	httpd.Run()
}
