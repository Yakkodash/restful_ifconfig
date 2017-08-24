# Description
RESTful API server to list network interfaces of the host machine

# Synopsis
* `api_address/list[?json=true|verbose=true]` -- get list of network interfaces of the host machine
* `api_address/help` -- show readme file
If URL doesn't match any of the described paths, returns readme.

# Supported `GET` parameters
* `json=true` -- return list of network devices in a JSON format
* `verbose=true` -- if not set, returns only names of interfaces, full info otherwise

# Host machine usage
* `make build` -- build in `docker`
* `make run` -- run in `docker`

# Docker pull command
* docker pull yakkodash/restful_ifconfig
