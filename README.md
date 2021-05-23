# Spotify Media Server

SMS downloads Spotify playlists in `mp3` format and keeps them up to date!

## Installation

Download the project via
```sh
cd $GOPATH
go get github.com/lepasq/spotify-media-server
cd src/github.com/lepasq/spotify-media-server
```


Then configure your `$SPOTIFY_ID` and your `$SPOTIFY_SECRET` inside `Dockerfile`.  
You can get both of these values by creating a spotify application [here](https://developer.spotify.com/dashboard/login).  
Leave `$GOSMS` unchanged.


Next, you should add the ids of the playlists that you want to keep track of to `config.txt` in the format:  
```
id1
id2
id3
```

Finally, build and run the Docker container with   
```sh
docker build -t spotify-media-server .
docker run -it -v "${PWD}/music:/app/music" --rm --name sms-app spotify-media-server
```
Replace `${PWD}/music` with the folder, in which you want to store your music.



## Setup on Remote Machine
I'd recommend running SMS on a server, which you can also use for hosting a Media Server.
My preferred front-end is [Plex](https://www.plex.tv/).  

I followed [this guide](https://pimylifeup.com/raspberry-pi-plex-server/) in order to set up Plex on my raspberry Pi.


## Common Errors
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