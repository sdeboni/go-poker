package poker

type validCase struct {
	description string
	input       []string
	expected    []string
}

type invalidCase struct {
	description string
	input       []string
	errContains []string
}

var validCases = []validCase{
	{
		description: "Returns best hand with one hand",
		input:       []string{"2♢ 2♡ 3♡ 4♡ 5♡"},
		expected:    []string{"2♢ 2♡ 3♡ 4♡ 5♡"},
	},
	{
		description: "Ignores extra spaces in input",
		input:       []string{"  2♢ 2♡ 3♡   4♡  5♡  "},
		expected:    []string{"2♢ 2♡ 3♡ 4♡ 5♡"},
	},
	{
		description: "Does not reorder cards",
		input:       []string{"5♢ 2♡ 8♡ A♡ J♡"},
		expected:    []string{"5♢ 2♡ 8♡ A♡ J♡"},
	},
	{
		description: "Returns hand with highest card of two high card hands",
		input:       []string{"5♢ 2♡ 8♡ 7♡ J♡", "5♡ 2♢ 8♢ A♢ J♢"},
		expected:    []string{"5♡ 2♢ 8♢ A♢ J♢"},
	},
	{
		description: "Same scoring HighCards are returned",
		input:       []string{"5♢ 2♡ 8♡ 7♡ J♡", "2♢ J♤ 7♤ 5♡ 8♧", "2♤ 7♧ 3♤ 4♡ 5♧"},
		expected:    []string{"5♢ 2♡ 8♡ 7♡ J♡", "2♢ J♤ 7♤ 5♡ 8♧"},
	},
	{
		description: "Pair beats HighCard",
		input:       []string{"5♢ 2♡ 8♡ 7♡ J♡", "2♢ 2♤ 3♡ 4♡ 5♧"},
		expected:    []string{"2♢ 2♤ 3♡ 4♡ 5♧"},
	},
	{
		description: "Same ranked pairs are retured",
		input:       []string{"4♡ 3♤ 3♡ 2♡ 5♡", "8♢ 6♧ 2♧ 4♧ 5♧", "3♢ 3♧ 2♤ 4♤ 5♤"},
		expected:    []string{"4♡ 3♤ 3♡ 2♡ 5♡", "3♢ 3♧ 2♤ 4♤ 5♤"},
	},
	{
		description: "Higher pair of two pairs is retured",
		input:       []string{"4♤ 3♤ 3♧ 2♡ 7♡", "A♢ 3♡ 2♧ 4♢ 5♧", "6♢ 4♡ 2♤ 4♧ 5♤"},
		expected:    []string{"6♢ 4♡ 2♤ 4♧ 5♤"},
	},
	{
		description: "Same pairs ranked by next highest card",
		input:       []string{"4♡ 3♤ 3♧ 2♡ 5♡", "3♢ 3♡ 2♤ 4♧ 6♧"},
		expected:    []string{"3♢ 3♡ 2♤ 4♧ 6♧"},
	},
	{
		description: "TwoPair beats one pair",
		input:       []string{"4♡ A♤ A♡ 2♡ 5♡", "3♢ 2♧ 2♤ 3♡ 6♧"},
		expected:    []string{"3♢ 2♧ 2♤ 3♡ 6♧"},
	},
	{
		description: "Same ranked TwoPair hands are returned",
		input:       []string{"3♡ 3♤ 2♢ 2♧ 5♡", "5♢ 2♡ 2♤ 3♢ 3♧"},
		expected:    []string{"3♡ 3♤ 2♢ 2♧ 5♡", "5♢ 2♡ 2♤ 3♢ 3♧"},
	},
	{
		description: "Three of a kind beats TwoPair",
		input:       []string{"3♡ 3♤ 3♢ 2♧ 4♡", "Q♢ A♡ A♤ K♢ K♧"},
		expected:    []string{"3♡ 3♤ 3♢ 2♧ 4♡"},
	},
	{
		description: "Higher three of a kind beats lower three of a kind",
		input:       []string{"3♡ 3♤ 3♢ 2♢ 4♧", "4♡ 4♤ 4♢ 2♧ 5♡"},
		expected:    []string{"4♡ 4♤ 4♢ 2♧ 5♡"},
	},
	{
		description: "Stright beats three of a kind",
		input:       []string{"2♡ 3♤ 4♢ 5♢ 6♧", "A♡ A♤ A♢ K♧ Q♡"},
		expected:    []string{"2♡ 3♤ 4♢ 5♢ 6♧"},
	},
	{
		description: "Higher straight beats lower straight",
		input:       []string{"2♡ 3♤ 4♢ 5♢ 6♧", "3♡ 4♤ 5♡ 6♢ 7♡"},
		expected:    []string{"3♡ 4♤ 5♡ 6♢ 7♡"},
	},
	{
		description: "Two equal straights are both returned",
		input:       []string{"A♡ K♤ Q♢ J♢ 10♧", "10♡ J♤ Q♧ K♧ A♢"},
		expected:    []string{"A♡ K♤ Q♢ J♢ 10♧", "10♡ J♤ Q♧ K♧ A♢"},
	},
	{
		description: "flush beats straight",
		input:       []string{"A♤ K♤ Q♢ J♢ 10♧", "2♡ 3♡ Q♡ K♡ A♡"},
		expected:    []string{"2♡ 3♡ Q♡ K♡ A♡"},
	},
	{
		description: "higher flush beats lower flush",
		input:       []string{"2♤ 3♤ Q♤ K♤ A♤", "2♡ 4♡ Q♡ K♡ A♡"},
		expected:    []string{"2♡ 4♡ Q♡ K♡ A♡"},
	},
	{
		description: "full house beats flush",
		input:       []string{"2♤ 2♡ 3♤ 3♡ 3♢", "6♡ 7♡ Q♡ K♡ A♡"},
		expected:    []string{"2♤ 2♡ 3♤ 3♡ 3♢"},
	},
	{
		description: "higher full house beats lower full house",
		input:       []string{"A♤ A♡ 2♤ 2♡ 2♢", "3♢ 3♧ 4♤ 4♡ 4♢"},
		expected:    []string{"3♢ 3♧ 4♤ 4♡ 4♢"},
	},
	{
		description: "four of a kind beats full house",
		input:       []string{"2♤ 3♧ 3♤ 3♡ 3♢", "K♤ K♡ A♤ A♡ A♢"},
		expected:    []string{"2♤ 3♧ 3♤ 3♡ 3♢"},
	},
	{
		description: "higher four of a kind beats lower four of a kind",
		input:       []string{"A♤ 3♧ 3♤ 3♡ 3♢", "2♤ 4♧ 4♤ 4♡ 4♢"},
		expected:    []string{"2♤ 4♧ 4♤ 4♡ 4♢"},
	},
	{
		description: "straight flush beatsf four of a kind",
		input:       []string{"K♤ A♧ A♤ A♡ A♢", "2♤ 3♤ 4♤ 5♤ 6♤"},
		expected:    []string{"2♤ 3♤ 4♤ 5♤ 6♤"},
	},
	{
		description: "higher straight flush beats lower straight flush",
		input:       []string{"2♤ 3♤ 4♤ 5♤ 6♤", "3♡ 4♡ 5♡ 6♡ 7♡"},
		expected:    []string{"3♡ 4♡ 5♡ 6♡ 7♡"},
	},
	{
		description: "equal straight flush hands are both returned",
		input:       []string{"2♤ 3♤ 4♤ 5♤ 6♤", "K♤ A♧ A♤ A♡ A♢", "2♡ 3♡ 4♡ 5♡ 6♡"},
		expected:    []string{"2♤ 3♤ 4♤ 5♤ 6♤", "2♡ 3♡ 4♡ 5♡ 6♡"},
	},
}

