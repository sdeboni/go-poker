package poker

import "cmp"

type straight struct {
	str   string
	cards []Card
}

func newStraight(hand string, cards []Card) Hand {
	prevRank := cards[0].rank
	for _, card := range cards[1:] {
		if card.rank != prevRank+1 {
			return nil
		}
		prevRank = card.rank
	}

	return &straight{hand, cards}
}

func (s *straight) Cards() []Card {
	return s.cards
}

func (*straight) Rank() HandRank {
	return STRAIGHT
}

func (s *straight) String() string {
	return s.str
}

func (s *straight) Compare(h Hand) int {
	other, ok := h.(*straight)
	if !ok {
		return cmp.Compare(s.Rank(), h.Rank())
	}

	return cmp.Compare(s.cards[0].rank, other.Cards()[0].rank)
}
