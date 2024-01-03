# Change these variables as necessary.
MAIN_PACKAGE_PATH := ./cmd/neuro-news
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
	go run ./cmd/neuro-news/main.go

## build: build app to ops/production/ansible
.PHONY: build
build:
	go build -o ./ops/production/ansible/ ./cmd/neuro-news/

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