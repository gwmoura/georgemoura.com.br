#!/bin/bash
# source scripts/install_go_appengine.sh

# install Go component
./google-cloud-sdk/bin/gcloud components install app-engine-go -q
./google-cloud-sdk/platform/google_appengine/appcfg.py --oauth2_refresh_token=$GAE_OAUTH -v --version=`date '+%Y%m%dt%H%M%S'` --application="george-moura-site" update ./app.yaml
