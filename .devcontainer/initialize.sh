#!/usr/bin/env bash
WORKSPACE=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/.." &> /dev/null && pwd )

# Make sure some files/directory exists before docker bindmounts them
touch ${HOME}/.netrc
touch ${HOME}/.gitconfig
touch ${HOME}/.Xauthority
# touch ${HOME}/.git-credentials
# touch ${HOME}/.tmux.conf
mkdir -p /tmp/.X11-unix
mkdir -p ~/.azure
mkdir -p ~/.ssh

# Prepare some environment variables
HOST_USER_NAME="$(id -un)"
HOST_USER_UID="$(id -u)"
HOST_USER_GID="$(id -g)"
# Match the docker setup in the host system
HOST_DOCKER_GID="$(stat /var/run/docker.sock --format="%g")"
HOST_DOCKER_VERSION="$(docker version --format '{{.Server.Version}}')"
HOST_DOCKER_BUILDX_VERSION="$(docker buildx version | sed -ne 's|.*v\(\([0-9]*\.\?\)*\).*|\1|p')"
HOST_DOCKER_COMPOSE_VERSION="$(docker compose version | sed -ne 's|.*v\(\([0-9]*\.\?\)*\).*|\1|p')"

if [ ! -f ${WORKSPACE}/.devcontainer/compose.user.yml ]; then
    # User compose file missing create the empty template
    cat <<EOF >${WORKSPACE}/.devcontainer/compose.user.yml
services:
  instance:
    environment: []
    volumes: []
EOF
fi

cat <<EOF >${WORKSPACE}/.devcontainer/compose.gen.yml
services:
  instance:
    user: ${HOST_USER_NAME}
    build:
      args:
        HOST_USER_NAME: ${HOST_USER_NAME}
        HOST_USER_UID: ${HOST_USER_UID}
        HOST_USER_GID: ${HOST_USER_GID}
        HOST_DOCKER_GID: ${HOST_DOCKER_GID}
        HOST_DOCKER_VERSION: ${HOST_DOCKER_VERSION}
        HOST_DOCKER_BUILDX_VERSION: ${HOST_DOCKER_BUILDX_VERSION}
        HOST_DOCKER_COMPOSE_VERSION: ${HOST_DOCKER_COMPOSE_VERSION}
    environment:
      WORKSPACE: ${WORKSPACE}
    volumes:
      # The workspace mount
      - type: bind
        source: ${WORKSPACE}
        target: ${WORKSPACE}
EOF

# If user has a tmux configuration in the HOST system them mount that
# over the default devcontainer tmux config
if [ -f ${HOME}/.tmux.conf ]; then
    cat <<EOF >>${WORKSPACE}/.devcontainer/compose.gen.yml
      # For now dont use system tmux config
      - type: bind
        source: ${HOME}/.tmux.conf
        target: ${HOME}/.tmux.conf
EOF
fi

# If the user has a .git-credentials file in the system then mount that
# into the devcontainer as well
if [ -f ${HOME}/.git-credentials ]; then
    cat <<EOF >>${WORKSPACE}/.devcontainer/compose.gen.yml
      # Mount user git credentials
      - type: bind
        source: ${HOME}/.git-credentials
        target: ${HOME}/.git-credentials
EOF
fi

# If user has a tmux configuration in the HOST system them mount that
# over the default devcontainer tmux config
if [ -d ${HOME}/.config/nvim ]; then
    cat <<EOF >>${WORKSPACE}/.devcontainer/compose.gen.yml
      # For now dont use system tmux config
      - type: bind
        source: ${HOME}/.config/nvim
        target: ${HOME}/.config/nvim
EOF
fi
