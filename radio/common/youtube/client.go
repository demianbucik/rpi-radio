package youtube

import (
	"errors"
	"sort"

	"github.com/kkdai/youtube/v2"
)

type Client struct {
	client *youtube.Client
}

func NewClient() *Client {
	return &Client{
		client: &youtube.Client{},
	}
}

func (c *Client) Client() *youtube.Client {
	return c.client
}

func (c *Client) GetVideoTitle(url string) (string, error) {
	video, err := c.client.GetVideo(url)
	if err != nil {
		return "", err
	}
	return video.Title, nil
}

func (c *Client) GetBestAudioStreamURL(url string) (string, error) {
	video, err := c.client.GetVideo(url)
	if err != nil {
		return "", err
	}
	var formats []youtube.Format
	var videoFormats []youtube.Format
	for _, format := range video.Formats {
		if format.AudioChannels == 0 {
			continue
		}
		if format.FPS == 0 && format.Width == 0 && format.Height == 0 {
			formats = append(formats, format)
		} else {
			videoFormats = append(videoFormats, format)
		}
	}
	if len(formats) == 0 {
		if len(videoFormats) == 0 {
			return "", errors.New("no audio formats")
		}
		formats = videoFormats
	}

	sort.Slice(formats, func(i, j int) bool {
		return formats[i].Bitrate > formats[j].Bitrate
	})

	return c.client.GetStreamURL(video, &formats[0])
}
