
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
    echo "TODO: gitws clone\n"
}

#==============================================================================
# gitws add <branch>
#
#   This function adds a new branch into the gitws
#   Note: It will perform a fetch if the branch is not found locally
function __gitws_add {
    printf "TODO: gitws add\n"
}

#==============================================================================
# gitws rm <branch>
#
#   This function will remove a branch and perform a prune
function __gitws_rm {
    printf "TODO: gitws rm\n"
}

#==============================================================================
# gitws list
#
#   This function will remove a branch and perform a prune
function __gitws_list {
    printf "TODO: gitws list\n"
}
#==============================================================================
# Common functions
function __private_gitws_interactive {
    printf "TODO: interactive mode\n"
}

#==============================================================================
# gitws
#
#   The main function of gitws
function gitws {
    # Go into interactive mode if no subcommand is chosen
    if [ $# -eq 0 ]; then
        __private_gitws_interactive
        return
    fi
    
    # Restrict valid subcommands 
    valid_subcommands=(
        clone
        add
        rm
        list
    )
    local cmdname=$1; shift
    if ! [[ ${valid_subcommands[*]} =~ (^|[[:space:]])"$cmdname"($|[[:space:]]) ]]; then
        printf "\e[7mError:\e[0m Invalid sub command.\n\n"
        __gitws_help
        return
    fi
    if type "__gitws_$cmdname" >/dev/null 2>&1; then
        "__gitws_$cmdname" "$@"
    fi
}
