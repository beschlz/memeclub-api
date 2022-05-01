FROM golang:1.18 as builder

WORKDIR /go/src/memeclub-api
COPY . .
RUN rm .env && mv .env.docker .env
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM gcr.io/distroless/static
COPY --from=builder /go/src/memeclub-api /
CMD ["/app"]