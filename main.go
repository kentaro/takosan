package main

import (
	"flag"
	"os"
)

var host string
var port int

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "host")
	flag.IntVar(&port, "port", 4979, "port number")
}

func main() {
	slack := NewSlack(os.Getenv("SLACK_API_TOKEN"))
	MessageBus.Subscribe(slack)
	httpd := NewHttpd(host, port)
	httpd.Run()
}
