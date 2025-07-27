#!/usr/bin/env bash
# Ownership fixes for mounts
function chown_tree {
    dir="$(realpath $1)"
    [[ ! -e $dir ]] && return 0
    while [[ "$dir" =~ ^$HOME.* ]] ; do
        sudo chown ${HOST_USER_NAME}:${HOST_USER_NAME} $dir || true
        dir=${dir%/*}
    done
}
chown_tree $(pwd)
chown_tree ~/.cache
chown_tree ~/.config/nvim

# Prepare command history
if [ ! -d ~/.cache/commandhistory ]; then
    mkdir ~/.cache/commandhistory
fi
touch ~/.cache/commandhistory/.bash_history
sed -i "/^export PROMPT_COMMAND=/{h;s/=.*/='history -a'/};\${x;/^$/{s//export PROMPT_COMMAND='history -a'/;H};x}" ~/.bashrc
sed -i "/^export HISTFILE=/{h;s/=.*/=~\/.cache\/commandhistory\/.bash_history/};\${x;/^$/{s//export HISTFILE=~\/.cache\/commandhistory\/.bash_history/;H};x}" ~/.bashrc

# Install pre commits
pre-commit install
