<div align="center">
    <h1>Go-pool</h1>
    <h5>
        Daily tasks from bootcamp in School 21 written in Go
    </h5>
</div>

## Navigation
* **[Day00](#day00)**
* **[Day01](#day01)**
* **[Day02](#day02)**
* **[Day03](#day03)**
* **[Day04](#day04)**
* **[Day05](#day05)**
* **[Day06](#day06)**
* **[Day07](#day07)**
* **[Day08](#day08)**
* **[Day09](#day09)**
* **[Team00](#team00)**
  

## Day00
### Task 
Implement application for read large set of integers between `-100000` and `100000`, separated by newlines, from Stdin. 
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

## Day03
### Task 

1. Use [Elasticsearch](https://www.elastic.co/downloads/elasticsearch) as database to provide the ability to search for things on version 7.9.2. Dataset of restaurants (taken from an Open Data portal) consists of more than 13 thousands of restaurants in the area of Moscow, Russia. Every entry has:
- **ID**
- **Name**
- **Address**
- **Phone**
- **Longitude**
- **Latitude**
2. Create an index and a mapping (use "places" as a name for an index and "place" as a name for an entry). You can create an index using cURL like this:
```
~$ curl -XPUT "http://localhost:9200/places"
```
but in this task you should use Go Elasticsearch bindings to do the same thing. Next thing you have to do is to provide type mappings for our data. With cURL it will look like this:
```
~$ curl -XPUT http://localhost:9200/places/place/_mapping?include_type_name=true -H "Content-Type: application/json" -d @"schema.json"
```
where `schema.json` looks like this:

```
{
  "properties": {
    "name": {
        "type":  "text"
    },
    "address": {
        "type":  "text"
    },
    "phone": {
        "type":  "text"
    },
    "location": {
      "type": "geo_point"
    }
  }
}
```
3. Create an HTML UI for our database. Abstract your database behind an interface. To just return the list of entries and be able to [paginate](https://www.elastic.co/guide/en/elasticsearch/reference/current/paginate-search-results.html) through them, this interface is enough:
```
type Store interface {
    // returns a list of items, a total number of hits and (or) an error in case of one
    GetPlaces(limit int, offset int) ([]types.Place, int, error)
}
```
HTTP application should run on port 8888, responding with a list of restaurants and providing a simple pagination over it. So. when querying "http://127.0.0.1:8888/?page=2" (mind the 'page' GET param) you should be getting a page like this:

```console
$ curl -s -XGET "http://127.0.0.1:8888/?page=2"
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Places</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>

<body>
<h5>Total: 13649</h5>
<ul>
    <li>
        <div>Sushi Wok</div>
        <div>gorod Moskva, prospekt Andropova, dom 30</div>
        <div>(499) 754-44-44</div>
    </li>
    <li>
        <div>Ryba i mjaso na ugljah</div>
        <div>gorod Moskva, prospekt Andropova, dom 35A</div>
        <div>(499) 612-82-69</div>
    </li>
    <li>
        <div>Hleb nasuschnyj</div>
        <div>gorod Moskva, ulitsa Arbat, dom 6/2</div>
        <div>(495) 984-91-82</div>
    </li>
</ul>
<a href="/?page=1">Previous</a>
<a href="/?page=3">Next</a>
<a href="/?page=1364">Last</a>
</body>
</html>
```

A "Previous" link should disappear on page 1 and "Next" link should disappear on last page.
Also, in case 'page' param is specified with a wrong value (outside [0..last_page] or not numeric) your page should return HTTP 400 error and plain text with an error description:

```
Invalid 'page' value: 'foo'
```

4. Implement handler which responds with `Content-Type: application/json` and JSON version of previous task (example for `http://127.0.0.1:8888/api/places?page=3`):

```console
$ curl -s -XGET "http://127.0.0.1:8888/api/places?page=3"
{
  "name": "Places",
  "total": 13649,
  "places": [
    {
      "id": 65,
      "name": "AZERBAJDZhAN",
      "address": "gorod Moskva, ulitsa Dem'jana Bednogo, dom 4",
      "phone": "(495) 946-34-30",
      "location": {
        "lat": 55.769830485601204,
        "lon": 37.486914061171504
      }
    },
    {
      "id": 69,
      "name": "Vojazh",
      "address": "gorod Moskva, Beskudnikovskij bul'var, dom 57, korpus 1",
      "phone": "(499) 485-20-00",
      "location": {
        "lat": 55.872553383512496,
        "lon": 37.538326789741
      }
    },
  ],
  "prev_page": 2,
  "next_page": 4,
  "last_page": 1364
}
```

Also, in case 'page' param is specified with a wrong value (outside [0..last_page] or not numeric) your API should respond with a corresponding HTTP 400 error and similar JSON:
```
{
    "error": "Invalid 'page' value: 'foo'"
}
```
5. Implement searching for *three* closest restaurants, configure sorting for your query:

```
"sort": [
    {
      "_geo_distance": {
        "location": {
          "lat": 55.674,
          "lon": 37.666
        },
        "order": "asc",
        "unit": "km",
        "mode": "min",
        "distance_type": "arc",
        "ignore_unmapped": true
      }
    }
]
```
where "lat" and "lon" are your current coordinates. 
So, for an URL http://127.0.0.1:8888/api/recommend?lat=55.674&lon=37.666 your application should return JSON like this:

```console
$ curl -s -XGET "http://127.0.0.1:8888/api/recommend?lat=55.674&lon=37.666"
{
  "name": "Recommendation",
  "places": [
    {
      "id": 30,
      "name": "Ryba i mjaso na ugljah",
      "address": "gorod Moskva, prospekt Andropova, dom 35A",
      "phone": "(499) 612-82-69",
      "location": {
        "lat": 55.67396575768212,
        "lon": 37.66626689310591
      }
    },
    {
      "id": 3348,
      "name": "Pizzamento",
      "address": "gorod Moskva, prospekt Andropova, dom 37",
      "phone": "(499) 612-33-88",
      "location": {
        "lat": 55.673075576456,
        "lon": 37.664533747576
      }
    },
    {
      "id": 3347,
      "name": "KOFEJNJa «KAPUChINOFF»",
      "address": "gorod Moskva, prospekt Andropova, dom 37",
      "phone": "(499) 612-33-88",
      "location": {
        "lat": 55.672865251005106,
        "lon": 37.6645689561318
      }
    }
  ]
}
```

6. Provide some simple form of authentication with protecting `/api/recommend` endpoint with a [JWT](https://jwt.io/introduction/) middleware, that will check the validity of this token.

Implement an API endpoint `http://127.0.0.1:8888/api/get_token` which sole purpose will be to generate a token and return it, like this (this is an example, your token will likely be different):

```console
$ curl -s -XGET "http://127.0.0.1:8888/api/get_token"
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjAxOTc1ODI5LCJuYW1lIjoiTmlrb2xheSJ9.FqsRe0t9YhvEC3hK1pCWumGvrJgz9k9WvhJgO8HsIa8"
}
```
So by default when querying this API from the browser it should now fail with an **HTTP 401** error, but work when `Authorization: Bearer <token>` header is specified by the client.

### Lessons Learnt

1. Worked with **ElasticSearh**, **JWT**.
2. Implemented simple API with AUTH.

## Day04
### Task 

1. Recreate the server:
```yaml
---
swagger: '2.0'
info:
  version: 1.0.0
  title: Candy Server
paths:
  /buy_candy:
    post:
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - in: body
          name: order
          description: summary of the candy order
          schema:
            type: object
            required:
              - money
              - candyType
              - candyCount
            properties:
              money:
                description: amount of money put into vending machine
                type: integer
              candyType:
                description: kind of candy
                type: string
              candyCount:
                description: number of candy
                type: integer
      operationId: buyCandy
      responses:
        201:
          description: purchase succesful
          schema:
              type: object
              properties:
                thanks:
                  type: string
                change:
                  type: integer
        400:
          description: some error in input data
          schema:
              type: object
              properties:
                error:
                  type: string
        402:
          description: not enough money
          schema:
              type: object
              properties:
                error:
                  type: string
```

Every candy buyer puts in money, chooses which kind of candy to purchase and how many. This data is being sent over to the server via HTTP and JSON and then:

1) If the sum of candy prices (see Chapter 1) is smaller or equal to the amount of money the buyer gave to a machine, the server responds with **HTTP 201** and returns a JSON with two fields - `"thanks"** saying "Thank you!" and "change"` being the amount of change the machine has to give back the customer.
2) If the sum is larger that the amount of money provided, the server responds with **HTTP 402** and an error message in JSON saying:
`"You need {amount} more money!"`,
where `{amount}` is the difference between the provided and expected.
4) If the client provided a negative candyCount or wrong candyType (all five candy types are encoded by two letters, so it's one of "CE", "AA", "NT", "DE" or "YR", all other cases are considered non-valid) then the server respond with **HTTP 400** and an error inside JSON describing what had gone wrong.

> **Hint:** all data from both client and server should be in JSON, so you can test like this, for example:

```console
$ curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": 1}' http://127.0.0.1:3333/buy_candy

{"change":5,"thanks":"Thank you!"}

$ curl -XPOST -H "Content-Type: application/json" -d '{"money": 46, "candyType": "YR", "candyCount": 2}' http://127.0.0.1:3333/buy_candy
{"change":0,"thanks":"Thank you!"}
```

2. Implement a certificate authentication for the server as well as a test client which will be able to query your API using a self-signed certificate and a local security authority to "verify" it on both sides.

Need a local "certificate authority" to manage certificates e.g.: [minica](https://github.com/jsha/minica).

So, because we're talking a full-blown mutual TLS authentication, you'll have to generate two cert/key pairs - one for the server and one for the client. Minica will also generate a CA file called `minica.pem` for you which you'll need to plug into your client somehow (your auto-generated server should already support specifying CA file as well as `key.pem` and `cert.pem` through command line parameters).
Also, generating certificate may require you to use a domain instead of an IP address, so in examples below we will use "candy.tld". For it to work on a local machine you can put it into '/etc/hosts' file.

Your test client should support flags '-k' (accepts two-letter abbreviation for the candy type), '-c' (count of candy to buy) and '-m' (amount of money you "gave to machine"). So, the "buying request" should look like this:

```
~$ ./candy-client -k AA -c 2 -m 50
Thank you! Your change is 20
```

3. Import C to Go with for **Cow say**

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

unsigned int i;
unsigned int argscharcount = 0;

char *ask_cow(char phrase[]) {
  int phrase_len = strlen(phrase);
  char *buf = (char *)malloc(sizeof(char) * (160 + (phrase_len + 2) * 3));
  strcpy(buf, " ");

  for (i = 0; i < phrase_len + 2; ++i) {
    strcat(buf, "_");
  }

  strcat(buf, "\n< ");
  strcat(buf, phrase);
  strcat(buf, " ");
  strcat(buf, ">\n ");

  for (i = 0; i < phrase_len + 2; ++i) {
    strcat(buf, "-");
  }
  strcat(buf, "\n");
  strcat(buf, "        \\   ^__^\n");
  strcat(buf, "         \\  (oo)\\_______\n");
  strcat(buf, "            (__)\\       )\\/\\\n");
  strcat(buf, "                ||----w |\n");
  strcat(buf, "                ||     ||\n");
  return buf;
}

int main(int argc, char *argv[]) {
  for (i = 1; i < argc; ++i) {
    argscharcount += (strlen(argv[i]) + 1);
  }
  argscharcount = argscharcount + 1;

  char *phrase = (char *)malloc(sizeof(char) * argscharcount);
  strcpy(phrase, argv[1]);

  for (i = 2; i < argc; ++i) {
    strcat(phrase, " ");
    strcat(phrase, argv[i]);
  }
  char *cow = ask_cow(phrase);
  printf("%s", cow);
  free(phrase);
  free(cow);
  return 0;
}
```

```console
$ curl -s --key cert/client/key.pem --cert cert/client/cert.pem --cacert cert/minica.pem -XPOST -H "Content-Type: application/json" -d '{"candyType": "NT", "candyCount": 2, "money": 34}' "https://candy.tld:3333/buy_candy"
{"change":0,"thanks":" ____________\n< Thank you! >\n ------------\n        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||\n"}

```

> **Hint:** > For creating `candy.tld:3333` instead of `127.0.0.1:3333` need to take following steps:
> 1) sudo nano /etc/hosts
> 2) add "127.0.0.1 candy.tld"

### Lessons Learnt

1. Worked with TLS server/client.
2. Worked with C from Go.


## Day05
### Task 

1. Write a function `areToysBalanced` which will receive a pointer to a tree root as an argument. The point is to spit out a true/false boolean value depending on if left subtree has the same amount of toys as the right one. The value on the root itself can be ignored. 
Function should return `true` for such trees (0/1 represent false/true, equal amount of 1's on both subtrees):
   
[Binary tree](https://en.wikipedia.org/wiki/Binary_tree) node:

```go
type TreeNode struct {
    HasToy bool
    Left *TreeNode
    Right *TreeNode
}
```

```
    0
   / \
  0   1
 / \
0   1
```
and `false` for such trees (non-equal amount of 1's on both subtrees):

```
  1
 / \
1   0
```

2.  Write function called `unrollGarland()`, which also receives a pointer to a root node. The idea is to go top down, layer by layer, going right on even horisontal layers and going left on every odd. The returned value of this function should be a slice of bools. So, for this tree:

```
    1
   /  \
  1     0
 / \   / \
1   0 1   1
```
The answer will be [true, true, false, true, true, false, true] (root is true, then on second level we go from left to right, and then on third from right to left, like a zig-zag).

3. Implement a function `getNCoolestPresents()`, that, given an unsorted slice of Presents and an integer `n`, will return a sorted slice (desc) of the "coolest" ones from the list.
It should use the PresentHeap data structure inside and return an error if `n` is larger than the size of the slice or is negative.

```go
type Present struct {
    Value int
    Size int
}
```

So, if we represent each Present by a tuple of two numbers (Value, Size), then for this input:

```
(5, 1)
(4, 5)
(3, 1)
(5, 2)
```
the two "coolest" Presents would be [(5, 1), (5, 2)], because the first one has the smaller size of those two with Value = 5.

4. Implement a classic dynamic programming algorithm, also known as "Knapsack Problem". Input is almost the same, as in the previous task - you have a slice of Presents, each with Value and Size, but this time you also have a hard drive with a limited capacity. So, you have to pick only those presents, that fit into that capacity and maximize the resulting value.
Write a function `grabPresents()`, that receives a slice of Present instances and a capacity of your hard drive. As an output, this function should give out another slice of Presents, which should have a maximum cumulative Value that you can get with such capacity.

**Run tests**
```shell
make test
```

### Lessons Learnt

1. Wordked with data structure (Binary tree, Heap).
2. Worked with [DFS](https://en.wikipedia.org/wiki/Depth-first_search), [BFS](https://en.wikipedia.org/wiki/Breadth-first_search), [Knapstack problem](https://en.wikipedia.org/wiki/Knapsack_problem).


## Day06
### Task 

1. Generate 300x300px PNG file with name 'amazing_logo.png'. Image should appear in the same directory as the launched binary executable after compiling.
2. Create a website (blog where everybody will be able to read ideas on the world improvement). Here is a list of features it should have:

- Database: [PostgreSQL](https://www.postgresql.org/);
- Admin panel: (on '/admin' endpoint) where only you can login with just a form for posting new articles;
- Basic markdown support (so it can at least show "###" headers and links in generated HTML);
- Pagination (show no more than 3 thoughts on one page for people to not get too much of your awesomeness);
- Application UI should use port 8888.
- Admin credentials for posting access (login and password) and database credentials (database name and user) should be submitted separately as well in a file called *admin_credentials.txt*.
- When clicking a link to article, user should get to a page with a rendered markdown text and a single "Back" link which brings him/her back to main page.
- Protect from 
3. Implement rate limiting, so if ther are more than a hundred clients per second trying to access it, they should get a **429 Too Many Requests** response.


**Run app**
```shell
make run-srv
```

## Day07
### Task 

1. Implement `mincoins` algorithm.
2. Test it with `_test.go`.

## Day08
### Task 

1. Working with `reflect`

## Day09
### Task 

1. Working with channels. `Fan-in` && `Fan-out` patterns.
