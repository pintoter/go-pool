# go-pool
Daily tasks from bootcamp in School 21 written in Go

## Day00
### Task 
Implement application for read large set of integers between-100000 and 100000, separated by newlines, from Stdin. 
After that count four major statistical metrics and print all of them as a result.
Additionally, handle flags to display specials of the statistical metrics.

### Example
```console
$ make build
$ ./main
Mean: 8.2
Median: 9.0
Mode: 3
SD: 4.35

$ make build
$ ./main -median -mean 
Mean: 8.2
Median: 9.0
```

## Day01
### Task 

1. Implement interface `DBReader` for encoding recipes of cakes from JSON or XML.
2. Compare original and stolen databases (should work and both formats JSON/XML).
3. Compare server filesystem backups between original and stolen databases.

### Examples

```console
$ ./readDB -f original_database.xml
$ ./readDB -f stolen_database.json
.
.
.
$ ./compareDB --old original_database.xml --new stolen_database.json
ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee beans" for cake  "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake  "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake  "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake
.
.
.
$ ./compareFS --old snapshot1.txt --new snapshot2.txt
ADDED /etc/systemd/system/very_important/stash_location.jpg
REMOVED /var/log/browser_history.txt
```

### Lessons Learnt

1. Worked with encoding JSON/XML.
2. Improved understanding of interfaces

## Day02
### Task 

1. Implement utility `find` with set of command-line options to be able to locate different types of entries (directories as `-d`, regular files as `-f` and symbolic links as `-sl`).
2. Implement option `-ext` (works only when -f is specified) for user to be able to print only files with a certain extension. 
3. Implement utility `wc` to gather basic statistics about our files with flags: `-l` for counting lines, `-m` for counting characters and `-w` for counting words. For some files need utilize goroutines to process them concurrently.
4. Implement utility `xargs`.
5. Implement log rotation tool `myRotate`. "Log rotation" is a process when the old log file is archived and put away for storage so logs wouldn't pile up in a single file indefinitely.

### Examples

```console
# Finding all files/directories/symlinks recursively in directory /foo
$ ./myFind /foo
/foo/bar
/foo/bar/baz
/foo/bar/baz/deep/directory
/foo/bar/test.txt
/foo/bar/buzz -> /foo/bar/baz
/foo/bar/broken_sl -> [broken]

$ ./myFind -f -ext 'go' /go
/go/src/github.com/mycoolproject/main.go
/go/src/github.com/mycoolproject/magic.go

# Counting words in file input.txt
$ ./myWc -w input.txt
777 input.txt
# Counting lines in files input2.txt and input3.txt
$ ./myWc -l input2.txt input3.txt
42 input2.txt
53 input3.txt
# Counting characters in files input4.txt, input5.txt and input6.txt
$ ./myWc -m input4.txt input5.txt input6.txt
1337 input4.txt
2664 input5.txt
3991 input6.txt

$ ./myFind -f -ext 'log' /path/to/some/logs | ./myXargs ./myWc -l

$ ./myRotate /path/to/logs/some_application.log
```

### Lessons Learnt

1. Worked with operation system with packages: "os", "path/filepath".
2. Worked with goroutines and package "sync".

<!--
Day03

Create candy.tld:3333 instead of 127.0.0.1:3333
1) sudo nano /etc/hosts
2) add "127.0.0.1 candy.tld"


Day04

Create candy.tld:3333 instead of 127.0.0.1:3333
1) sudo nano /etc/hosts
2) add "127.0.0.1 candy.tld"

Example request: 
curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": 1}' http://candy.tld:3333/buy_candy
-->
