package main

import (
	"github.com/philippta/franky"
	"github.com/philippta/franky/poll"
)

func main() {
	c := franky.NewClient("franky")
	c.Addr = "irc.freenode.org:6667"
	c.Channel = "#frankytest"
	c.Debug = true
	// c.UseTLS = true
	// c.Password = "password"

	// Handle command to join a new channel
	c.HandleMessage(`franky join (#\w+)`, func(e *franky.Event) {
		e.Client.Join(e.Matches[1])
	})

	// Handle command to leave a channel
	c.HandleMessage(`franky leave (#\w+)`, func(e *franky.Event) {
		e.Client.Part(e.Matches[1])
	})

	// Quit the connection to IRC server once a quit command is received
	quit := make(chan int)
	c.HandleMessage(`franky quit`, func(e *franky.Event) {
		e.Client.Quit()
		quit <- 1
	})

	// Register all poll handlers
	c.HandleMessage(poll.NewPollPattern, poll.NewPollHandler)
	c.HandleMessage(poll.NewOptionPattern, poll.NewOptionHandler)
	c.HandleMessage(poll.VotePattern, poll.VoteHandler)
	c.HandleMessage(poll.VotesPattern, poll.VotesHandler)

	// Start the connection to the IRC Server
	c.Connect()

	// Wait for the quit command before exiting
	<-quit
}
