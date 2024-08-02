#!/bin/bash
w1=$1
# Echoing numbers 1 to 5
for i in {1..5}
do
    mkdir HW0${w1}_${i}

    # rm -r HW0${w1}_${i}

    cp -r ./HW03_4/test ./HW0${w1}_${i}
done
