FROM golang:1.8

# Get and install glide
# RUN curl https://glide.sh/get | sh
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy file to container
RUN mkdir -p /go/src/foreign_currency
ADD . /go/src/foreign_currency
# Go to main folder and install dependencies
RUN cd /go/src/foreign_currency && \ 
  dep ensure
# Go to main folder and build executable file
RUN cd /go/src/foreign_currency && \
  GOPATH=$GOPATH:$(pwd)/vendor:/go/src/  go build foreign_currency

ENTRYPOINT ["/go/src/foreign_currency/foreign_currency"]

# Document that the service listens on port 7001.
EXPOSE 7001