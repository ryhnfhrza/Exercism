package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    cardValue := 0
	switch{
        case card == "ace":
        	cardValue = 11
        case card == "two":
        	cardValue = 2
        case card == "three":
        	cardValue = 3
        case card == "four":
        	cardValue = 4
        case card == "five":
        	cardValue = 5
        case card == "six":
        	cardValue = 6
        case card == "seven":
        	cardValue = 7
        case card == "eight":
        	cardValue = 8
        case card == "nine":
        	cardValue = 9
        case card == "ten":
        	cardValue = 10
        case card == "jack":
        	cardValue = 10
        case card == "queen":
        	cardValue = 10
        case card == "king":
        	cardValue = 10
        default:
        	cardValue = 0
    }
    return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var decision string

	card1Value := ParseCard(card1);
	card2Value := ParseCard(card2);

    dealerCardValue := ParseCard(dealerCard)

    cardValue := card1Value + card2Value
    
	if card1 == "ace" && card2 == "ace"{
        decision = "P"
    }else if cardValue == 21{
        if dealerCardValue >= 10  {
            decision = "S"
        }else{
            decision = "W"
        }
    }else{
        switch{
            case cardValue >= 17 && cardValue <= 20 :
            	decision = "S"
            case cardValue >= 12 && cardValue <= 16 && dealerCardValue < 7 :
            	decision = "S"
            case cardValue >= 12 && cardValue <= 16 && dealerCardValue >= 7 :
            	decision = "H"
            case cardValue <= 11 :
            	decision = "H"
        }
    }
    return decision
}
