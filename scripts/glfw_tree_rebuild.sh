#!/bin/sh

# This script computes the hash of the directory called "glfw", relative to $PWD
# at time of invocation, and writes it into the upstreamTreeSHA constant in
# glfw_tree_rebuild.go. Invoke this in a directory containing
# glfw_tree_rebuild.go.

set -e

export GIT_INDEX_FILE=$(mktemp)
rm $GIT_INDEX_FILE
git update-index
git add glfw
HASH=$(git write-tree)
rm $GIT_INDEX_FILE

sed -Ei 's/const upstreamTreeSHA = "[0-9a-f]+"/const upstreamTreeSHA = "'${HASH}'"/' glfw_tree_rebuild.go
