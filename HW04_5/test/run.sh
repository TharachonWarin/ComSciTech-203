#!/usr/bin/env bash

valid=( $(cat output.txt) )
result=( $(cat out.txt) )


for i in ${#valid}
do
	diff <(echo ${valid[$i]}) <(echo ${result[$i]}) || true
	echo $valid[$i] $result[$i]
done
