#!/bin/bash

DIRECTORY=/opt/rpgapp
EXEC=/opt/rpgapp/rpgApp
LOCALSERVICECONFIG=/lib/systemd/system/rpgapp.service
SERVICECONFIG=data/systemd/rpgapp.service
SERVER=rpgapp
INSTALLDIR=/opt/rpgapp/app
ASSETINSTALLDIR=/opt/rpgapp/app/assets
CONFINSTALLDIR=/opt/rpgapp/app/conf
APPDIR=./ui/app
COMPDIR=./ui/app/components
STATICDIR=./ui/app/static
CSSDIR=./ui/app/assets/css
IMGDIR=./ui/app/assets/img
JSDIR=./ui/app/assets/js
CONFDIR=./data/conf

if (( $(ps -ef | grep -v grep | grep $SERVER | wc -l) > 0)); then
  echo "$SERVER server is running. Stopping it now."
  sudo systemctl stop rpgapp
fi

echo "Building RPG Toolkit Server."
go build
echo "Writing conf file."
echo -e "{\n\t\"Dir\": \"$(pwd)\"\n}" > $CONFDIR/conf.json
echo "Installing RPG Toolkit Server."

if [ ! -d "$DIRECTORY" ]; then
  sudo mkdir $DIRECTORY
fi

if [ ! -f "$EXEC" ]; then
  sudo cp rpgApp $DIRECTORY
else
  sudo rm -f $EXEC
  sudo cp rpgApp $DIRECTORY
fi

if [ ! -d "$INSTALLDIR" ]; then
  sudo mkdir $INSTALLDIR
fi

if [ ! -d "$ASSETINSTALLDIR" ]; then
  sudo mkdir $ASSETINSTALLDIR
fi

sudo cp -ru "$(pwd)/$CSSDIR" $ASSETINSTALLDIR
sudo cp -ru "$(pwd)/$IMGDIR" $ASSETINSTALLDIR
sudo cp -ru "$(pwd)/$JSDIR" $ASSETINSTALLDIR
sudo cp -ru "$(pwd)/$COMPDIR" $INSTALLDIR
sudo cp -ru "$(pwd)/$STATICDIR" $INSTALLDIR
sudo cp -ru "$(pwd)/$CONFDIR" $INSTALLDIR

if [ ! -f "$LOCALSERVICECONFIG" ]; then
  sudo rm -f $LOCALSERVICECONFIG
fi

sudo cp $SERVICECONFIG /lib/systemd/system/
echo "Cleaning up Install."

if [ ! -f "bin/rpgApp" ]; then
  rm -rf bin/rpgApp
fi

mv rpgApp bin/rpgApp
echo "Starting up server as a daemon..."
sudo systemctl start rpgapp
sudo journalctl -f -u rpgapp
