FROM golang:1.26.2-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/cmd/metricservice/exe /app/cmd/metricservice

FROM alpine:3.23 AS runner
WORKDIR app/
COPY --from=builder /app/cmd/metricservice/exe /app
ENV PROJECT_ROOT=/app
ENV LOGGER_FOLDER=/app

CMD ["/app/exe"]