# Find GLESv1
#
# GLESv1_INCLUDE_DIR
# GLESv1_LIBRARY
# GLESv1_FOUND

find_path(GLESv1_INCLUDE_DIR NAMES GLES/gl.h)

set(GLESv1_NAMES ${GLESv1_NAMES} GLESv1_CM)
find_library(GLESv1_LIBRARY NAMES ${GLESv1_NAMES})

include(FindPackageHandleStandardArgs)
find_package_handle_standard_args(GLESv1 DEFAULT_MSG GLESv1_LIBRARY GLESv1_INCLUDE_DIR)

mark_as_advanced(GLESv1_INCLUDE_DIR GLESv1_LIBRARY)

