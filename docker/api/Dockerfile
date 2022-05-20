FROM golang:1.18 as builder

WORKDIR /go/src/memeclub-api
COPY . .
RUN rm .env && \
    mv .env.docker .env && \
    go mod download && \
     CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ./memeclub

FROM gcr.io/distroless/static@sha256:d6fa9db9548b5772860fecddb11d84f9ebd7e0321c0cb3c02870402680cc315f
USER nonroot:nonroot
COPY --from=builder --chown=nonroot:nonroot /go/src/memeclub-api /
CMD ["/app"]