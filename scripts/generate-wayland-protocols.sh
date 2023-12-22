#!/bin/sh

set -e

TMP_CLONE_DIR="$(mktemp -d)"
GLGLFW_PATH="$1"

if [ "$GLGLFW_PATH" = "" ]; then
    echo "no glfw destination path set."
    echo "sample: generate-wayland-protocols.sh ../v3.2/glfw/glfw/src"
    exit 1
fi

git clone https://gitlab.freedesktop.org/wayland/wayland-protocols.git $TMP_CLONE_DIR

generate() {
  HEADER=$1
  VER=$2

  if [ "$VER" = "stable" ]; then
    NAME="$HEADER"
    GROUP="stable"
  else
    NAME="$HEADER"-unstable-$VER
    GROUP="unstable"
  fi

  WAYLAND_CLIENT_C="$GLGLFW_PATH/wayland-$NAME-client-protocol.c"
  WAYLAND_CLIENT_H="$GLGLFW_PATH/wayland-$NAME-client-protocol.h"

  rm -f "$WAYLAND_CLIENT_C" "$WAYLAND_CLIENT_H"

  wayland-scanner private-code $TMP_CLONE_DIR/"$GROUP"/"$HEADER"/"$NAME".xml "$WAYLAND_CLIENT_C"
  wayland-scanner client-header $TMP_CLONE_DIR/"$GROUP"/"$HEADER"/"$NAME".xml "$WAYLAND_CLIENT_H"

  # Go modules don't support symbolic links.
  # This removes the "wayland-xdg-decoration-client-protocol.h"
  # link and moves the original file in place.
  if [ "$HEADER" = "xdg-decoration" ]; then
    mv "$WAYLAND_CLIENT_H" "$GLGLFW_PATH/wayland-xdg-decoration-client-protocol.h"
  fi
}

generate "pointer-constraints" "v1"
generate "relative-pointer" "v1"
generate "idle-inhibit" "v1"
generate "xdg-shell" "stable"
generate "xdg-decoration" "v1"
generate "viewporter" "stable"

rm -rf "${TMP_CLONE_DIR}"
