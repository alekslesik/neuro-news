# Change these variables as necessary.
MAIN_PATH := ./cmd/neuro-news
BINARY_NAME := neuro-news

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run: run app
.PHONY: run
run:
	go run $(MAIN_PATH)/main.go

## run-build: run and build app local
.PHONY: run-build
run-build: build
	./$(BINARY_NAME)

## build: build app local
.PHONY: build
build: tidy
	go build -o ./ $(MAIN_PATH)/

## build-ansible: build app to ops/production/ansible
.PHONY: build-ansible
build-ansible:
	go build -o ./ops/production/ansible/ $(MAIN_PATH)/

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

# ==================================================================================== #
# TESTING
# ==================================================================================== #
## test-cover: test cover of whole project
.PHONY: test-cover
test-cover:
	go test -v -race -coverprofile=./tests/coverage.out ./...
	go tool cover -html=./tests/coverage.out -o ./tests/coverage.html

# ==================================================================================== #
# MYSQL
# ==================================================================================== #

## mysql-root: connect to mySQL by root user
.PHONY: mysql-root
mysql-root:
	mysql -u root -p

# ==================================================================================== #
# DEPLOY
# ==================================================================================== #

## deploy: deploy to host by ansible
.PHONY: deploy
deploy: build
	ansible-playbook -i ops/production/ansible/hosts.ini ops/production/ansible/dpl.yml -vv

# ==================================================================================== #
# OTHER
# ==================================================================================== #

## go-update: update Golang, use "v" for version (go-update v=1.21.5)
.PHONY: go-update
go-update:
	sudo rm -rf go$(v).linux-amd64.tar.gz
	wget https://go.dev/dl/go$(v).linux-amd64.tar.gz
	sudo rm -rf /usr/local/go
	sudo tar -C /usr/local -xzf go$(v).linux-amd64.tar.gz
	sudo rm -rf go$(v).linux-amd64.tar.gz
	go version