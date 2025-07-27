ARG BASEIMAGE=required
FROM ${BASEIMAGE}
#
ARG SYSTEM=linux
ARG ARCHITECTURE=amd64
ARG DEBIAN_FRONTEND=noninteractive
ARG HOST_USER_NAME=devuser \
    HOST_USER_UID=1234 \
    HOST_USER_GID=1234 \
    HOST_DOCKER_GID=999
ENV HOST_USER_NAME=${HOST_USER_NAME} \
    HOST_USER_UID=${HOST_USER_UID} \
    HOST_USER_GID=${HOST_USER_GID} \
    HOST_DOCKER_GID=${HOST_DOCKER_GID}
#
RUN rm -f /etc/apt/apt.conf.d/docker-clean; echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache
#
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    apt-get update && apt-get install -y --no-install-recommends \
        git \
        wget \
        curl \
        tree \
        nano \
        sudo \
        gnupg2 \
        locales \
        pkg-config \
        libssl-dev \
        bsdmainutils \
        ca-certificates \
        build-essential \
        python3 \
        python3-pip \
        tmux \
        jq \
    && echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen \
    && locale-gen \
    && groupadd --gid $HOST_USER_GID $HOST_USER_NAME \
    && useradd -s /bin/bash --uid $HOST_USER_UID --gid $HOST_USER_NAME -m $HOST_USER_NAME \
    && echo $HOST_USER_NAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$HOST_USER_NAME \
    && chmod 0440 /etc/sudoers.d/$HOST_USER_NAME \
    && groupadd -g $HOST_DOCKER_GID docker -f \
    && usermod -aG docker $HOST_USER_NAME

ARG DOCKER_BASE_URL="https://download.docker.com/linux/ubuntu/dists/jammy/pool/stable/amd64"
ARG HOST_DOCKER_VERSION="28.3.2"
ARG HOST_DOCKER_BUILDX_VERSION="0.25.0"
ARG HOST_DOCKER_COMPOSE_VERSION="2.38.2"
RUN curl -o docker-ce-cli.deb ${DOCKER_BASE_URL}/docker-ce-cli_${HOST_DOCKER_VERSION}-1~ubuntu.22.04~jammy_amd64.deb && \
    curl -o docker-buildx-plugin.deb ${DOCKER_BASE_URL}/docker-buildx-plugin_${HOST_DOCKER_BUILDX_VERSION}-1~ubuntu.22.04~jammy_amd64.deb && \
    curl -o docker-compose-plugin.deb  ${DOCKER_BASE_URL}/docker-compose-plugin_${HOST_DOCKER_COMPOSE_VERSION}-1~ubuntu.22.04~jammy_amd64.deb&& \
    dpkg -i ./docker-ce-cli.deb ./docker-buildx-plugin.deb ./docker-compose-plugin.deb && \
    rm -fr docker-*.deb

# Install github cli
ARG GH_VERSION="2.76.1"
RUN curl -sSfL https://github.com/cli/cli/releases/download/v${GH_VERSION}/gh_${GH_VERSION}_linux_amd64.tar.gz | tar -xz --strip-components=1 -C /usr

# Install Golang and tools
ARG GOLANG_VERSION="1.24.5"
RUN curl -sSfL https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz
ENV GOPATH=/home/${HOST_USER_NAME}/go
ENV PATH=$PATH:/usr/local/go/bin:${GOPATH}/bin

# Install some precommits framework
RUN python3 -m pip install pre-commit==4.2.0

USER ${HOST_USER_NAME}
WORKDIR /home/$HOST_USER_NAME
# Make sure USER is defined
ENV USER ${HOST_USER_NAME}

# Install Azure CLI
RUN curl -sSfL https://aka.ms/InstallAzureCLIDeb | sudo bash

# Install Rust
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
# Install cargo make
ENV PATH="/home/${HOST_USER_NAME}/.cargo/bin:$PATH"
RUN cargo install cargo-make

# Install UV
ARG UV_VERSION="0.8.3"
RUN curl --proto '=https' --tlsv1.2 -LsSf https://github.com/astral-sh/uv/releases/download/${UV_VERSION}/uv-installer.sh | sh

# Install ripgrep
ARG RIPGREP_VERSION="14.1.1"
RUN curl -sSfL https://github.com/BurntSushi/ripgrep/releases/download/${RIPGREP_VERSION}/ripgrep-${RIPGREP_VERSION}-x86_64-unknown-linux-musl.tar.gz | tar -zxvf - --strip-components=1 -C /home/$HOST_USER_NAME/.local/bin --wildcards "*/rg"

# Install neovim
ARG NEOVIM_VERSION="0.11.3"
RUN curl -sSfL https://github.com/neovim/neovim/releases/download/v${NEOVIM_VERSION}/nvim-linux-x86_64.tar.gz | tar zxvf - --strip-components=1 -C /home/$HOST_USER_NAME/.local

# Install starhip
RUN curl -sSfL https://starship.rs/install.sh | sh -s -- --yes

# Install autoenv
RUN git clone 'https://github.com/hyperupcall/autoenv' ~/.autoenv

COPY config/.bashrc /home/$HOST_USER_NAME/.bashrc
COPY config/.tmux.conf /home/$HOST_USER_NAME/.tmux.conf
COPY config/nvim /home/$HOST_USER_NAME/.config/nvim
COPY config/starship.toml /home/$HOST_USER_NAME/.config/starship.toml
