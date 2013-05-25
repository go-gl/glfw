unsigned int GetGammaAtIndex(unsigned short *color, int i) {
	return color[i];
}

void SetGammaAtIndex(unsigned short *color, int i, unsigned short value) {
	color[i] = value;
}