package poker

import "cmp"

type fourOfAKind struct {
	str        string
	cards      []Card
	quadRank   CardRank
	singleRank CardRank
}

func newFourOfAKind(hand string, cards []Card) Hand {
	counts := make(map[CardRank]int)
	for _, card := range cards {
		if _, ok := counts[card.rank]; ok {
			counts[card.rank]++
		} else {
			counts[card.rank] = 1
		}
	}

	var quadRank CardRank
	var singleRank CardRank

	for rank, count := range counts {
		if count == 4 {
			quadRank = rank
		} else if count == 1 {
			singleRank = rank
		}
	}

	if quadRank == 0 || singleRank == 0 {
		return nil
	}
	return &fourOfAKind{hand, cards, quadRank, singleRank}
}

func (f *fourOfAKind) Cards() []Card {
	return f.cards
}

func (*fourOfAKind) Rank() HandRank {
	return FOUR_OF_A_KIND
}

func (f *fourOfAKind) String() string {
	return f.str
}

func (f *fourOfAKind) Compare(h Hand) int {
	other, ok := h.(*fourOfAKind)
	if !ok {
		return cmp.Compare(f.Rank(), h.Rank())
	}

	if c := cmp.Compare(f.quadRank, other.quadRank); c != 0 {
		return c
	}
	return cmp.Compare(f.singleRank, other.singleRank)
}
