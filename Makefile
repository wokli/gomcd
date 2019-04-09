VERSION         :=      $(shell cat ./VERSION)
IMAGE_NAME      :=      gomcd
GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

all: install

install:
	go install -v

test: $(GO_TEST_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		go test -v ; \
		popd > /dev/null ; \
	done;

fmt: $(GO_SRC_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		go fmt ; \
		popd > /dev/null ; \
	done;

image:
	docker build -t $(IMAGE_NAME):$(VERSION) .

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	goreleaser --rm-dist

show:
	@echo "SRC  = $(GO_SRC_DIRS)"
	@echo "TEST = $(GO_TEST_DIRS)"
	@echo "VERSION = $(VERSION)"
	@echo "IMAGE_NAME = $(IMAGE_NAME)"
.PHONY: install test fmt release