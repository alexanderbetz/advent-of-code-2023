package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardStrengths = "J23456789TQKA"

type Hand struct {
	hand string
	bid  int
}

func (hand Hand) compareTo(compare Hand) bool {
	h1 := hand.calcHandValue()
	h2 := compare.calcHandValue()

	if h1 != h2 {
		return h1 < h2
	}

	// compare card by card
	for i := 0; i < 5; i++ {
		if hand.hand[i] == compare.hand[i] {
			continue
		}
		card1Strength := strings.Index(cardStrengths, string(hand.hand[i]))
		card2Strength := strings.Index(cardStrengths, string(compare.hand[i]))
		return card1Strength < card2Strength
	}

	return true
}

func (hand Hand) calcHandValue() int {
	var cardCounts []int

	for i := 1; i < len(cardStrengths); i++ {
		occurrences := strings.Count(hand.hand, string(cardStrengths[i]))
		if occurrences > 0 {
			cardCounts = append(cardCounts, occurrences)
		}
	}

	// hand full of jokers
	if len(cardCounts) == 0 {
		return 6
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cardCounts)))

	jokerCount := strings.Count(hand.hand, "J")
	cardCounts[0] += jokerCount

	handValue := 0
	if len(cardCounts) == 1 {
		// five of a kind
		handValue = 6
	} else if len(cardCounts) == 2 {
		if cardCounts[0] == 4 {
			// four of a kind
			handValue = 5
		} else if cardCounts[0] == 3 && cardCounts[1] == 2 {
			// full house
			handValue = 4
		}
	} else if len(cardCounts) == 3 {
		if cardCounts[0] == 3 {
			// three of a kind
			handValue = 3
		} else if cardCounts[0] == 2 && cardCounts[1] == 2 {
			// two pair
			handValue = 2
		}
	} else if len(cardCounts) == 4 {
		handValue = 1
	}

	return handValue
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var hands []Hand

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		hand := new(Hand)
		hand.hand = line[0]
		n, _ := strconv.Atoi(line[1])
		hand.bid = n
		hands = append(hands, *hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compareTo(hands[j])
	})

	var rankSums int = 0
	for i := 0; i < len(hands); i++ {
		rankSums += (i + 1) * hands[i].bid
	}

	fmt.Printf("(Challenge 2) Sum of ranks and bids: %d\n", rankSums)
}

func Min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func Max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
