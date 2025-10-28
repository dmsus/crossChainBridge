# System Architecture (C4 Model)

## Level 1: System Context Diagram

![C4 Level 1](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/dmsus/crossChainBridge/main/docs/architecture/diagrams/c4-level1.puml)

**Description:**
- **Users** interact with the bridge via dApp frontend or directly with smart contracts
- **System Operators** monitor bridge health and perform administrative functions
- **Bridge System** coordinates cross-chain transfers between Ethereum Sepolia and Polygon Amoy
- **External Systems** provide blockchain connectivity and node services

## Level 2: Container Diagram

![C4 Level 2](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/dmsus/crossChainBridge/main/docs/architecture/diagrams/c4-level2.puml)

**Description:**
- **Web Frontend**: React application for user interactions
- **Go Relay Service**: Core backend processing events and transactions
- **PostgreSQL**: Persistent storage for idempotency and state management
- **Smart Contracts**: Solidity contracts deployed on both networks
- **External Dependencies**: Blockchain networks and node providers

## Level 3: Component Diagram (Go Relay Service)

![C4 Level 3](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/dmsus/crossChainBridge/main/docs/architecture/diagrams/c4-level3.puml)

**Description:**
- **API Server**: RESTful interface for monitoring and administration
- **Event Listener**: WebSocket subscription to blockchain events
- **Transaction Processor**: Core business logic for cross-chain transfers
- **Signer Service**: Cryptographic operations for transaction security
- **Database Client**: Data persistence and idempotency management
- **Metrics Collector**: Observability and monitoring instrumentation