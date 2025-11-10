package main

func main() {
	// Map of ProductID -> RatingID
	productRatingsMap := map[string]string{
		"PD01": "R001",
		"PD02": "R002",
	}

	// Create two Rating structs
	rating1 := Rating{RatingID: productRatingsMap["PD01"]}
	rating2 := Rating{RatingID: productRatingsMap["PD02"]}

	// Add ratings dynamically
	rating1.addRating("User1", 5, "Excellent product!")
	rating1.addRating("User3", 4, "Good quality overall.")

	rating2.addRating("User2", 2, "Not happy with the performance.")
	rating2.addRating("User4", 3, "Okayish product.")

	// Print all rating details
	rating1.printRatingDetails("PD01")
	rating2.printRatingDetails("PD02")
}
