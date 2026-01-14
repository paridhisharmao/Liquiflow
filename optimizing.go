package main

import "fmt"

func main() {
	var safetyLimit float64
	var currentBalance float64

	fmt.Println("🏦 --- LiquiFlow Manual Configuration ---")

	// 1. Get the Safety Limit from the user
	fmt.Print("Enter your preferred Safety Limit (e.g., 5000): ₹")
	fmt.Scanln(&safetyLimit)

	// 2. Get the current Bank Balance
	fmt.Print("Enter your current HDFC Balance: ₹")
	fmt.Scanln(&currentBalance)

	// 3. The Logic runs based on YOUR input
	fmt.Println("\n--- Running Optimization ---")
	if currentBalance > safetyLimit {
		surplus := currentBalance - safetyLimit
		fmt.Printf("✅ Success! You have a surplus of ₹%.2f\n", surplus)
		fmt.Printf("💡 Recommendation: Invest this ₹%.2f in a Liquid Fund.\n", surplus)
	} else {
		gap := safetyLimit - currentBalance
		fmt.Printf("⚠️ Warning: You are below your limit by ₹%.2f. Save more!\n", gap)
	}
}
