FROM golang:1.17-alpine

WORKDIR /go/src/app

#
#RUN ls -la
#RUN pwd
#
RUN go get -d -v
RUN go install -v
#
CMD ["app"]