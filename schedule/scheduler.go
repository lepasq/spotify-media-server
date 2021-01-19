package schedule

import (
	"fmt"
	"time"

	"github.com/lepasq/spotify-media-server/config"
	"github.com/lepasq/spotify-media-server/download"

	"github.com/zmb3/spotify"
)

// Scheduler struct for scheduling updates.
type Scheduler struct {
	Client   *spotify.Client
	Playlist *config.Playlists
}

// Watch fetches playlist updates with duration d in between
func Watch(d time.Duration) error {
	for {
		client, err := setupClient()
		if err != nil {
			return err
		}

		playlist, err := setupPlaylists(client)
		if err != nil {
			return err
		}

		s := Scheduler{Client: client, Playlist: playlist}
		if err := s.Fetch(); err != nil {
			return err
		}
		fmt.Println("Done for today!")
		time.Sleep(d)
	}
}

// Fetch starts the download process
func (s *Scheduler) Fetch() error {
	for k, v := range s.Playlist.Playlists {
		if err := download.Download(s.Client, &k, &v); err != nil {
			return err
		}
	}
	return nil
}

func setupClient() (*spotify.Client, error) {
	client, err := config.Authenticate()
	if err != nil {
		fmt.Printf("%v\nDid you add the environment variables $SPOTIFY_ID and $SPOTIFY_SECRET?\nYou can get them from https://developer.spotify.com/dashboard/login.\n", err)
		return nil, err
	}
	return client, nil
}

func setupPlaylists(client *spotify.Client) (*config.Playlists, error) {
	var playlist config.Playlists
	if err := playlist.ProcessPlaylists(client); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &playlist, nil
}
