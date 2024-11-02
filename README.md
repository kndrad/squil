# squil

A command-line tool for managing dogs in a shelter database (please don't bring dogs to a shelter unless necessary). Built to learn and practice Go tools like sqlc, golang-migrate, docker, compose, and others. It's a simple CRUD app, but I've learned a lot during the process.

## Tools & Technologies Used

- **PostgreSQL** - Database
- **sqlc** - SQL to Go code generator
- **golang-migrate** - Database migrations
- **Docker & Docker Compose** - Containerization

## What does it do?

Squil lets you:
- Add dogs to shelter database
- List all dogs
- Read specific dog details
- Update dog information
- Delete dogs from database
- Check database connection


## Installation and Setup

1. Clone the repository
2. Create a `.env` file with database settings:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db_name
```
3. Run `docker-compose up`
4. Use the CLI:
```bash
squil dogs create --name miki --breed kundelek

squil dogs

squil dogs read --name miki

squil dogs update --id 1 --name miki --breed "belgian shepherd"

squil dogs delete --name miki
```

## License

MIT License
