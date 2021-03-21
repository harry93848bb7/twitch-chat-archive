package archiver

import (
	"log"
	"net/http"

	"github.com/harry93848bb7/twitch-chat-archiver/badges"
	"github.com/harry93848bb7/twitch-chat-archiver/emotes"
	"github.com/harry93848bb7/twitch-chat-archiver/messages"
	"github.com/harry93848bb7/twitch-chat-archiver/protobuf"
)

// Emotes ...
func Emotes(client *http.Client, channelID string) ([]*protobuf.Emote, error) {
	var emote = []*protobuf.Emote{}

	emotesClient := emotes.NewClient(client)

	bttvglobal, err := emotesClient.BTTVGlobal()
	if err != nil {
		return nil, err
	}
	emote = append(emote, bttvglobal...)
	bttvuser, err := emotesClient.BTTVChannel(channelID)
	if err != nil {
		return nil, err
	}
	emote = append(emote, bttvuser...)
	ffzglobal, err := emotesClient.FFZGlobal()
	if err != nil {
		return nil, err
	}
	emote = append(emote, ffzglobal...)
	ffzuser, err := emotesClient.FFZChannel(channelID)
	if err != nil {
		return nil, err
	}
	emote = append(emote, ffzuser...)
	twitchglobal, err := emotesClient.TwitchGlobal()
	if err != nil {
		return nil, err
	}
	emote = append(emote, twitchglobal...)
	robot, err := emotesClient.Robot()
	if err != nil {
		return nil, err
	}
	emote = append(emote, robot...)
	channel, err := emotesClient.Channel(channelID)
	if err != nil {
		return nil, err
	}
	emote = append(emote, channel...)

	// Provide warning logs if emotes are duplicate across different sources
	for i := range emote {
		emoteFound := 0
		for d := range emote {
			if emote[i].Code == emote[d].Code {
				emoteFound++
			}
		}
		if emoteFound > 1 {
			log.Printf(`Warning: Duplicate emote "%s" found across sources`, emote[i].Code)
		}
	}
	return emote, nil
}

// Badges ...
func Badges(client *http.Client, channelID string) ([]*protobuf.Badge, error) {
	var badge = []*protobuf.Badge{}

	badgesClient := badges.NewClient(client)

	global, err := badgesClient.TwitchGlobal()
	if err != nil {
		return nil, err
	}
	channel, err := badgesClient.Channel(channelID)
	if err != nil {
		return nil, err
	}

	// Remove twitch global badges when a channel badge exists

	// - A channel can replace the global subscription badge with custom channel subscription badges
	// - A channel can replace the global bit badge with custom channel bit badges
	// - A channel can replace the global cheer badge with custom channel cheer badges

	deleted := 0
	for g := range global {
		for c := range channel {
			i := g - deleted

			if (global[i].Code == channel[c].Code) && (global[i].Version == channel[c].Version) {
				global = append(global[:i], global[i+1:]...)
				deleted++
			}
		}
	}
	log.Printf("Replace %d twitch badges in favor of channel badges\n", deleted)

	badge = append(badge, global...)
	badge = append(badge, channel...)

	// Provide warning logs if badges are duplicate across different sources
	for i := range badge {
		badgeFound := 0
		for d := range badge {
			if (badge[i].Code == badge[d].Code) && (badge[i].Version == badge[d].Version) {
				badgeFound++
			}
		}
		if badgeFound > 1 {
			log.Printf(`Warning: Duplicate badge "%s" version "%s" found across sources`, badge[i].Code, badge[i].Version)
		}
	}
	return badge, nil
}

// Messages ...
func Messages(client *http.Client, vodID, clientID string) ([]*protobuf.Message, error) {
	var message = []*protobuf.Message{}

	messagesClient := messages.NewClient(client, clientID)

	// Continue to move through the message chunks until next is not present
	next := ""
	for {
		chunk, err := messagesClient.GetMessageChunk(next, vodID)
		if err != nil {
			return nil, err
		}
		for _, c := range chunk.Comments {
			var badges = []*protobuf.MessageBadge{}

			for _, b := range c.Message.UserBadges {
				badges = append(badges, &protobuf.MessageBadge{
					Code:    b.ID,
					Version: b.Version,
				})
			}
			message = append(message, &protobuf.Message{
				ContentOffset: c.ContentOffsetSeconds,
				DisplayName:   c.Commenter.DisplayName,
				DisplayColor:  c.Message.UserColor,
				Content:       c.Message.Body,
				Badges:        badges,
			})
		}
		if chunk.Next == "" {
			break
		}
		next = chunk.Next
	}
	return message, nil
}
