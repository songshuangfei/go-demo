package mp

import (
	"fmt"
)

//Player play functiong interface
type Player interface {
	Play(source string)
}

//Play play function
func Play(source, musicType string) {
	var p Player //接口变量，用于复制实例对象中该接口有的方法
	switch musicType {
	case "MP3":
		//这里
		p = &MP3Player{} //从对象实例复制接口，前提是对象实例实现了该接口所有方法
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsopported music type ", musicType)
		return
	}
	p.Play(source)
}
