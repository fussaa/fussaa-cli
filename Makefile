APP_VERSION := $(shell git tag --sort=-version:refname | head -n 1)

ifeq ($(APP_VERSION),)
APP_VERSION := v0.1.0
endif

bin:
	GOOS=linux GOARCH=amd64 go build \
		-ldflags "-X main.Version=${APP_VERSION}" \
		-tags release \
		-o build/fss_${APP_VERSION}_linux_amd64 .
	upx build/fss_${APP_VERSION}_linux_amd64
