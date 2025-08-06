# Monitoring Stack with Grafana, Loki and Promtail

This project integrated monitoring stack included:
- **Grafana**: Platform visualization for logs and metrics
- **Loki**: Log aggregation system
- **Promtail**: Log collection agent

## How to start

### 1. Start monitoring stack

```bash
./bin.sh monitoring start
```

After starting success, we could access to:
- **Grafana**: http://localhost:3000 (admin/admin)
- **Loki**: http://localhost:3100

### 2. Other monitor command

```bash
# Stop monitoring stack
./bin.sh monitoring stop

# Restart
./bin.sh monitoring restart

# View log
./bin.sh monitoring logs

# View logs of a specific service
./bin.sh monitoring logs grafana
./bin.sh monitoring logs loki
./bin.sh monitoring logs promtail
```


## Customization

### Add new service to the monitoring
1. Update `build/promtail/promtail-config.yaml` to add new job
2. Add logger for new service `utils.NewLogger()`
3. Add logging middleware to the service

### Add new dashboard
1. Create dashboard JSON file in `build/grafana/dashboards/`
2. Dashboard will automatically import into Grafana

### Config retention
Update `build/loki/loki-config.yaml` to change retention policy for logs. 