#!/bin/bash
# Tharachon Warintaweewat (Tonnam)
# 660510657
# HW05_2
# 204203 sec 001

local_maxima() {
    declare -a arr
    for i in $(echo $1 | tr " " "\n") 
    do
        arr+=($i)
    done

    count=0

    for i in $(seq 1 1 $((${#arr[@]}-2)))
    do
        if [[ ${arr[$i]} -gt ${arr[$(($i-1))]} && ${arr[$i]} -gt ${arr[$(($i+1))]} ]]; then
            count=$(($count+1))
        fi
        # echo ${arr[$i]}
    done 
    # if [[ $(($1))]];then
    
    # fi
    echo $count
}

# Main script
main() {
    result=$(local_maxima "$1")
    echo $result
}


# Execute main function if this script is run directly
if [ "$0" = "${BASH_SOURCE[0]}" ]; then
    while read -r line; do
        # echo $line
        line="${line%$'\n'}"
        [[ -n $line ]] && main "$line"
    done < /dev/stdin
fi