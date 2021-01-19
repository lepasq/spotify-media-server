package main

import (
	"fmt"
	"time"

	"github.com/lepasq/spotify-media-server/schedule"
)

func main() {
	if err := schedule.Watch(time.Hour * 24); err != nil {
		fmt.Println(err)
	}
}
