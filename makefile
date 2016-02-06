.PHONY: build all clean test

all: clean build test

build:
	go build -o packaged.so -buildmode=c-shared

clean:
	-rm packaged.so
	-rm packaged.h

test:
	ruby hello.rb
