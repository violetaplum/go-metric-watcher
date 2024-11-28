
# 🐱 go-metric-watcher

Real-time System Metrics Monitoring and Alert Platform

## 📊 Key Features

### System Monitoring
- Real-time monitoring of CPU, memory, and disk usage
- Collection of API response times and performance metrics
- Tracking of error rates and system status

### Dashboard Features
- Collection and monitoring of CPU, Memory, Disk and Network metrics
- Metric storage through Prometheus
- Dashboard visualization through Grafana

### Dashboard Panels

1. CPU Monitoring
   - CPU Usage (%)
   - CPU Core Count

2. Memory Monitoring
   - Memory Usage Overview
   - Memory Usage (%)

3. Disk Monitoring
   - Disk Usage (%)
   - Disk I/O

4. Network Monitoring
   - Network Traffic (Bytes)
   - Network Packets

### Alert System
- Threshold-based alert configuration
- Slack/Email notification integration
- Alert history management

## 🛠 Tech Stack
- **Language:** Go
- **Framework:** gRPC
- **Database:** Prometheus
- **Visualization:** Grafana
- **Monitoring:** Custom Metrics Collector
- **Infrastructure:** Docker, Docker Compose

## 📐 Architecture

```mermaid
graph TD
    A[Metrics Collector] --> B[gRPC Server]
    B --> C[Prometheus]
    C --> D[Grafana]
    B --> E[Alert Manager]
    E --> F[Notification Service]
```

## 📁 Project Structure

```
.
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── collector/
│       └── main.go
├── deployments/
│   └── grafana/
│       ├── provisioning/
│       │   └── init-scripts/
│       │       └── 01-init.sql
│       ├── docker-compose.yml
│       └── Dockerfile
├── internal/
│   ├── model/
│   │   └── metric.go
│   ├── repository/
│   │   └── timescaledb.go
│   └── service/
│       ├── collector.go
│       └── metrics.go
├── pkg/
│   ├── grpcutil/
│   │   └── server.go
│   ├── logger/
│   │   └── logger.go
│   ├── monitoring/
│   │   ├── cpu.go
│   │   ├── disk.go
│   │   └── memory.go
│   └── notifier/
│       ├── email.go
│       └── slack.go
└── prometheus/
    └── prometheus.yml
```

## 📌 API Documentation
### Metrics API
- `CollectMetrics`: Collect system metrics
- `StreamMetrics`: Real-time metrics streaming
- `GetMetricHistory`: Query historical metric data

### Alert API
- `ConfigureAlert`: Configure alert rules

## 🔧 Configuration


### Prometheus Configuration
```yaml
# prometheus/prometheus.yml
scrape_configs:
  - job_name: 'go-metric-watcher'
    static_configs:
      - targets: ['metrics-collector:2112']
```

### Alert Configuration
```yaml
alerts:
  cpu_usage:
    threshold: 80
    duration: 5m
    severity: warning
```

## 🚀 Getting Started

### Prerequisites

```bash
- Go 1.23 or higher
- Docker
- Docker Compose
```

### Installation

```bash
# Clone repository
git clone https://github.com/violetaplum/go-metric-watcher

# Install dependencies
go mod download

# Build
go build -o cmd/api/main.go
go build -o cmd/collector/main.go
```

### Running the Application

```bash
# Set up infrastructure
make build

# Direct execution
go build -o api cmd/api/main.go
go build -o collector cmd/collector/main.go

./api
./collector

# Run tests
go test ./...
```

## 🌐 Accessing the Dashboard

### Access Information
- Grafana: http://localhost:3000
- Prometheus: http://localhost:9090

### Login Credentials
- Username: admin
- Password: admin

### Available Services
1. Grafana Dashboard
    - System metrics visualization
    - Custom dashboard templates
    - Read-time monitoring
2. Prometheus
    - Metrics database
    - Query interface
    - Data exploration

### Simple metric api call
```bash
curl http://localhost:8080/api/v1/health
curl http://localhost:8080/api/v1/metrics
```


