FROM golang:1.19-alpine

WORKDIR /app

# Install dependencies
ADD go.mod .
ADD go.sum .
RUN go mod download

# Build
ADD *.go ./
RUN go build -o /compute

CMD ["/compute"]
