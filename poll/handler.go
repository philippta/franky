package poll

import "github.com/philippta/franky"

const (
	NewPollPattern   = `^franky new poll (.*)$`
	NewOptionPattern = `^franky new option (.*)$`
	VotePattern      = `^franky vote (.*)$`
	VotesPattern     = `^franky votes$`
)

var poll *Poll

func NewPollHandler(e *franky.Event) {
	question := e.Matches[1]

	poll = New(question, []string{})
	e.Client.Privmsgf(e.Arguments[0], "I've created a new poll: %v", poll.Question)
}

func NewOptionHandler(e *franky.Event) {
	if poll == nil {
		return
	}
	option := e.Matches[1]
	channel := e.Arguments[0]

	poll.Votes[option] = 0
	e.Client.Privmsgf(channel, "Option %v was added to the poll.", option)
}

func VoteHandler(e *franky.Event) {
	if poll == nil {
		return
	}
	vote := e.Matches[1]
	channel := e.Arguments[0]

	if err := poll.Vote(vote, e.User); err == ErrorVoterAlreadyVoted {
		e.Client.Privmsgf(channel, "Sorry %v, you already voted on this poll. :(", e.Nick)
		return
	} else if err == ErrorOptionNotFound {
		e.Client.Privmsgf(channel, "Sorry %v, this is not a valid option. :(", e.Nick)
		return
	}
	e.Client.Privmsgf(channel, "Thanks for your vote, %v.", e.Nick)
}

func VotesHandler(e *franky.Event) {
	if poll == nil {
		return
	}
	channel := e.Arguments[0]

	e.Client.Privmsgf(channel, "Current poll: %v", poll.Question)
	for option, votes := range poll.Votes {
		e.Client.Privmsgf(channel, "    %v: %d", option, votes)
	}
}
