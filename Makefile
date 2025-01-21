.PHONY: build push release tag

build:
	docker build -t ghcr.io/haleyrc/tings .

push:
	docker push ghcr.io/haleyrc/tings:latest

tag:
	# docker tag ghcr.io/haleyrc/tings ghcr.io/haleyrc/tings:$(git rev-parse --abbrev-ref HEAD)-$(git rev-parse --short HEAD)

release: build tag push
