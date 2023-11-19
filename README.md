# Docker Go Lang

## Introduction

This project serves as a learning exercise for understanding the basics of Docker and Go (Golang). It provides a simple demonstration of containerizing a Go application, highlighting key concepts and practices in both technologies.

## Prerequisites

Before running this project, ensure that you have the following installed on your system:

- Docker: [Docker Installation Guide](https://docs.docker.com/get-docker/)
- Go (Golang): [Golang Installation Guide](https://golang.org/doc/install) 1.19 or higher

## Project Structure

- `main.go`: a simple http server written in Go Lang.
- `Dockerfile`: step to build a very basic docker image of our application
- `Dockerfile.multistage`: More optimized multistage Dockerfile.

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/nguyenanhhao221/docker-golang.git
   ```
2. **Move into our project:**

   ```bash
   cd docker-golang
   ```

3. **Build Docker Image**
   - For basic image
       ```bash
       docker build -t docker-golang:latest .
       ```
   - For multistage image
       ```bash
        docker build . \
        -t docker-golang:multistage \
        -f Dockerfile.multistage
       ```

4. **Run Docker Image**
    - For basic image
    ```bash
        docker run \
        -p 8080:8080 \
        --name docker-golang \
        -d \
        docker-golang:latest
    ```

   - For multistage image
    ```bash
        docker run \
        -p 8080:8080 \
        --name docker-golang-container \
        -d \
        docker-golang:multistage
    ```
5. **See how it work**

    Once your docker image is running, there is a http server listening on port 8080, you can make a request to it to see the response.
   - Make request

    ```bash
    curl http://localhost:8080
    ```

