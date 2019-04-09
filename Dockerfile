FROM golang:alpine AS builder

# Add all the source code (except what's ignored
# under `.dockerignore`) to the build context.
ADD ./ /go/src/github.com/wokli/gomcd
RUN set -ex && \
  cd /go/src/github.com/wokli/gomcd/ && \       
  CGO_ENABLED=0 go build -tags netgo -v -a -ldflags '-extldflags "-static"' cmd/main.go && \
  mv ./main /usr/bin/main

FROM alpine

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/main /usr/local/bin/main

EXPOSE 8080

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "main" ]