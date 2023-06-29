.PHONY: build
build:
	docker build  -t test:latest .

.PHONY: start
start: build
	docker run -d -p 9000:8080 test:latest
