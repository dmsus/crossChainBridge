-- Rollback cross-chain bridge database schema
DROP TABLE IF EXISTS nonce_tracking;
DROP TABLE IF EXISTS processed_events;
DROP TABLE IF EXISTS transactions;
