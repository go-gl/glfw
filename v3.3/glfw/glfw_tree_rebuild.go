package glfw

//go:generate ../../scripts/glfw_tree_rebuild.sh

// upstreamTreeSHA is a recursive hash of the full contents of the upstream
// glfw, as generated by git (doesn't need to be committed) when you run `go
// generate` on this package. This exists to invalidate the build cache (see
// https://github.com/go-gl/glfw/issues/269), which is unaffected by C source
// inputs.
//lint:ignore U1000 ^
const upstreamTreeSHA = "a4a3bc8c0f4e37695a4de3f51cd7123bc186043b" 
