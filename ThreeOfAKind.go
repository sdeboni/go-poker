package poker

import (
	"cmp"
	"fmt"
)

type threeOfAKind struct {
	str         string
	cards       []Card
	tripletRank CardRank
}

func newThreeOfAKind(hand string, cards []Card) Hand {
	counts := make(map[CardRank]int)
	for _, card := range cards {
		if _, ok := counts[card.rank]; ok {
			counts[card.rank]++
		} else {
			counts[card.rank] = 1
		}
	}

	var tripletRank CardRank

	for rank, count := range counts {
		if count == 3 {
			tripletRank = rank
			break
		}
	}

	if tripletRank == 0 {
		return nil
	}
	return &threeOfAKind{hand, cards, tripletRank}
}

func (t *threeOfAKind) Cards() []Card {
	return t.cards
}

func (*threeOfAKind) Rank() HandRank {
	return THREE_OF_A_KIND
}

func (t *threeOfAKind) String() string {
	return t.str
}

func (t *threeOfAKind) Compare(h Hand) int {
	other, ok := h.(*threeOfAKind)
	if !ok {
		return cmp.Compare(t.Rank(), h.Rank())
	}

	if len(other.cards) != len(t.cards) {
		panic(fmt.Sprintf("invalid number of cards found in pair hand: %d vs %d", len(other.cards), len(t.cards)))
	}

	return cmp.Compare(t.tripletRank, other.tripletRank)
}
