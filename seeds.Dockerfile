#FROM ubuntu:latest
#LABEL authors="miro"
#
#ENTRYPOINT ["top", "-b"]

# Initial stage
FROM golang:1.19-alpine as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download


# Build stage
FROM golang:1.19-alpine as builder

COPY --from=modules /go/pkg /go/pkg

RUN apk update && apk add --no-cache ca-certificates tzdata
RUN adduser -Du 10001 botviewer

RUN cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime
RUN echo "Europe/Moscow" > /etc/timezone

RUN mkdir -p /application
ADD . /application
WORKDIR /application
RUN go mod vendor

# change amd64 to necessary architecture
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -mod=vendor -o ./bin/main.out ./cmd/fill_db.go


# Final stage
#FROM scratch // TODO scratch, remove develop.env dependence
FROM golang:1.19-alpine

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo/
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
ENV TZ=Europe/Moscow

COPY --from=builder /application/bin/main.out /application/bin/main.out
USER botviewer
ADD ./develop.env /application

WORKDIR /application
CMD ["./bin/main.out"]
