# Simple HTTP Service that forwards incoming requests

This is a service that sends incoming request to another URL specified via
`HTTP_FORWARD_BASEURL` environment variable and returns back the response from
there.

## Building the Docker Image
1. Build the Docker image:
    ```sh
    docker build -t http-forward .
    ```

## Running the Docker Container
1. Run the Docker container:
    ```sh
    docker run --rm -p 8080:8080 -e HTTP_FORWARD_BASEURL="http://example.com" http-forward
    ```
