package utils

// if position (posX, posY) valid in 2d matrix
func IsPosValid(posX, posY, length, width int) bool {
	return posX >= 0 && posY >= 0 && posX < length && posY < width
}
