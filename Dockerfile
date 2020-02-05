# Start by building the application.
FROM golang:1.13-alpine as build

WORKDIR /go/src/hello-go
ADD . /go/src/hello-go

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/hello-go

# Now copy it into our base image.
FROM scratch
COPY --from=build /go/bin/hello-go /
CMD ["/hello-go"]
