help:
	@echo "scooter"
	@echo "----------------------------------------------------------------------------------------------------------"
	@echo "List of available targets:"
	@echo "  api-docs                       - generates Restful API documentation.                                   "
	@echo "  container-up                   - spins up application docker-containers."
	@echo "  container-down                 - kills application docker-containers."
	@echo "  daemon                         - installs a go daemon (go mod OFF), for dev."
	@echo "  deps                           - installs all go dependencies from go mod."
	@echo "  dist                           - builds the binary of this service."
	@echo "  help                           - shows this dialog."
	@echo "  install                        - installs project in GOPATH/bin and cache all non-main packages, for dev."
	@echo "  migrate                        - runs migrations of this service."
	@echo "  simulation                     - simulates the connection and disconnection of a scooter."
	@echo "  test-unit                      - runs unit tests of the application."
	@exit 0

.PHONY: \
	daemon \
	deps \
	dist \
	install \
	test-unit

api-docs:
	@printf "$(OK_COLOR)==> Generating API endpoints documentation$(NO_COLOR)\n"
	${DOCKER_RAML} -o ${WEB_DIR}/api.html
	@printf "$(OK_COLOR)==> Finished!$(NO_COLOR), API docs are available here: ${WEB_DIR}/api.html\n"

container-up:
	@printf "$(OK_COLOR)==> Spinning UP containers$(NO_COLOR)\n"
	@docker-compose up -d
	@make
	@printf "$(OK_COLOR)==> Finished!$(NO_COLOR), containers are up and running."

container-down:
	@printf "$(OK_COLOR)==> Killing containers$(NO_COLOR)\n"
	@docker-compose down
	@printf "$(OK_COLOR)==> Finished!$(NO_COLOR), containers are down."

daemon:
	@printf "$(OK_COLOR)==> Installing CompileDaemon$(NO_COLOR)\n"
	@GO111MODULE=off go get github.com/githubnemo/CompileDaemon

deps:
	@printf "$(OK_COLOR)==> Installing go.mod$(NO_COLOR)\n"
	@go mod vendor -v

dist:
	@printf "$(OK_COLOR)==> Building binary$(NO_COLOR)\n"
	@go build -mod=vendor -o ${DIR_OUT}/${BINARY_NAME} ${GO_LINKER_FLAGS} ${SRC}

install:
	@printf "$(OK_COLOR)==> Installing project$(NO_COLOR)\n"
	@go install -mod vendor -v $(SRC)

migrate:
	@printf "$(OK_COLOR)==> Migrating$(NO_COLOR)\n"
	@docker-compose exec app ./docker/scripts/migrate.sh

simulation:
	@printf "$(OK_COLOR)==> Starting simulation$(NO_COLOR)\n"
	@./docker/scripts/simulation.sh $(SRC)
	@printf "$(OK_COLOR)==> Simulation finished successfully$(NO_COLOR)\n"

test-unit:
	@printf "$(OK_COLOR)==> Unit Testing$(NO_COLOR)\n"
	@go test -v -mod=vendor ./internal/...

include docker/env/.env.makefile
