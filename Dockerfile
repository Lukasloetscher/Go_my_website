# syntax=docker/dockerfile:1

FROM golang:1.23.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod downlaod

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
#note that i will probably also need to copy the content ogf the fiolder webpages, as soon as i load them
#or maybe it is cleaner to instead mounbt them to the container?
#TODO try out and find out:)
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /my_application

CMD ["/my_application"]
