  
FROM golang:1.16

COPY collector collector

RUN apt-get -qy \
	-o "Dpkg::Options::=--force-confdef" \
	-o "Dpkg::Options::=--force-confold" upgrade && \
	apt-get -qy autoclean

WORKDIR /go/collector

RUN go build -v ./...

CMD ["bash"]
