package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	//panic("Please implement the ParseCard function")
    switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	//panic("Please implement the FirstTurn function")
    // Special case: pair of aces -> split
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	playerSum := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	// Check if player has a blackjack
	if playerSum == 21 {
		if dealerValue == 10 || dealerValue == 11 {
			return "S" // Stand if dealer has 10 or Ace
		}
		return "W" // Win otherwise
	}

	// Apply standard decision rules
	switch {
	case playerSum >= 17:
		return "S" // Stand
	case playerSum <= 11:
		return "H" // Hit
	case playerSum >= 12 && playerSum <= 16:
		if dealerValue >= 7 {
			return "H" // Hit
		}
		return "S" // Stand
	}

	return "" // fallback (shouldn't happen)
}
