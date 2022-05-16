# quotes
A tiny dockerized application to serve famous quotes over HTTP GET.

## Usage
```bash
# Build container
docker build . -t quotes
# Run container
docker run -p 8080:8080 quotes
# Get a quote
curl localhost:8080/quote
```