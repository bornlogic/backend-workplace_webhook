# WIW - Workplace Integration Webhooks

Server for integration of [Workplace](https://www.workplace.com/) with Webhooks

## Setup

Export `WORKPLACE_ACCESS_TOKEN` with your api token givened by workplace

This access token will be used for send message in group, without this, you can't connect with workplace  
see: [Workplace Generate New Token](./doc/WORKPLACE_CREATE_APP.md)

```sh
export WORKPLACE_ACCESS_TOKEN="< api token >"
```

## Run

After export the env you need to start the server in ./cmd/server/main.go

Examples:  

run on default port (:3000)

``` sh
go run ./cmd/server/main.go
```

run on port 8080
``` sh
go run ./cmd/server/main.go -p ":8080"
```


## Server Configuration

Server contains the webhooks mapped internally, from now, it contains:
 - Github, events: (issues, push on master, pull_request) [see](https://developer.github.com/webhooks/#events)

For any webhook, you have a path for use it, so is listenning on `<ip>:<port>/<service>`, eg: `localhost:3000/github`


### Github

You have an url like `localhost:3000/github` waiting for webhooks comming from github  

Now you need configure in the repository for receive the webhook  
if you are running localhost, you can use [ngrop](https://ngrok.com/download) for expose the service  

Inside your repository on github in configuration you have the option `Webhooks` ([verify](https://developer.github.com/webhooks/) if is available)
Example: https://github.com/bornlogic/wiw/settings/hooks

You will click `Add webhook` and put the url with service and groupID configured for receive the messages.
Example: https://eaa141a6.ngrok.io/github/<groupID>

make sure to enable events `issues`, `pull_request` and `push`.

Now you have github sending webhooks for your server, and can read the messages on your workplace group.  

Open a issue to test this.


## Cli tool

### Group Send

cli used for send message to a given group

For more help try `groupSend -h`:
``` sh
Usage of groupSend:
  -access-token WORKPLACE_ACCESS_TOKEN
        access token used to connect with workplace api, if empty it will use the env WORKPLACE_ACCESS_TOKEN
  -f string
        formatting of message, eg: PLAINTEXT, MARKDOWN(shorthand) (default "PLAINTEXT")
  -formatting string
        formatting of message, eg: PLAINTEXT, MARKDOWN (default "PLAINTEXT")
  -g string
        group id of group for send the message(shorthand)
  -group-id string
        group id of group for send the message
  -m string
        message to send in given group(shorthand)
  -message string
        message to send in given group
  -t WORKPLACE_ACCESS_TOKEN
        access token used to connect with workplace api, if empty it will use the env WORKPLACE_ACCESS_TOKEN(shorthand)
  -v    prints feedback of operations(shorthand)
  -verbose
        prints feedback of operations
```

Examples:

``` sh
go run cmdPath --verbose \
	--access-token <accessToken> \
	--group-id <groupId> \
	--formatting MARKDOWN \
	--message "OLA MUNDO"
```
if `WORKPLACE_ACCESS_TOKEN` was setted you don't need pass the flag `--access-token`
``` sh
export WORKPLACE_ACCESS_TOKEN=<accessToken>
go run cmdPath -g <groupId> -f MARKDOWN -m "OLA MUNDO"
```

## Development

Some useful explains about development

### Requirements

- [go](https://golang.org/)
  - [httprouter](http://github.com/julienschmidt/httprouter)
- [docker](http://docker.com/)

### Test

run all tests
``` sh
make test
```

run specific with args
``` sh
make test args="server/handlers/github/github_test.go -run=TestGithubServe/invalid_status_from -v"
```

### Integration Test

For integration test you need to set `WORKPLACE_GROUP_ID_TEST` and `WORKPLACE_ACCESS_TOKEN` env for test if message is sended

``` sh

```

