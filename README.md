franky
====

Franky is an IRC bot designed to easily handle commands based on regular expressions.

## Getting Started

### Installing

To start using franky, install Go and run `go get`:
```sh
$ go get github.com/philippta/franky
```

### Connecting to an IRC Server

To connect to an IRC Server, simply instantiate a new Client and use the `Connect()` function:
```go
package main

import "github.com/philippta/franky"

func main() {
	// Instantiate a new Client
	c := franky.NewClient("franky")
	c.Addr = "irc.freenode.org:6667"
	c.Channel = "#franky"

	// Connect to the IRC Server
	c.Connect()

	// Sleep indefinitely
	select {}
}
```

### Handling Messages

To handle messages, register a handler with a regular expression.
When a message on an IRC channel matches the given regular expression, the handler will be called with an Event.
The Event contains data about the message, the sender and the channel it occured on and also the client which received the message.
It also contains a `[]string` of the matches parsed by the regular expression.

```go
c.HandleMessage(`franky join (#\w+)`, func(e *franky.Event) {
	// log the actually message
	log.Println(e.Message())

	// join a new IRC channel
	channel := e.Matches[1] // []string{"franky join #channel", "#channel"}
	e.Client.Join(channel)
})
```

### Executing IRC specific commands

Franky is a wrapper around the https://github.com/thoj/go-ircevent library.
Therefore it contains a number of functions to interact with the IRC server.
Here are a few examples.

```go
c := franky.NewClient("franky")
// further client configuration...

c.Join("#channel") // Join a channel
c.Part("#channel") // Leave a channel
c.Notice("nickname", "message") // Send a notification to a nickname
c.Action("nickname", "message") // Send an action to a nickname
c.Privmsg("nickname/#channel", "message") // Send a private message to a nick or channel
c.Kick("user", "#channel", "message") // Kick a user from a channel
c.Nick("franky2") // Change franky's nickname
c.Disconnect() // Disconnets from an IRC server
```
