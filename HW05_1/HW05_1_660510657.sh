#!/bin/bash
# Tharachon Warintaweewat (Tonnam)
# 660510657
# HW05_1
# 204203 sec 001

transform_name() {
    first=$(echo $1 | cut -d " " -f 2)
    last=$(echo $1 | cut -d " " -f 1)
    first=${first^}
    last=${last^}
    first_len=9
    last_len=5
    # echo $first.$last
    if [[ ${#last} -lt 5 ]]; then
        if [[ ${#first} -gt 9 ]]; then
            first_len=$((first_len+5-${#last}))
            # echo "[PASS] "
        fi
    else
        if [[ ${#first} -lt 9 ]]; then
            last_len=$((last_len+9-${#first}))
            # echo "[NOT PASS] "
        fi
    fi

    first=${first::first_len}
    last=${last::last_len}
    echo $first.$last
}

# Main script
main() {
    result=$(transform_name "$1")
    echo $result
}

# Execute main function if this script is run directly
if [ "$0" = "${BASH_SOURCE[0]}" ]; then
    while read -r line; do
        # echo $line
        line="${line%$'\n'}"
        [[ -n $line ]] && main "$line"
    done </dev/stdin
fi
