# Group Send

Cli used for send message to a given group

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
