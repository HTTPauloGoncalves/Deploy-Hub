#!/bin/bash
set -e

APP_NAME="deployhub"
URL="https://github.com/HTTPauloGoncalves/Deploy-Hub/releases/download/v0.1.1/deployhub-linux-amd64"
CONFIG_DIR="${XDG_CONFIG_HOME:-$HOME/.config}/deployhub"

curl -L "$URL" -o "$APP_NAME"
chmod +x "$APP_NAME"
sudo mv "$APP_NAME" /usr/local/bin/deployhub
mkdir -p "$CONFIG_DIR"

echo "DeployHub instalado!"
echo "Configuracao padrao em: $CONFIG_DIR/deploy.yaml"
deployhub --help
