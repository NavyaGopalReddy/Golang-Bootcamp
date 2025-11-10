package main

import (
	"fmt"
	"time"
)

// CommentDetail holds the comment text and when it was created
type CommentDetail struct {
	CreatedTime time.Time
	Comment     string
}

// RatingDetail stores individual user feedback
type RatingDetail struct {
	User    string
	Rating  int
	Comment CommentDetail
}

// Rating represents the overall rating info for a product
type Rating struct {
	RatingID    string
	AllFeedback []RatingDetail
}

// addRating adds a new user feedback entry
func (r *Rating) addRating(user string, rating int, comment string) {
	newFeedback := RatingDetail{
		User:   user,
		Rating: rating,
		Comment: CommentDetail{
			CreatedTime: time.Now(),
			Comment:     comment,
		},
	}
	r.AllFeedback = append(r.AllFeedback, newFeedback)
}

// getAverageRating computes the average rating
func (r *Rating) getAverageRating() float64 {
	if len(r.AllFeedback) == 0 {
		return 0
	}
	total := 0
	for _, f := range r.AllFeedback {
		total += f.Rating
	}
	return float64(total) / float64(len(r.AllFeedback))
}

// printRatingDetails displays full rating info
func (r *Rating) printRatingDetails(productID string) {
	fmt.Printf("Product ID: %s\n", productID)
	fmt.Printf("Rating ID: %s\n", r.RatingID)
	fmt.Printf("Average Rating: %.1f\n", r.getAverageRating())
	fmt.Println("Feedbacks:")
	for _, f := range r.AllFeedback {
		fmt.Printf(" - %s rated %d at %s: \"%s\"\n",
			f.User, f.Rating, f.Comment.CreatedTime.Format("2006-01-02 15:04:05"), f.Comment.Comment)
	}
	fmt.Println("====================================")
}

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
