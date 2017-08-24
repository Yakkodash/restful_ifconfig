# Description
RESTful API server to list network interfaces of the host machine

# Synopsis
* `api_address/list[?json=true|verbose=true]` -- get list of network interfaces of the host machine;
* `api_address/help` -- show help info.
If URL doesn't match any of the described paths, returns help message.

# Supported `GET` parameters
* `json=true` -- return list of network devices in a JSON format;
* `verbose=true` -- if not set, returns only names of interfaces, full info otherwise.

