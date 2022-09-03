# Simple file archiver
It can pack and unpack files with shannon-fano algo

## Usage
```bash
make build
./archiver pack -m shannon_fano file.txt
./archiver unpack -m shannon_fano file.vlc 
```
### Run unit-tests
```bash
make tests
```
### Run integration-tests
```bash
make all
```

## TODO
1. Refactoring (project structure mainly)
2. Rewrite tests for vlc
3. Adding logic with static encoding table