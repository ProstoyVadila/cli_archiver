build:
	@echo "\n > Building binary..."
	go build
tests:
	@echo "\n > Running unit tests..."
	go test archiver/lib/compression/vlc
run:
	@echo "\n > Archiving and unarchiving test file..."
	./archiver pack -m vlc examples/text.txt
	./archiver unpack -m vlc text.vlc
clean:
	@echo "\n > Removing temp test files..."
	rm archiver text.*

all: build tests run clean