#ifdef _GLFW_WIN32

	// Bug fix for 6l/8l on Windows, see:
	//  https://github.com/go-gl/glfw3/issues/82#issuecomment-53859967
	#include <stdlib.h>
	#include <string.h>

	#ifndef strdup
	char *strdup (const char *str) {
		char *new = malloc(strlen(str));
		strcpy(new, str);
		return new;
	}
	#endif
	
	// http://msdn.microsoft.com/en-us/library/571yb472.aspx
	#include <stdio.h>
	#ifndef _get_output_format
	unsigned int _get_output_format(void) {
		return 0;
	};
	#endif

	#include "glfw/src/win32_init.c"
#endif

