package poker

import (
  "cmp"
  "fmt"
)

type twoPair struct {
  str string
  cards []Card
  pairRank [2]CardRank
}

func newTwoPair(hand string, cards []Card) Hand {
  counts := make(map[CardRank]int)
  for _, card := range cards {
    if _, ok := counts[card.rank]; ok {
      counts[card.rank]++
    } else {
      counts[card.rank] = 1
    }
  }

  var pairRank [2]CardRank
  var pairs int

  for rank, count := range counts {
    if count == 2 {
      pairRank[pairs] = rank
      pairs++
    }
  }

  if pairRank[0] < pairRank[1] {
    pairRank[0], pairRank[1] = pairRank[1], pairRank[0]
  } 

  if pairs == 2 {
    return &twoPair{hand, cards, pairRank}
  }
  return nil
}

func (p *twoPair) Cards() []Card {
  return p.cards
}

func (*twoPair) Rank() HandRank {
  return TWO_PAIR
}

func (p *twoPair) String() string {
  return p.str
}

func (p *twoPair) Compare(h Hand) int {
  other, ok := h.(*twoPair)
  if !ok {
    return cmp.Compare(p.Rank(), h.Rank())
  }

  if len(other.cards) != len(p.cards) {
    panic(fmt.Sprintf("invalid number of cards found in pair hand: %d vs %d", len(other.cards), len(p.cards)))
  }

  if c := cmp.Compare(p.pairRank[0], other.pairRank[0]); c != 0 {
    return c
  } 

  if c := cmp.Compare(p.pairRank[1], other.pairRank[1]); c != 0 {
    return c
  }

  for i, card := range p.cards {
    if c := cmp.Compare(card.rank, other.cards[i].rank); c != 0 {
      return c
    }
  }
  return 0
}
