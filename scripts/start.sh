#!/bin/bash
./cmd/crons/main > ./crons.log 2>&1 &
./cmd/api/main > ./api.log 2>&1 &
tail -f *.log
