package main

import (
	"broadband-speed-twitter-bot/configs"
	"broadband-speed-twitter-bot/pkg"
	"fmt"
	"log"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	config := new(configs.Configuration).Load()

	for {
		user, _ := speedtest.FetchUserInfo()

		serverList, _ := speedtest.FetchServerList(user)
		targets, _ := serverList.FindServer([]int{})

		for _, s := range targets {
			s.PingTest()
			s.DownloadTest(true)
			s.UploadTest(true)

			var cc string
			if config.BotConfig.BroadbandProviderTwitterID != "" &&
				(config.BotConfig.ExpectedDownloadSpeed >= 2*int(s.DLSpeed) ||
					config.BotConfig.ExpectedUploadSpeed >= 2*int(s.ULSpeed)) {
				cc = fmt.Sprintf("cc %s", config.BotConfig.BroadbandProviderTwitterID)
			}

			message := fmt.Sprintf("Broadband speed test for provider: %s\n\n"+
				"Expected measures (Mbit/s):\n"+
				"* Download: %d\n"+
				"* Upload: %d\n\n"+
				"Actual measures (Mbit/s):\n"+
				"* Download: %f\n"+
				"* Upload: %f\n"+
				"* Ping: %s\n"+
				"* Country: %s\n"+
				"%s",
				config.BotConfig.BroadbandProvider,
				config.BotConfig.ExpectedDownloadSpeed,
				config.BotConfig.ExpectedUploadSpeed,
				s.DLSpeed,
				s.ULSpeed,
				s.Latency,
				s.Country,
				cc)

			client, err := pkg.GetTwitterClient(config)
			if err != nil {
				log.Println("Error getting Twitter Client")
				log.Println(err)
			}
			tweet, resp, err := client.Statuses.Update(message, nil)
			if err != nil {
				log.Println(err)
			}
			log.Printf("%+v\n\n", resp)
			log.Printf("%+v\n\n", tweet)
			log.Println(message)
		}

		time.Sleep(time.Duration(config.BotConfig.MeasureFrequency) * time.Second)
		config = new(configs.Configuration).Load()
	}
}
