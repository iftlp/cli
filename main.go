package main

import (
	"encoding/json"
	"fmt"
	"log"

	"main/pkg"
)

func main() {
	datas, err := pkg.ProcessFile("./musics/test_sound_1.mp3")
	if err != nil {
		log.Fatal("error during process -> ", err)
	}

	jsonDatas, err := json.Marshal(datas)
	if err != nil {
		log.Fatal("error while mashalling -> ", err)
	}

	fmt.Println(string(jsonDatas))
}
