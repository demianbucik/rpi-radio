#!/bin/bash

iptables -t nat -L --line-numbers

# Delete a rule
# iptables -t nat -D PREROUTING $RULENUMBER
