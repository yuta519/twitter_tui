<div align="center">
	<h1>Twitter TUI</h1>
	<p><b>Fetching tweets from Twitter API in CLI</b></p>
	<br>
</div>


## About
This is a Twitter Client in CLI. When running this app, you could fetch new tweets.


## Techniques
- Go 1.18

- Using libraries:
  - [anaconda](https://github.com/ChimeraCoder/anaconda)

- Using external service
  - [Twitter API v1.1](https://developer.twitter.com/en/docs/twitter-api/v1)


## Usage
This section expains how to use this library.

1. Download the library.
```bash
 $ go get github.com/yuta519/twitter_tui
 $ go mod tidy
```

2. Set you Twitter API information at `.ennrc` (below information).
- API Key
- API Secrret
- Access Key
- Access Secret

```bash
$ vim .envrc
```

```
export API_KEY="xxxxxxxxxxxxxxxxxxxxxxx"
export API_SECRET="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
export ACCESS_KEY="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
export ACCESS_SECRET="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
```

3. Run `cmd/timeline/main.go`.
```bash
$ go run cmd/timeline/main.go
```

## Sample
### How works
<img src="https://raw.githubusercontent.com/yuta519/assets/master/twitter_tui/twitter_tui_sample_video.gif" alt="How to work">


### Usage sample
If you use tmux, you could fetch tweets on a pane during coding.
<img src="https://raw.githubusercontent.com/yuta519/assets/master/twitter_tui/twitter_tui_sample_screenshot.png" alt="How to work">


## WHY I created this library
I wanted to check twitter during coding.

<!-- # Architecture -->

<!-- # Upcoming Features -->

## License
Copyright 2022 Yuta Kawamura

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
