#!/bin/bash

rm -rf ./users/*


curl -i localhost:8080/user/shivang
curl -i localhost:8080/user/roz3x
curl -i localhost:8080/user/ki2


curl -i localhost:8080/frnd/roz3x/shivang
curl -i localhost:8080/fnrd/shivang/roz3x


curl -i localhost:8080/post/roz3x/shivang/hellobrother
curl -i localhost:8080/post/shivang/roz3x/supbro
