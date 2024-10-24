
<img align="center" src="https://github.com/user-attachments/assets/d5f5fce7-afcb-4136-a7a8-a5e8a7e7d7c4" alt="a_cool_futuristic_hd_mascot_for_the" width="300"/>

[_*Mascot Generated by MetaAI_](https://www.meta.ai/)

# MigrateX
## A Database Migration Tool

A command-line tool for applying SQL migrations to a PostgreSQL database using Go and Cobra.

## Features

- Applies SQL migration files in order
- Tracks applied migrations to ensure each is run only once
- Configurable via command-line flags

## Usage

### Running Locally

Clone the repository, build the project, and run it:

```bash
go build -o migrate
./migrate --dbHost="localhost" --dbPort="5432" --dbUser="postgres" --dbPassword="password" --dbName="mydb" --migrationDir="./migrations" --sslMode="disable"
```

### Command-Line Flags:

| Flag             | Default      | Description                          |
|------------------|--------------|--------------------------------------|
| `--dbHost`       | `localhost`  | PostgreSQL host                      |
| `--dbPort`       | `5432`       | PostgreSQL port                      |
| `--dbUser`       | `postgres`   | PostgreSQL user                      |
| `--dbPassword`   | `password`   | PostgreSQL password                  |
| `--dbName`       | `postgres`   | PostgreSQL database name             |
| `--migrationDir` | `migrations` | Path to the migration files directory |
| `--sslMode`      | `disable`    | SSL Mode                             |

### Running with Docker

Build the Docker image and run the container:

```bash
docker build -t db-migration-tool .
docker run --rm \
  -e DB_HOST="localhost" \
  -e DB_PORT="5432" \
  -e DB_USER="postgres" \
  -e DB_PASSWORD="password" \
  -e DB_NAME="mydb" \
  -v /path/to/migrations:/migrations \
  db-migration-tool
```
