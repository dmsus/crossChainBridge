# User Stories

## User Perspectives

### US.001 - Locking tokens for cross-chain transfer
**As a user**  
**I want to** lock tokens in Ethereum Sepolia  
**So that I can** initiate a transfer to Polygon Amoy

**Acceptance Criteria:**
- User can call lockTokens function in the smart contract
- Balance is validated before locking tokens
- TokensLocked event is emitted with correct parameters
- User receives transaction confirmation

### US.002 - Receiving tokens in target network
**As a user**  
**I want to** receive equivalent tokens in Polygon Amoy  
**So that I can** use them in the destination network

**Acceptance Criteria:**
- Tokens are automatically minted in target network
- Amount matches locked tokens minus fees
- Transaction completes within 30 minutes maximum
- User can view token balance in target network

### US.003 - Monitoring transfer status
**As a user**  
**I want to** track my cross-chain transfer status  
**So that I can** understand current operation state

**Acceptance Criteria:**
- Real-time status updates (pending, processing, completed, failed)
- Transaction hash and block confirmations visible
- Estimated completion time provided
- Failure reasons clearly explained

### US.004 - Secure transaction processing
**As a system operator**  
**I want to** ensure each transaction processes only once  
**So that I can** maintain bridge reliability

**Acceptance Criteria:**
- Database tracks processed event nonces
- Duplicate events are detected and skipped
- Idempotency mechanisms prevent double spending
- Audit logs record all processing attempts

## Operator Perspectives

### US.005 - System health monitoring
**As a system operator**  
**I want to** monitor bridge health and performance  
**So that I can** ensure system availability

**Acceptance Criteria:**
- Real-time metrics dashboard
- Alerting for critical failures
- Node connectivity status
- Transaction success rate monitoring

### US.006 - Emergency response
**As a system operator**  
**I want to** pause bridge operations in emergencies  
**So that I can** prevent fund loss during critical issues

**Acceptance Criteria:**
- Emergency pause function available
- Graceful handling of in-progress transactions
- Clear communication to users during downtime
- Controlled resumption of operations

### US.007 - Fee management
**As a system operator**  
**I want to** configure and update bridge fees  
**So that I can** maintain sustainable operations

**Acceptance Criteria:**
- Dynamic fee adjustment based on network conditions
- Transparent fee structure visible to users
- Revenue tracking and reporting
- Competitive fee positioning