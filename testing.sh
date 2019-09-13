#!/bin/bash

rm -rf ./users/*


curl -i localhost:8080/user/create/shivang
curl -i localhost:8080/user/create/roz3x


curl -i localhost:8080/frnd/roz3x/shivang
curl -i localhost:8080/fnrd/shivang/roz3x


curl -i localhost:8080/post/roz3x/shivang/hellobrother
curl -i localhost:8080/post/shivang/roz3x/supbro
