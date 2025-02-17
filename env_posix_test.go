//go:build !windows
// +build !windows

// Copyright 2025 The globalenv Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package globalenv

import (
	"fmt"
	"testing"
)

func TestParseEnv(t *testing.T) {
	fmt.Println(query([]byte(testset), "PATH"))
	fmt.Println(unset([]byte(testset), "PATH"))
}

const testset = `
# ~/.bashrc file - Example with all sorts of configurations

# User specific environment and path settings
export PATH=$PATH:$HOME/bin
export EDITOR=nano
export LANG=en_US.UTF-8

# Set terminal title (for bash prompt)
case $TERM in
  xterm*|rxvt*)
    PROMPT_COMMAND='echo -ne "\033]0;${USER}@${HOSTNAME}: ${PWD}\007"'
    ;;
esac

# Custom prompt (PS1)
export PS1="\[\e[1;32m\]\u@\h\[\e[00m\]:\[\e[1;34m\]\w\[\e[00m\]$ "

# Define some useful aliases
alias ll='ls -l'
alias la='ls -a'
alias ls='ls --color=auto'
alias ..='cd ..'
alias ...='cd ../..'

# Enable color support for ls and grep
export CLICOLOR=1
export LSCOLORS=GxFxCxDxBxegedabagaced

# Set history options (size and file)
export HISTSIZE=1000
export HISTFILESIZE=2000
export HISTCONTROL=ignoredups:erasedups
shopt -s histappend

# Set bash options
shopt -s autocd
shopt -s globstar

# Enable vi mode in the shell
set -o vi

# Load virtualenvwrapper if it's installed
if [ -f "$HOME/.virtualenvs" ]; then
  export WORKON_HOME=$HOME/.virtualenvs
  export VIRTUALENVWRAPPER_PYTHON=$(which python3)
  source /usr/local/bin/virtualenvwrapper.sh
fi

# Custom function to show system uptime
uptime_info() {
  echo "System Uptime: $(uptime -p)"
}

# Custom function to extract .tar, .zip, .gz, and .bz2 files
extract() {
  if [ -f "$1" ]; then
    case "$1" in
      *.tar.gz)  tar -xzvf "$1" ;;
      *.tar.bz2) tar -xjvf "$1" ;;
      *.tar.xz)  tar -xJvf "$1" ;;
      *.zip)     unzip "$1" ;;
      *.tar)     tar -xvf "$1" ;;
      *) echo "Unsupported file type"; return 1 ;;
    esac
  else
    echo "File does not exist!"
    return 1
  fi
}

# Function to clean up Python cache files
cleanup_python_cache() {
  find . -type f -name "*.pyc" -exec rm -f {} \;
}

# Exporting custom paths
export PATH="$PATH:$HOME/scripts"
export PYTHONPATH="$HOME/python_modules"

# Load additional shell scripts (optional)
if [ -f "$HOME/.bashrc_custom" ]; then
  source "$HOME/.bashrc_custom"
fi

# Disable bell sound (optional)
xset -b

# Show some info when starting the terminal
echo "Welcome to your Bash shell, $USER!"
echo "Current directory: $PWD"
echo "System: $(uname -s) $(uname -r)"
echo "Kernel uptime: $(uptime -p)"

export PATH="$PATH:$HOME/bin;$HOME/scripts:$HOME/tools"
export PATH="$PATH:$HOME/scripts"; export EDITOR=nano; export EDITOR=nano
`
