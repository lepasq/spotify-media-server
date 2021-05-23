FROM golang:1.16-buster

RUN mkdir /app
ADD . /app
WORKDIR /app

ENV SPOTIFY_ID ""
ENV SPOTIFY_SECRET ""
ENV GOSMS "music/"

RUN apt-get update && \
    apt-get install python3-pip -y

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .

RUN wget -c https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz && \
        tar -xvf ffmpeg-release-amd64-static.tar.xz && \
        rm ffmpeg-release-amd64-static.tar.xz && \
        cd ./ffmpeg-4.4-amd64-static && \
        ln -s "${PWD}/ffmpeg" /usr/local/bin/ && \
        ln -s "${PWD}/ffprobe" /usr/local/bin/ && \
        cd ..

RUN pip3 install spotdl

CMD ["/app/main"]