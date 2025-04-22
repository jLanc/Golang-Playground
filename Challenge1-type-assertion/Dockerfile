FROM golang:tip-bookworm
WORKDIR /app

# Install make
RUN apt-get update && apt-get install make -y

# Copy in the source code
COPY Makefile main.go go.mod ./

CMD ["make"]