# gitws

## Installation

Gitws requires `fzy` to be installed.

```
# Install requirements
sudo apt install fzy

# Install GitWS
wget -O - https://gitlab.com/johnmperg/gitws/-/raw/v0.0.2/setup.bash | bash
```

## Usage

```bash
# Clone a repo and prepare a worktree
gitws clone <remote>

# Add a new branch into gitws
gitws add <branch>

# Switch between branches
gitws

# Remove a branch that is not longer needed
gitws rm
```
