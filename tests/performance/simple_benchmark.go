package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Подключаемся к тестовой БД
	connStr := "postgres://bridge_user:bridge_password@localhost:5432/bridge_local?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀 Starting Performance Benchmark...")
	fmt.Println("=====================================")

	// Тест 1: Вставка транзакций
	start := time.Now()
	_, err = db.Exec(`
		INSERT INTO transactions 
		(event_nonce, user_address, amount, source_chain_id, target_chain_id, target_address, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, 1001, "0x742E6D8eFa6c6F0b6d2a76d78b6f5F2416Df8B5C", "1000000000000000000", 
	   11155111, 80002, "0x89205A3E8A6E8a5e5F6D2a6b7C8D9E0F1A2B3C4D5", "pending")
	
	if err != nil {
		fmt.Printf("❌ Insert Transaction: Error - %v\n", err)
	} else {
		fmt.Printf("✅ Insert Transaction: %v\n", time.Since(start))
	}

	// Тест 2: Поиск по статусу
	start = time.Now()
	rows, err := db.Query(`
		SELECT COUNT(*) FROM transactions WHERE status = $1
	`, "pending")
	if err != nil {
		fmt.Printf("❌ Count by Status: Error - %v\n", err)
	} else {
		var count int64
		rows.Scan(&count)
		rows.Close()
		fmt.Printf("✅ Count by Status: %v (Rows: %d)\n", time.Since(start), count)
	}

	// Тест 3: Поиск по пользователю
	start = time.Now()
	rows, err = db.Query(`
		SELECT * FROM transactions 
		WHERE user_address = $1 
		ORDER BY created_at DESC 
		LIMIT 10
	`, "0x742E6D8eFa6c6F0b6d2a76d78b6f5F2416Df8B5C")
	if err != nil {
		fmt.Printf("❌ Find by User: Error - %v\n", err)
	} else {
		var count int64
		for rows.Next() {
			count++
		}
		rows.Close()
		fmt.Printf("✅ Find by User: %v (Rows: %d)\n", time.Since(start), count)
	}

	fmt.Println("\n🎯 Benchmark completed!")
}
