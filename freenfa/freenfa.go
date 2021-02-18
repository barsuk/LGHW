package freenfa

import (
	"fmt"
)

type FARule struct {
	State     int
	Character interface{}
	NextState int
}

func (rule FARule) AppliesTo(state int, character interface{}) bool {
	return rule.State == state && rule.Character == character
}

func (rule FARule) Follow() int {
	return rule.NextState
}

func (rule FARule) String() {
	fmt.Printf("#<FARule %d --%s--> %d", rule.State, rule.Character, rule.NextState)
}

type RuleBook struct {
	Rules []FARule
}

type States map[int]bool

func (b RuleBook) followFreeMoves(states States) States {
	moreStates := b.nextStates(states, nil)
	if isSubset(moreStates, states) {
		return states
	}
	union := make(States)
	for st := range states {
		union[st] = false
	}
	for st := range moreStates {
		union[st] = false
	}
	return b.followFreeMoves(union)
}

func (b RuleBook) nextStates(states States, character interface{}) States {
	set := make(States)
	for state := range states {
		for _, newState := range b.followRulesFor(state, character) {
			set[newState] = false
		}
	}
	return set
}

func (b RuleBook) followRulesFor(state int, character interface{}) []int {
	var follows []int
	for _, rule := range b.rulesFor(state, character) {
		follows = append(follows, rule.Follow())
	}
	return follows
}

func (b RuleBook) rulesFor(state int, character interface{}) []FARule {
	var selectedRules []FARule
	for _, rule := range b.Rules {
		if rule.AppliesTo(state, character) {
			selectedRules = append(selectedRules, rule)
		}
	}

	if len(selectedRules) < 1 {
		//panic("нет такого правила")
	}

	return selectedRules
}

type NFA struct {
	currentStates States
	AcceptStates  []int
	RuleBook      RuleBook
}

func (nfa *NFA) SetCurrentStates(states States) {
	nfa.currentStates = states
}

func (nfa *NFA) CurrentStates() States {
	return nfa.RuleBook.followFreeMoves(nfa.currentStates)
}

func (nfa *NFA) accepting() bool {
	for _, v := range nfa.AcceptStates {
		if _, ok := nfa.CurrentStates()[v]; ok {
			return true
		}
	}
	return false
}

func (nfa *NFA) readCharacter(character string) {
	nfa.SetCurrentStates(nfa.RuleBook.nextStates(nfa.CurrentStates(), character))
}

func (nfa *NFA) readString(str string) {
	for _, char := range str {
		nfa.readCharacter(string(char))
	}
}

type Design struct {
	StartState   int
	AcceptStates []int
	RuleBook     RuleBook
}

func (n *Design) toNFA() NFA {
	nfa := NFA{
		AcceptStates:  n.AcceptStates,
		RuleBook:      n.RuleBook,
	}
	nfa.SetCurrentStates(map[int]bool{n.StartState: false})
	return nfa
}

func (n *Design) accepts(str string) bool {
	nfa := n.toNFA()
	nfa.readString(str)
	return nfa.accepting()
}

func isSubset(needle, haystack States) bool {
	for st := range needle {
		//fmt.Printf("итерируем хэш: ключик %[1]T => %[1]v\n", st)
		if _, ok := haystack[st]; !ok {
			return false
		}
	}
	return true
}
