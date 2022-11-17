
#==============================================================================
# gitws help
function __gitws_help {
    echo "Git WS is a workspace manager for git. By simplifying the usage"
    echo "of the git worktree command."
    echo ""
    echo "Not using a subcommand will activate interactive branch selection"
    echo ""
    echo "Usage: gitws [option]"
    echo ""
    echo "Options:"
    echo "    help  - This help message"
    echo "    clone - Clone a new repo and setup for gitws"
    echo "    add   - Add a new branch for the current gitws"
    echo "    rm    - Remove a branch from gitws"
    echo "    list  - List the current available branches"
}

#==============================================================================
# gitws clone <repo>
#
#   This function will clone a repo into the directory <repo>/<main-branch>/
function __gitws_clone {
    echo "TODO: gitws clone"
}

#==============================================================================
# gitws add <branch>
#
#   This function adds a new branch into the gitws
#   Note: It will perform a fetch if the branch is not found locally
function __gitws_add {
    echo "TODO: gitws add"
}

#==============================================================================
# gitws rm <branch>
#
#   This function will remove a branch and perform a prune
function __gitws_rm {
    echo "TODO: gitws rm"
}

#==============================================================================
# gitws list
#
#   This function will remove a branch and perform a prune
function __gitws_list {
    echo "TODO: gitws list"
}
#==============================================================================
# Common functions
function __private_gitws_interactive {
    echo "TODO: interactive mode"
}

#==============================================================================
# gitws
#
#   The main function of gitws
function gitws {
    if [ $# -eq 0 ]; then
        __private_gitws_interactive
        return
    fi
    local cmdname=$1; shift
    if type "__gitws_$cmdname" >/dev/null 2>&1; then
        "__gitws_$cmdname" "$@"
    fi
}
