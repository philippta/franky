// Package poll provides an easy way to create polls
package poll

import "errors"

var (
	ErrorOptionNotFound    = errors.New("Option not found")
	ErrorVoterAlreadyVoted = errors.New("Voter already voted")
)

// A poll represents a poll containing a question and options to answer
type Poll struct {
	Question string
	Votes    map[string]int
	Voters   []string
}

// New instantiates a new poll with 0 votes
func New(question string, options []string) *Poll {
	var p Poll
	p.Question = question
	p.Votes = make(map[string]int)
	for _, o := range options {
		p.Votes[o] = 0
	}
	return &p
}

// Vote increases the votes on poll and adds the voter to the list of voters
func (p *Poll) Vote(option string, voter string) error {
	if hasVoter(p, voter) {
		return ErrorVoterAlreadyVoted
	}
	if _, ok := p.Votes[option]; !ok {
		return ErrorOptionNotFound
	}
	p.Votes[option]++
	p.Voters = append(p.Voters, voter)
	return nil
}

// hasVoter checks if a voter has already votes on a poll
func hasVoter(p *Poll, voter string) bool {
	for _, v := range p.Voters {
		if v == voter {
			return true
		}
	}
	return false
}
