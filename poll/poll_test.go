package poll_test

import (
	"testing"

	"github.com/philippta/franky/poll"
)

func TestNew(t *testing.T) {
	p := poll.New("dogs or cats?", []string{
		"dogs",
		"cats",
	})
	if p.Question != "dogs or cats?" {
		t.Fatal("unexpected question: %v", p.Question)
	}
	if p.Votes["dogs"] != 0 {
		t.Fatalf("unexpected votes for dogs: %d", p.Votes["dogs"])
	}
	if p.Votes["cats"] != 0 {
		t.Fatalf("unexpected votes for cats: %d", p.Votes["cats"])
	}
}

func TestVote(t *testing.T) {
	p := poll.New("test", []string{"option"})
	if err := p.Vote("option", "voter"); err != nil {
		t.Fatalf("unexpected error: %v", err.Error())
	}
	if err := p.Vote("option", "voter"); err != poll.ErrorVoterAlreadyVoted {
		t.Fatalf("expected error: %v", poll.ErrorVoterAlreadyVoted)
	}
}
