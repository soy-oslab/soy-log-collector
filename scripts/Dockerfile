FROM golang:1.16

COPY collector collector


RUN apt-get -qy \
	-o "Dpkg::Options::=--force-confdef" \
	-o "Dpkg::Options::=--force-confold" upgrade && \
	apt-get -qy autoclean && \
	apt-get -y update && \
	apt-get -y install redis-server && \
	

WORKDIR /go/collector

RUN go build -v ./...

CMD service redis-server restart; go run cmd/server/server.go
