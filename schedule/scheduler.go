package schedule

import (
	"spotify-media-server/config"
	"spotify-media-server/download"
	"time"

	"github.com/zmb3/spotify"
)

// Scheduler struct for scheduling updates.
type Scheduler struct {
	Client   *spotify.Client
	Playlist *config.Playlists
}

// Watch fetches playlist updates with duration d in between
func (s *Scheduler) Watch(d time.Duration) error {
	for {
		err := s.Fetch()
		if err != nil {
			return err
		}
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
