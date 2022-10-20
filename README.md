# go-petstore
First petstore with Golang with focus on one-button deployment and industrial-level package - logging/tracing/metrics


# Testing

0. `cd dev_infra && docker compose up -d` to init dev infrastructure - grafana, prometheus and jaeger
1. `docker-compose up` to init DB and engage the migrations
2. `make test`
