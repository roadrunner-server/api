#!/usr/bin/make
# Makefile readme (ru): <http://linux.yaroslavl.ru/docs/prog/gnu_make_3-79_russian_manual.html>
# Makefile readme (en): <https://www.gnu.org/software/make/manual/html_node/index.html#SEC_Contents>

SHELL = /bin/sh

test_coverage:
	rm -rf coverage-ci
	mkdir ./coverage-ci
	go test -v -race -cover -tags=debug -coverpkg=./... -failfast -coverprofile=./coverage-ci/pipeline_jobs.out -covermode=atomic ./plugins/jobs/pipeline
	echo 'mode: atomic' > ./coverage-ci/summary.txt
	tail -q -n +2 ./coverage-ci/*.out >> ./coverage-ci/summary.txt

test: ## Run application tests
	go test -v -race -tags=debug ./plugins/jobs/pipeline

generate-proto:
	protoc -I./proto/jobs/v1 --go_out=./proto/jobs/v1 jobs.proto
	protoc -I./proto/kv/v1 --go_out=./proto/kv/v1 kv.proto
	protoc -I./proto/websockets/v1 --go_out=./proto/websockets/v1 websockets.proto
	protoc -I./proto/cache/v1beta --go_out=./proto/cache/v1beta response.proto
	protoc -I./proto/http/v1beta --go_out=./proto/http/v1beta http.proto
	protoc -I./proto/service/v1 --go_out=./proto/service/v1 service.proto
