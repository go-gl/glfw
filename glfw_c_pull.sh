#!/bin/bash
rm -rf glfw/
git clone https://github.com/glfw/glfw
git add glfw/
cd glfw/
git rev-parse HEAD > ../GLFW_C_REVISION.txt
