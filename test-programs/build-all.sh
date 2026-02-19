#!/usr/bin/env bash

# build-all.sh - compile nyassembler and run it on all .nyan programs in this directory
# Usage: simply execute from the test-programs directory (it builds all .nyan files)

set -euo pipefail

# locate the nyassembler directory relative to this script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
NYASM_DIR="${SCRIPT_DIR}/../nyassembler"

# build the assembler
echo "Building nyassembler in ${NYASM_DIR}..."
cd "${NYASM_DIR}"
go build .

# path to assembler executable
NYASM_EXEC="${NYASM_DIR}/nyassembler"

# go back to test-programs directory
cd "${SCRIPT_DIR}"

echo "Processing .nyan files in ${SCRIPT_DIR}..."

shopt -s nullglob
for src in *.nyan; do
    # skip the script itself if somehow named .nyan
    if [[ "$src" == "build-all.sh" ]]; then
        continue
    fi
    out="${src%.nyan}.nyobj"
    echo "  assembling $src -> $out"
    "${NYASM_EXEC}" --input "$src" --output "$out"
done

echo "Done."
