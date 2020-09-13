FROM golang:1.14-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/linkpower/main.go ./
RUN go mod download
COPY . .

# Build binary
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /app/main .
COPY resources /app/resources

ENTRYPOINT ["/main"]