# Takosan

Takosan is a simple Web interface to Slack ([Ikachan](https://github.com/yappo/p5-App-Ikachan) for Slack).

![](./takosan.jpg)

_Illustrated by [@demiflare168](https://twitter.com/demiflare168)_

## Installing

### For Users

You can choose and get binaries from the [releases](https://github.com/kentaro/takosan/releases) like below:

```
$ wget https://github.com/kentaro/takosan/releases/download/v1.0.3/takosan_linux_amd64 -O takosan
$ chmod +x takosan
```

### For Developers

Just `go get` as below:

```
$ go get github.com/kentaro/takosan
```

#### Dependencies

You have to also `go get` dependencies as below:

```
$ go get github.com/go-martini/martini
$ go get github.com/martini-contrib/binding
$ go get github.com/nlopes/slack
```

### Deploy to Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Usage

First, set your Slack API token.

```
$ export SLACK_API_TOKEN="YOUR SLACK API TOKEN"
```

Then, execute `takosan` command like below:

```
$ takosan [-host string] [-port int] [-name string] [-icon string]
```

## Options

### `-host` (default: "127.0.0.1")

The interface which `takosan` binds.

### `-port` (default: 4979)

The port to which `takosan` listens.

### `-name` (default: "takosan")

The name which you want to display on Slack for this bot.

### `-icon` (default: the URL of the image above)

The icon URL which you want to display on Slack for this bot.

## API

### `/notice`
### `/privmsg`

```
$ curl -d "channel=#channel&message=test message" localhost:4979/privmsg
```

You can use both of the endpoints to send messages to Slack. No change can be seen on Slack, though.

### `/join`
### `/leave`

When you post requests to these endpoints, the server always returns `404`. Which is because you don't need to join/leave groups on Slack explicitely.

## License

MIT

## Author

[Kentaro Kuribayashi](http://kentarok.org)
