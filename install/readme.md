# ssh-notify install

This simple go program has 4 functions;

- Writing the config file
- Copying the binary notify
- Configure the system pluggable authentication module

# Configuration

the ssh-notify install program is configured through
environment variables, that are prefixed with SSHNOTIFY
for example `SSHNOTIFY_WEBHOOK` configures the slack webhook url
for here on in the document it can be assumed that all environment
variables are prefixed with `SSHNOTIFY_` and all in uppercase

# Install directory
By default the binary and configuration file are installed in the
same directory, but if the DIR environment variable is set to
anything but full path (e.g. not starting with a forward slash)
then other file based configuration options will need to be full pathed
themselves. This would be applicable when you wanted to put the config
file and the binary elsewhere. for example;

```
export SSHNOTIFY_DIR="_"
export SSHNOTIFY_BINARY="/usr/bin/ssh-notify"
export SSHNOTIFY_CONFIG="/etc/ssh-notify/config.yaml"
```

# Functions
## Writing the config file

Creates a yaml file based upon environment variables that match the
prefix and available options, arbitary environment variables are
not passed through.

## Copy the notification binary

If the destination binary filename `BINARY` exists then leave as is
otherwise copy a file from another place in the filesystem `SOURCE`

## Configure the system pluggable authentication module

This ensures that the system login file `PAM_CONFIG` contains a line
that matches the [PAM Configuration File Format](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/managing_smart_cards/pam_configuration_files) the environment parameters

`AUTH_INTERFACE` by default "`session`" such that when added to the
 pam configuration file used by the ssh daemon will evaluated
 when a user session is created (e.g. upon login)
`AUTH_CONTROL` by default "`optional`" so that
 if the notification does not succeed then the login is still allowed,
 set to "required" to prevent login if notification fails
`AUTH_EXEC` by default "`pam_exec.so`"
 the execution plugin that will recieve the parameters of the
 place the destination is copied to and the configuration file written

by default /etc/pam.d/systemd-user will be searched for;
`session optional pam_exec.so /opt/ssh-notify/notify /opt/ssh-notify/config.yaml`
if found then the file will be left as is, otherwise the install program will
append the the line to the file.

# Example implementation

```
export SSHNOTIFY_PAM_CONFIG="/etc/pam/ssh-server"
export SSHNOTIFY_AUTH_INTERFACE="account"
export SSHNOTIFY_AUTH_CONTROL="required"
export SSHNOTIFY_AUTH_EXEC="pam_run.so"
export SSHNOTIFY_=""
export SSHNOTIFY_=""
export SSHNOTIFY_=""
export SSHNOTIFY_SOURCE="/usr/local/bin/ssh-notify"
export SSHNOTIFY_DIR="/var/lib/ssh-notify"
export SSHNOTIFY_BINARY="send"
export SSHNOTIFY_CONFIG="config"
