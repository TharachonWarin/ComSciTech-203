#!/usr/bin/bash
# Techapat Wongsawek
# 660510651
# HW05_3
# 204203 Sec002

array_lines=()
if [ "$0" = "${BASH_SOURCE[0]}" ]; then
    while read -r line; do
        line="${line%$'\n'}"
        echo $line
        array_lines+=("$line")
    done </dev/stdin
fi

# echo $maxLen
declare -A result
echo ${array_lines[@]}

for i in $(seq 1 1 $((${#array_lines[@]} + 1))); do
    # echo ${array_lines[i]}
    poke=$(echo ${array_lines[i]} | tr ", " "\n")
    poke=( $(IFS=' ' echo $poke) )
    # echo ${#poke[@]}
    for j in $(seq 0 1 $((${#poke[@]} - 1))); do
        # echo ${poke[j]}
        result["${poke[j]}"]+="(v$i $(($j+1))) "
    done
done  

# echo ${!result[@]}
for p in ${!result[@]}; do
    echo "$p: ${result[$p]}"
done 
# echo ${result[@]}