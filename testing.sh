#!/bin/bash

rm -rf ./users/*


curl -i localhost:8080/user/peka/123
curl -i localhost:8080/user/roz3x/157
curl -i localhost:8080/user/maria/894


curl -i localhost:8080/frnd/roz3x/peka
curl -i localhost:8080/frnd/peka/roz3x
curl -i localhost:8080/frnd/maria/peka

curl -i localhost:8080/post/roz3x/peka/hellobrother
curl -i localhost:8080/post/peka/roz3x/supbro
curl -i localhost:8080/post/maria/roz3x/hii