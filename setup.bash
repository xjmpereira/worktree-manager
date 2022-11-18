#!/usr/bin/env bash

if [ ! -d ~/.config/gitws ]; then
    mkdir -p ~/.config/gitws
fi

wget -O ~/.config/gitws/gitws.bash https://gitlab.com/johnmperg/gitws/-/raw/v1.0.0/src/gitws.bash

if [ -z "$(cat ~/.bashrc | grep 'source ~/.config/gitws/gitws.bash')" ]; then
    echo 'source ~/.config/gitws/gitws.bash' >> ~/.bashrc
fi
source ~/.config/gitws/gitws.bash
