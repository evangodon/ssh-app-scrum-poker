#!/bin/bash

USERNAME=$1

while true
do
  clear
  echo "Connecting..."
  ssh "$USERNAME"@127.0.0.1 -p 8080
  sleep 2
done


