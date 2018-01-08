#!/bin/bash

DIRECTORY=/opt/rpgapp
EXEC=/opt/rpgapp/rpgApp
LOCALSERVICECONFIG=/lib/systemd/system/rpgapp.service
SERVICECONFIG=data/systemd/rpgapp.service
SERVER=rpgapp

if (( $(ps -ef | grep -v grep | grep $SERVER | wc -l) > 0)); then
  echo "$SERVER server is running. Stopping it now."
  sudo systemctl stop rpgapp
fi

echo "Building RPG Toolkit Server."
go build
echo "Installing RPG Toolkit Server."

if [ ! -d "$DIRECTORY" ]; then
  sudo mkdir $DIRECTORY
fi

if [ ! -f "$EXEC" ]; then
  sudo mv rpgApp $DIRECTORY
else
  sudo rm -f $EXEC
  sudo cp rpgApp $DIRECTORY
fi

if [ ! -f "$LOCALSERVICECONFIG" ]; then
  sudo rm -f $LOCALSERVICECONFIG
fi

sudo cp $SERVICECONFIG /lib/systemd/system/
echo "Cleaning up Install."
rm rpgApp
echo "Starting up server as a daemon..."
sudo systemctl start rpgapp
sudo journalctl -f -u rpgapp
