package poker

import (
  "testing"
  "strings"
)

func TestValidCases(t *testing.T) {
	for _, tc := range validCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := BestHand(tc.input)
			if err != nil {
				t.Errorf("\nunexpected error: %s", err.Error())
			} else {
        if len(result) != len(tc.expected) {
          t.Errorf("expected %d hands, got %d", len(tc.expected), len(result))
        }
				for i, got := range result {
					expected := tc.expected[i]
					if expected != got {
						t.Errorf("\nexpected: %s\ngot     : %s", expected, got)
					}
				}
			}
		})
	}
}

func TestInvalidCases(t *testing.T) {
	for _, tc := range invalidCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := BestHand(tc.input)
			if err == nil {
				t.Errorf("\nexpected error not returned, got result: %v", result)
			} else {
				e := err.Error()
				for _, expected := range tc.errContains {
					if !strings.Contains(e, expected) {
						t.Errorf("\nexpected error to reference: '%s'\ngot: '%s'", expected, e)
					}
				}
			}
		})
	}
}
