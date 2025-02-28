What's my IP
============

This utility will track your *public* IP address changing. It will notify you
on Slack when it changes.

::

    SLACK_WEBHOOK_URL=https://hooks.slack.com/services/.../L... go run cmd/whatismyip/main.go

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
