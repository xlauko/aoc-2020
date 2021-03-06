package main

import (
	"fmt"
	"strconv"
	"utils/utils"
)

type Deck []int

func (self *Deck) enqueue(card int) {
	*self = append(*self, card)
}

func (self *Deck) dequeue() int {
	card := (*self)[0]
	*self = (*self)[1:]
	return card
}

func (self *Deck) copy(n int) Deck {
	deck := make([]int, n)
	copy(deck, (*self)[:n])
	return deck
}

func (self *Deck) score() int {
	res := 0
	length := len(*self)
	for i, card := range *self {
		res += card * (length - i)
	}
	return res
}

func game(p1, p2 Deck, recurse bool) (int, Deck) {
	seen := make(map[int]bool)
	for len(p1) != 0 && len(p2) != 0 {
		state := p1.score() * p2.score()
		if seen[state] {
			return 1, p1
		}
		seen[state] = true

		a := p1.dequeue()
		b := p2.dequeue()

		winner := 1
		if recurse && a <= len(p1) && b <= len(p2) {
			winner, _ = game(p1.copy(a), p2.copy(b), recurse)
		} else if b > a {
			winner = 2
		}

		if winner == 1 {
			p1.enqueue(a)
			p1.enqueue(b)
		} else {
			p2.enqueue(b)
			p2.enqueue(a)
		}
	}

	if len(p1) == 0 {
		return 2, p2
	}
	return 1, p1
}

func deal(file string) []Deck {
	var decks []Deck
	utils.ScanGroup(file, func(group []string) {
		var deck Deck
		for _, line := range group[1:] {
			card, _ := strconv.Atoi(line)
			deck.enqueue(card)
		}
		decks = append(decks, deck)
	})

	return decks
}

func main() {
	decks := deal("big.txt")
	_, winner := game(decks[0], decks[1], false)
	fmt.Println("Part 1: ", winner.score())
	_, winner = game(decks[0], decks[1], true)
	fmt.Println("Part 2: ", winner.score())
}
