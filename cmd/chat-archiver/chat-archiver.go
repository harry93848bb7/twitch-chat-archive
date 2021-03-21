package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/harry93848bb7/chat-archiver/archiver"
	"github.com/harry93848bb7/chat-archiver/messages"
	"github.com/harry93848bb7/chat-archiver/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type excludeFlags struct {
	flags map[string]bool
}

func (i *excludeFlags) String() string {
	return ""
}

func (i *excludeFlags) Set(value string) error {
	switch value {
	case "messages":
		i.flags[value] = true
		return nil
	case "badges":
		i.flags[value] = true
		return nil
	case "emotes":
		i.flags[value] = true
		return nil
	default:
		return fmt.Errorf(`The possible values to exclude are "messages", "badges" or "emotes"`)
	}
}

func main() {
	var (
		vodID        string
		clientID     string
		outputDir    string
		removeUnused bool
	)
	exclude := excludeFlags{
		flags: make(map[string]bool),
	}
	flag.StringVar(&vodID, "vod_id", "", `Target VOD ID to archive`)
	flag.StringVar(&clientID, "client_id", "", `Twitch Developer Application Client ID`)
	flag.StringVar(&outputDir, "output_dir", ".", `Where to output archive files`)
	flag.BoolVar(&removeUnused, "remove_unused", true, `Remove dangling emotes / badges which are not used throughout the entire VOD`)
	flag.Var(&exclude, "exclude", `What to exclude from the archival. Possible values "messages", "badges" or "emotes"`)
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

	// Custom httpClient with timeouts because we are preforming many requests
	// to different API's which we can't guarantee will be supported in the future.
	var httpClient = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
		Timeout: time.Second * 10,
	}
	client := messages.NewClient(httpClient, clientID)

	// Basic VOD information is provided for all archives
	log.Println("Getting basic VOD information")
	vod, err := client.GetVODInfo(vodID)
	if err != nil {
		log.Println("Failed to get VOD information", err)
		os.Exit(1)
	}
	vID, err := strconv.ParseInt(vodID, 10, 64)
	if err != nil {
		os.Exit(1)
	}
	archive := &protobuf.Archive{
		VodId:      vID,
		Title:      vod.Title,
		Category:   vod.Game,
		Length:     timestamppb.New(time.Unix(int64(vod.Length), 0)),
		RecordedAt: timestamppb.New(vod.RecordedAt),
		Channel: &protobuf.Channel{
			Name: vod.Channel.DisplayName,
			Id:   int64(vod.Channel.ID),
		},
	}
	if !exclude.flags["badges"] {
		log.Println("Getting all global and channel specific badges")
		badge, err := archiver.Badges(httpClient, fmt.Sprintf("%d", vod.Channel.ID))
		if err != nil {
			os.Exit(1)
		}
		archive.Badges = badge
	}
	if !exclude.flags["emotes"] {
		log.Println("Getting all global, thirdparty and channel subscription emotes")
		emote, err := archiver.Emotes(httpClient, fmt.Sprintf("%d", vod.Channel.ID))
		if err != nil {
			os.Exit(1)
		}
		archive.Emotes = emote
	}
	if !exclude.flags["messages"] {
		log.Println("Getting all messages sent during the VOD. This may take a while...")
		message, err := archiver.Messages(httpClient, vodID, clientID)
		if err != nil {
			os.Exit(1)
		}
		archive.Messages = message
	}
	if removeUnused {
		// Remove badges which are not being used in any of the chat messages to reduce file size as they are redundant
		deleted := 0
		for g := range archive.Badges {
			i := g - deleted
			badgeOccurance := 0

			for _, m := range archive.Messages {
				for _, messageBadge := range m.Badges {
					if (messageBadge.Code == archive.Badges[i].Code) && (messageBadge.Version == archive.Badges[i].Version) {
						badgeOccurance++
					}
				}
			}
			if badgeOccurance == 0 {
				archive.Badges = append(archive.Badges[:i], archive.Badges[i+1:]...)
				deleted++
			}
		}
		if deleted != 0 {
			log.Printf("Successfully remove %d badges not being used in the archive messages", deleted)
		}
	}
	if removeUnused {
		// Remove emotes which are not being used in any of the chat messages to reduce file size as they are redundant
		deleted := 0
		for g := range archive.Emotes {
			i := g - deleted
			emoteOccurance := 0

			for _, m := range archive.Messages {
				if strings.Index(m.Content, archive.Emotes[i].Code) != -1 {
					emoteOccurance++
				}
			}
			if emoteOccurance == 0 {
				archive.Emotes = append(archive.Emotes[:i], archive.Emotes[i+1:]...)
				deleted++
			}
		}
		if deleted != 0 {
			log.Printf("Successfully remove %d emotes not being used in the archive messages", deleted)
		}
	}
	// Finally marshal write the archive json and proto files to the directory
	b, err := proto.Marshal(archive)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf(outputDir+"/%d", archive.VodId), b, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to write and save json archive file")
		os.Exit(1)
	}
	b, err = json.MarshalIndent(archive, "", "   ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf(outputDir+"/%d.json", archive.VodId), b, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to write and save protocol buffer archive file")
		os.Exit(1)
	}
}
