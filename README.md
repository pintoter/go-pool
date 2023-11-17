# go-pool
Daily tasks from bootcamp in School 21 written in Go

## Day00
### Task 
Implement application for read large set of integers between-100000 and 100000, separated by newlines, from Stdin. 
After that count four major statistical metrics and print all of them as a result.

### Example
```console
$ make build
$ ./main
Mean: 8.2
Median: 9.0
Mode: 3
SD: 4.35
```
Additionally, add flags processing to display specials of the statistical metrics.
```console
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
