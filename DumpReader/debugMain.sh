#!/usr/bin/env bash

go build -gcflags='main=-N-l'   -o App
dlv --listen=:12345 --headless=true --api-version=2 exec ./App -- --useUi=true
rm App