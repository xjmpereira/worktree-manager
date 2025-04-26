# WS

## Installation

```
# Install WS
wget -O - https://gitlab.com/johnmperg/ws/-/raw/main/install.sh | bash

# Required to source bashrc (or create a new terminal)
source ~/.bashrc
```

## Usage

```bash
# Clone a repo and prepare a worktree
ws clone <remote>

# Add a new branch into gitws
ws add [branch]

# Remove a branch that is not longer needed
ws rm [branch]

# List branches available in gitws
ws list

# Create new branch locally using gitws
ws create <branch>

# Switch between branches
ws
```
