-include .env
export

export PROJECT_ROOT=$(shell pwd)
export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs

app-run:
	@go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/metricservice/main.go

metric-deploy-start:
	@docker compose up -d --build metric-app

metric-deploy-stop:
	@docker compose down metric-app

metric-deploy-check:
	@docker compose up -d --build metric-app
	@for i in 1 2 3 4 5; do \
		if docker run --network host fullstorydev/grpcurl -plaintext -connect-timeout 2 localhost:50051 list; then \
            exit 0; \
        fi; \
		echo "App is not ready"; \
		sleep 2; \
  	done && \
  	echo "App failed check" && \
  	exit 1

test-metric:
	@for i in 1 2 3 4 5; do \
		if grpcurl -plaintext -connect-timeout 2  "$(IP_SERVER_TEST)" list; then \
			exit 0; \
		fi; \
		echo "App is not ready"; \
		sleep 2; \
  	done && \
  	echo "App failed check" && \
  	exit 1