
![a_cool_futuristic_mascot_for_the_migratex](https://github.com/user-attachments/assets/8f1b6fb2-fbff-4b56-8cd6-d29d82d93efd)

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
./migrate --dbHost="localhost" --dbPort="5432" --dbUser="postgres" --dbPassword="password" --dbName="mydb" --migrationDir="./migrations"
```

### Command-Line Flags:

| Flag            | Default       | Description                              |
|-----------------|---------------|------------------------------------------|
| `--dbHost`      | `localhost`   | PostgreSQL host                          |
| `--dbPort`      | `5432`        | PostgreSQL port                          |
| `--dbUser`      | `postgres`    | PostgreSQL user                          |
| `--dbPassword`  | `password`    | PostgreSQL password                      |
| `--dbName`      | `postgres`    | PostgreSQL database name                 |
| `--migrationDir`| `migrations`  | Path to the migration files directory    |

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
