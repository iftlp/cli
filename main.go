package main

import (
	"fmt"
	"log"

	"main/pkg"
)

func main() {
	datas, err := pkg.ProcessFile("./musics/test_sound_1.mp3")
	if err != nil {
		log.Fatal("error during process -> ", err)
	}
	fmt.Println(datas)
}
