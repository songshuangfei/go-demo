package main

import (
	"bufio"
	"fmt"
	"go-demo/music/mlib"
	"go-demo/music/mp"
	"os"
	"strconv"
	"strings"
)

var lib *mlib.MusicManager //将用于指向musicManager的指针
var id int = 1             //id,储存音乐时需要
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&mlib.MusicEntry{
				strconv.Itoa(id),
				tokens[2],
				tokens[3],
				tokens[4],
				tokens[5],
			})
		} else {
			fmt.Println("USEAGE: libadd <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USEAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command: ", tokens[1])
	}
}

func handelPlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USEAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])

	if e == nil {
		fmt.Println("The music ", tokens[1], " dose not exist")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
			Enter following Commands to control the player
		----------------------------------------------------
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add music
		lib remove <name> -- Remove music by name
		play <name> -- Play music 
	`)

	lib = mlib.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command >")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "e" || line == "q" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handelPlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command: ", tokens[0])
		}
	}
}
