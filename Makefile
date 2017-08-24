TARGET = server
DOCK = sudo docker

all: build 

build: 
	$(DOCK) build -t $(TARGET) $(CURDIR)

run: build
	$(DOCK) run --net=host -it $(TARGET) -rm
	
.PHONY: clean
clean:
	$(DOCK) rmi -f $(shell sudo docker images -q)
	$(DOCK) rm -f $(shell sudo docker ps -a -q -f status=exited)

