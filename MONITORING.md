# Monitoring Stack với Grafana, Loki và Promtail

Project này đã được tích hợp với monitoring stack bao gồm:
- **Grafana**: Platform visualization cho logs và metrics
- **Loki**: Log aggregation system
- **Promtail**: Log collection agent

## Cách sử dụng

### 1. Khởi động monitoring stack

```bash
./bin.sh monitoring start
```

Sau khi khởi động thành công, bạn có thể truy cập:
- **Grafana**: http://localhost:3000 (admin/admin)
- **Loki**: http://localhost:3100

### 2. Các lệnh monitoring khác

```bash
# Dừng monitoring stack
./bin.sh monitoring stop

# Khởi động lại
./bin.sh monitoring restart

# Xem logs của monitoring stack
./bin.sh monitoring logs

# Xem logs của service cụ thể
./bin.sh monitoring logs grafana
./bin.sh monitoring logs loki
./bin.sh monitoring logs promtail
```

### 3. Khởi động toàn bộ hệ thống (infra + monitoring + API)

```bash
# Khởi động infrastructure và monitoring
./bin.sh infra up -d
./bin.sh monitoring start

# Khởi động API services
./bin.sh api start
```

## Cấu hình

### Promtail Configuration
Promtail được cấu hình để collect logs từ:
- `/var/log/apps/auth-service/*.log` - Logs từ auth-service
- `/var/log/apps/rest-service/*.log` - Logs từ rest-service
- `/var/log/syslog` - System logs
- `/var/lib/docker/containers/*/*-json.log` - Docker container logs

### Grafana Dashboard
Dashboard mặc định bao gồm:
- **Log Volume by Service**: Hiển thị số lượng logs theo service
- **Recent Logs**: Logs gần đây từ tất cả services
- **Error Logs**: Chỉ hiển thị error logs

## Tích hợp với Go Services

### Sử dụng Logger trong Go services

```go
import "github.com/ngoctb13/seta-train/shared-modules/utils"

// Tạo logger cho service
logger := utils.NewLogger("auth-service")

// Ghi logs
logger.Info("Service started on port %s", port)
logger.Error("Database connection failed: %v", err)
logger.Debug("Processing request: %s", requestID)
```

### Sử dụng Logging Middleware

```go
import "github.com/ngoctb13/seta-train/shared-modules/middleware"

// Tạo logging middleware
loggingMiddleware := middleware.NewLoggingMiddleware("auth-service")

// Thêm vào Gin router
router.Use(loggingMiddleware.LoggingMiddleware())
router.Use(loggingMiddleware.ErrorLogging())
```

## Troubleshooting

### 1. Promtail không collect được logs
- Kiểm tra file logs có tồn tại không: `ls -la logs/`
- Kiểm tra permissions của thư mục logs
- Xem logs của Promtail: `./bin.sh monitoring logs promtail`

### 2. Grafana không kết nối được Loki
- Kiểm tra Loki có đang chạy không: `docker ps | grep loki`
- Kiểm tra network connectivity giữa Grafana và Loki
- Xem logs của Grafana: `./bin.sh monitoring logs grafana`

### 3. Không thấy logs trong Grafana
- Kiểm tra datasource Loki đã được cấu hình chưa
- Kiểm tra query syntax trong Grafana
- Xem logs của Loki: `./bin.sh monitoring logs loki`

## Cấu trúc thư mục

```
build/
├── docker-compose.monitoring.yaml    # Monitoring stack compose
├── grafana/
│   ├── provisioning/
│   │   ├── datasources/loki.yaml     # Loki datasource config
│   │   └── dashboards/dashboard.yaml # Dashboard provisioning
│   └── dashboards/
│       └── app-logs.json             # Default dashboard
├── loki/
│   └── loki-config.yaml              # Loki configuration
└── promtail/
    └── promtail-config.yaml          # Promtail configuration

logs/                                 # Application logs directory
├── auth-service/
│   ├── info-2024-01-01.log
│   ├── error-2024-01-01.log
│   └── debug-2024-01-01.log
└── rest-service/
    ├── info-2024-01-01.log
    ├── error-2024-01-01.log
    └── debug-2024-01-01.log
```

## Customization

### Thêm service mới vào monitoring
1. Cập nhật `build/promtail/promtail-config.yaml` để thêm job mới
2. Tạo logger cho service mới sử dụng `utils.NewLogger()`
3. Thêm logging middleware vào service

### Tạo dashboard mới
1. Tạo file JSON dashboard trong `build/grafana/dashboards/`
2. Dashboard sẽ tự động được import vào Grafana

### Cấu hình retention
Cập nhật `build/loki/loki-config.yaml` để thay đổi retention policy cho logs. 