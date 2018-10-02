OUT							:= articleapp-api
VERSION 				:= $(shell git describe --always --long --dirty)
SHORTNAME				:= api
LDFLAGS					:= -ldflags "-X main.BuildNumber=${BUILD_NUMBER} -X main.Version=${VERSION}"
PKG_API					:= articleapp

.PHONY: build
build: clean ## Prepare a build for a linux environment
	go build -v -o ${OUT}-v${VERSION} ${LDFLAGS} ${PKG_API}
	$(MAKE) build-dir
	mv ${OUT}-v${VERSION} build
	cd build && ln -sf ${OUT}-v${VERSION} ${OUT}

.PHONY: clean
clean: ## Remove all the temporary and build files
	rm -rf build
	go clean

.PHONY: build-dir
build-dir:
	mkdir -p build

PHONY: build-linux
build-linux: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${OUT}-linux-v${VERSION} ${LDFLAGS} ${PKG_API}
	$(MAKE) build-dir
	mv ${OUT}-linux-v${VERSION} build
	cd build && ln -sf ${OUT}-linux-v${VERSION} ${OUT}


.PHONY: container
container: clean build-linux ## Build the container
# container: ## Build the container
	docker build -t ${OUT}:${VERSION} -f Dockerfile .
	docker tag ${OUT}:${VERSION} ${OUT}:latest

.PHONY: container-start
container-start: container ## Start container
	docker run -d -p 3000:8090 --name ${SHORTNAME}-linux  ${OUT}:${VERSION}

.PHONY: container-stop
container-stop: ## Stop container
	docker rm -f ${SHORTNAME}-linux

.PHONY: start
start: build ## Start the app
	cd build && ./${OUT} server
