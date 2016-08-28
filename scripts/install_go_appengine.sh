#! /bin/bash
wget "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-$APPENGINE_VERSION.zip" -nv
unzip -q "go_appengine_sdk_linux_amd64-$APPENGINE_VERSION.zip"
export GO_APPENGINE="$(pwd)/go_appengine"
