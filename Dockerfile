# syntax=docker/dockerfile:1

FROM golang:1.24.0 AS build-stage

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
#note this works but is not so nice as it also copies test files. 
#todo find better solution
COPY . ./

# Build
#CGO_ENABLED=0 -> runs without external dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o /my_app ./cmd/web

#copied from tutorial, check if this is the right place, and more importantly
# what happens if the test fails.
FROM build-stage AS run-test-stage
RUN go test -v ./...


#copy only the relevant files from the build into the release stage:
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /my_app /my_app

#expose a port:
#sadly even so the documentations says, this is optional, i need to do this here, and can not do this "only" in the docker desktop add on windows.
EXPOSE 8080

#We also need to Copy the tmpl files...
#maybe check if we want those int he dockerfile, or rather in its own filesystem mounted to docker...
ADD webpages/. webpages/.
 
#makes it so the program is not run as root(?)
USER nonroot:nonroot

CMD ["/my_app"]

