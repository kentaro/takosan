package main

import (
	"flag"
	"os"
)

var host string
var port int
var name string
var icon string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "host")
	flag.IntVar(&port, "port", 4979, "port number")
	flag.StringVar(&name, "name", "takosan", "bot name")
	flag.StringVar(&icon, "icon", "https://raw.githubusercontent.com/kentaro/takosan/master/takosan.jpg", "icon for takosan")
	flag.Parse()
}

func main() {
	slack := NewSlack(name, icon, os.Getenv("SLACK_API_TOKEN"))
	MessageBus.Subscribe(slack)
	httpd := NewHttpd(host, port)
	httpd.Run()
}
