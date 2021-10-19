ARG BUILDER_IMAGE=golang:alpine

FROM ${BUILDER_IMAGE} AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /usr/src/postservice

RUN apk update && \
    apk add --no-cache git ca-certificates tzdata curl && \
    update-ca-certificates

# Install air for hot reloading
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY go.mod .

RUN go mod download
RUN go mod verify

COPY . .

EXPOSE 8080

CMD [ "air", "-c", "internal/apps/postservice/.air.toml" ]