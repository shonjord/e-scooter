# colors
INFO_COLOR=\\e[34m
NO_COLOR=\\033[0m
OK_COLOR=\\e[32m
ERROR_COLOR=\\e[31m

# variables for Makefile
BINARY_API=scooter
BINARY_NAME?=scooter
DIR_OUT=$(CURDIR)/bin
# https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
GO_LINKER_FLAGS=-ldflags "-s"
REPO=github.com/shonjord/e-scooter
SRC=$(REPO)/cmd/scooter
DOCKER_RAML=@docker run --rm -v `pwd`:/data letsdeal/raml2html\:6.2 -i "docs/api.raml"
WEB_DIR=web
