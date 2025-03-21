package main

import (
	"encoding/json"
	"fmt"
	"log"

	"main/client"
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

	client := client.CreateClient("https://localhost:7073")
	resp, err := client.SendMusicByte(jsonDatas)
	if err != nil {
		log.Fatal("error while sending http request -> ", err)
	}

	fmt.Println(resp)
}
