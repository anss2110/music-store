// main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Album struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseYear int     `json:"release_year"`
	Price       float64 `json:"price"`
	ArtistID    int     `json:"artist_id"`
}

type Song struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AlbumID  int    `json:"album_id"`
	Duration string `json:"duration"`
	Year     int    `json:"year"`
}

type Artist struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	FormedYear int    `json:"formed_year"`
}

var db *sql.DB

func initDB() (*sql.DB, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Get database connection parameters from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Set default values if environment variables are not set
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbName == "" {
		dbName = "musicstore"
	}

	// Create the connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// Open database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	return db, nil
}

func main() {
	// Initialize database
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Router setup
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/albums", getAllAlbums).Methods("GET")
	router.HandleFunc("/albums/{id}", getAlbumByID).Methods("GET")
	router.HandleFunc("/songs", getAllSongs).Methods("GET")
	router.HandleFunc("/songs/{id}", getSongByID).Methods("GET")
	router.HandleFunc("/artists", getAllArtists).Methods("GET")
	router.HandleFunc("/artists/{id}", getArtistByID).Methods("GET")
	router.HandleFunc("/songs/year/{year}", getSongsByYear).Methods("GET")

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// [Rest of the handler functions remain the same as in the previous version]
func getAllAlbums(w http.ResponseWriter, r *http.Request) {
	var albums []Album
	rows, err := db.Query("SELECT id, title, release_year, price, artist_id FROM albums")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.ReleaseYear, &album.Price, &album.ArtistID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		albums = append(albums, album)
	}

	json.NewEncoder(w).Encode(albums)
}

func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var album Album
	err := db.QueryRow("SELECT id, title, release_year, price, artist_id FROM albums WHERE id = ?", id).
		Scan(&album.ID, &album.Title, &album.ReleaseYear, &album.Price, &album.ArtistID)

	if err == sql.ErrNoRows {
		http.Error(w, "Album not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(album)
}

func getAllSongs(w http.ResponseWriter, r *http.Request) {
	var songs []Song
	rows, err := db.Query("SELECT id, title, album_id, duration, year FROM songs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.ID, &song.Title, &song.AlbumID, &song.Duration, &song.Year); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		songs = append(songs, song)
	}

	json.NewEncoder(w).Encode(songs)
}

func getSongByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var song Song
	err := db.QueryRow("SELECT id, title, album_id, duration, year FROM songs WHERE id = ?", id).
		Scan(&song.ID, &song.Title, &song.AlbumID, &song.Duration, &song.Year)

	if err == sql.ErrNoRows {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(song)
}

func getAllArtists(w http.ResponseWriter, r *http.Request) {
	var artists []Artist
	rows, err := db.Query("SELECT id, name, country, formed_year FROM artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var artist Artist
		if err := rows.Scan(&artist.ID, &artist.Name, &artist.Country, &artist.FormedYear); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		artists = append(artists, artist)
	}

	json.NewEncoder(w).Encode(artists)
}

func getArtistByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var artist Artist
	err := db.QueryRow("SELECT id, name, country, formed_year FROM artists WHERE id = ?", id).
		Scan(&artist.ID, &artist.Name, &artist.Country, &artist.FormedYear)

	if err == sql.ErrNoRows {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(artist)
}

func getSongsByYear(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	year := params["year"]

	var songs []Song
	rows, err := db.Query("SELECT id, title, album_id, duration, year FROM songs WHERE year = ?", year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.ID, &song.Title, &song.AlbumID, &song.Duration, &song.Year); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		songs = append(songs, song)
	}

	json.NewEncoder(w).Encode(songs)
}
