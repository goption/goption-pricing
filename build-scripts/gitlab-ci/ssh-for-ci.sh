#!/usr/bin/env bash
# Copyright (C) 2018 Brian Hazeltine <onwsk8r@gmail.com> https://wasthat.me

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published
# by the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.

# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.
#
# Set up SSH for use in CI
#
# Configure the SSH agent, known_hosts, and a private key for fetching
# dependencies from a host. Note this makes `go get` use SSH as well.
#
# This script will
# - Add known hosts specified in the $SSH_KNOWN_HOSTS environment variable
# - Append several sensible options and maybe one other option to .ssh/default
# - Optionally update the global git config to redirect all HTTPS requests to
#   SSH for a particular host ($GITLAB_HOST). One effect of this is that running
#   `go get` will prepend the import path with `ssh://` instead of `https://`
# - Start the SSH agent and add a private key via the environment variable
#   `$SSH_PRIVATE_KEY`. If no key is specified, the agent will not be started.
#
# This script references a total of seven variables, five of which are defined
# below. The script also checks for the existence of two environment variables:
# `$SSH_KNOWN_HOSTS` and `$SSH_PRIVATE_KEY`. The content of those variables is
# sent to the known_hosts file and SSH agent respectively.
#
# `$SSH_KNOWN_HOSTS` and `$SSH_PRIVATE_KEY` are defined in the goption group
# in Gitlab. The key is the goption-deploy key, which should provide read-only
# access to all goption repositories (though permission must be granted on a
# per-repo basis), and the known_host is the signature of git.dev.vatik.link.
set +eux

# These shouldn't need to change
export SSH_DIR=${SSH_DIR:-$HOME/.ssh}
export KNOWN_HOSTS_FILE=${KNOWN_HOSTS_FILE:-$SSH_DIR/known_hosts}
export SSH_UNSAFE=${SSH_UNSAFE:-no} # Turn off strict host key checking
export SSH_AUTH_SOCK=${SSH_AUTH_SOCK:-$SSH_DIR/socket}
export SSH_INITIALIZED=${SSH_INITIALIZED:-false}

# Use SSH when cloning or go getting (cloning) from GITLAB_DOMAIN
export LOCAL_SSH=${LOCAL_SSH:-yes}
export GITLAB_DOMAIN=${GITLAB_DOMAIN:-$(echo $CI_PROJECT_URL | awk -F[/:] '{print $4}')}

# Whether to use the GOPATH or not
export USE_GOPATH=${USE_GOPATH:-false}
export GOPATH_DIR=/go/src/${CI_PROJECT_URL#*//}

function init_ssh {
    # Create and chmod ~/.ssh
    mkdir -pv $SSH_DIR
    chmod -v 700 $SSH_DIR

    # Create, chmod, and potentially add known hosts
    echo "Checking for SSH_KNOWN_HOSTS..."
    touch $KNOWN_HOSTS_FILE
    if [ -n "$SSH_KNOWN_HOSTS" ]; then
        echo "Adding known hosts:"
        echo "$SSH_KNOWN_HOSTS" | tee -a $KNOWN_HOSTS_FILE
    else
        echo "SSH_KNOWN_HOSTS is not set. Skipping."
    fi
    chmod -v 600 $KNOWN_HOSTS_FILE

    # Add some useful .ssh/config options
    echo "Configuring SSH config..."
    echo -e "AddKeysToAgent yes\nConnectTimeout 3\nVisualHostKey yes\nPreferredAuthentications publickey" | tee -a $SSH_DIR/config

    # Don't set this to yes.
    if [ "$SSH_UNSAFE" == "yes" ]; then
        echo "\033[1;31mSSH_UNSAFE is set. Turning off StrictHostKeyChecking.\033[0m"
        echo "StrictHostKeyChecking no" | tee -a ~/.ssh/config
    fi

    # Tell Git to use SSH in lieu of HTTPS for `go get`ting because we have
    # no way to use HTTP authentication, but we do have a way to use SSH auth
    if [ "$LOCAL_SSH" == "yes" ]; then
        echo "Using SSH to clone repositories from local server"
        git config --global url."git@${GITLAB_DOMAIN}:".insteadOf "https://${GITLAB_DOMAIN}/"
        git config --global -l
    fi
}

function start_agent {
    echo "Starting SSH agent..."
    eval $(ssh-agent -s -a $SSH_AUTH_SOCK)
}

if [ "$SSH_INITIALIZED" != "true" ]; then
    init_ssh;
    SSH_INITIALIZED=true;
fi

if [ -n "$SSH_PRIVATE_KEY" ]; then
    # This is a simplified way of checking for the SSH agent. For better options see
    # https://superuser.com/questions/141044/sharing-the-same-ssh-agent-among-multiple-login-sessions
    # ps -efC ssh-agent | grep $SSH_AUTH_SOCK > /dev/null || start_agent;
    start_agent
    echo "Adding private key..."
    echo "$SSH_PRIVATE_KEY" | ssh-add -
else
    echo "\033[1;31m*** WARNING: NO SSH PRIVATE KEY GIVEN\033[0m"
fi

if [[ "$USE_GOPATH" = "true" && $PWD != $GOPATH_DIR ]]; then
    echo "Preparing to use GOPATH"
    mkdir -pv $(dirname $GOPATH_DIR)
    ln -sfv $PWD $GOPATH_DIR
    cd $GOPATH_DIR # this works if you source this script
fi
