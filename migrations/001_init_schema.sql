-- Cross-Chain Bridge Database Schema
-- Migration 001: Initial Schema Setup

-- Таблица транзакций
CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    event_nonce BIGINT NOT NULL UNIQUE,
    user_address VARCHAR(42) NOT NULL,
    amount NUMERIC(78, 0) NOT NULL,
    source_chain_id INTEGER NOT NULL,
    target_chain_id INTEGER NOT NULL,
    target_address VARCHAR(42) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    signature VARCHAR(132)
);

-- Таблица для идемпотентности
CREATE TABLE processed_events (
    id BIGSERIAL PRIMARY KEY,
    event_hash VARCHAR(66) UNIQUE NOT NULL,
    transaction_id BIGINT REFERENCES transactions(id),
    block_number BIGINT NOT NULL,
    log_index INTEGER NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    processed_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(block_number, log_index, contract_address)
);

-- Таблица для защиты от replay-атак
CREATE TABLE nonce_tracking (
    id BIGSERIAL PRIMARY KEY,
    user_address VARCHAR(42) NOT NULL,
    chain_id INTEGER NOT NULL,
    last_nonce BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_address, chain_id)
);

-- Базовые индексы
CREATE INDEX idx_transactions_status ON transactions(status);
CREATE INDEX idx_transactions_user ON transactions(user_address);
CREATE INDEX idx_events_hash ON processed_events(event_hash);

COMMENT ON TABLE transactions IS 'Cross-chain transaction tracking';
COMMENT ON TABLE processed_events IS 'Idempotency guarantee for blockchain events';
COMMENT ON TABLE nonce_tracking IS 'Replay attack protection';
