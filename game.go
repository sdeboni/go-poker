package poker

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type CardRank int

const (
	TWO CardRank = iota + 2
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

type Suit int

const (
	HEARTS Suit = iota + 1
	CLUBS
	SPADES
	DIAMONDS
)

type HandRank int

const (
	HIGH_CARD HandRank = iota + 1
	PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
)

type Hand interface {
	Compare(Hand) int
	Cards() []Card
	Rank() HandRank
	String() string
}

type Card struct {
	rank CardRank
	suit Suit
}

func (c *Card) String() string {
	return cardRankToString(c.rank) + string(suitToRune(c.suit))
}

func BestHand(str []string) ([]string, error) {
	hands, err := parseHands(str)
	if err != nil {
		return nil, err
	}

	cardCounts := getCountsByCard(hands)
	for card, count := range cardCounts {
		if count > 1 {
			return nil, fmt.Errorf("card %s used %d times", card.String(), count)
		}
	}

	slices.SortStableFunc(hands, func(a, b Hand) int {
		return -a.Compare(b)
	})

	result := []string{hands[0].String()}

	for _, hand := range hands[1:] {
		if hand.Compare(hands[0]) == 0 {
			result = append(result, hand.String())
		} else {
			break
		}
	}

	return result, nil
}

func parseHands(arr []string) ([]Hand, error) {
	hands := make([]Hand, 0, len(arr))

	for _, str := range arr {
		if hand, err := parseHand(str); err != nil {
			return nil, err
		} else {
			hands = append(hands, hand)
		}
	}
	return hands, nil
}

func parseHand(str string) (Hand, error) {
	str = strings.TrimSpace(str)

	cards, err := parseCards(str)
	if err != nil {
		return nil, err
	}

	if len(cards) != 5 {
		return nil, fmt.Errorf("invalid hand '%s': expected 5 cards, found: %d", str, len(cards))
	}

	unsortedNormalFormHand := normalFormHand(cards)

	slices.SortFunc(cards, func(a, b Card) int {
		return cmp.Compare(a.rank, b.rank)
	})

	if hand := newStraightFlush(unsortedNormalFormHand, cards); hand != nil {
		return hand, nil
	}
	if hand := newFourOfAKind(unsortedNormalFormHand, cards); hand != nil {
		return hand, nil
	}
	if hand := newFullHouse(unsortedNormalFormHand, cards); hand != nil {
		return hand, nil
	}
	if hand := newFlush(unsortedNormalFormHand, cards); hand != nil {
		return hand, nil
	}
	if hand := newStraight(unsortedNormalFormHand, cards); hand != nil {
		return hand, nil
	}
	if hand := newThreeOfAKind(unsortedNormalFormHand, cards); hand != nil {
		return hand, nil
	}
	if hand := newTwoPair(unsortedNormalFormHand, cards); hand != nil {
		return hand, nil
	}
	if hand := newPair(unsortedNormalFormHand, cards); hand != nil {
		return hand, nil
	}
	return newHighCard(unsortedNormalFormHand, cards), nil
}

func parseCards(hand string) ([]Card, error) {
	strCards := strings.Split(hand, " ")
	cards := make([]Card, 0, len(strCards))

	cardCounts := make(map[Card]int)

	for _, str := range strCards {
		if isEmpty(str) {
			continue
		}
		card, err := parseCard(str)
		if err != nil {
			return nil, err
		}
		if _, ok := cardCounts[card]; ok {
			return nil, fmt.Errorf("invalid hand %s: duplicate card %s found", hand, str)
		} else {
			cardCounts[card] = 1
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func normalFormHand(cards []Card) string {
	var sb strings.Builder
	sb.WriteString(cards[0].String())
	for _, card := range cards[1:] {
		sb.WriteString(" ")
		sb.WriteString(card.String())
	}
	return sb.String()
}

func isEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func parseCard(str string) (card Card, err error) {
	var rank CardRank
	var suit Suit

	chars := []rune(str)

	switch len(chars) {
	case 2:
		if rank, err = getCardRank(chars[0:1]); err == nil {
			suit, err = getSuit(chars[1])
		}
	case 3:
		if rank, err = getCardRank(chars[0:2]); err == nil {
			suit, err = getSuit(chars[2])
		}
	default:
		err = fmt.Errorf("invalid card: '%s'", str)
	}

	card = Card{rank, suit}
	return
}

func getCardRank(chars []rune) (rank CardRank, err error) {
	if len(chars) == 2 {
		if chars[0] == '1' && chars[1] == '0' {
			rank = TEN
		} else {
			err = fmt.Errorf("invalid card rank (length %d): '%s'", len(chars), string(chars))
		}
		return
	} else if len(chars) != 1 {
		err = fmt.Errorf("invalid card rank (length %d) : '%s'", len(chars), string(chars))
		return
	}

	switch chars[0] {
	case '2':
		rank = TWO
	case '3':
		rank = THREE
	case '4':
		rank = FOUR
	case '5':
		rank = FIVE
	case '6':
		rank = SIX
	case '7':
		rank = SEVEN
	case '8':
		rank = EIGHT
	case '9':
		rank = NINE
	case 'J':
		rank = JACK
	case 'Q':
		rank = QUEEN
	case 'K':
		rank = KING
	case 'A':
		rank = ACE
	default:
		err = fmt.Errorf("invalid card rank: %c", chars[0])
	}
	return
}

func cardRankToString(rank CardRank) string {
	switch rank {
	case TWO:
		return "2"
	case THREE:
		return "3"
	case FOUR:
		return "4"
	case FIVE:
		return "5"
	case SIX:
		return "6"
	case SEVEN:
		return "7"
	case EIGHT:
		return "8"
	case NINE:
		return "9"
	case TEN:
		return "10"
	case JACK:
		return "J"
	case QUEEN:
		return "Q"
	case KING:
		return "K"
	case ACE:
		return "A"
	default:
		panic("invalid CardRank")
	}
}

func getSuit(char rune) (suit Suit, err error) {
	switch char {
	case '♤':
		suit = SPADES
	case '♡':
		suit = HEARTS
	case '♧':
		suit = CLUBS
	case '♢':
		suit = DIAMONDS
	default:
		err = fmt.Errorf("invalid suit: %c", char)
	}
	return
}

func suitToRune(suit Suit) rune {
	switch suit {
	case SPADES:
		return '♤'
	case HEARTS:
		return '♡'
	case CLUBS:
		return '♧'
	case DIAMONDS:
		return '♢'
	default:
		panic("invalid suit")
	}
}

func getCountsByCard(hands []Hand) map[Card]int {
	counts := make(map[Card]int)
	for _, hand := range hands {
		for _, card := range hand.Cards() {
			if _, ok := counts[card]; ok {
				counts[card]++
			} else {
				counts[card] = 1
			}
		}
	}
	return counts
}
