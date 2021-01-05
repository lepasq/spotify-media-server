package download

import (
	"bufio"
	"fmt"
	"os/exec"
	"time"

	"github.com/zmb3/spotify"
)

// Download downloads the new songs from playlists
func Download(c *spotify.Client) {
	tracks, err := c.GetPlaylistTracks("3YSMjeMVYTKX8HlWoPsmz5")
	if err != nil {
		fmt.Println(err)
	}

	for page := 1; ; page++ {
		err := c.NextPage(tracks)
		if err == spotify.ErrNoMorePages {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		for i, track := range tracks.Tracks {
			trackDate, err := time.Parse("2006-01-02", string(track.AddedAt[:10]))
			if err != nil {
				fmt.Println(err)
			}
			t := time.Now()
			fmt.Printf("%v: %v::%v::%v::%v\n", i+1+page*100, track.Track.Name, track.Track.Artists, track.AddedAt, track.Track.ID)
			if t.Year() == trackDate.Year() && t.YearDay() == trackDate.YearDay() {
				downloadSong(track.Track.ID.String())
			}
		}

	}
}

func downloadPlaylist(playlistID string) {
	download(playlistID, "playlist")
}

func downloadSong(trackID string) {
	download(trackID, "track")
}

func download(id string, contentType string) {
	cmd := exec.Command("spotdl", fmt.Sprintf("https://open.spotify.com/%v/%v", contentType, id))
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		fmt.Printf("Couldn't start download of track %v\n", id)
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}
