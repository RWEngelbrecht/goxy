# jwter

Pronounced ʤɒtə | jot-ter.
A reverse proxy that adds a JWT header to an HTTP request.

## Setup
### For dev:
1. You need the Go compiler set up
2. Clone this repo into go/src/ (location depends on how you set Go up)
3. Run `$> make certs` to generate signing and public certificates
4. Run `$> make conf` and answer the prompts (to skip interactive mode, set each var in terminal, e.g. `$> proxy_port=1337 ... make conf`) 
	- Alternatively, create a .env file with the following vars in jwter/ directory:
	```
	PROXY_PORT=<port to run jwter on>
	APP_URL=<url:port to proxy requests to>
	SIGNING_CERT=<absolute path to cert used for signing jwt>
	PUBLIC_CERT=<absolute path to cert used for decoding jwt>
	USER_ID=<user id>
	INSIGHT_URL=<url:port that insight-proxy runs on> 
	```
5. To run in Docker, run `$> make docker-up`

### Just for use:
1. Download jwter executable or clone repo - whatever floats your boat
2. Set up .env as seen above in same directory as executable
3. In terminal run /path/to/jwter
4. To run in Docker, run `$> make docker-up`

## TODO
- refactor / modularize
- improve logging output
- get some sleep
- make this a CLI?
