#!/bin/sh

seq 0 1000000 | parallel --nice 19 'echo {}; ssh-keygen -o -a 100 -t ed25519 -P changeme -q -f ./id_ed25519_{} -C "m@42.am"'
