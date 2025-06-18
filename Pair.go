package poker

import (
	"cmp"
	"fmt"
)

type pair struct {
	str      string
	cards    []Card
	pairRank CardRank
}

func newPair(hand string, cards []Card) Hand {
	counts := make(map[CardRank]int)
	for _, card := range cards {
		if _, ok := counts[card.rank]; ok {
			counts[card.rank]++
		} else {
			counts[card.rank] = 1
		}
	}

	var pairRank CardRank
	var pairs int

	for rank, count := range counts {
		if count == 2 {
			pairs++
			pairRank = rank
		}
	}

	if pairs == 1 {
		return &pair{hand, cards, pairRank}
	}
	return nil
}

func (p *pair) Cards() []Card {
	return p.cards
}

func (*pair) Rank() HandRank {
	return PAIR
}

func (p *pair) String() string {
	return p.str
}

func (p *pair) Compare(h Hand) int {
	other, ok := h.(*pair)
	if !ok {
		return cmp.Compare(p.Rank(), h.Rank())
	}
	if len(other.cards) != len(p.cards) {
		panic(fmt.Sprintf("invalid number of cards found in pair hand: %d vs %d", len(other.cards), len(p.cards)))
	}

	c := cmp.Compare(p.pairRank, other.pairRank)
	if c != 0 {
		return c
	}

	for i, card := range p.cards {
		c := cmp.Compare(card.rank, other.cards[i].rank)
		if c != 0 {
			return c
		}
	}
	return 0
}
