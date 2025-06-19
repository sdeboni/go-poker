package poker

import "cmp"

type straightFlush struct {
	straight Hand
}

func newStraightFlush(hand string, cards []Card) Hand {
	if f := newFlush(hand, cards); f == nil {
		return nil
	}

	if s := newStraight(hand, cards); s == nil {
		return nil
	} else {
		return &straightFlush{s}
	}
}

func (s *straightFlush) Cards() []Card {
	return s.straight.Cards()
}

func (*straightFlush) Rank() HandRank {
	return STRAIGHT_FLUSH
}

func (s *straightFlush) String() string {
	return s.straight.String()
}

func (s *straightFlush) Compare(h Hand) int {
	other, ok := h.(*straightFlush)
	if !ok {
		return cmp.Compare(s.Rank(), h.Rank())
	}

	return s.straight.Compare(other.straight)
}
