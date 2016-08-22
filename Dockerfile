FROM golang

RUN mkdir -p /home/go/src/github.com/stinkyfingers/netstuff
ADD . /home/go/src/github.com/stinkyfingers/netstuff
RUN GOPATH=/home/go/src
RUN apt-get install libpcap-dev #fuck



RUN go get github.com/google/gopacket
RUN go get github.com/google/gopacket/pcap
RUN go install github.com/stinkyfingers/netstuff
ENTRYPOINT /go/bin/netstuff



