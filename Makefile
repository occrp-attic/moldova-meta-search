# Makefile for deployment

PYTHON        = python

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  go  					    starts go application on port 8080"
	@echo "  gulp                       starts angular application on port 3000"
	@echo "  gulp-build                 prepares an distribution package from FE sources"


go: ; vagrant ssh -c 'cd /vagrant/server/ && go run main.go'
gulp: ; vagrant ssh -c 'cd /vagrant/client/ && gulp serve'
gulp-build: ; vagrant ssh -c 'cd /vagrant/client/ && gulp'
