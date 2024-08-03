#!/bin/bash

set -e

URL=https://github.com/Jonath-z/commit-ai
INSTALL_DIR="/usr/local/bin"
BINARY_NAME=commit-ai

echo "Fetching lastest version..."
curl -lo $BINARY_NAME $URL

echo "Making $BINARY_NAME executable..."
chmod +x $BINARY_NAME

# Move the binary to the install directory
echo "Installing $BINARY_NAME to $INSTALL_DIR..."
sudo mv $BINARY_NAME $INSTALL_DIR


# Verify installation
echo "Verifying installation..."
if command -v $BINARY_NAME &> /dev/null
then
    echo "$BINARY_NAME installed successfully!"
else
    echo "Installation failed!"
    exit 1
fi

# Print success message
echo "Installation completed successfully!"