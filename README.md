# Gator

A simple RSS feed aggregator built with Go and PostgreSQL.

## Prerequisites

Before getting started, make sure you have the following installed:

- Go
- PostgreSQL

## Installation

Clone the repository and install the application:

```bash
git clone <repository-url>
cd <repository-name>
go install
```

This will compile the program and install the executable to your Go binary directory.

## Configuration

Create a configuration file named `.gatorconfig.json` in your home directory.

Example:

```json
{
  "db_url": "postgres://username:password@localhost:5432/database_name?sslmode=disable",
  "current_user_name": ""
}
```

Replace the values with your PostgreSQL connection information.

## Usage

### Register a User

Create a new user:

```bash
gator register <name>
```

Example:

```bash
gator register alice
```

### Aggregate Posts

Fetch and aggregate posts from all followed feeds:

```bash
gator agg
```

### Other Commands

Gator includes additional commands for managing users, feeds, and posts. Run the application with the appropriate command to explore its functionality.

## Features

- User registration and management
- RSS feed subscription
- Feed aggregation
- PostgreSQL-backed storage
- Simple command-line interface

## Tech Stack

- Go
- PostgreSQL
- RSS Feeds

## License

This project is open source and available under the MIT License.
