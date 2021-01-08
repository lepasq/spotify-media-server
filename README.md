# Spotify Media Server

SMS downloads Spotify playlists in `mp3` format and keeps them up to date!

## Installation

```sh
cd $GOPATH
go get github.com/lepasq/spotify-media-server
cd src/github.com/lepasq/spotify-media-server
go get ./...
go build 
```

Please set `$GOSMS` to the path, where you'd like to download all of your playlists!  
<br>
In addition you'll need to configure your `$SPOTIFY_ID` and your `$SPOTIFY_SECRET`.  
You can get both of these values by creating a spotify application [here](https://developer.spotify.com/dashboard/login).  


Next, you should add the ids of the playlists that you want to keep track of to `config.txt` in the format:  
```
id1
id2
id3
```


## Dependencies
* golang
* spotdl (python3, pytube)
* ffmpeg

```sh
pip install spotdl

sudo pacman -S ffmpeg # Arch Linux (btw)
sudo apt-get install ffpmeg # Debian
brew install ffmpeg # macOS
```

Windows users can install ffmpeg [here](https://ffmpeg.org/download.html#build-windows).  
<br>

To launch SMS, simply run
```sh
./spotify-media-server
```


## Setup on Raspberry Pi
I'd recommend running SMS on a server, which you can also use for hosting a Media Server.
My preferred front-end is [Plex](https://www.plex.tv/).  

To setup Plex on your Raspberry Pi, please follow [this guide](https://pimylifeup.com/raspberry-pi-plex-server/).  


## Common Errors
> Spotdl throws the following: `Download: KeyError: playNavigationEndpoint`  

* Reinstall `spotdl` as suggested [here](https://github.com/spotDL/spotify-downloader/issues/1038):
```sh
pip install pip-autoremove
pip-autoremove spotdl
pip uninstall spotdl # run this line if the command above didn't work
pip install https://codeload.github.com/spotDL/spotify-downloader/zip/master
pip uninstall pytube
pip install git+https://github.com/nficano/pytube
```

> SMS takes really long to start when downloading playlists  

* When downloading a playlist for the first time, SMS will run `spotdl $playlist`. The above command will start by searching the corresponding YouTube video for <b>each</b> track in the playlist. If you are using a large playlist, this process may take quite long.

If it takes more than 5 minutes, you should try running `spotdl $playlist` on its own. If `spotdl` is stuck at `Fetching playlist`, it's very likely that your version of `pytube` is outdated. Try to do reinstall pytube:
```sh
pip uninstall pytube
pip uninstall spotdl

python -m pip install git+https://github.com/nficano/pytube # you might have to use python3 instead of python here
pip install https://github.com/k0mat/spotify-downloader/archive/next-rel-dev.zip
```  

> The folder doesn't contain all songs
* Unfortunately, there's a small chance that `spotdl` won't find an equivalent youtube video for the song. In this case, no song is downloaded.