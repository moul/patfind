#!/bin/sh

cat *.pub| awk '{print $2}' | patfind | sort -k 4 -n | tail -n 100
