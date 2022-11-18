#==================================================================================================
function __gitws_version {
    printf "Git WS version: v1.0.0\n"
}
#==================================================================================================
#   ######   #### ######## ##      ##  ######     ##     ## ######## ##       ########
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##     ## ##       ##       ##     ##
#  ##         ##     ##    ##  ##  ## ##          ##     ## ##       ##       ##     ##
#  ##   ####  ##     ##    ##  ##  ##  ######     ######### ######   ##       ########
#  ##    ##   ##     ##    ##  ##  ##       ##    ##     ## ##       ##       ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##     ## ##       ##       ##
#   ######   ####    ##     ###  ###   ######     ##     ## ######## ######## ##
#==================================================================================================
# gitws help
function __gitws_help {
    if [ $# -gt 0 ]; then
        printf "\e[7mError:\e[0m Too many arguments specified.\n\n"
        __gitws_help
        return 1
    fi

    printf "Git WS is a workspace manager for git. By simplifying the usage\n"
    printf "of the git worktree command.\n"
    printf "\n"
    printf "Not using a subcommand will activate interactive branch selection\n"
    printf "\n"
    printf "Usage: gitws [option]\n"
    printf "\n"
    printf "Options:\n"
    printf "    help  - This help message\n"
    printf "    clone - Clone a new repo and setup for gitws\n"
    printf "    add   - Add a new branch for the current gitws\n"
    printf "    rm    - Remove a branch from gitws\n"
    printf "    list  - List the current available branches\n"
}

function __gitws_root {
    DIR=$(pwd)
    while [ ! -z "$DIR" ] && [ ! -f "$DIR/.gitws" ]; do
        DIR="${DIR%\/*}"
    done
    echo $DIR
}
#==================================================================================================
#   ######   #### ######## ##      ##  ######      ######  ##        #######  ##    ## ########
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##    ## ##       ##     ## ###   ## ##
#  ##         ##     ##    ##  ##  ## ##          ##       ##       ##     ## ####  ## ##
#  ##   ####  ##     ##    ##  ##  ##  ######     ##       ##       ##     ## ## ## ## ######
#  ##    ##   ##     ##    ##  ##  ##       ##    ##       ##       ##     ## ##  #### ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##    ## ##       ##     ## ##   ### ##
#   ######   ####    ##     ###  ###   ######      ######  ########  #######  ##    ## ########
#==================================================================================================
# gitws clone <repo>
#
#   This function will clone a repo into the directory <repo>/<main-branch>/
#
function __gitws_clone_help {
    printf "Usage: gitws clone <repo>\n"
}

function __gitws_clone {
    if [ $# -eq 0 ]; then
        printf "\e[7mError:\e[0m No arguments specified.\n\n"
        __gitws_clone_help
        return 1
    elif [ $# -ge 3 ]; then
        printf "\e[7mError:\e[0m Too many arguments specified.\n\n"
        __gitws_clone_help
        return 1
    fi

    # Get args
    GITWS_REMOTE=$1
    GITWS_ROOT_PREFIX=$2

    # Get the repo name so we can create a folder with same name
    REMOTE_NAME_WITH_EXT=${GITWS_REMOTE##*/}
    REMOTE_NAME=${REMOTE_NAME_WITH_EXT%.*}
    GITWS_ROOT_DIR=$(pwd)/${REMOTE_NAME}

    # Verify that we are not inside a gitws already
    __GITWS_ROOT_DIR=$(__gitws_root)
    if ! [ -z ${__GITWS_ROOT_DIR} ]; then
        printf "\e[7mError:\e[0m Already inside a gitws workspace.\n\n"
        return 1
    fi

    # Verify if root GITWS_ROOT_DIR is valid
    if [ -d ${GITWS_ROOT_DIR} ]; then
        printf "\e[7mError:\e[0m Path already exists: ${GITWS_ROOT_DIR}\n"
        return 1
    fi

    # Clone a temporary version of the repo
    rm -frd /tmp/setup_gitws || true
    git clone ${GITWS_REMOTE} /tmp/setup_gitws

    # Query which is the current branch (which should be the default one)
    # This is required to know which is the Main Git worktree directory
    GITWS_ROOT_BRANCH=$(git -C /tmp/setup_gitws branch --show-current)
    GITWS_GIT_DIR=${GITWS_ROOT_DIR}/${GITWS_ROOT_BRANCH}/${GITWS_ROOT_PREFIX}

    # Create the final directory for the final root directory of gitws for this repo
    mkdir -p ${GITWS_GIT_DIR}
    mv /tmp/setup_gitws/{,.[^.]}* ${GITWS_GIT_DIR}

    # Root git directory has been prepared
    # Set up the require metadata file for GITWS in its root dir
    cat <<EOF > ${GITWS_ROOT_DIR}/.gitws
GITWS_REMOTE=${GITWS_REMOTE}
GITWS_GIT_DIR=${GITWS_GIT_DIR}
GITWS_ROOT_DIR=${GITWS_ROOT_DIR}
GITWS_ROOT_BRANCH=${GITWS_ROOT_BRANCH}
GITWS_ROOT_PREFIX=${GITWS_ROOT_PREFIX}
EOF
    # Set .gitws as read only (this file should be immutable from moment of creation)
    chmod 0444 ${GITWS_ROOT_DIR}/.gitws
}


#==================================================================================================
#   ######   #### ######## ##      ##  ######        ###    ########  ########
#  ##    ##   ##     ##    ##  ##  ## ##    ##      ## ##   ##     ## ##     ##
#  ##         ##     ##    ##  ##  ## ##           ##   ##  ##     ## ##     ##
#  ##   ####  ##     ##    ##  ##  ##  ######     ##     ## ##     ## ##     ##
#  ##    ##   ##     ##    ##  ##  ##       ##    ######### ##     ## ##     ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##     ## ##     ## ##     ##
#   ######   ####    ##     ###  ###   ######     ##     ## ########  ########
#==================================================================================================
# gitws add <branch>
#
#   This function will add a branch into the directory <GITWS_ROOT>/<branch>/<GITWS_PREFIX>
#
function __gitws_add_help {
    printf "Usage: gitws add <branch>\n"
}

function __gitws_add {
    # Verify that we are inside a gitws directory
    __GITWS_ROOT_DIR=$(__gitws_root)
    if [ -z ${__GITWS_ROOT_DIR} ]; then
        printf "\e[7mError:\e[0m Not inside a gitws workspace.\n\n"
        return 1
    fi

    # Setup variables from gitws workspace
    source ${__GITWS_ROOT_DIR}/.gitws

    if [ $# -eq 0 ]; then
        # Get branch lists available on remote
        BRANCH_LIST=$(git -C ${GITWS_GIT_DIR} branch --remotes --format '%(refname)' | sed 's|^refs/remotes/origin/||g' | grep -v '^HEAD$' | grep -v "^${GITWS_ROOT_BRANCH}$")
        __gitws_select_branch
        [ $? -ne 0 ] && return 1
        BRANCH_TO_ADD=$SELECTED_LINE
    elif [ $# -eq 1 ]; then
        BRANCH_TO_ADD=$1
    elif [ $# -ge 2 ]; then
        printf "\e[7mError:\e[0m Too many arguments specified.\n\n"
        __gitws_add_help
        return 1
    fi

    # Add the new branch into worktree
    git -C ${GITWS_GIT_DIR} worktree add ${GITWS_ROOT_DIR}/${BRANCH_TO_ADD}/${GITWS_ROOT_PREFIX} ${BRANCH_TO_ADD}
    cd ${GITWS_ROOT_DIR}/${BRANCH_TO_ADD}/${GITWS_ROOT_PREFIX}
}

#==================================================================================================
#   ######   #### ######## ##      ##  ######     ########  ##     ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##     ## ###   ###
#  ##         ##     ##    ##  ##  ## ##          ##     ## #### ####
#  ##   ####  ##     ##    ##  ##  ##  ######     ########  ## ### ##
#  ##    ##   ##     ##    ##  ##  ##       ##    ##   ##   ##     ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##    ##  ##     ##
#   ######   ####    ##     ###  ###   ######     ##     ## ##     ##
#==================================================================================================
function __gitws_rm_help {
    printf "Usage: gitws rm <branch>\n"
}

function __gitws_rm {
    # Verify that we are inside a gitws directory
    __GITWS_ROOT_DIR=$(__gitws_root)
    if [ -z ${__GITWS_ROOT_DIR} ]; then
        printf "\e[7mError:\e[0m Not inside a gitws workspace.\n\n"
        return 1
    fi

    # Setup variables from gitws workspace
    source ${__GITWS_ROOT_DIR}/.gitws

    if [ $# -eq 0 ]; then
        # Get branch lists available to gitws
        BRANCH_LIST=$(__gitws_list)
        __gitws_select_branch
        [ $? -ne 0 ] && return 1
        BRANCH_TO_REMOVE=$(echo "$SELECTED_LINE" | awk '{gsub(/\[|\]/, "", $3); print $3}')
    elif [ $# -eq 1 ]; then
        BRANCH_TO_REMOVE=$1
    elif [ $# -ge 2 ]; then
        printf "\e[7mError:\e[0m Too many arguments specified.\n\n"
        __gitws_rm_help
        return 1
    fi

    printf "Selected for removal: ${BRANCH_TO_REMOVE}\n"
    if [ ! -z "$(echo "$(pwd)" | grep "^${GITWS_ROOT_DIR}/${BRANCH_TO_REMOVE}" )" ]; then
        printf "Branch selected is in use.!\n"
    fi

    if [ "${GITWS_ROOT_BRANCH}" == "${BRANCH_TO_REMOVE}" ]; then
        printf "\e[7mError:\e[0m Removal of root branch not allowed\n\n"
        return 1
    fi

    read -p " > Confirm removal (y/n)? " yn
    case $yn in 
        [yY] ) ;;
        [nN] ) printf "Operation canceled\n"; return 0;;
        * ) printf "Invalid choice\n"; return 1;;
    esac

    if [ ! -z "$(echo "$(pwd)" | grep "^${GITWS_ROOT_DIR}/${BRANCH_TO_REMOVE}" )" ]; then
        # Change to git worktree root
        cd ${GITWS_GIT_DIR}
    fi

    # Remove the branch from worktree
    printf "Removing branch from worktree\n"
    git -C ${GITWS_GIT_DIR} worktree remove ${GITWS_ROOT_DIR}/${BRANCH_TO_REMOVE}/${GITWS_ROOT_PREFIX}

    # Clean the empty directories
    printf "Cleaning up directories on workspace\n"
    DIR=${GITWS_ROOT_DIR}/${BRANCH_TO_REMOVE}
    while [ ! -z "$DIR" ] && [ ! -f "$DIR/.gitws" ]; do
        if [ ! -z "$(find ${DIR} -type d -empty 2>/dev/null)" ]; then
            # Sanity check the DIR were are going to remove is inside the GITWS_ROOT_DIR
            if [ ! -z "$(echo "${DIR}" | grep "^${GITWS_ROOT_DIR}" )" ]; then
                rm -fd ${DIR}
            else
                printf "\e[7mWarning:\e[0m There may be some empty directories leftover.\n\n"
                return 1
            fi
        fi
        DIR="${DIR%\/*}"
    done
}

#==================================================================================================
#   ######   #### ######## ##      ##  ######     ##       ####  ######  ########
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##        ##  ##    ##    ##
#  ##         ##     ##    ##  ##  ## ##          ##        ##  ##          ##
#  ##   ####  ##     ##    ##  ##  ##  ######     ##        ##   ######     ##
#  ##    ##   ##     ##    ##  ##  ##       ##    ##        ##        ##    ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##        ##  ##    ##    ##
#   ######   ####    ##     ###  ###   ######     ######## ####  ######     ##
#==================================================================================================
# gitws list
#
#   This function will remove a branch and perform a prune
function __gitws_list_help {
    printf "Usage: gitws list\n"
}

function __gitws_list {
    # Verify that we are inside a gitws directory
    __GITWS_ROOT_DIR=$(__gitws_root)
    if [ -z ${__GITWS_ROOT_DIR} ]; then
        printf "\e[7mError:\e[0m Not inside a gitws workspace.\n\n"
        return 1
    fi

    # Setup variables from gitws workspace
    source ${__GITWS_ROOT_DIR}/.gitws

    if [ $# -gt 0 ]; then
        printf "\e[7mError:\e[0m Too many arguments specified.\n\n"
        __gitws_list_help
        return 1
    fi

    git -C ${GITWS_GIT_DIR} worktree list
}

#==================================================================================================
#   ######   #### ######## ##      ##  ######     ##     ## ######## ##    ## ##     ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ###   ### ##       ###   ## ##     ##
#  ##         ##     ##    ##  ##  ## ##          #### #### ##       ####  ## ##     ##
#  ##   ####  ##     ##    ##  ##  ##  ######     ## ### ## ######   ## ## ## ##     ##
#  ##    ##   ##     ##    ##  ##  ##       ##    ##     ## ##       ##  #### ##     ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##    ##     ## ##       ##   ### ##     ##
#   ######   ####    ##     ###  ###   ######     ##     ## ######## ##    ##  #######
#==================================================================================================
# Common functions
function __gitws_select_branch {
    if [ ! -z "$(command -v fzy)" ]; then
        # FZY command is supported
        SELECTED_LINE=$(echo "$BRANCH_LIST" | fzy)
    else
        printf "\e[7mError:\e[0m Please install FZY. [sudo apt install fzy]\n\n"
        return 1
    fi
}

function __gitws_menu {
    # Verify that we are inside a gitws directory
    __GITWS_ROOT_DIR=$(__gitws_root)
    if [ -z ${__GITWS_ROOT_DIR} ]; then
        printf "\e[7mError:\e[0m Not inside a gitws workspace.\n\n"
        return 1
    fi

    # Setup variables from gitws workspace
    source ${__GITWS_ROOT_DIR}/.gitws

    # Get branch lists available to gitws
    BRANCH_LIST=$(__gitws_list)
    __gitws_select_branch
    [ $? -ne 0 ] && return 1
    
    SELECTED_BRANCH=$(echo "$SELECTED_LINE" | awk '{gsub(/\[|\]/, "", $3); print $3}')
    cd ${GITWS_ROOT_DIR}/${SELECTED_BRANCH}/${GITWS_ROOT_PREFIX}
}


#==================================================================================================
#   ######   #### ######## ##      ##  ######
#  ##    ##   ##     ##    ##  ##  ## ##    ##
#  ##         ##     ##    ##  ##  ## ##
#  ##   ####  ##     ##    ##  ##  ##  ######
#  ##    ##   ##     ##    ##  ##  ##       ##
#  ##    ##   ##     ##    ##  ##  ## ##    ##
#   ######   ####    ##     ###  ###   ######
#==================================================================================================
# gitws
#
#   The main function of gitws
function gitws {
    # Go into interactive mode if no subcommand is chosen
    if [ $# -eq 0 ]; then
        __gitws_menu
        return 0
    fi

    # Restrict valid subcommands
    valid_subcommands=(
        help
        clone
        add
        rm
        list
        version
    )
    local cmdname=$1; shift
    if ! [[ ${valid_subcommands[*]} =~ (^|[[:space:]])"$cmdname"($|[[:space:]]) ]]; then
        printf "\e[7mError:\e[0m Invalid sub command.\n\n"
        __gitws_help
        return 1
    fi
    if type "__gitws_$cmdname" >/dev/null 2>&1; then
        "__gitws_$cmdname" "$@"
    fi
}
