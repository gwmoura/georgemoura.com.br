from subprocess import call
call(["$HOME/appengine/go_appengine/appcfg.py", "--version=`date '+%Y%m%dt%H%M%S'`","update","."])
