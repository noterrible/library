#!/bin/sh
echo =================================
echo  自动化部署脚本启动
echo =================================

echo 停止原来运行中的工程
APP_NAME=/home/library/lib

tpid=`ps -ef|grep $APP_NAME|grep -v grep|grep -v kill|awk '{print $2}'`
if [ ${tpid} ]; then
    echo 'Stop Process...'
    kill -15 $tpid
fi
sleep 2
tpid=`ps -ef|grep $APP_NAME|grep -v grep|grep -v kill|awk '{print $2}'`
if [ ${tpid} ]; then
    echo 'Kill Process!'
    kill -9 $tpid
else
    echo 'Stop Success!'
fi

echo 准备从Git仓库拉取最新代码
cd /home/library
echo 开始从Git仓库拉取最新代码
git pull
echo 拉取完毕
echo 开始打包,启动文件为lib
go build -o lib
echo 正在打包···
echo 启动项目
nohup /home/library/lib &> /www/server/other_project/logs/libraryManagementSystem.log 2>&1 &
echo 项目启动完成
