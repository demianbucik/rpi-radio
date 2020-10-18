#!/bin/bash

URL=${1:-"localhost:8080"}

# Create a playlist
PLAYLIST_ID=$(http POST "$URL"/playlists <<< '{"name":"Matter"}' | jq -j ".id")

# Create tracks
JIUJITSU_ID=$(http POST "$URL"/tracks <<< '{"name":"Matter ~ Jiu Jitsu", "url":"https://www.youtube.com/watch?v=INekeaoUR4M", "thumbnail":""}' | jq -j ".id")
MANA_ID=$(http POST "$URL"/tracks <<< '{"name":"Matter/persons from porlock ~ MANA", "url":"https://www.youtube.com/watch?v=i2NlVQi9XUE", "thumbnail":""}' | jq -j ".id")

# Add tracks to the start of the playlist
http POST "$URL"/playlists/"$PLAYLIST_ID"/tracks tracks:="[$JIUJITSU_ID,$MANA_ID]"

