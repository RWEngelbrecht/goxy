# jwter

Pronounced ʤɒtə | jot-ter.
A reverse proxy that adds Services Portal-esque headers to an HTTP request.

## Setup
### For dev:
1. You need the Go compiler set up
2. Clone this repo into go/src/ (location depends on how you set Go up)
3. Create a .env file with the following vars in jwter/ directory:
```
PORT=<port to run jwter on>
URL=<url to proxy to>
CERT_URL=<url to SP cert endpoint> (not in use atm) Note: might not need this here, as jwter only needs to encode jwt, not decode it for now
SIGNING_CERT=<absolute path to cert used for signing jwt>
USER_ID=<user id>
```

### Just for use:
1. Download jwter executable or clone repo - whatever floats your boat
2. Set up .env as seen above in same directory as executable
3. In terminal run /path/to/jwter

## TODO
- refactor / modularize
- improve logging output
- get some sleep
- make this a CLI? 
- switch to using local vars?
