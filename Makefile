PROJECTNAME=$(shell basename "$(PWD)")
FILENAME=text2
FILEPATH=examples

build:
	@echo "\n > Building binary for $(PROJECTNAME)..."
	go build
tests:
	@echo "\n > Running unit tests..."
	go test archiver/lib/compression/vlc
test-run:
	@echo "\n > Archiving and unarchiving test file $(FILENAME).txt ..."
	./archiver pack -m vlc $(FILEPATH)/$(FILENAME).txt
	./archiver unpack -m vlc $(FILENAME).vlc
	@echo "\n > Comparing unpacked file to original"
	@cmp --silent $(FILEPATH)/$(FILENAME).txt $(FILENAME).txt || echo "there is some difference"
clean:
	@echo "\n > Removing temp test files..."
	rm archiver $(FILENAME).*

all: build tests test-run clean