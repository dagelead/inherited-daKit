#!/bin/bash
date +"%Y-%m-%d %H:%M:%S" > /workspaces/system/lastexectime.log
sudo supervisord -c /workspaces/daTools/supervisor.conf &

