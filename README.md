# moldova-meta-search
Meta-search work at the Moldova hackathon

Project installation:

- install the Go language (https://golang.org/doc/install);
- run `go get github.com/gin-gonic/gin` to install the required framework;
- check the config folder (there is an apache vhost and an example of supervisord setup);
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
  * virtualbox (version 4.x recommended, not 5 which is the latest one and has a lot of bugs)
  * vagrant (>=1.7.4)
  * ansible (for UNIX systems only)

2. clone this repo, `cd` into project folder and start vagrant
  * `cd moldova-meta-search`
  * `vagrant up`

   Now wait for 20+ minutes. If something goes wrong - don't panic - it happens.
   Just try to run `vagrant provision`. The problem is that sometimes the network is bad,
   other times servers are to busy and so on.

3. Start the Go application with
  * `make go`

4. Open the new terminal window, go to the project folder and run the frontend app with
  * `make gulp`

  The app is now available at http://192.168.33.110:3000/

  Enjoy!


### Local

in progress


## License

GNU GPLv3
