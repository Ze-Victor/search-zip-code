# Search Zip Code

This is a project to retrieve address information based on a zip code.

## Prerequisites

- Go 1.21.1 installed
- Project dependencies installed

## Installation

To install the project, follow these steps:

1. Clone this repository and navigate to the main directory:
   ```bash
    $ git clone https://github.com/Ze-Victor/search-zip-code.git
    $ cd search-zip-code
   ```

2. Install project dependencies:

```bash
    $ go mod tidy
```

## Usage

1. Start the server:
```bash
    $ make
```

2. Make requests to the API using an HTTP client or web browser.

```bash
$ curl -X GET http://localhost:8001/api/v1/cep -H "Content-Type: application/json" -d '{"cep": "02010010"}'
```

## Documentation

API Documentation

The API documentation is available at [Swagger UI](http://localhost:8001/api/v1/swagger/index.html) after starting the server.

## Unit Tests

To run the project tests, navigate to the main directory and use the command:

```bash
    $ cd search-zip-code
    $ make test
```
## End-to-end Test

End-to-end testing requires API initiation. To run these tests, start the server, open a terminal, navigate to the _"test"_ directory, and run the following command:

```bash
    $ cd test
    $ go test -v
```