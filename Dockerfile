FROM golang:1.18-alpine

LABEL maintainer="Rafid Aslam"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /golangSimpleCrudServer

EXPOSE 80

ENV PORT 80
env GIN_MODE release

# The following env vars can be set:
#
# env MONGODB_HOST "mongo"
# env MONGODB_USERNAME "root"
# env MONGODB_PASSWORD "example"
# env JWT_SECRET_KEY "verysecret12345"
# env JWT_ISSUER "golangSimpleCrud"

CMD [ "/golangSimpleCrudServer" ]
