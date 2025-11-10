package main

import "fmt"

// Ratings holds product feedback details
type Ratings struct {
	ProductID string
	Rating    int
	Comment   string
	User      string
}

// printRatings displays the details and feedback message
func printRatings(r Ratings) {
	fmt.Println("User:", r.User)
	fmt.Println("Product ID:", r.ProductID)
	fmt.Println("Rating:", r.Rating)
	fmt.Println("Comment:", r.Comment)

	if r.Rating >= 5 {
		fmt.Println("Thank you for your feedback! We hope you're happy with the purchase!")
	} else {
		fmt.Println("Oh no! Sorry for your experience. We'll try to improve!")
	}

	fmt.Println("////////////////////////////////")
}

func main() {
	// Creating two product ratings
	apple := Ratings{
		ProductID: "PD01",
		Rating:    5,
		Comment:   "iphone share your patent rights",
		User:      "User1",
	}

	vanillaIceCream := Ratings{
		ProductID: "PD02",
		Rating:    0,
		Comment:   "Android share your patent rights",
		User:      "User2",
	}

	// Printing both ratings
	printRatings(apple)
	printRatings(vanillaIceCream)
}
