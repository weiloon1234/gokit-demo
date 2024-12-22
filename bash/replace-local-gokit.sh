#!/bin/bash

# Check if GOKIT_PATH is set
if [ -z "$GOKIT_PATH" ]; then
    echo "Error: GOKIT_PATH is not set. Please set it to your local gokit path."
    exit 1
fi

# Expand GOKIT_PATH environment variable
GOKIT_PATH_EXPANDED=$(eval echo $GOKIT_PATH)

# Check if GOKIT_PATH exists
if [ ! -d "$GOKIT_PATH_EXPANDED" ]; then
    echo "Error: GOKIT_PATH ($GOKIT_PATH_EXPANDED) does not exist."
    exit 1
fi

# Get the cached path of gokit
GOKIT_CACHE_PATH=$(go list -m -json github.com/weiloon1234/gokit | jq -r .Dir)

# Get the Go module cache directory
GO_MODCACHE=$(go env GOMODCACHE)

# Check if GOKIT_CACHE_PATH is inside the Go module cache directory
if [[ "$GOKIT_CACHE_PATH" != "$GO_MODCACHE/"* ]]; then
    echo "Error: Detected GOKIT_CACHE_PATH ($GOKIT_CACHE_PATH) is not inside the Go module cache ($GO_MODCACHE)."
    echo "Aborting to prevent overwriting your development folder."
    exit 1
fi

# Proceed to update the cache
echo "Found gokit cache at: $GOKIT_CACHE_PATH"

# Copy the local gokit folder into the module cache
echo "Syncing gokit from $GOKIT_PATH_EXPANDED to $GOKIT_CACHE_PATH..."
rsync -a --delete "$GOKIT_PATH_EXPANDED/" "$GOKIT_CACHE_PATH/"

echo "Successfully updated gokit in the Go module cache."