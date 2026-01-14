package main

import (
	"fmt"
	"time"
)

// Transaction represents a single money movement
type Transaction struct {
	ID        int
	Description string
	Amount    float64
	Category  string    // e.g., "Food", "Salary", "Rent"
	Type      string    // "INCOME" or "EXPENSE"
	Timestamp time.Time
}

// Ledger holds all transactions
var Ledger []Transaction

func addTransaction(desc string, amount float64, category string, t string) {
	newTxn := Transaction{
		ID:          len(Ledger) + 1,
		Description: desc,
		Amount:      amount,
		Category:    category,
		Type:        t,
		Timestamp:   time.Now(),
	}
	Ledger = append(Ledger, newTxn)
	fmt.Printf("✅ Recorded: %s (₹%.2f)\n", desc, amount)
}

func showSummary() {
	var totalIncome, totalExpense float64
	
	fmt.Println("\n--- Financial Summary ---")
	for _, txn := range Ledger {
		if txn.Type == "INCOME" {
			totalIncome += txn.Amount
		} else {
			totalExpense += txn.Amount
		}
		fmt.Printf("[%s] %-12s: ₹%8.2f (%s)\n", 
			txn.Timestamp.Format("15:04"), txn.Description, txn.Amount, txn.Type)
	}
	
	fmt.Printf("\nTotal Income:   ₹%.2f\n", totalIncome)
	fmt.Printf("Total Expenses: ₹%.2f\n", totalExpense)
	fmt.Printf("Net Savings:    ₹%.2f\n", totalIncome-totalExpense)
}

func main() {
	// Let's simulate some data
	addTransaction("Monthly Salary", 50000, "Work", "INCOME")
	addTransaction("Groceries", 2500, "Food", "EXPENSE")
	addTransaction("Amazon Shop", 1200, "Shopping", "EXPENSE")
	addTransaction("Freelance Gig", 5000, "Work", "INCOME")

	showSummary()
}
