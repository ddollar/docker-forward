.PHONY: all build

default: build

build:
	rm -rf build
	gox -osarch="linux/amd64" -output="build/{{.Dir}}"
	docker build -t ddollar/docker-forward .

release: build
	docker push ddollar/docker-forward

run: build
	docker run ddollar/docker-forward
