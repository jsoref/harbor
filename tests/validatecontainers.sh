#!/bin/sh
set -e

if docker ps --filter "status=restarting" | grep 'vmware'; then
  echo "container is restarting, fail CI."
  exit 1
fi
