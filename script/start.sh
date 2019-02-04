#!/bin/sh
#compile go build
go build .

#check if exville process is running
ps -ef | grep ./exville
if [ $? -eq 0 ]
then
  echo "exville Running..."
  echo "killing process..."
  ps -ef | grep ./exville | grep -v grep | awk '{print $2}' | xargs kill
  if [ $? -eq 0 ]
    then
    echo "exville process killed!"
  else
    echo "could not kill exville process" >&2
    exit 1
  fi
else
  echo "process not running"
fi

nohup ./exville &
if [ $? -eq 0 ]
then
  echo "exville process restarted"
  tail -f nohup.out
else
  echo "failed to start exville process" >&2
fi
