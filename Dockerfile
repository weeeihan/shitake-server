FROM golang:1.21.4-alpine

WORKDIR /app
## We copy everything in the root directory
## into our /app directory
COPY . . 

RUN go mod download

RUN go build -o /shitake-server
## Our start command which kicks off
EXPOSE 8080

## our newly created binary executable
CMD ["/shitake"]