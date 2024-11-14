# --- Build stage ---
# specify the base image
FROM golang:1.23-alpine3.20 AS builder
# declare the current working directory inside the image
WORKDIR /app
# copy all into the image working directory above from the current folder where we call the run `docker build` command (will be the root folder in our case). 
# first dot is the `from` directory (our root directory here) and second dot is the `to` directory (here the working dir inside the image, i.e /app)
COPY . .
# build to a single binary executable file. `o` stands for `output`, `main` here is the name of the output binary file, and finally passing the main entry point of our app (main.go)
RUN go build -o main main.go

# --- Run stage ---
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

# informs docker that the container listens on a specified network port at runtime
# please note that this doesn't publish the port, it only functions as a documentation.
EXPOSE 8080
# define the default command to run when the container starts
CMD [ "/app/main"]