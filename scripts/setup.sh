#!/bin/bash

# SmoothSSH Setup Script
# Copies config to user's config directory

set -e

CONFIG_DIR="$HOME/.config/smoothssh"
mkdir -p "$CONFIG_DIR"

cp "config.yaml" "$CONFIG_DIR/config.yaml"

echo "Config copied to $CONFIG_DIR/config.yaml"
echo ""
echo "Please edit the file to configure your servers and AI settings."
