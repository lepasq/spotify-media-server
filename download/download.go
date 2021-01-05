package download

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/zmb3/spotify"
)

// Download downloads the new songs from playlists
func pageDownload(c *spotify.Client, playlist *string, startIndex *int) error {
	options := spotify.Options{Offset: startIndex}
	tracks, err := c.GetPlaylistTracksOpt(spotify.ID(*playlist), &options, "")
	if err != nil {
		return err
	}

	for page := 0; ; page++ {
		for i, track := range tracks.Tracks {
			fmt.Printf("%v: %v::%v::%v\n", *startIndex+i+1+page*100, track.Track.Name, track.AddedAt, track.Track.ID)
			if err := downloadSong(track.Track.ID.String()); err != nil {
				return err
			}
		}

		err := c.NextPage(tracks)
		if err == spotify.ErrNoMorePages {
			break
		}

		if err != nil {
			return err
		}
	}
	return nil
}

// Download sets up and starts the downloading of the playlist tracks
func Download(c *spotify.Client, id *string, playlistName *string) error {
	startIndex, err := getPlaylistIndex(playlistName)
	if err != nil {
		return err
	}

	if startIndex == 0 {
		if err := downloadPlaylist(id); err != nil {
			return err
		}
		return nil
	}

	startIndex = startIndex - 1
	if err := pageDownload(c, id, &startIndex); err != nil {
		return err
	}
	return nil
}

func downloadPlaylist(playlistID *string) error {
	if err := download(playlistID, "playlist"); err != nil {
		return err
	}
	return nil
}

func downloadSong(trackID string) error {
	if err := download(&trackID, "track"); err != nil {
		return err
	}
	return nil
}

func download(id *string, contentType string) error {
	fmt.Printf("Downloading https://open.spotify.com/%v/%v\n", contentType, *id)
	cmd := exec.Command("spotdl", fmt.Sprintf("https://open.spotify.com/%v/%v", contentType, *id))
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("\rOn %v/10\n", scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func getPlaylistIndex(playlistName *string) (int, error) {
	path, _ := os.LookupEnv("GOSMS")
	if err := os.Chdir(path); err != nil {
		return 0, err
	}

	if err := os.Chdir(*playlistName); err != nil {
		return 0, err
	}

	files, err := filepath.Glob("*")
	if err != nil {
		return 0, err
	}
	return len(files), nil
}
