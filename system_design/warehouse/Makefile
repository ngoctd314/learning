gomodule:=warehouse

testfunc:
	go test -v -count=1 -run ${func} ${gomodule}/${pkg}

bruteforce:
	go test -v -count=1 -run Test_data ${gomodule}/${pkg}

bitmap:
	go test -v -count=1 -run Test_bitmapOr ${gomodule}/${pkg}

.PHONY: benchmark
benchmark:
	go test -bench "^${func}" ${gomodule}/${pkg} -benchmem -v

.PHONY: protogen
protogen:
	protoc --proto_path=proto proto/message.proto \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative
