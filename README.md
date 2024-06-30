# Notification Service

## Overview

This project implements a notification service in Go, with rate limiting to prevent recipients from receiving too many emails. Notifications include various types such as status updates, daily news, and marketing messages, each with their own rate limit rules. The service ensures that the specified rate limits are respected for each recipient.

## Features

- Rate limiting for different types of notifications (status updates, daily news, marketing messages, etc.).
- Implementation follows SOLID principles.
- Tests are included and can be run in a Docker container.
- Docker and Docker Compose setup for easy deployment and testing.

## Technologies

- Go (Golang)
- Docker
- Docker Compose

# Getting Started

## Building and running the application:

1. Build the Docker image:

```
docker compose build
```

2. Run the application:

```
docker compose run app
```

## Running tests:

1. Build the Docker image:

```
docker compose build
```

2. Run the tests:

```
docker compose run tests
```
