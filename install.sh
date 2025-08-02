#!/usr/bin/env bash

INSTALL_APP=gitws
INSTALL_VERSION=1.0.0
INSTALL_FZF_VERSION=0.65.0

check_cmd() {
    command -v "$1" > /dev/null 2>&1
    return $?
}

if ! check_cmd git; then
    echo "Required 'git' is missing. Installing it..."
    sudo apt update
    sudo apt install -y git
fi

if ! check_cmd fzf; then
    echo "Required 'fzf' is missing. Installing it..."
    curl -sSfL https://github.com/junegunn/fzf/releases/download/v${INSTALL_FZF_VERSION}/fzf-${INSTALL_FZF_VERSION}-linux_amd64.tar.gz | tar zxvf - -C $HOME/.local/bin
fi

if [ ! -d ${HOME}/.config/gitws ]; then
    echo "Creating ~/.config/gitws folder..."
    mkdir -p ${HOME}/.config/gitws
fi

# Actually Installation is a copy of file
echo "Installing gitws..."
curl -sSfL https://github.com/xjmpereira/worktree-manager/releases/download/v${INSTALL_VERSION}/gitws.sh -o ${HOME}/.config/gitws/gitws.sh

if [ -z "$(cat ~/.bashrc | grep 'source ${HOME}/.config/gitws/gitws.sh')" ]; then
    echo "Modifying ~/.bashrc..."
    echo 'source ${HOME}/.config/gitws/gitws.sh' >> ~/.bashrc
fi

echo "Make sure to re-source ~/.bashrc" 
echo "Installation complete"
