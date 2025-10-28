# Use Case: Cross-Chain Token Transfer

## Description
A user initiates token transfer from Ethereum Sepolia to Polygon Amoy via bridge smart contract. The relay system detects lock events, processes them, and initiates unlocks in target network.

## Actors
- **Primary Actor**: User
- **Secondary Actor**: System Operator  
- **System**: Relay Service (Go)
- **External Systems**: Smart Contracts, Blockchain Nodes, Database

## Trigger
User calls `lockTokens` function in `BridgeEthereum.sol` contract.

## Preconditions
- User holds sufficient token balance in source network
- User has approved token spending for bridge contract
- Relay service is operational and monitoring
- Blockchain nodes for both networks are accessible
- Database connection is established

## Postconditions
- Tokens are locked in Ethereum Sepolia contract
- Equivalent tokens are unlocked in Polygon Amoy
- Transaction recorded in processing journal with status
- System metrics updated accordingly

## Main Success Scenario

1. User calls `lockTokens(amount, targetAddress)` in BridgeEthereum.sol
2. Smart contract validates user balance and allowances
3. Contract transfers tokens from user to contract address
4. Contract emits `TokensLocked(user, amount, nonce, targetAddress, targetChain)` event
5. Relay service detects event via WebSocket subscription
6. Relay validates event structure and required parameters
7. Relay checks database for duplicate nonce processing
8. Relay generates cryptographic signature for the transfer
9. Relay submits `releaseTokens(amount, signature, sourceTxHash)` to BridgePolygon.sol
10. Polygon contract verifies signature validity
11. Polygon contract mints equivalent tokens to targetAddress
12. Relay updates transaction status to "completed" in database
13. System increments success metrics and updates monitoring

## Alternative Scenarios

### AS1: Duplicate Event Detection
1. Relay detects event with already processed nonce
2. System logs warning with event details
3. Event is skipped without further processing
4. Metrics track duplicate detection rate

### AS2: Signature Verification Failure
1. Polygon contract rejects transaction due to invalid signature
2. Relay retries signature generation (max 3 attempts)
3. If all attempts fail, transaction marked as "signature_failed"
4. Operator alerted for manual investigation

### AS3: Insufficient Gas
1. Transaction fails due to insufficient gas
2. Relay detects failure and recalculates required gas
3. Transaction resubmitted with adjusted gas limit
4. Process repeats until success or maximum retries

## Exception Scenarios

### EX1: Source Network Node Unavailable
1. WebSocket connection to Ethereum node drops
2. Relay switches to backup RPC endpoint
3. System alerts operator of primary node failure
4. Missed events caught via block polling until reconnection

### EX2: Target Network Congestion
1. Polygon network experiences high congestion
2. Relay detects pending transactions exceeding timeout
3. Gas price adjustment algorithm activates
4. Transactions resubmitted with optimized gas pricing

### EX3: Database Connection Loss
1. Database becomes unreachable
2. Relay continues processing but queues persistence
3. In-memory cache maintains recent transaction state
4. Batch write operations resume when connection restored

### EX4: Smart Contract Paused
1. Bridge contract emergency pause activated
2. Transactions fail with "contract_paused" error
3. Relay halts new transaction processing
4. Operator notified for immediate attention

## Security Considerations
- Signature replay protection across chains
- Rate limiting per user address
- Maximum transfer amount limits
- Regular security audits and monitoring