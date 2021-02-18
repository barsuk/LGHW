package freenfa

import (
	"fmt"
	"testing"
)

var rulebook = RuleBook{Rules: []FARule{
	{1, nil, 2},
	{1, nil, 4},

	{2, "a", 3},
	{3, "a", 2},

	{4, "a", 5},
	{5, "a", 6},
	{6, "a", 4},
}}

func TestNFARulebook(t *testing.T) {
	fmt.Println(rulebook.nextStates(map[int]bool{1: false}, nil))
}

func TestFollowFree(t *testing.T) {
	s := map[int]bool{1:false}
	fmt.Printf("following %[1]v: %v\n", s, rulebook.followFreeMoves(s))
}

func TestFree(t *testing.T) {
	nfades := Design{
		StartState:   1,
		AcceptStates: []int{2, 4},
		RuleBook:     rulebook,
	}
	for _, c := range []string{
		"aa",
		"aaa",
		"aaaa",
		"aaaaa",
		"aaaaaaa",
		"aaaaaaaaaaa",
		"aaaaaaaaaaaa",
		"aaaaaaaaaaaaa",
	} {
		fmt.Printf("str: %s, len: %d, accepted: %t\n", c, len(c), nfades.accepts(c))
	}
}
