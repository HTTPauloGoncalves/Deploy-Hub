#!/bin/bash
set -e

APP_NAME="deployhub"
URL="https://github.com/HTTPauloGoncalves/Deploy-Hub/releases/latest/download/deployhub-linux-amd64"

curl -L "$URL" -o "$APP_NAME"
chmod +x "$APP_NAME"
sudo mv "$APP_NAME" /usr/local/bin/deployhub

echo "DeployHub instalado!"
deployhub --help