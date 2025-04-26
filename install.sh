#!/usr/bin/env sh
GITLAB_NAME=ws
GITLAB_PROJECT=41089011
GITWS_BINARY=ws-cmd
GITWS_VERSION=main
GITWS_CONFIG_DIR=.config/gitws

if [ ! -d ~/${GITWS_CONFIG_DIR} ]; then
    mkdir -p ~/${GITWS_CONFIG_DIR}
fi

echo " : Downloading ws shell script"
wget -qO ~/${GITWS_CONFIG_DIR}/ws.sh https://gitlab.com/johnmperg/gitws/-/raw/main/ws.sh

echo " : Downloading ws binary"
wget -qO ~/${GITWS_CONFIG_DIR}/${GITWS_BINARY} https://gitlab.com/api/v4/projects/$GITLAB_PROJECT/packages/generic/$GITLAB_NAME/$GITWS_VERSION/$GITWS_BINARY

echo " : Using sudo to chmod binary"
sudo chmod +x ~/${GITWS_CONFIG_DIR}/${GITWS_BINARY}

echo " : Adding ws script source to .bashrc"
if [ -z "$(cat ~/.bashrc | grep "source ~/${GITWS_CONFIG_DIR}/ws.sh")" ]; then
    echo "source ~/${GITWS_CONFIG_DIR}/ws.sh" >> ~/.bashrc
fi

echo "Installation completed"
echo "Create new terminal or re-source ~/.bashrc"
