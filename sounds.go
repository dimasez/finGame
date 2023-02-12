package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func moneySound() {
	f, err := os.Open("money.mp3")
	errCheck(err)

	streamer, format, err := mp3.Decode(f)
	errCheck(err)
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	//select {}
}
