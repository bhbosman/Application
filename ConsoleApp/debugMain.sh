#!/usr/bin/env bash

go build -gcflags='main=-N-l'  main.go
dlv --listen=:12345 --headless=true --api-version=2 exec ./main
rm main