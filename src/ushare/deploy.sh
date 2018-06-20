#!/bin/bash

# 项目名称
PROJECT_NAME="ushare"
# 项目分支
BRANCH="master"

pull() {
    git checkout ${BRANCH}
    git pull
    echo "------------------"
    echo "Pull success..."
    echo "------------------"
}

build() {
    result=`go build ${PROJECT_NAME}`
    if [ -z "$result" ]; then
        echo "------------------"
        echo "Build success..."
        echo "------------------"
    fi
}

init() {
    pull
    build
}

start_service() {
    log_path=${PROJECT_NAME}.log
    echo "Log path:$log_path"
#    nohup ./${PROJECT_NAME} 2>&1 >> ${log_path} 2>&1 /dev/null &
    ./${PROJECT_NAME}
}

stop_service() {
    killall ${PROJECT_NAME}
}

start() {
    init
    start_service
    echo "------------------"
	echo "Service started..."
	echo "------------------"
}

stop() {
    stop_service
    echo "------------------"
	echo "Service stopped..."
	echo "------------------"
}

restart() {
    stop_service
	start_service
	echo "------------------"
	echo "Service restarted..."
	echo "------------------"
}

other() {
    echo "$0 {pull|build|start|stop|restart}"
	exit 4
}

case $1 in
    pull)
        pull
    ;;
    build)
        build
    ;;
	start)
		start
	;;
	stop)
		stop
	;;
	restart)
		restart
	;;
	*)
		other
	;;
esac



