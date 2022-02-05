build:
	go build -ldflags="-w -s -H windowsgui" .

run:
	bkrs.exe

br: build run
