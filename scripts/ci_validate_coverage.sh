#!/bin/bash

set -e

PERCENTAGE=$(awk '/statements/{print $NF}' coverage/coverage.txt)
PERCENTAGE_THRESHOLD=80.0
PERCENTAGE_NUMBER=${PERCENTAGE/\%/}

echo ""
if (( $(echo "$PERCENTAGE_NUMBER $PERCENTAGE_THRESHOLD" | awk '{print ($1 > $2)}') )); then
    echo "PASSED - Coverage: $PERCENTAGE";
else
    echo "FAILED - Coverage: $PERCENTAGE < 80.0%";
    exit 1
fi