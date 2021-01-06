package main

import (
	"fmt"
	"time"

	"github.com/lepasq/spotify-media-server/config"
	"github.com/lepasq/spotify-media-server/schedule"
)

func main() {
	client, err := config.Authenticate()
	if err != nil {
		fmt.Println(err)
		return
	}

	var playlist config.Playlists
	if err := playlist.ProcessPlaylists(client); err != nil {
		fmt.Println(err)
		return
	}

	s := schedule.Scheduler{Client: client, Playlist: &playlist}
	if err := s.Watch(time.Hour * 24); err != nil {
		fmt.Println(err)
	}
}
