
# Deployment Architecture

## Development Environment

![Development Deployment](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/dmsus/crossChainBridge/main/docs/architecture/diagrams/deployment.puml)

**Components:**
- **Developer Machine**: Local development environment
- **Go Relay Service**: Docker container for bridge processing
- **PostgreSQL**: Local database instance
- **Hardhat Node**: Local blockchain simulation
- **Infura**: Development node access for testnets

**Connectivity:**
- Developer interacts with relay service on localhost:8080
- Relay service connects to local PostgreSQL on port 5432
- Local testing via Hardhat node on localhost:8545
- External testnet access via Infura WebSocket/HTTPS

## Production Deployment

**Architecture:**
- **Load Balancer**: Distributes traffic across multiple relay instances
- **Relay Services**: Horizontally scaled instances in Docker/K8s
- **PostgreSQL Cluster**: Primary for writes, replicas for reads
- **Monitoring Stack**: Prometheus for metrics, Grafana for visualization

**Network Ports:**
- Relay Service API: 8080 (HTTP/REST)
- Metrics: 9090 (Prometheus)
- Health Checks: 8081
- PostgreSQL: 5432
- Grafana: 3000

## Security Considerations

### Network Segmentation
- **Public Facing**: Load balancer and API endpoints
- **Internal Network**: Database and service-to-service communication
- **External Services**: Blockchain node providers

### Secret Management
- **Private Keys**: Encrypted storage with key rotation
- **Database Credentials**: Environment variables or secret manager
- **API Keys**: Secure credential storage

## Scaling Strategy

### Horizontal Scaling
- **Stateless Relay Services**: Multiple instances behind load balancer
- **Database Read Replicas**: Offload read operations
- **Connection Pooling**: Efficient database resource utilization

### Monitoring and Alerting
- **Resource Metrics**: CPU, memory, network utilization
- **Business Metrics**: Transaction volume, success rates, latency
- **Health Checks**: Automated service health validation