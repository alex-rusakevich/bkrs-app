build:
	go build -ldflags="-s -w -H windowsgui" .

build-dev:
	go build .

run:
	bkrs.exe

br: build run

dbr: build-dev run
