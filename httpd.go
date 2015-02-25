package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

type Httpd struct {
}

type Param struct {
	Channel string `form:"channel" binding:"required"`
	Message string `form:"message"`
}

func NewHttpd() *Httpd {
	return &Httpd{}
}

func (h *Httpd) Run() {
	m := martini.Classic()
	m.Post("/notice", binding.Bind(Param{}), messageHandler)
	m.Post("/privmsg", binding.Bind(Param{}), messageHandler)
	m.RunOnAddr(":8080")
}

func messageHandler(p Param) string {
	go MessageBus.Publish(NewMessage(p.Channel, p.Message))
	return fmt.Sprintf("message sent channel: %s %s", p.Channel, p.Message)
}
