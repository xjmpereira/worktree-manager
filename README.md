# gitws

## Installation

Gitws requires `fzy` to be installed.

```
# Install requirements
sudo apt install fzy

# Install GitWS
wget -O - https://gitlab.com/johnmperg/gitws/-/raw/v1.2.2/setup.bash | bash

# Required to source bashrc (or create a new terminal)
source ~/.bashrc
```

## Usage

```bash
# Clone a repo and prepare a worktree
gitws clone <remote>

# Add a new branch into gitws
gitws add [branch]

# Remove a branch that is not longer needed
gitws rm [branch]

# List branches available in gitws
gitws list

# Create new branch locally using gitws
gitws create <branch>

# Switch between branches
gitws
```

```bash
WS_ROOT_DIR=~/workspaces/test/.ws-root
mkdir -p ${WS_ROOT_DIR}
git -C ${WS_ROOT_DIR} init --bare
git -C ${WS_ROOT_DIR} remote add origin https://github.com/xjmpereira/test-worktree-manager.git
DEFAULT_BRANCH=$(git -C ${WS_ROOT_DIR} ls-remote --symref | sed -n -E 's|^ref: refs/heads/([a-z]+).*|\1|p')
git -C ${WS_ROOT_DIR} fetch origin $DEFAULT_BRANCH
git -C ${WS_ROOT_DIR} worktree add ~/workspaces/test/$DEFAULT_BRANCH $DEFAULT_BRANCH
cd ~/workspaces/test/$DEFAULT_BRANCH
```
