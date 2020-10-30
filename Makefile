all:
	go build -o esewebcache bin/*.go


windows:
	GOOS=windows GOARCH=amd64 \
            go build \
	    -o eseparser.exe ./bin/*.go

generate:
	cd parser/ && binparsegen conversion.spec.yaml > ese_gen.go


test:
	go test ./...
