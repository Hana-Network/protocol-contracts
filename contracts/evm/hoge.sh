#!/bin/bash

for file in *Zeta*.sol; do
    if [[ -f "$file" ]]; then
        new_file="${file/Zeta/Hana}"
        mv "$file" "$new_file"
        echo "Renamed $file to $new_file"
    fi
done
