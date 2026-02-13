
# medown

medown is a simple cross‑platform CLI to download videos from popular platforms by just pasting the URL.

- Native Go downloader for YouTube (using kkdai/youtube).
- Twitter / X downloads delegated to yt-dlp.
- Progress bar for YouTube downloads.
- Default output directory is your Videos folder; can be overridden per command.

Project status: personal/experimental. APIs and sites can change at any time.

---

## Features

- Detects platform from URL (YouTube, X/Twitter).
- Downloads best available muxed (video + audio) format for YouTube.
- Shows a progress bar while downloading YouTube videos.
- Uses yt-dlp under the hood for X/Twitter.
- Saves files by default to $HOME/Videos, or to a custom directory via -o/--output.

---

## Requirements

- Go 1.21+ installed.
- On Windows, macOS or Linux, a working terminal/PowerShell.

For X / Twitter support:

- yt-dlp available in your PATH.

Example (Windows with winget):

```powershell
winget install yt-dlp.yt-dlp
yt-dlp --version
```
---
## Installation
Assuming your module path is github.com/somnifobia/medown:
```
go install github.com/somnifobia/medown@latest
```
This will install the medown binary into $GOPATH/bin (or $GOBIN if set).  
Make sure that directory is on your PATH so you can call medown directly.
- - - 
## Usage
Basic YouTube download
```
medown https://www.youtube.com/watch?v=GZYm3sLuvwY
```
-   Downloads the best muxed (video + audio) MP4.   
-   Shows a progress bar during download.    
-   Saves the file to your $HOME/Videos folder by default.

Download to a custom directory
```
medown -o "D:\videos" https://www.youtube.com/watch?v=GZYm3sLuvwY
```
Download from X / Twitter
```
medown https://x.com/user/status/1234567890123456789
# or
medown https://twitter.com/user/status/1234567890123456789
```
-   medown detects the domain and calls yt-dlp with an output template like twitter_<id>.(ext).
-   All progress and extra logs for X/Twitter come from yt-dlp.
---
## How it works (high level)

-   cmd/root.go handles CLI flags and arguments (built with Cobra).
    
-   internal/app contains a router that:
    
    -   Parses the URL.
        
    -   Chooses the correct downloader based on the host (YouTube vs X/Twitter).
        
-   internal/ytdl:
    
    -   Uses github.com/kkdai/youtube/v2 to fetch video metadata and formats.
        
    -   Prefers muxed MP4 formats that contain both video and audio.
        
    -   Streams the response to disk while updating a progress bar from github.com/schollz/progressbar/v3.
        
-   internal/twitterdl:
    
    -   Extracts the tweet ID from the URL.
        
    -   Calls the external yt-dlp binary via os/exec with a predefined output pattern.
        

----------

## TODO / Ideas

-   --audio-only flag for YouTube (download only audio stream, e.g. .m4a).
    
-   --quality flag (e.g. best, 1080p, 720p).
    
-   Forward more options to yt-dlp for X/Twitter (custom format selectors, audio-only, etc.).
    
-   Add support for more platforms by:
    
    -   Implementing native downloaders in Go, or
        
    -   Delegating to yt-dlp when that makes more sense.
        

----------

## Disclaimer

This project is for personal and educational use only.

Make sure you respect the terms of service and copyright rules of each platform you download from.  
The author is not responsible for how this tool is used.
