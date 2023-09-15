all:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/handle_failover functions/handle_failover.go
