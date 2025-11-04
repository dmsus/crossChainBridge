.PHONY: benchmark analyze setup

# Настройка окружения для performance тестов
setup:
	@echo "🔧 Setting up performance testing environment..."
	docker cp ../scripts/performance/db-analysis-fixed.sql test-postgres:/tmp/
	docker cp ../scripts/performance/add-performance-indexes.sql test-postgres:/tmp/
	docker exec test-postgres psql -U bridge_user -d bridge_local -c "CREATE EXTENSION IF NOT EXISTS pg_stat_statements;"

# Запуск анализа БД
analyze:
	@echo "📊 Analyzing database performance..."
	docker exec test-postgres psql -U bridge_user -d bridge_local -f /tmp/db-analysis-fixed.sql

# Добавление performance индексов
indexes:
	@echo "⚡ Adding performance indexes..."
	docker exec test-postgres psql -U bridge_user -d bridge_local -f /tmp/add-performance-indexes.sql

# Запуск бенчмарка
benchmark:
	@echo "🚀 Running performance benchmark..."
	cd .. && go run tests/performance/run_benchmark.go

# Полный performance тест
all: setup indexes analyze benchmark
