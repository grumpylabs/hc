FROM golang:1.12-alpine
#
RUN apk add --no-cache git
#
# Set the Current Working Directory inside the container
WORKDIR /app/hcf
#
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
#
RUN go mod download
#
COPY . .
#
# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/hcf .
