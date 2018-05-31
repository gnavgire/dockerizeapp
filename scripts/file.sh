#!/bin/bash

if [ -f $1 ]; then
	file $1  |  awk '{print $2" "$3" "$6}' | sed 's/,//'
	#file $1  |  awk '{print $2" "$3" "$6}' 
else
	echo $1 " No such file exists "
fi
