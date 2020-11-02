#!/bin/bash

set -e

export RADIO_ENV=prod
export CONF_DIR=/app/radio/config

exec /app/radio/radio @$