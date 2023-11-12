# dupdup

Duplicate file finder (and some more)

It can also search for empty files/dictionaries.

## Installation

```sh
$ git clone https://github.com/luka-hash/dupdup
$ cd dupdup/
$ go mod tidy
$ go build
```

## Usage

```sh
Usage of ./dupdup:
  -directory string
    	directory to scan (default "./")
  -empty
    	scan for empty files
  -verbose
    	output additional information
```

*NOTE*: For now, the preffered workflow is to pipe the output to the file, like so:

```sh
$ ./dupdup -directory ~/Pictures/ >duplicates.txt
```

## Is it any good?

Yes.

## TODO

- [ ] Make it interactive.
    - For start, go over each collision, and ask what file(s) to save.
- [ ] Make it faster.
    - Caching is a thing.
- [ ] Make it handle all edge cases correctly
    - Add tests.
    - Re-do directory handling.
    - Add some more tests.

## Licence

This code is licensed under MIT licence (see LICENCE for details).

