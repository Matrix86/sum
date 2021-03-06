SHELL := bash
.PHONY: all clients godep golint gomegacheck deps test codecov html_coverage benchmark
.PHONY: clean reset_env profile run build_docker run_docker pycli phpcli

#
# Config
#
GRPC_PATH=/opt/grpc/bins/opt
GRPC_PHP_PLUGIN=${GRPC_PATH}/grpc_php_plugin
GRPC_PROTOC=${GRPC_PATH}/protobuf/protoc

SUMD_DATAPATH=/tmp/sumd

PACKAGES=storage wrapper service

#
# Main actions
#
all: sumd clients test codecov benchmark docker

server_deps: deps proto/sum.pb.go

clients: pycli phpcli 

godep:
	@go get -u github.com/golang/dep/...

deps: godep golint gomegacheck
	@dep ensure

proto/sum.pb.go:
	@${GRPC_PROTOC} -I. --go_out=plugins=grpc:. proto/sum.proto

sumd: server_deps
	@go build -o sumd .

run: reset_env sumd
	@./sumd -datapath "${SUMD_DATAPATH}"

clean:
	@rm -rf sumd
	@rm -rf *.profile
	@rm -rf *.profile.html
	@rm -rf "${SUMD_DATAPATH}"

reset_env: clean
	@mkdir -p "${SUMD_DATAPATH}"
	@mkdir -p "${SUMD_DATAPATH}/data"
	@mkdir -p "${SUMD_DATAPATH}/oracles"

install:
	@mkdir -p /var/lib/sumd/data
	@mkdir -p /var/lib/sumd/oracles
	@cp sumd /usr/local/bin/
	@cp sumd.service /etc/systemd/system/
	@systemctl daemon-reload

#
# Testing and benchmarking
#
golint:
	@go get github.com/golang/lint/golint

gomegacheck:
	@go get honnef.co/go/tools/cmd/megacheck

# Go 1.9 doesn't support test coverage on multiple packages, while
# Go 1.10 does, let's keep it 1.9 compatible in order not to break
# travis
test: server_deps gomegacheck golint
	@echo "mode: atomic" > coverage.profile
	@for pkg in $(PACKAGES); do \
		go vet ./$$pkg ; \
		golint -set_exit_status ./$$pkg ; \
		megacheck ./$$pkg ; \
		go test -race ./$$pkg -coverprofile=$$pkg.profile -covermode=atomic; \
		tail -n +2 $$pkg.profile >> coverage.profile && rm $$pkg.profile ; \
	done
	
codecov: test
	@bash <(curl -s https://codecov.io/bash)

html_coverage: test
	@go tool cover -html=coverage.profile -o coverage.profile.html

profile: reset_env sumd
	@./sumd -datapath "${SUMD_DATAPATH}" -cpu-profile cpu.profile -mem-profile mem.profile

benchmark: server_deps
	@go test ./... -v -run=doNotRunTests -bench=. -benchmem

follow:
	@pidstat --human -l -u -p `pidof sumd` 1

#
# Docker stuff
#
docker:
	@docker build -t sumd:latest .

run_docker: docker
	@docker run -it -p 50051:50051 sumd:latest

#
# Client code generation related stuff.
#
pycli:
	@python -m grpc_tools.protoc \
		-Iproto \
		--python_out=clients/python/proto \
		--grpc_python_out=clients/python/proto \
		proto/sum.proto

phpcli:
	@${GRPC_PROTOC} -I. --proto_path=proto \
		--php_out=clients/php \
		--grpc_out=clients/php \
		--plugin=protoc-gen-grpc=${GRPC_PHP_PLUGIN} \
		proto/sum.proto
