# go-pool
Tasks from bootcamp in School 21 written in Go


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