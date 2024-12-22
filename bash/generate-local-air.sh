#!/bin/bash

echo "Generating air file"

# Define source and target files based on HOT_RELOAD_MODULE
if [ "$HOT_RELOAD_MODULE" == "cli" ]; then
    SOURCE_FILE="./air-cli.toml"
    TARGET_FILE="./.air-cli-local.toml"
    MAIN_FILE_PATH="./cmd/cli"
    REPLACEMENT_CMD='./bash/replace-local-gokit.sh && go build -o ./tmp/cli ./cmd/cli'
else
    SOURCE_FILE="./.air.toml"
    TARGET_FILE="./.air-local.toml"
    MAIN_FILE_PATH="./cmd/web"
    REPLACEMENT_CMD='./bash/replace-local-gokit.sh && go build -o ./tmp/web ./cmd/web'
fi

# Ensure the source file exists
if [ ! -f "$SOURCE_FILE" ]; then
    echo "Error: $SOURCE_FILE does not exist."
    exit 1
fi

# Check if GOKIT_PATH is set and valid
if [ -z "$GOKIT_PATH" ]; then
    echo "Error: GOKIT_PATH is not set. Please set it to your local gokit path."
    exit 1
fi

# Expand GOKIT_PATH to handle ~ and check if it exists
GOKIT_PATH_EXPANDED=$(eval echo "$GOKIT_PATH")
if [ ! -d "$GOKIT_PATH_EXPANDED" ]; then
    echo "Error: GOKIT_PATH ($GOKIT_PATH_EXPANDED) does not exist."
    exit 1
fi

# Create or update the symlink for gokit
if [ -L ./gokit ]; then
    # Check if the symlink points to the correct path
    CURRENT_TARGET=$(readlink ./gokit)
    if [ "$CURRENT_TARGET" != "$GOKIT_PATH_EXPANDED" ]; then
        echo "Symlink ./gokit exists but points to $CURRENT_TARGET, recreating it."
        rm ./gokit
        ln -s "$GOKIT_PATH_EXPANDED" ./gokit
    else
        echo "Symlink ./gokit already points to $GOKIT_PATH_EXPANDED, no changes needed."
    fi
else
    echo "Creating symlink ./gokit -> $GOKIT_PATH_EXPANDED"
    ln -s "$GOKIT_PATH_EXPANDED" ./gokit
fi

# Copy the source file to the target file
cp "$SOURCE_FILE" "$TARGET_FILE"

# Debug: Check if the cmd line exists
if ! grep -q "^[[:space:]]*cmd = " "$TARGET_FILE"; then
    echo "Error: 'cmd = ' line not found in $TARGET_FILE"
    exit 1
fi

# Use sed to replace the cmd line in the target file
sed -i.bak "/^[[:space:]]*cmd = /c\\
  cmd = \"${REPLACEMENT_CMD}\"\\
" "$TARGET_FILE"

# Clean up backup file created by sed (for macOS compatibility)
if [ -f "${TARGET_FILE}.bak" ]; then
    rm "${TARGET_FILE}.bak"
fi

# Confirm updates
echo "$TARGET_FILE has been updated:"
echo "- Include GOKIT_PATH: $GOKIT_PATH_EXPANDED"
echo "- cmd updated to: $REPLACEMENT_CMD"