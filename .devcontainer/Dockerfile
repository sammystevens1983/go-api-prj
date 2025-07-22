# 1) Base image & OS‑level deps
FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu-22.04"

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# 1) Create vscode user & give it sudo
RUN groupadd --gid $USER_GID $USERNAME \
 && useradd --uid $USER_UID --gid $USER_GID -m -s /usr/bin/zsh $USERNAME \
 && apt-get update \
 && apt-get install -y sudo \
 && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME

RUN apt-get update \
 && apt-get install -y \
       wget curl git build-essential ca-certificates \
       zsh fonts-powerline \
 && rm -rf /var/lib/apt/lists/*

# # 2) Install Go (layer is only rebuilt if version or URL changes)
# ARG GO_VERSION=1.21.0
# RUN wget -q "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
#  && rm -rf /usr/local/go \
#  && tar -C /usr/local -xzf "go${GO_VERSION}.linux-amd64.tar.gz" \
#  && rm "go${GO_VERSION}.linux-amd64.tar.gz"

# # 3) Oh My Zsh & plugins
# RUN export RUNZSH=no \
#  && sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" \
#  && git clone https://github.com/zsh-users/zsh-autosuggestions   ~/.oh-my-zsh/custom/plugins/zsh-autosuggestions \
#  && git clone https://github.com/zsh-users/zsh-syntax-highlighting ~/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting


# COPY dotfiles/.zshrc /home/$USERNAME/.zshrc
# COPY dotfiles/.zshrc.local /home/$USERNAME/.zshrc.local
# RUN chown -R $USERNAME:$USERNAME /home/$USERNAME

# # 5) Final defaults
# ENV GOROOT=/usr/local/go \
#     GOPATH=/home/$USERNAME/go \
#     PATH="$PATH:/usr/local/go/bin:/home/$USERNAME/go/bin"

USER $USERNAME
WORKDIR /workspace
