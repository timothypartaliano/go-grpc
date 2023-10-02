1. generate : protoc --go_out=. --go-grpc_out=. pb/product.proto
2. run server(product service) : go run product_service/main.go
3. run client(api) : go run api/main.go
