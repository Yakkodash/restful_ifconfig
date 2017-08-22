TARGET = server
DOCK = sudo docker

all: build 

build: 
	$(DOCK) build -t $(TARGET) $(CURDIR)

run: build
	$(DOCK) run -it $(TARGET)
	
.PHONY: clean
clean:
	$(DOCK) rm $(shell sudo docker ps -a -q -f status=exited)
	$(DOCK) rmi $(shell sudo docker images -q)