var invalidCases = []invalidCase{
	{
		description: "Recognizes invalid card rank",
		input:       []string{"2♢ 2♡ 3♡ 4♡ 11♡"},
		errContains: []string{
			"card",
			"rank",
		},
	},
	{
		description: "Recognizes invalid suit",
		input:       []string{"2♢ 2♡ 3♡ 4x 11♡"},
		errContains: []string{
			"suit",
			"x",
		},
	},
	{
		description: "Invalid hand: duplicate card",
		input:       []string{"5♢ 2♡ 2♡ 4♡ 10♡"},
		errContains: []string{
			"duplicate",
			"2♡",
		},
	},
	{
		description: "Invalid hand: too many cards",
		input:       []string{"2♢ 3♢ 4♡ 5♡ 6♡ 7♡"},
		errContains: []string{
			"hand",
			"6",
			"5",
		},
	},
	{
		description: "Invalid hand: too few cards",
		input:       []string{"2♢ 3♢ 4♡"},
		errContains: []string{
			"hand",
			"5",
			"3",
		},
	},
	{
		description: "Hands must have unique cards",
		input:       []string{"3♡ 3♤ 2♢ 2♧ 5♡", "5♡ 2♡ 2♤ 3♢ 3♧"},
		errContains: []string{
			"card",
			"5♡",
		},
	},
}
