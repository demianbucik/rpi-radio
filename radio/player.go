package main

import (
	"log"

	vlc "github.com/adrg/libvlc-go/v3"
)

func play() {
	// Initialize libVLC. Additional command line arguments can be passed in
	// to libVLC by specifying them in the Init function.
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	// Create a new player.
	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	// Add a media file from path or from URL.
	// Set player media from path:
	// media, err := player.LoadMediaFromPath("localpath/test.mp4")
	// Set player media from URL:
	media, err := player.LoadMediaFromURL("https://r3---sn-uxax3vhna-aw0l.googlevideo.com/videoplayback?expire=1596250428&ei=3IQkX-auDZ7n1gLhmZvgCA&ip=89.142.35.212&id=o-AAj3iKfW9GfiXRq8WeYjvzwQx5xur9r4lvJDeWL2h7Fd&itag=251&source=youtube&requiressl=yes&mh=j-&mm=31%2C29&mn=sn-uxax3vhna-aw0l%2Csn-hpa7znsz&ms=au%2Crdu&mv=m&mvi=3&pl=18&initcwndbps=805000&vprv=1&mime=audio%2Fwebm&gir=yes&clen=3903075&dur=244.001&lmt=1500105863731857&mt=1596228691&fvip=3&keepalive=yes&fexp=23883098&c=WEB&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cvprv%2Cmime%2Cgir%2Cclen%2Cdur%2Clmt&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AG3C_xAwRAIgU3z0xJLS4Q9u62cBqE_qpmrQrzagYcC-lJJhCV_kMUkCIF6z7otDHjvizADwYf9ON8eN2d5sYohTzkU_jiAOO6Yy&sig=AOq0QJ8wRAIgbOpiRHzVNOiqC3lpOVoZ2Kjq8oP7VnaWCpFuGIggYt0CIBPLi0gAWwtgZnto_jscSxEWJ38IfZ8VWJ1PUixEsqla&ratebypass=yes")
	// media, err := player.LoadMediaFromURL("https://www.youtube.com/watch?v=aIHF7u9Wwiw")
	if err != nil {
		log.Fatal(err)
	}
	defer media.Release()

	// Start playing the media.
	err = player.Play()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve player event manager.
	manager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	// Register the media end reached event with the event manager.
	quit := make(chan struct{})
	eventCallback := func(event vlc.Event, userData interface{}) {
		close(quit)
	}

	eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Detach(eventID)

	<-quit
}
