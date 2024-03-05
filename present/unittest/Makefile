gomodule:=unittest

# test function in a package
# Usage: make testbyfunc name=<func_name> pkg=<pkg_name>.
# Parameters: name: name of function want to test
testfunc:
	go test -v -count=1 -run ${func} ${gomodule}/${pkg}

testpkg:
	go test -v -shuffle=on -count=1 ./${pkg}

testall:
	go test -v ./...

testtag:
	go test --tags=${tag} -v -count=1 ./...

testcoveragepkg:
	go test -cover -covermode=count -coverprofile=cover.txt ${gomodule}/${pkg} -count=1  && \
	go tool cover -html=cover.txt && \
	rm cover.txt

mockgen:
	mockgen -source=${file}.go -destination=${file}_mock.go -package=${pkg}
