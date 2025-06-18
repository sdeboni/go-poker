package poker

import (
	"cmp"
	"fmt"
)

type highCard struct {
	str   string
	cards []Card
}

func newHighCard(str string, cards []Card) *highCard {
	return &highCard{str, cards}
}

func (hc *highCard) Cards() []Card {
	return hc.cards
}

func (*highCard) Rank() HandRank {
	return HIGH_CARD
}

func (hc *highCard) String() string {
	return hc.str
}

func (hc *highCard) Compare(h Hand) int {
	other, ok := h.(*highCard)
	if !ok {
		return cmp.Compare(hc.Rank(), h.Rank())
	}
	if len(other.cards) != len(hc.cards) {
		panic(fmt.Sprintf("invalid number of cards found in highCard hand: %d vs %d", len(other.cards), len(hc.cards)))
	}
	for i, card := range hc.cards {
		c := cmp.Compare(card.rank, other.cards[i].rank)
		if c != 0 {
			return c
		}
	}
	return 0
}
