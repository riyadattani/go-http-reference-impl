FROM golang:1.16.4 as deps

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Copy and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY models/ ./models/

FROM deps as build

# Build the application and copy somewhere convienient
RUN go build -o main ./cmd/web/*.go
WORKDIR /dist
RUN cp /build/main .

# create our new image with just the stuff we need
FROM scratch
WORKDIR /root/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /dist/* ./
EXPOSE 8080
CMD ["./main"]