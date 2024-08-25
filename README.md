# Go Weather App

This is a Go-based web application that provides weather forecasts based on latitude and longitude. The application is built using the Fiber web framework and fetches weather data from an external API.

## Prerequisites

Before running the application, make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.20 or higher)
- [Docker](https://www.docker.com/get-started) (if you want to run the app in a container)

## Project Structure

go_weather_app/
├── cmd/
│ └── server/
│ └── main.go
├── internal/
│ ├── api/
│ │ ├── handlers.go
│ ├── forecast/
│ │ ├── client.go
│ ├── models/
│ │ └── models.go
│ └── utils/
│ ├── temperature.go
├── Dockerfile
├── Makefile
├── go.mod
└── go.sum


## Installation and Setup

### 1. Clone the Repository

```bash
git clone https://github.com/cheikhsidi/weatherApp
cd go_weather_app
```


## Build and Run Locally
```bash
make build
make run
```


## Run with Docker
```bash
make docker-build
make docker-run
```

## Stop the Docker Container
```bash
make docker-stop
```

# Usage

Accessing the Application

You can access the application through your web browser or any HTTP client (e.g., curl, Postman).
Example Usage

To get the weather forecast for a specific location, provide the latitude and longitude as query parameters:

URL Format:
```bash
http://localhost:3001/forecast?latitude=<latitude>&longitude=<longitude>
```

Example:
```bash
http://localhost:3001/forecast?latitude=<latitude>&longitude=<longitude>
```

## Response

The application will return a string response containing a summary of the weather forecast, including temperature and a short forecast description.
API Endpoints

    GET /: Returns "Hello, World!" to verify the server is running.
    GET /forecast: Accepts latitude and longitude as query parameters and returns the weather forecast.

## Testing

```bash
make test
```

## Cleaning Up
To remove the binary, stop the Docker container, and remove the Docker image:

```bash
make clean
```
