package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var balance float64
var safetyLimit float64

func getFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "liquiflow_save.txt") // New filename to start fresh
}

func saveData() {
	path := getFilePath()
	content := fmt.Sprintf("%f,%f", balance, safetyLimit)
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println("❌ Save Error:", err)
	}
}

func loadData() bool {
	path := getFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		return false // No file found
	}

	parts := strings.Split(strings.TrimSpace(string(data)), ",")
	if len(parts) == 2 {
		b, _ := strconv.ParseFloat(parts[0], 64)
		s, _ := strconv.ParseFloat(parts[1], 64)

		// If the file has 0 inside it, we need to know
		if b == 0 && s == 0 {
			return false
		}

		balance = b
		safetyLimit = s
		return true
	}
	return false
}

func main() {
	fmt.Println("🔍 Checking for saved data...")

	loaded := loadData()

	if loaded {
		fmt.Printf("✅ Success! Found Balance: ₹%.2f and Limit: ₹%.2f\n", balance, safetyLimit)
	} else {
		fmt.Println("📂 No previous data found. Let's set up your account.")
		fmt.Print("Enter starting balance: ₹")
		fmt.Scanln(&balance)
		fmt.Print("Set your Safety Limit: ₹")
		fmt.Scanln(&safetyLimit)
		saveData()
	}

	for {
		fmt.Printf("\n--- 🌊 Wallet: ₹%.2f | Limit: ₹%.2f ---\n", balance, safetyLimit)
		fmt.Println("1. Add Income  2. Add Expense  3. AI Optimizer  4. Exit")
		fmt.Print("Choice: ")

		var choice string // Changed to string to prevent "0" errors on accidental letters
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var inc float64
			fmt.Print("Amount: ")
			fmt.Scanln(&inc)
			balance += inc
			saveData()
		case "2":
			var exp float64
			fmt.Print("Amount: ")
			fmt.Scanln(&exp)
			balance -= exp
			saveData()
		case "3":
			runAI()
		case "4":
			saveData()
			fmt.Println("Final Save Complete. Closing...")
			os.Exit(0)
		default:
			fmt.Println("❌ Invalid Choice, use 1-4")
		}
	}
}

func runAI() {
	if balance > safetyLimit {
		fmt.Printf("\n🤖 AI: Invest ₹%.2f for better returns!\n", balance-safetyLimit)
	} else {
		fmt.Println("\n🤖 AI: Keep building your savings!")
	}
}
