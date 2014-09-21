#!/bin/bash
cd ./glfw/
git pull --ff-only
git rev-parse HEAD > ../GLFW_C_REVISION.txt
