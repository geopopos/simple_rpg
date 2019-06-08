# Build Binary
FROM golang:1.12 as build

ARG VERSION

RUN apt update && apt install -y git tree

RUN mkdir -p "$GOPATH/src/github.com/geopopos/simple_rpg/services/player-service"

WORKDIR $GOPATH/src/github.com/geopopos/simple_rpg/services/player-service

COPY main.go .
COPY Gopkg.* ./
COPY pkg/ ./pkg
COPY vendor/ ./vendor
COPY proto/ ./proto

RUN ls ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.VERSION=$VERSION" -o player-service main.go  && \
     cp player-service /player-service

# Inject Binary into container
FROM scratch
COPY --from=build /player-service /usr/local/bin/player-service
ENTRYPOINT ["/usr/local/bin/player-service"]
