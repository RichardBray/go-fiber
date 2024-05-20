#!/bin/bash

git fetch

git diff HEAD..FETCH_HEAD

read -p "Do you want to merge? (Y/n) " answer

if [[ -z "$answer" ]]; then
    answer="y"
fi

case ${answer:0:1} in
    y|Y )
        git merge FETCH_HEAD
        echo "Thank you for merging."
    ;;
    * )
        echo "Exiting without merging."
    ;;
esac

echo "You're such a junior!!"