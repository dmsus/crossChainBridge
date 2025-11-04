package main

import (
	"database/sql"
	"fmt"
	"log"

	"crosschain-bridge/tests/performance"

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

	// Запускаем бенчмарк
	results := performance.RunDatabaseBenchmark(db)
	performance.PrintResults(results)

	fmt.Println("\n🎯 Benchmark completed!")
}
