# WIW - Workplace Integration Webhooks

Server for integration of [Workplace](https://www.workplace.com/) with Webhooks

## Setup

Export `WORKPLACE_ACCESS_TOKEN` with your api token givened by workplace

This access token will be used for send message in group, without this, you can't connect with workplace  
see: [Workplace Generate New Token](./doc/WORKPLACE_GENERATE_TOKEN.md)

```sh
$ export WORKPLACE_ACCESS_TOKEN="< api token >"
```

## Build

For build you can use:

``` sh
$ make build
```

it will generate two binaries:
- `wiw` the server used for start
- `wSendToGroup` used to send message to a given group by clim

## Install

You can use:
``` sh
$ sudo make install
```
it will build and copy the two binaries generated for `/usr/local/bin` folder  

For custom path specify `bin` var  
You can use:
``` sh
$ make install bin=$HOME/bin
```

Or if you prefer you can get by `go`:

``` sh
$ go get github.com/bornlogic/wiw/cmd/...
```

## Run

After export the env you need to start the server in ./cmd/server/main.go

Examples:  

run on default port (:3000)

``` sh
$ wiw
```

run on port 8080
``` sh
$ wiw -p ":8080"
```

## Find GroupID

Inside a group you can check in url of workspace the groupID [see](https://developers.facebook.com/docs/workplace/reference/graph-api/group/):  
Example:  
url of group: https://enterprise.workplace.com/chat/t/123456789026103  
groupID: `123456789026103`  

## Server Configuration

Server contains the webhooks mapped internally, from now, it contains:
 - Github, events: (issues, push on master, pull_request) [see](https://developer.github.com/webhooks/#events)

For any webhook, you have a path for use it `{ip}:{port}/{service}/{groupID}`, eg: `localhost:3000/github/1203140219421`

### Github

You have an url like `localhost:3000/github/<groupID>` waiting for webhooks comming from github  

#### Configuration

if you are running localhost, you can use [ngrok](https://ngrok.com/download) for expose the service  

Inside your repository in configuration you have the option `Webhooks` ([verify](https://developer.github.com/webhooks/) if is available)
Example: https://github.com/bornlogic/wiw/settings/hooks

You will click `Add webhook` and put the url with service and groupID configured for receive the messages.
Example: `https://eaa141a6.ngrok.io/github/<groupID>`

make sure to enable events `issues`, `pull_request` and `push`.

Now you have github sending webhooks for your server, and can read the messages in your workplace group.  

Open an issue to test this.

## Cli tool

### Group Send

cli used for send message to a given group

For more help try `wSendToGroup -h`:
``` sh
$ wSendToGroup -h
Usage of wSendToGroup:
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

verbose mode
``` sh
$ wSendToGroup --verbose \
	--access-token <accessToken> \
	--group-id <groupId> \
	--formatting MARKDOWN \
	--message "HELLO WORLD"
```
if `WORKPLACE_ACCESS_TOKEN` was setted you don't need pass the flag `--access-token`
``` sh
$ export WORKPLACE_ACCESS_TOKEN=<accessToken>
$ wSendToGroup -g <groupId> -f MARKDOWN -m "HELLO WORLD"
```

## Development

Some useful explains about development

### Requirements

- [go](https://golang.org/) (for run without docker, outside the makefile)
  - [httprouter](http://github.com/julienschmidt/httprouter)
- [docker](http://docker.com/) (for run without go, inside the makefile)

### Test

run all unitary tests
``` sh
make test
```

run specific with args
``` sh
make test args="server/handlers/github/github_test.go -run=TestGithubServe/invalid_status_from -v"
```

#### Integration Test

Integration tests are disabled by default  

For integration test you need to set `WORKPLACE_GROUP_ID_TEST` and `WORKPLACE_ACCESS_TOKEN` env for test if message is sended  

you need specify the build tag `integration` for run integration tests  

Example:

run all with integration tests included
``` sh
make test args="./... -tags=integration"
```

run all with integration tests of cmd
``` sh
make test args="./cmd/... -tags=integration"
```

run specific integration test
``` sh
make test args="./cmd/... "
```

