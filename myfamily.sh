#!/bin/bash

#curl -s https://01.tomorrow-school.ai/assets/superhero/all.json | jq ' .[] | select ( .id == $HERO_ID) | .connections.relatives' | tr -d "\""

curl -s https://01.tomorrow-school.ai/assets/superhero/all.json | jq --arg hero_id "$HERO_ID" '.[] | select(.id == ($hero_id | tonumber)) | .connections.relatives' | tr -d '"'
