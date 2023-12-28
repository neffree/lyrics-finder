package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ApiResponse struct {
	Result []struct {
		Song       string `json:"song"`
		SongLink   string `json:"song-link"`
		Artist     string `json:"artist"`
		ArtistLink string `json:"artist-link"`
		Album      string `json:"album"`
		AlbumLink  string `json:"album-link"`
	} `json:"result"`
}

func searchLyricsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // This is not recommended for production use
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Pre-flight request
	if r.Method == "OPTIONS" {
		return
	}
	searchTerm := r.URL.Query().Get("term")
	apiUrl := "https://www.stands4.com/services/v2/lyrics.php"
	uid := os.Getenv("UID")
	tokenid := os.Getenv("TOKENID")

	// Make the API request
	resp, err := http.Get(apiUrl + "?uid=" + uid + "&tokenid=" + tokenid + "&term=" + searchTerm + "&format=json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode the API response
	var apiResp ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response back to the frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResp)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/searchLyrics", searchLyricsHandler)
	http.ListenAndServe(":8080", nil)
}
