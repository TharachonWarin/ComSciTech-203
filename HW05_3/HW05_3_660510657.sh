#!/bin/bash
# Tharachon Warintaweewat (Tonnam)
# 660510657
# HW05_3
# 204203 sec 001

declare -a input

if [ "$0" = "${BASH_SOURCE[0]}" ]; then
    while read -r line; do
        line="${line%$'\n'}"
        input+=("$line")
    done < /dev/stdin
fi

echo ${input[@]}

declare -A vote

for index in $(seq 1 1 $((${#input[@]}-1)))
do
    count=1
    for name in $(echo ${input[index]} | tr ", " "\n") 
    do
        vote[$name]+="(v${index} $count) "
        count=$(($count+1))
    done
done


for result in "${!vote[@]}" 
do
    rank="${vote[$result]}"
    echo $result: $rank
done | sort

