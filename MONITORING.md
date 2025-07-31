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