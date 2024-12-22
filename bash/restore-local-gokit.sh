#!/bin/bash

# Get the cache path for gokit
GOKIT_CACHE_PATH=$(go list -m -json github.com/weiloon1234/gokit | jq -r .Dir)

# Check if the cache path exists
if [ -d "$GOKIT_CACHE_PATH" ]; then
    echo "Found gokit cache at: $GOKIT_CACHE_PATH"
    echo "Removing gokit cache..."
    rm -rf "$GOKIT_CACHE_PATH"
    echo "Successfully removed gokit cache."
else
    echo "No gokit cache found."
fi

# Re-fetch gokit from GitHub
echo "Fetching gokit from GitHub..."
GOPROXY=direct go get -u github.com/weiloon1234/gokit
echo "Successfully fetched gokit from GitHub."