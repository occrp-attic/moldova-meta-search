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

### Local




### Vagrant

#### Prerequisites

1) Ensure you have installed:
    * virtualbox (version 4.x recommended, not 5 which is latest and have a lot of bugs)
    * vagrant (>=1.7.4)
    * ansible (for UNIX systems only)

2) cd into project folder and start vagrant

    * `cd moldova-meta-search`
    * `vagrant up`

3) Start Go application
    * `make go`



## License

GNU GPLv3
