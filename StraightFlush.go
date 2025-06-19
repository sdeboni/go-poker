package poker

import "cmp"

type straightFlush struct {
	str   string
	cards []Card
}

func newStraightFlush(hand string, cards []Card) Hand {
	prevRank := cards[0].rank

	for _, card := range cards[1:] {
		if card.rank != prevRank+1 {
			return nil
		}
		if card.suit != cards[0].suit {
			return nil
		}
		prevRank = card.rank
	}

	return &straightFlush{hand, cards}
}

func (s *straightFlush) Cards() []Card {
	return s.cards
}

func (*straightFlush) Rank() HandRank {
	return STRAIGHT_FLUSH
}

func (s *straightFlush) String() string {
	return s.str
}

func (s *straightFlush) Compare(h Hand) int {
	other, ok := h.(*straightFlush)
	if !ok {
		return cmp.Compare(s.Rank(), h.Rank())
	}

	return cmp.Compare(s.cards[0].rank, other.Cards()[0].rank)
}
