FROM golang:1.9.2-alpine AS build
RUN apk add --no-cache git && \
    go get github.com/Masterminds/glide
WORKDIR /go/src/github.com/bluehawk27/assignment
COPY glide.yaml glide.lock ./
RUN glide install
COPY . .
RUN go build -o $(pwd)/build/assignmentd && ls -R
FROM alpine:3.6
COPY --from=build /go/src/github.com/bluehawk27/assignment/build/assignmentd /bin/
COPY config.yml /tmp/
RUN apk add --no-cache ca-certificates
ENTRYPOINT ["assignmentd"]
