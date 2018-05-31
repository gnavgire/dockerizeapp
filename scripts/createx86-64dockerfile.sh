#!/bin/bash

if [ -f $1 ]; then
	binname=$1
	#echo "U gave ${binname}"
	cp -f ${binname} workdir/testapp
	cp -f template/Dockerfile_x86-64  workdir/Dockerfile.testapp
	sed -i "s/COPY bin/COPY testapp/" workdir/Dockerfile.testapp
else
	echo $1 " No such file exists "
fi
