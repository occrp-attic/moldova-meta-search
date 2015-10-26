# moldova-meta-search
Meta-search work at the Moldova hackathon

Project installation:

- install go language (https://golang.org/doc/install);
- run `go get github.com/gin-gonic/gin` to install required framework;
- check the config folder (there is the apache vhost and an example of supervisord setup);
- setup the reverse proxy and enable apache host;

Front-end stuff:

- gulp, npm and bower are required;
- run `npm install && bower install`
... currently working on seting up a stable build.

:)


## Development

### Vagrant
#### Steps

1. Ensure you have installed:
  * virtualbox (version 4.x recommended, not 5 which is latest and have a lot of bugs)
  * vagrant (>=1.7.4)
  * ansible (for UNIX systems only)

2. clone this repo, cd into project folder and start vagrant
  * `cd moldova-meta-search`
  * `vagrant up`

   now wait for 20+ minutes. If something fails in process - don't panic - it happens,
   just try to run `vagrant provision`. The problem is that sometimes the network is bad,
   other times servers are to busy and so on.

3. Start Go application with
  * `make go`

4. Open new terminal window go in the project folder and run the frontend app with
  * `make gulp`



### Local

in progress


## License

GNU GPLv3
