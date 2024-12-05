GOPATH ?= $(HOME)/go
GOBIN   = $(GOPATH)/bin

$(AIR):
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

air: $(AIR)
	$(GOBIN)/air

start: air

build:
	@go build -v -o ./tmp/server ./main.go
