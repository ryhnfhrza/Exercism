package techpalace

import "strings"


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {

    
    
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {

    stringStar := strings.Repeat("*",numStarsPerLine)

    result := stringStar + "\n" + welcomeMsg + "\n"+stringStar

    return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	replaceStar := strings.ReplaceAll(oldMsg, "*", "")
	trimSpace := strings.TrimSpace(replaceStar)

    return trimSpace
}



