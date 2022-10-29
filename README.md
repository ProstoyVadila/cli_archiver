# A simple file archiver
It can pack and unpack files with shannon-fano algo.
Based on Nikolay Tuzov youtube lessons.

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
1. Refactor (project structure mainly).
2. Rewrite tests for vlc.
3. Add logic with static encoding table.
4. Add support for another algorithms.
