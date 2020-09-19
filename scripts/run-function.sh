#!/bin/bash

set -eou pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $SCRIPT_DIR/../

# First ensure you have a proper .env file (in the same format as .env.sample) with the deployment values filled out
# Load settings from .env file
if [[ -f ".env" ]]; then
    export $(grep -v '^#' .env | xargs)
fi

func host start