LOCAL_BIN = $(CURDIR)/bin

generate-structs:
	mkdir -p backend/src/pkg/model pinger/src/pkg/model
	protoc --go_out=./backend/src/pkg --go_opt=paths=source_relative model/ping.proto
	protoc --go_out=./backend/src/pkg --go_opt=paths=source_relative model/container.proto
	protoc --go_out=./pinger/src/pkg --go_opt=paths=source_relative model/ping.proto
	protoc --go_out=./pinger/src/pkg --go_opt=paths=source_relative model/container.proto
