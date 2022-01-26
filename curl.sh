#!/bin/bash

curl -X POST -d @message1.json 192.168.0.22:30312/api/wordfind
curl -X POST -d @message2.json 192.168.0.22:30312/api/wordfind
curl -X POST -d @message2.json 192.168.0.22:30312/api/wordfind
curl -X POST -d @message1.json 192.168.0.22:30312/api/wordfind
