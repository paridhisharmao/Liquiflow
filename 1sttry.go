package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID          int
	Description string
	Amount      float64
	Type        string // "INCOME" or "EXPENSE"
	Timestamp   time.Time
}

var Ledger []Transaction

func addTransaction(desc string, amount float64, t string) {
	Ledger = append(Ledger, Transaction{
		ID: len(Ledger) + 1, Description: desc, Amount: amount, Type: t, Timestamp: time.Now(),
	})
}

// --- NEW: The AI Advisor Logic ---
func analyzeOpportunity() {
	var totalIncome, totalExpense float64
	for _, txn := range Ledger {
		if txn.Type == "INCOME" {
			totalIncome += txn.Amount
		} else {
			totalExpense += txn.Amount
		}
	}

	savings := totalIncome - totalExpense

	// Let's say a Savings Account gives 3% and a Liquid Fund gives 7%
	// The "Gap" is 4% per year
	yearlyLoss := savings * 0.04
	monthlyLoss := yearlyLoss / 12

	fmt.Println("\n--- 💡 LiquiFlow Intelligence ---")
	fmt.Printf("Idle Savings: ₹%.2f\n", savings)

	if savings > 10000 {
		fmt.Printf("⚠️ Alert: You are losing ₹%.2f every month in potential interest!\n", monthlyLoss)
		fmt.Println("👉 Suggestion: Move surplus to a Liquid Fund immediately.")
	} else {
		fmt.Println("✅ Your balance is within the safety limit. No optimization needed.")
	}
}

func main() {
	addTransaction("Salary", 60000, "INCOME")
	addTransaction("Rent", 15000, "EXPENSE")
	addTransaction("Food", 5000, "EXPENSE")

	analyzeOpportunity()
}
