## How to start

Start infra - local environment

```bash
./bin.sh infra up
```

Start API server (include migration) - local environment

```bash
./bin.sh api start
```

## Monitoring & Logging

Start monitoring stack (Grafana, Loki, Promtail)

```bash
./bin.sh monitoring start
```

Access monitoring:
- **Grafana**: http://localhost:3000 (admin/admin)
- **Loki**: http://localhost:3100

For detailed monitoring documentation, see [MONITORING.md](MONITORING.md)