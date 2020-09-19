#!/bin/bash

set -eou pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Build go code
cd ${SCRIPT_DIR}/../
go build -o azure-playground-generator
chmod +x azure-playground-generator

echo "Build complete!"