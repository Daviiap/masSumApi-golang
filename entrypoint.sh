#!/bin/sh
if [ "$1" = "start" ]; then
  exec ./build
else
  exec "$@"
fi