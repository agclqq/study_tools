GO_BIN_DIR=bin

.PHONY: http-local
httpd-local: 
	go build -o $(GO_BIN_DIR)/$@/$@ -v cmd/$@/*.go
	mkdir -p $(GO_BIN_DIR)/$@
	cp .env* $(GO_BIN_DIR)/$@
	

httpd-linux: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(GO_BIN_DIR)/httpd/httpd -v cmd/httpd/*.go
	mkdir -p $(GO_BIN_DIR)/$@
	cp .env* $(GO_BIN_DIR)/$@/
	

clean: 
	rm -rf bin
	

