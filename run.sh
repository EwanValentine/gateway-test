docker run -d -p 8500:8500 consul || true
go run main.go --registry=mdns --server_address=localhost:9090
