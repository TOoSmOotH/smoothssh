#!/bin/bash

echo "Setting up SmoothSSH config..."

mkdir -p ~/.config/smoothssh
cp /home/mreeves/Projects/Personal/smoothssh/config.yaml ~/.config/smoothssh/config.yaml

echo "Config copied to ~/.config/smoothssh/config.yaml"
echo "Please edit the file to add your SSH servers and AI settings."
