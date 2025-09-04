# daKit





on: .devcontainer/devcontainer.json

postStartCommand": "bash /workspaces/system/autostart.sh"



sudo apt update
sudo apt install supervisor



/workspaces/system/autostart.sh

#!/bin/bash
date +"%Y-%m-%d %H:%M:%S" > /workspaces/system/lastexectime.log
sudo supervisord -c /workspaces/daTools/supervisor.conf &



supervisor.conf

[supervisord]
logfile=/var/log/supervisord.log
pidfile=/var/log/supervisord.pid
nodaemon=true

[program:daKit]
command=/workspaces/daTools/daKit
autostart=true
autorestart=true
stdout_logfile=/var/log/daKit.out.log
stderr_logfile=/var/log/daKit.err.log
user=root





