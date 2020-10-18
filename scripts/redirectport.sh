#!/bin/bash

P=$1

iptables -t nat -I PREROUTING -p tcp --dport 80 -j REDIRECT --to-ports $P

