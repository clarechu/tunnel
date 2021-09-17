#!/usr/bin/env bash


go build -ldflags "-X pkg/version/GitBranch=$()" -o tunnel .
