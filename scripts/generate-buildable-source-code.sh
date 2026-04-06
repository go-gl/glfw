#!/usr/bin/env bash

# Generate vendored GLFW C sources for a specific go-gl target directory.
# Example:
#   scripts/generate-buildable-source-code.sh v3.4 "$(cat v3.4/glfw/GLFW_C_REVISION.txt)"

set -euo pipefail

EXEC="$0"

usage() {
    echo "usage: $EXEC <go_gl_target_dir> <glfw_revision>"
    exit "$1"
}

TARGET_DIR="${1:-}"
GLFW_REVISION="${2:-}"
if [ -z "$TARGET_DIR" ] || [ -z "$GLFW_REVISION" ]; then
    usage 1
fi

TMP_DIR="$(mktemp -d)"
cleanup() {
    rm -rf "$TMP_DIR"
}
trap cleanup EXIT

generate_dummy_go_files() {
    local glfw_root="$1"
    local deps_import_root="$2"

    find "$glfw_root/deps" "$glfw_root/include" "$glfw_root/src" -type d -print0 |
        while IFS= read -r -d $'\0' d; do
            cat > "$d"/dummy.go <<'EOF'
//go:build required
// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy
EOF
        done

    {
        cat <<'EOF'
//go:build required
// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy
EOF
        echo
        echo "import ("
        find "$glfw_root/deps" -mindepth 1 -maxdepth 1 -type d | sort |
            while IFS= read -r dep_dir; do
                dep_name="$(basename "$dep_dir")"
                echo "	_ \"$deps_import_root/$dep_name\""
            done
        echo ")"
    } > "$glfw_root/deps/dummy.go"
}

generate_wayland_protocol_headers() {
    local upstream_root="$1"
    local include_dir="$2"
    local scanner

    scanner="$(command -v wayland-scanner || true)"
    if [ -z "$scanner" ]; then
        echo "$EXEC: wayland-scanner is required but was not found in PATH" >&2
        exit 1
    fi

    local protocol_files=(
        wayland.xml
        viewporter.xml
        xdg-shell.xml
        idle-inhibit-unstable-v1.xml
        pointer-constraints-unstable-v1.xml
        relative-pointer-unstable-v1.xml
        fractional-scale-v1.xml
        xdg-activation-v1.xml
        xdg-decoration-unstable-v1.xml
    )

    for protocol in "${protocol_files[@]}"; do
        local protocol_path="$upstream_root/deps/wayland/$protocol"
        local protocol_base="${protocol%.xml}"

        "$scanner" client-header "$protocol_path" \
            "$include_dir/${protocol_base}-client-protocol.h"
        "$scanner" private-code "$protocol_path" \
            "$include_dir/${protocol_base}-client-protocol-code.h"
    done
}

WORK_DIR="$TMP_DIR/work"
UPSTREAM_SRC="$WORK_DIR/glfw-src"
AGGREGATE_DIR="$WORK_DIR/glfw-aggregate"

mkdir -p "$UPSTREAM_SRC"
curl -fsSL "https://github.com/glfw/glfw/archive/${GLFW_REVISION}.tar.gz" |
    tar xz --strip-components=1 --directory="$UPSTREAM_SRC"
mkdir -p "$AGGREGATE_DIR/include"

cp -r "$UPSTREAM_SRC/src" "$AGGREGATE_DIR/src"
cp -r "$UPSTREAM_SRC/include/"* "$AGGREGATE_DIR/include"
cp -r "$UPSTREAM_SRC/deps" "$AGGREGATE_DIR/deps"
cp "$UPSTREAM_SRC/LICENSE.md" "$AGGREGATE_DIR/LICENSE.md"
generate_wayland_protocol_headers "$UPSTREAM_SRC" "$AGGREGATE_DIR/include"

# Keep parity with files historically excluded in go-gl vendoring scripts.
rm -f "$AGGREGATE_DIR"/src/CMakeLists.txt "$AGGREGATE_DIR"/src/*.in

GLFW_DIR="$TARGET_DIR/glfw/glfw"
rm -rf "$GLFW_DIR"
mv "$AGGREGATE_DIR" "$GLFW_DIR"

generate_dummy_go_files "$GLFW_DIR" "github.com/go-gl/glfw/$TARGET_DIR/glfw/glfw/deps"
