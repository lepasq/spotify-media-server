package config

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zmb3/spotify"
)

// Playlists is a struct that contains the playlists to be downloaded
type Playlists struct {
	Playlists map[string]string
}

// ProcessPlaylists parses the config file and creates folders for the playlistss
func (p *Playlists) ProcessPlaylists(spotifyClient *spotify.Client) error {
	if err := p.parsePlaylists(spotifyClient); err != nil {
		return err
	}
	path, present := os.LookupEnv("GOSMS")
	if !present {
		return fmt.Errorf("Please set $GOSMS to the path where you'd like to download all of your playlists")
	}
	if err := p.createFolders(&path); err != nil {
		return err
	}
	return nil
}

func (p *Playlists) parsePlaylists(spotifyClient *spotify.Client) error {
	file, err := os.Open("config.txt")
	if err != nil {
		file.Close()
		return err
	}

	if err := p.createMap(spotifyClient, file); err != nil {

	}
	return nil
}

func (p *Playlists) createMap(spotifyClient *spotify.Client, file *os.File) error {
	scanner := bufio.NewScanner(file)
	p.Playlists = make(map[string]string)
	for scanner.Scan() {
		playlistID := scanner.Text()
		playlistSpotify, err := spotifyClient.GetPlaylist(spotify.ID(playlistID))
		if err != nil {
			return err
		}
		p.Playlists[playlistID] = playlistSpotify.Name

	}
	file.Close()
	return nil
}

func (p *Playlists) createFolders(path *string) error {
	if err := os.Chdir(*path); err != nil {
		return err
	}

	for _, v := range p.Playlists {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			os.Mkdir(v, 0777)
			fmt.Printf("Creating folder: %v\n", v)
		}
	}

	return nil
}
