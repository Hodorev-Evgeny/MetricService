include .env
export

export PROJECT_ROOT=$(shell pwd)
export LOGGER_FOLDER = ${PROJECT_ROOT}/out/logs

app-run:
	@go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/metricservice/main.go
