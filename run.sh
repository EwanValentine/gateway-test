docker run -d -p 8500:8500 consul || true
go run main.go --server_address=0.0.0.0:50051
