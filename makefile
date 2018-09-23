
.PHONY: build
build: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 go build -a -installsuffix cgo -o articleService
	./articleService

.PHONY: clean
clean: ## Remove all the temporary and build files
	go clean
