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
	AvgRating   float64
	AllFeedback []RatingDetail
}

func main() {
	// Map of ProductID -> RatingID
	productRatingsMap := map[string]string{
		"PD01": "R001",
		"PD02": "R002",
	}

	// Slice of feedback for Product 1
	product1Feedback := []RatingDetail{
		{
			User:   "User1",
			Rating: 5,
			Comment: CommentDetail{
				CreatedTime: time.Now(),
				Comment:     "Excellent product!",
			},
		},
		{
			User:   "User3",
			Rating: 4,
			Comment: CommentDetail{
				CreatedTime: time.Now(),
				Comment:     "Good quality overall.",
			},
		},
	}

	// Slice of feedback for Product 2
	product2Feedback := []RatingDetail{
		{
			User:   "User2",
			Rating: 2,
			Comment: CommentDetail{
				CreatedTime: time.Now(),
				Comment:     "Not happy with the performance.",
			},
		},
		{
			User:   "User4",
			Rating: 3,
			Comment: CommentDetail{
				CreatedTime: time.Now(),
				Comment:     "Okayish product.",
			},
		},
	}

	// Create Rating structs
	rating1 := Rating{
		RatingID:    productRatingsMap["PD01"],
		AvgRating:   calculateAverage(product1Feedback),
		AllFeedback: product1Feedback,
	}

	rating2 := Rating{
		RatingID:    productRatingsMap["PD02"],
		AvgRating:   calculateAverage(product2Feedback),
		AllFeedback: product2Feedback,
	}

	// Print results
	printRatingDetails("PD01", rating1)
	printRatingDetails("PD02", rating2)
}

// calculateAverage computes average rating from all feedback
func calculateAverage(details []RatingDetail) float64 {
	total := 0
	for _, d := range details {
		total += d.Rating
	}
	if len(details) == 0 {
		return 0
	}
	return float64(total) / float64(len(details))
}

// printRatingDetails displays full rating info
func printRatingDetails(productID string, rating Rating) {
	fmt.Printf("Product ID: %s\n", productID)
	fmt.Printf("Average Rating: %.1f\n", rating.AvgRating)
	fmt.Println("Feedbacks:")
	for _, f := range rating.AllFeedback {
		fmt.Printf(" - %s rated %d at %s: \"%s\"\n",
			f.User, f.Rating, f.Comment.CreatedTime.Format("2006-01-02 15:04:05"), f.Comment.Comment)
	}
	fmt.Println("====================================")
}
