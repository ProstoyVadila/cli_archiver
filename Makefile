build:
	go build
tests:
	go test archiver/lib/compression/vlc
run:
	./archiver pack -m vlc examples/text.txt
	./archiver unpack -m vlc text.vlc
cleanup:
	rm archiver text.*

all: build tests run cleanup