# Music Store REST API

A very simple REST API, for managing offline/online music store database, built with Go and MySQL/MariaDB. This API provides endpoints to access information about albums, songs, and artists.

## REST API Features

- Get all albums and album details
- Get all songs and song details
- Get all artists and artist details
- Filter songs by release year

## Technical Features

- MySQL/MariaDB database integration
- Environment variable configuration
- Proper error handling

## Prerequisites

- Go 1.16 or higher
- MySQL/MariaDB
- Git

## Required Dependencies

```bash
github.com/gorilla/mux
github.com/go-sql-driver/mysql
github.com/joho/godotenv
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd music-store-api
```

2. Install dependencies:
```bash
go mod init musicstore
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
go get github.com/joho/godotenv
```

3. Create a `.env` file in the project root:
```env
DB_USER=your_username
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=musicstore
PORT=8080
```

4. Set up the database:
```bash
# Log into MySQL
mysql -u root -p

# Create the database
CREATE DATABASE musicstore;

# Select the database
USE musicstore;

# Run the schema file
source schema.sql

# Run the seed file
source seed.sql
```

## Database Schema

The database consists of three main tables:

### Artists
- id (PRIMARY KEY)
- name
- country
- formed_year

### Albums
- id (PRIMARY KEY)
- title
- release_year
- price
- artist_id (FOREIGN KEY)

### Songs
- id (PRIMARY KEY)
- title
- album_id (FOREIGN KEY)
- duration
- year

## API Endpoints

### Albums
- GET `/albums` - Get all albums
- GET `/albums/{id}` - Get album by ID

### Songs
- GET `/songs` - Get all songs
- GET `/songs/{id}` - Get song by ID
- GET `/songs/year/{year}` - Get songs by year

### Artists
- GET `/artists` - Get all artists
- GET `/artists/{id}` - Get artist by ID

## Running the Application

1. Start the server:
```bash
go run main.go
```

2. The API will be available at `http://localhost:8080`

## Example API Requests

### Get all albums
```bash
curl http://localhost:8080/albums
```

### Get album by ID
```bash
curl http://localhost:8080/albums/1
```

### Get songs from a specific year
```bash
curl http://localhost:8080/songs/year/1987
```

## Response Formats

### Album Response
```json
{
    "id": 1,
    "title": "The Joshua Tree",
    "release_year": 1987,
    "price": 29.99,
    "artist_id": 1
}
```

### Song Response
```json
{
    "id": 1,
    "title": "With or Without You",
    "album_id": 1,
    "duration": "4:56",
    "year": 1987
}
```

### Artist Response
```json
{
    "id": 1,
    "name": "U2",
    "country": "Ireland",
    "formed_year": 1976
}
```

## Error Handling

The API returns appropriate HTTP status codes:

- 200: Success
- 404: Resource not found
- 500: Internal server error

## How to Contribute

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This simple project is licensed under the MIT License, see the LICENSE file for details.