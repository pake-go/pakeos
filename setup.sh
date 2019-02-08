#!/usr/bin/env bash

ln -s $(pwd)/.hooks/govet-check $(pwd)/.git/hooks/pre-commit
chmod u+x $(pwd)/.hooks/govet-check
