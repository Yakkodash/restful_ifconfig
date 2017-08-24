TARGET = restful_ifconfig
DOCK = sudo docker
USERNAME = yakkodash

all: build 

build: 
	$(DOCK) build -t $(USERNAME)/$(TARGET):latest $(CURDIR)

run: build
	$(DOCK) run --net=host -it $(USERNAME)/$(TARGET) -rm
	
.PHONY: clean
clean:
	$(DOCK) rmi -f $(shell sudo docker images -q)
	$(DOCK) rm -f $(shell sudo docker ps -a -q -f status=exited)

