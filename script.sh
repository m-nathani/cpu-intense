#!/bin/bash

[[ -z "${FLENGTH}" ]] && flength='100000' || flength="${FLENGTH}"

[[ -z "${NGOROUTINES}" ]] && ngoroutines='10' || ngoroutines="${NGOROUTINES}"

go run main.go ${flength} ${ngoroutines}