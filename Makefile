build: 
	go build -o ccwc.exe cmd/ccwc/main.go

test:
	cd pkg/count
	go test