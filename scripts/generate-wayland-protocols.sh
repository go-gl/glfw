#!/bin/sh

TMP_CLONE_DIR="/tmp/wayland-protocols"
GLGLFW_PATH="$1"

if [ "$GLGLFW_PATH" = "" ]; then
    echo "no glfw destination path set."
    echo "sample: generate-wayland-protocols.sh ../v3.2/glfw/glfw/src"
    exit 1
fi

git clone https://github.com/wayland-project/wayland-protocols $TMP_CLONE_DIR

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

  rm -f "$GLGLFW_PATH/wayland-$NAME"-client-protocol.{h,c}

  wayland-scanner private-code $TMP_CLONE_DIR/"$GROUP"/"$HEADER"/"$NAME".xml "$GLGLFW_PATH"/wayland-"$NAME"-client-protocol.c
  wayland-scanner client-header $TMP_CLONE_DIR/"$GROUP"/"$HEADER"/"$NAME".xml "$GLGLFW_PATH"/wayland-"$NAME"-client-protocol.h
}

generate "pointer-constraints" "v1"
generate "relative-pointer" "v1"
generate "idle-inhibit" "v1"
generate "xdg-shell" "stable"
generate "xdg-decoration" "v1"
generate "viewporter" "stable"

# Patch for cgo
sed -i "s|types|wl_pc_types|g" "$GLGLFW_PATH"/wayland-pointer-constraints-unstable-v1-client-protocol.c
sed -i "s|types|wl_rp_types|g" "$GLGLFW_PATH"/wayland-relative-pointer-unstable-v1-client-protocol.c
sed -i "s|types|wl_ii_types|g" "$GLGLFW_PATH"/wayland-idle-inhibit-unstable-v1-client-protocol.c
