// Package franky provides an easy to confiure irc bot handling messages based
// on regular expressions
package franky

import (
	"regexp"

	irc "github.com/thoj/go-ircevent"
)

// Client is an IRC client wrapping the go-ircevent library.
// It connects to a given IRC server and authenticates using a nickname and a
// server password. Once authenticated it automatically joins a configured
// channel and listens on messages.
type Client struct {
	*irc.Connection

	Nick     string
	Addr     string
	Channel  string
	handlers []handler
}

// An Event represents a new message occured in a channel.
// It contains information like the message, sender and the channel it occured on.
// It also holds the client, the event occured on and the matches of a message
// extracted by a regular expression.
type Event struct {
	*irc.Event

	Client  *Client
	Matches []string
}

// A handler represents a function which is called when a message with a given
// regular expression occurs.
type handler struct {
	regexp *regexp.Regexp
	handle func(*Event)
}

// NewClient instantiates a new client with a nickname.
func NewClient(nick string) *Client {
	return &Client{
		Connection: irc.IRC(nick, nick),
		Nick:       nick,
	}
}

// Connect connects to a given IRC server and authenticates using nickname and password
func (c *Client) Connect() {
	c.Connection.Connect(c.Addr)
	c.AddCallback("001", c.welcomeHandler)
	c.AddCallback("PRIVMSG", c.privmsgHandler)
}

// welcomeHandler automatically joins a pre-defined channel once the client receives
// a 001 welcome message.
func (c *Client) welcomeHandler(e *irc.Event) {
	c.Join(c.Channel)
}

// privmsgHandler handles all messages occuring in private conversations or in channels.
// It goes through all registered handlers to process the message which has been sent,
// based on a regular expression.
func (c *Client) privmsgHandler(e *irc.Event) {
	for _, h := range c.handlers {
		matches := h.regexp.FindStringSubmatch(e.Message())
		if matches == nil {
			continue
		}
		h.handle(&Event{
			Event:   e,
			Client:  c,
			Matches: matches,
		})
	}
}

// HandleMessage registers new callback functions for message handling based on a regular
// expression.
func (c *Client) HandleMessage(pattern string, callback func(*Event)) error {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	c.handlers = append(c.handlers, handler{r, callback})
	return nil
}
