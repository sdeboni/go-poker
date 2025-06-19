package poker

import "cmp"

type fullHouse struct {
	str     string
	cards   []Card
	triplet CardRank
	pair    CardRank
}

func newFullHouse(hand string, cards []Card) Hand {
	counts := make(map[CardRank]int)
	for _, card := range cards {
		if _, ok := counts[card.rank]; ok {
			counts[card.rank]++
		} else {
			counts[card.rank] = 1
		}
	}

	var triplet CardRank
	var pair CardRank

	for rank, count := range counts {
		if count == 2 {
			pair = rank
		} else if count == 3 {
			triplet = rank
		}
	}

	if triplet == 0 || pair == 0 {
		return nil
	}

	return &fullHouse{hand, cards, triplet, pair}
}

func (f *fullHouse) Cards() []Card {
	return f.cards
}

func (*fullHouse) Rank() HandRank {
	return FULL_HOUSE
}

func (f *fullHouse) String() string {
	return f.str
}

func (f *fullHouse) Compare(h Hand) int {
	other, ok := h.(*fullHouse)
	if !ok {
		return cmp.Compare(f.Rank(), h.Rank())
	}

	if c := cmp.Compare(f.triplet, other.triplet); c != 0 {
		return c
	}
	return cmp.Compare(f.pair, other.pair)
}
