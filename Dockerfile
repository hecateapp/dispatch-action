# Build stage
FROM golang:1.12rc1-alpine3.9 AS build-stage

ADD . /src
RUN go get gopkg.in/go-playground/webhooks.v5/github
RUN cd /src && go build -o dispatch

# Release stage
FROM alpine:3.9
COPY --from=build-stage /src/dispatch /

LABEL version="1.0.0"
LABEL repository="https://github.com/hecateapp/dispatch-action"
LABEL homepage="https://github.com/hecateapp/dispatch-action"
LABEL maintainer="Hecate <hello@hecate.co>"

LABEL "com.github.actions.name"="hecateapp/dispatch-action"
LABEL "com.github.actions.description"="Sends merge emails to whoever needs them"
LABEL "com.github.actions.icon"="at-sign"
LABEL "com.github.actions.color"="purple"

ENTRYPOINT ["/dispatch"]