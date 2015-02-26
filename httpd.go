package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"log"
)

type Httpd struct {
	Host string
	Port int
}

type Param struct {
	Channel string `form:"channel" binding:"required"`
	Message string `form:"message"`
}

func NewHttpd(host string, port int) *Httpd {
	return &Httpd{
		Host: host,
		Port: port,
	}
}

func (h *Httpd) Run() {
	m := martini.Classic()
	m.Post("/notice", binding.Bind(Param{}), messageHandler)
	m.Post("/privmsg", binding.Bind(Param{}), messageHandler)
	m.RunOnAddr(fmt.Sprintf("%s:%d", h.Host, h.Port))
}

func messageHandler(p Param) (int, string) {
	ch := make(chan error)
	go MessageBus.Publish(NewMessage(p.Channel, p.Message, ch))
	err := <-ch

	if err != nil {
		message := fmt.Sprintf("Failed to send message to %s: %s\n", p.Channel, err)
		log.Printf(message)
		return 400, fmt.Sprintf("[error] %s", message)
	} else {
		return 200, fmt.Sprintf("message sent channel: %s", p.Channel)

	}
}
