What's my IP
============

This utility will track your *public* IP address changing. It will notify you
on Slack when it changes.

Build and run::

    # build
    go build -o whats-my-ip cmd/whatsmyip/main.go

    # run with correct webhook URL for your slack:
    SLACK_WEBHOOK_URL=https://hooks.slack.com/services/.../L... ./whats-my-ip


Settings
--------

SLACK_WEBHOOK_URL
~~~~~~~~~~~~~~~~~

The URL of the Slack webhook to send the notification to.

IFCONFIG_URL
~~~~~~~~~~~~

The URL of the ifconfig.me service to get the public IP address from. This
defaults to `http://ifconfig.me`.

STORAGE_FILE_PATH
~~~~~~~~~~~~~~~~~

The file and path of the IP address tracking DB. This uses SQLite3 for this. By
default, this will create a `storage.db` file in the current directory.


Development
-----------

To run the tests, run the following command:

::

    go test ./...


Linux Install
-------------

I run this as a cron job on a target computer. I configure a linux machine as
follows.

First, build the binary::

    go build -o whats-my-ip cmd/whatsmyip/main.go

Then add a user and group I'll run this utility as::

    # a non interactive user:
    sudo adduser --system tooling
    sudo addgroup tools
    sudo usermod -a -G tools tooling

Now, I manually install it into the /opt/whats-my-ip directory::

    sudo mkdir -p /opt/whats-my-ip/bin
    sudo mkdir -p /opt/whats-my-ip/db
    sudo cp scripts/whats-my-ip.sh /opt/whats-my-ip/bin/
    sudo cp whats-my-ip /opt/whats-my-ip/bin/

Next, I edit the environment variables in the script::

    sudo vi /opt/whats-my-ip/bin/whats-my-ip.sh

I need to give ownership of the directory to the tooling user::

    sudo chown -R tooling:tools /opt/whats-my-ip

Finally, I add a cron job to run the script every 30 minutes::

    sudo crontab -e -u tooling

    */30 * * * * /opt/whats-my-ip/bin/whats-my-ip.sh
