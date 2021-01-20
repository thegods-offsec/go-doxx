#!/bin/bash
if ! [ $(id -u) = 0 ]; then
   echo "This script must be run as root" 
   exit 1
fi


rm -rf go-doxx;

git clone https://github.com/TheG0ds/go-doxx.git;

mv go-doxx/doxx /usr/sbin;

rm -rf go-doxx;

echo 'doxx installed in /usr/sbin'