#!/bin/bash

pidfile="./tmp/pid.txt"

# Prevent duplicate events
if [ -f $pidfile ] && [ "$(( $(date +"%s") - $(stat -f "%m" $pidfile) ))" -lt "3" ]; then
  exit 0
fi
pid=$(cat $pidfile 2>/dev/null)
kill $pid 2>/dev/null
make run &
echo $! > $pidfile
