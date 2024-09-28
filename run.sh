#!/bin/bash

go build -o bin/SyncEnv
./bin/SyncEnv "$@"