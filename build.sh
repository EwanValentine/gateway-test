protoc -I/usr/local/include -I. \
  --go_out=plugins=micro:$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  proto/greeter/greeter.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:$GOPATH/src/github.com/ewanvalentine/gateway-test/api \
  proto/greeter/greeter.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:$GOPATH/src/github.com/ewanvalentine/gateway-test/api \
  proto/greeter/greeter.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:$GOPATH/src/github.com/ewanvalentine/gateway-test/api \
  proto/greeter/greeter.proto
