package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type LoadTestResult struct {
	TotalTransactions int
	SuccessCount      int
	ErrorCount        int
	TotalDuration     time.Duration
	TPS               float64
}

func runLoadTest(db *sql.DB, concurrentUsers int, transactionsPerUser int) LoadTestResult {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	
	startTime := time.Now()
	successCount := 0
	errorCount := 0

	for i := 0; i < concurrentUsers; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			
			for j := 0; j < transactionsPerUser; j++ {
				nonce := int64(userID*1000 + j)
				_, err := db.Exec(`
					INSERT INTO transactions 
					(event_nonce, user_address, amount, source_chain_id, target_chain_id, target_address, status)
					VALUES ($1, $2, $3, $4, $5, $6, $7)
				`, nonce, fmt.Sprintf("0xUser%08d", userID), "1000000000000000000", 
				   11155111, 80002, fmt.Sprintf("0xTarget%08d", j), "pending")
				
				mutex.Lock()
				if err != nil {
					errorCount++
				} else {
					successCount++
				}
				mutex.Unlock()
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(startTime)

	return LoadTestResult{
		TotalTransactions: concurrentUsers * transactionsPerUser,
		SuccessCount:      successCount,
		ErrorCount:        errorCount,
		TotalDuration:     duration,
		TPS:               float64(successCount) / duration.Seconds(),
	}
}

func main() {
	// Подключаемся к БД
	connStr := "postgres://bridge_user:bridge_password@localhost:5432/bridge_local?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀 Starting Load Testing...")
	fmt.Println("============================")

	// Тест 1: 10 пользователей, по 10 транзакций каждый
	fmt.Println("\n📊 Test 1: 10 concurrent users, 10 transactions each")
	result1 := runLoadTest(db, 10, 10)
	printLoadTestResult(result1)

	// Тест 2: 50 пользователей, по 20 транзакций каждый  
	fmt.Println("\n📊 Test 2: 50 concurrent users, 20 transactions each")
	result2 := runLoadTest(db, 50, 20)
	printLoadTestResult(result2)

	// Тест 3: 100 пользователей, по 10 транзакций каждый
	fmt.Println("\n📊 Test 3: 100 concurrent users, 10 transactions each")
	result3 := runLoadTest(db, 100, 10)
	printLoadTestResult(result3)
}

func printLoadTestResult(result LoadTestResult) {
	fmt.Printf("Total Transactions: %d\n", result.TotalTransactions)
	fmt.Printf("Success: %d\n", result.SuccessCount)
	fmt.Printf("Errors: %d\n", result.ErrorCount)
	fmt.Printf("Duration: %v\n", result.TotalDuration)
	fmt.Printf("TPS: %.2f transactions/second\n", result.TPS)
	fmt.Printf("Success Rate: %.2f%%\n", float64(result.SuccessCount)/float64(result.TotalTransactions)*100)
}
