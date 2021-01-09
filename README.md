# broadband-speed-twitter-bot

### Configuration:

`configs/config.yml`

```yml
botConfig:
  broadbandProvider: "Vodafone/UnityMedia"
  broadbandProviderTwitterID: "@vodafoneservice"
  # mbit/s
  expectedDownloadSpeed: 150
  expectedUploadSpeed: 10
  # seconds
  measureFrequency: 1800

twitterAPIConfig:
  consumerKey: ""
  consumerSecret: ""
  accessToken: ""
  accessTokenSecret: ""
```

### Usage:
`go build && ./broadband-speed-twitter-bot`

### Tweet sample:

![Tweet sample](tweet_sample.png?raw=true)
