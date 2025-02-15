# syntax=docker/dockerfile:1

FROM golang:1.24.0

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
#note that i will probably also need to copy the content ogf the fiolder webpages, as soon as i load them
#or maybe it is cleaner to instead mounbt them to the container?
#TODO try out and find out:)
COPY *.go ./

#for debugging, one may hope this works...
CMD ["ls"]
# Build
#CGO_ENABLED=0 -> runs without external dependencies
#RUN CGO_ENABLED=0 GOOS=linux go build -o my_application.exe ./cmd/web
#CMD ["/my_application"]
