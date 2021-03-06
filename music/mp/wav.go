package mp

import (
	"fmt"
	"time"
)

//WAVPlayer music player
type WAVPlayer struct {
	stat     int
	progress int
}

//Play play music
func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music ", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing ", source)
}
