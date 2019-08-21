#!/bin/bash
PATH_TO_EXAMPLES=examples
for e in $(ls $PATH_TO_EXAMPLES)
do
    echo "Example file: $e"
    echo "REAME.md:"
    cat $PATH_TO_EXAMPLES/$e/README.md
    sh $PATH_TO_EXAMPLES/$e/run.sh
    echo "=====\n\n"
done