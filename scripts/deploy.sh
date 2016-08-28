#!/bin/bash
$GO_APPENGINE/appcfg.py -v --version=`date '+%Y%m%dt%H%M%S'` --application="george-moura-site" update ./
