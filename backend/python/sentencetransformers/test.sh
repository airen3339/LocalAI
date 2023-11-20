#!/bin/bash
##
## A bash script wrapper that runs the sentencetransformers server with conda

# Activate conda environment
source activate sentencetransformers

# get the directory where the bash script is located
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

python -m unittest $DIR/test_sentencetransformers.py