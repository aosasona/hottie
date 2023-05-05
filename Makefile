start-dev:
	go run ./cmd -dir ./example -port 3000

release:
	hit p "Release v$(VERSION)"
	git tag -a v$(VERSION) -m "Release v$(VERSION)"
	git push origin v$(VERSION)
	goreleaser --rm-dist
