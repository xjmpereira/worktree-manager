# gitws

## Installation

```
# Install GitWS
curl -fsSL https://github.com/xjmpereira/worktree-manager/releases/download/v1.0.0/install.sh | sh

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
