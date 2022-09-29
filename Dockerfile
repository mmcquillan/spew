FROM golang:1.19.1-alpine3.15
WORKDIR /go/src/spew
COPY spew.go .
ENV GOBIN=/go/bin
RUN go install spew.go
CMD ["spew"]

