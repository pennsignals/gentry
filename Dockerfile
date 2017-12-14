FROM golang:alpine
WORKDIR /go/src/gentry
COPY . .
RUN go-wrapper download
RUN go-wrapper install
CMD [ "go-wrapper", "run" ]
