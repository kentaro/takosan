# Takosan

Takosan is a simple Web interface to Slack ([Ikachan](https://github.com/yappo/p5-App-Ikachan) for Slack).

## Installing

Just `go get` as below:

```
$ go get github.com/kentaro/takosan
```

## Dependencies

You have to also `go get` dependencies as below:

```
$ go get github.com/go-martini/martini
$ go get github.com/martini-contrib/binding
$ go get github.com/nlopes/slack
```

## Usage

First, set your Slack API token

```
$ export SLACK_API_TOKEN="YOUR SLACK API TOKEN"
```

Then, execute `takosan` command.

```
$ takosan [-host string] [-port int] [-name string]
```

## Options

### `-host` (default: "127.0.0.1")

The interface which `takosan` binds.

### `-port` (default: 4979)

The port to which `takosan` listens.

### `-name` (default: "takosan")

The name which you want to display on Slack for this bot.

## API

### `/notice`
### `/privmsg`

```
$ curl -d "channel=#channel&message=test message" localhost:4979/privmsg
```

You can use both of the endpoints to send messages to Slack. No change can be seen on Slack, though.

### `/join`
### `/leave`

When you post requests to these endpoints, the server always returns 404. Which is because you don't need to join/leave groups on Slack explicitely.

## License

MIT

## Author

[Kentaro Kuribayashi](http://kentarok.org)
