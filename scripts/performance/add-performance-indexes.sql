-- Performance Indexes для кросс-чейн моста

-- 1. Для поиска транзакций по статусу и времени (частая операция в API)
CREATE INDEX IF NOT EXISTS idx_transactions_status_created 
ON transactions(status, created_at DESC);

-- 2. Для поиска событий по блоку и контракту (оптимизация event listener)
CREATE INDEX IF NOT EXISTS idx_events_block_contract 
ON processed_events(block_number DESC, contract_address);

-- 3. Для проверки nonce по пользователю и сети (replay protection)
CREATE INDEX IF NOT EXISTS idx_nonce_user_chain 
ON nonce_tracking(user_address, chain_id, last_nonce DESC);

-- 4. Для поиска транзакций по пользователю и времени (история операций)
CREATE INDEX IF NOT EXISTS idx_transactions_user_created 
ON transactions(user_address, created_at DESC);

-- 5. Для поиска по source_chain_id и target_chain_id (аналитика)
CREATE INDEX IF NOT EXISTS idx_transactions_chains 
ON transactions(source_chain_id, target_chain_id);

-- Проверяем созданные индексы
SELECT 
    tablename,
    indexname,
    indexdef
FROM pg_indexes 
WHERE schemaname = 'public' 
AND indexname LIKE 'idx_%'
ORDER BY tablename, indexname;
