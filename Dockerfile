# GOPATH configured at /go.
FROM golang

ARG APP_NAME
ENV APP_NAME=$APP_NAME

COPY ./src/ ./src/

WORKDIR /go/src

RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get github.com/gorilla/handlers

RUN go build -o $APP_NAME

EXPOSE 8080



