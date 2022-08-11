FROM golang:1.16-alpine

# Set the Current Working Directory inside the container
WORKDIR /auth-server

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .
RUN go mod download

RUN go build -o ./AuthServer

EXPOSE 8001

CMD [ "./AuthServer" ]