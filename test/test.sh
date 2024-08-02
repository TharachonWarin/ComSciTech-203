#!/bin/bash
declare -A capitals
capitals["Thailand"]="Bangkok"
capitals["Japan"]="Tokyo"
capitals["Germany"]="Berlin"
capitals["France"]="Paris"

echo ${capitals[@]}

# Loop over the keys of the associative array
for country in "${!capitals[@]}"; do
capital="${capitals[$country]}"
echo "The capital of $country is $capital"
done | sort
