#!/usr/bin/env bash

if [ ! -d ~/.config/gitws ]; then
    mkdir -p ~/.config/gitws
fi

wget -O ~/.config/gitws/gitws.bash https://gitlab.com/johnmperg/gitws/-/raw/v0.0.2/src/gitws.bash

if [ ! -z "$(cat ~/.bashrc | grep 'source ~/.config/gitws/gitws.bash')" ]; then
    source ~/.config/gitws/gitws.bash
fi