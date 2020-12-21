package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// Authenticate assembles the spotify login data
func Authenticate() {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	fmt.Println(config.ClientID)
	fmt.Println(config.ClientSecret)
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	tracks, err := client.GetPlaylistTracks("6SlW2KgSAWzvOj6LA2tTKG")
	if err != nil {
		fmt.Println(err)
	}

	for _, track := range tracks.Tracks {
		var artist string
		for i, artists := range track.Track.Artists {
			artist += (artists.Name)
			if i < len(track.Track.Artists)-1 {
				artist += " "
			}
		}
		fmt.Printf("%v::%v::%v\n", track.Track.Name, artist, track.AddedAt)
	}
}
