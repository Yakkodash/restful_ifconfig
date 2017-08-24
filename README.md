[![Build Status](https://travis-ci.org/Yakkodash/restful_ifconfig.svg?branch=master)](https://travis-ci.org/Yakkodash/restful_ifconfig)

# Description
RESTful API server to list network interfaces of the host machine

# Synopsis
* `api_address:32432/list[?json=true|verbose=true]` -- get list of network interfaces of the host machine
* `api_address:32432/help` -- show readme file

If URL doesn't match any of the described paths, returns readme.

# Supported `GET` parameters
* `json=true` -- return list of network devices in a JSON format (uppercased)
* `verbose=true` -- if not set, returns only names of interfaces, full info otherwise

# Return codes
* `200` -- on success
* `520` -- on errors

# Usage example
`localhost:32432/list?json=true`

# Run on host
`docker run --net=host -it yakkodash/restful_ifconfig`

# Build on host
`docker build -t yakkodash/restful_ifconfig:latest .`

