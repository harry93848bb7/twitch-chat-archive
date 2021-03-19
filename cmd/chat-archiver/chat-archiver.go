package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	proto "github.com/golang/protobuf/proto"
	"github.com/harry93848bb7/chat-archiver/archiver"
)

func main() {
	var (
		vodID     string
		clientID  string
		outputDir string
	)
	flag.StringVar(&vodID, "vod_id", "", `Target VOD ID to archive`)
	flag.StringVar(&clientID, "client_id", "", `Twitch Developer Application Client ID`)
	flag.StringVar(&outputDir, "output_dir", ".", `Where to output archive files`)
	flag.Parse()

	if vodID == "" {
		fmt.Println("A VOD ID is required in order to archive")
		os.Exit(1)
	}
	if clientID == "" {
		fmt.Println("A Twitch Developer Application Client ID is required in order to archive. Register one here: https://dev.twitch.tv/console/apps")
		os.Exit(1)
	}
	d, err := os.Stat(outputDir)
	if os.IsNotExist(err) {
		fmt.Println(outputDir, "directory does not exist")
		os.Exit(1)
	} else if err != nil {
		panic(err)
	}
	if !d.IsDir() {
		fmt.Println(outputDir, "does not point to a directory")
		os.Exit(1)
	}
	r, err := archiver.Archive(vodID, clientID)
	if err != nil {
		panic(err)
	}
	b, err := proto.Marshal(r)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf(outputDir+"/%d", r.VodId), b, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to write and save json archive file")
		os.Exit(1)
	}
	b, err = json.MarshalIndent(r, "", "   ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf(outputDir+"/%d.json", r.VodId), b, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to write and save protocol buffer archive file")
		os.Exit(1)
	}
}
