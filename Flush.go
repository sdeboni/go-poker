package poker

import "cmp"

type flush struct {
	str   string
	cards []Card
}

func newFlush(hand string, cards []Card) Hand {
	for _, card := range cards[1:] {
		if card.suit != cards[0].suit {
			return nil
		}
	}

	return &flush{hand, cards}
}

func (f *flush) Cards() []Card {
	return f.cards
}

func (*flush) Rank() HandRank {
	return FLUSH
}

func (f *flush) String() string {
	return f.str
}

func (f *flush) Compare(h Hand) int {
	other, ok := h.(*flush)
	if !ok {
		return cmp.Compare(f.Rank(), h.Rank())
	}

	for i, card := range f.cards {
		c := cmp.Compare(card.rank, other.Cards()[i].rank)
		if c != 0 {
			return c
		}
	}
	return 0
}
