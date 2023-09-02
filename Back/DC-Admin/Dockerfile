FROM golang:1.18.5-alpine3.16 AS build
WORKDIR /go/src/myapp
COPY . .
RUN go mod download && go mod verify
RUN go build -o /go/bin/myapp

FROM alpine
COPY --from=build /go/bin/myapp /go/bin/myapp
WORKDIR /go/bin/assets
WORKDIR /go/bin
ENTRYPOINT ["/go/bin/myapp"]
EXPOSE 4000