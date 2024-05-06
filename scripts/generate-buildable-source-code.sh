#!/bin/env bash

EXEC=$0;

function usage() {
    echo "usage: $EXEC <glfw_tag_name>";
    exit $1;
}

TAG_NAME=$1;

if [ "x$TAG_NAME" = "x" ]; then
    usage 1;
fi

IFS='.' read -a TAG_PARTS <<< "$TAG_NAME"
TAG_DIR="v${TAG_PARTS[0]}.${TAG_PARTS[1]}"
echo TAG_DIR
echo $TAG_DIR

CMD="mktemp -d";
TMP_DIR=$($CMD);
EXIT_CODE=$?;
if [ $EXIT_CODE != 0 ]; then
    echo "$EXEC: \"$CMD\" failed to execute. exiting..." > /dev/stderr;
    exit 1;
fi

pushd "$TMP_DIR";

git clone --depth 1 --branch "$TAG_NAME" https://github.com/glfw/glfw.git
mkdir glfw/build
cd glfw/build
cmake ..
# TODO don't use this hacky way to generate the wayland protocol headers
make -f src/CMakeFiles/glfw.dir/build.make src/CMakeFiles/glfw.dir/depend # generates required header files in build process
BUILD_DIR=../../glfw-aggregate
mkdir "$BUILD_DIR"
mkdir "$BUILD_DIR/include"
cp src/*.h "$BUILD_DIR/include"

cd ..
BUILD_DIR="./build/$BUILD_DIR"

cp -r src "$BUILD_DIR/src"
cp -r include/* "$BUILD_DIR/include"
cp -r deps "$BUILD_DIR/deps"
cp LICENSE.md "$BUILD_DIR/LICENSE.md"
# TODO generate dummy.go files

popd;
BUILD_DIR="$TMP_DIR/glfw/$BUILD_DIR"

GLFW_DIR="$TAG_DIR/glfw/glfw"
rm -rf "$GLFW_DIR"
mv "$BUILD_DIR" "$GLFW_DIR"

rm -rf "$TMP_DIR";