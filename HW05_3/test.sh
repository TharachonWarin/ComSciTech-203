#!/bin/bash
# Tharachon Warintaweewat (Tonnam)
# 660510657
# HW05_2
# 204203 sec 001

vote() {
    echo $1
    # declare -A vote

    # for index in $(seq 1 1 $((${#input[@]}-1)))
    # do
    #     count=1
    #     for name in $(echo ${input[index]} | tr ", " "\n") 
    #     do
    #         vote[$name]+="(v${index} $count) "
    #         count=$(($count+1))
    #     done
    # done


    # for result in "${!vote[@]}" 
    # do
    #     rank="${vote[$result]}"
    #     echo $result: $rank
    # done | sort
}

# Main script
main() {
    result=$(vote ${1[@]})
    echo $result
}


# Execute main function if this script is run directly
if [ "$0" = "${BASH_SOURCE[0]}" ]; then
    while read -r line; do
        # echo $line
        line="${line%$'\n'}"
        declare -a input
        input+=("$line")
        [[ -n $line ]] && main ${input[@]}
    done < /dev/stdin
fi