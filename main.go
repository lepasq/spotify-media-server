package main

import (
	"fmt"
	"spotify-media-server/config"
	// "spotify-media-server/download"
)

func main() {
	client, err := config.Authenticate()
	if err != nil {
		fmt.Println(err)
		return
	}

	var p config.Playlists
	if err := p.ProcessPlaylists(client); err != nil {
		fmt.Println(err)
		return
	}
	// download.Download(client)
}
