#!/usr/bin/env bash

# Declare a map variable to store the word and count
declare -A map

while read line; do
    for word in $line; do
        ((map["$word"]=${map["$word"]:-0}+1))
    done
done < words.txt

for key in ${!map[@]}; do
    echo "$key ${map[$key]}"
done | sort -rn -k 2
