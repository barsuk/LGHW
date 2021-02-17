package nfa

import "LGHW/dka"

type RuleBook struct{
	Rules []dka.FARule
}

type States map[int]bool

func (b RuleBook) nextStates(states States, character string) States {
	set := make(States)
	for state, _ := range states {
		for _, newState := range b.followRulesFor(state, character) {
			set[newState] = false
		}
	}
	return set
}

func (b RuleBook) followRulesFor(state int, character string) []int {
	var follows []int
	for _, rule := range b.rulesFor(state, character) {
		follows = append(follows, rule.Follow())
	}
	return follows
}

func (b RuleBook) rulesFor(state int, character string) []dka.FARule {
	var selectedRules []dka.FARule
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
	CurrentStates States
	AcceptStates []int
	RuleBook RuleBook
}

func (nfa *NFA) accepting() bool {
	for _, v := range nfa.AcceptStates {
		if _, ok := nfa.CurrentStates[v]; ok {
			return true
		}
	}
	return false
}

func (nfa *NFA) readCharacter(character string) {
	nfa.CurrentStates = nfa.RuleBook.nextStates(nfa.CurrentStates, character)
}

func (nfa *NFA) readString(str string) {
	for _, char := range str {
		nfa.readCharacter(string(char))
	}
}