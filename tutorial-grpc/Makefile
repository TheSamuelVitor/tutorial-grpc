.PHONY: proto
proto:
	buf generate --template proto/buf.gen.yaml proto

.PHONY: runapi
runapi:
	cd cmd/api && go run .

.PHONY: rungateway
rungateway: 
	cd cmd/gateway && go run .