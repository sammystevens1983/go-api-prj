# ===== ZSH AND OH-MY-ZSH INSTALLATION START =====
echo "<<<<<<<<< Installing and configuring zsh and Oh My Zsh >>>>>>>>>>"

# Install zsh and required dependencies
sudo apt-get install -y zsh curl fonts-powerline

# Get current username
USERNAME=$(whoami)

# Install Oh My Zsh (unattended mode)
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended

# Install useful plugins
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

# Configure plugins in .zshrc
sed -i 's/plugins=(git)/plugins=(git zsh-autosuggestions zsh-syntax-highlighting)/g' ~/.zshrc

# Set the theme to agnoster (more robust pattern match)
sed -i 's/ZSH_THEME=.*/ZSH_THEME="agnoster"/g' ~/.zshrc

# Copy bash environment variables to zsh
cat << 'EOL' >> ~/.zshrc
# Environment variables
export PATH=$HOME/.local/bin:$PATH
export PATH=$PATH:$HOME/.temporalio/bin
export PATH=/usr/local/cuda/bin:$PATH
export LD_LIBRARY_PATH=/usr/local/cuda/lib64:$LD_LIBRARY_PATH
EOL

# Make zsh the default shell
sudo chsh -s /usr/bin/zsh $USERNAME

# Add zsh to /etc/shells if it's not already there
if ! grep -q "^/usr/bin/zsh$" /etc/shells; then
    echo "/usr/bin/zsh" | sudo tee -a /etc/shells
fi

# Create a custom .zshrc.local file with additional settings
cat > ~/.zshrc.local << 'EOL'
# Force the agnoster theme
ZSH_THEME="agnoster"
# Customize prompt segments for agnoster if desired
prompt_context() {
  if [[ "$USER" != "$DEFAULT_USER" || -n "$SSH_CLIENT" ]]; then
    prompt_segment black default "%(!.%{%F{yellow}%}.)$USER"
  fi
}
EOL

# Source the local file from .zshrc
echo "source ~/.zshrc.local" >> ~/.zshrc

# Fix permissions to avoid issues
sudo chown -R $USERNAME:$USERNAME ~

echo "zsh and Oh My Zsh have been successfully installed!"
# ===== ZSH AND OH-MY-ZSH INSTALLATION END =====

source ~/.bashrc
exec zsh -l
echo "Post-create script completed! Please restart your terminal to start using zsh."


# ─── 1) UPDATE & CORE TOOLS ─────────────────────────────
sudo apt-get update && sudo apt-get upgrade -y
sudo apt-get install -y wget curl git build-essential ca-certificates

# ─── 2) GO INSTALL ──────────────────────────────────────
GO_VERSION="1.21.0"
GO_TARBALL="go${GO_VERSION}.linux-amd64.tar.gz"

# download, unpack, then remove the tarball
wget -q "https://go.dev/dl/${GO_TARBALL}"
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf "${GO_TARBALL}"
rm -f "${GO_TARBALL}"