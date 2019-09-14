#!/bin/bash

rm -rf ./users/*


curl -i localhost:8080/user/peka
curl -i localhost:8080/user/roz3x
curl -i localhost:8080/user/maria


curl -i localhost:8080/frnd/roz3x/peka
curl -i localhost:8080/fnrd/peka/roz3x
curl -i localhost:8080/frnd/maria/peka

curl -i localhost:8080/post/roz3x/peka/hellobrother
curl -i localhost:8080/post/peka/roz3x/supbro
curl -i localhost:8080/post/maria/roz3x/hii