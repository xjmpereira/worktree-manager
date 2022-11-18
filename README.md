# gitws

## Installation

Gitws requires `fzy` to be installed.

```
# Install requirements
sudo apt install fzy

# Install GitWS
wget -O - https://gitlab.com/johnmperg/gitws/-/raw/v1.2.0/setup.bash | bash

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

# Switch between branches
gitws
```
