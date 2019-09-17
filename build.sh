#!/bin//bash
# git add * 
# git commit -m "commented via deploy.sh"

go build -o ./bin/server . 
./bin/server 
