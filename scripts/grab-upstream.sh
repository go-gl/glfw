#!/usr/bin/env bash

# This script fetches the glfw sources from upstream

# Files we don't want in the downstream go-gl/glfw repository.
# Users can get these from upstream.
excludes=(
  --exclude=glfw-*/.mailmap
  --exclude=glfw-*/CMake
  --exclude=glfw-*/CMakeLists.txt
  --exclude=glfw-*/README.md
  --exclude=glfw-*/*.in
  --exclude=glfw-*/docs
  --exclude=glfw-*/examples
  --exclude=glfw-*/src/CMakeLists.txt
  --exclude=glfw-*/src/*.in
  --exclude=glfw-*/tests
)

grab-upstream() {
  local godir=$1
  if [[ ! -d $godir ]]; then
    echo "Run with top-level directory as working directory." 1>&2
    exit 1
  fi
  local rev="$(cat $godir/GLFW_C_REVISION.txt)"
  local target=$godir/glfw

  rm -r $target
  mkdir $target

  curl -s -L "https://github.com/glfw/glfw/archive/${rev}.tar.gz" |
    tar xz \
      --strip-components=1 \
      --directory=$target \
      ${excludes[@]}
}

grab-upstream v3.3/glfw

if test -n "$(git status --porcelain)"; then

  git status

  {
    echo
    echo "Tree is dirty after fetching GLFW_C_REVISION from upstream."
    echo
    echo "To fix this:"
    echo "  1. Update GLFW_C_REVISION.txt, then:"
    echo "  2. Run scripts/grab-upstream.sh from the top-level directory"
    echo "  3. Submit the result as a pull request"
    echo
    echo "If you must apply patches, please do it in grab-upstream.sh so that"
    echo "these patches are described programmatically."
    echo
  } 1>&2
  exit 1
fi
