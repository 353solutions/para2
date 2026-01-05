tag = cars/httpd

build-docker:
	docker build -t $(tag) .


run-docker:
	docker run --rm -p 8080:8080 $(tag)
