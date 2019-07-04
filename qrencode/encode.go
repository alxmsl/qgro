// Package qrencode provides functions to generate QR codes.
package qrencode

// QR encode the content at the specified error correction level (ecLevel).
func Encode(content string, ecLevel ECLevel) (*Grid, error) {
	bits, version, err := stringContentBits(content, ecLevel)
	if err != nil {
		return nil, err
	}
	return encode(bits, version, ecLevel)
}

// QR encode the content at the specified error correction level (ecLevel).
func EncodeBytes(content []byte, ecLevel ECLevel) (*Grid, error) {
	bits, version, err := binaryContentBits(content, ecLevel)
	if err != nil {
		return nil, err
	}
	return encode(bits, version, ecLevel)
}

func encode(bits *vector, version versionNumber, ecLevel ECLevel) (*Grid, error) {
	bits = interleaveWithECBytes(bits, version, ecLevel)
	grid := NewGrid(version.dimension(), version.dimension())
	maskPattern := bestMaskPattern(bits, version, ecLevel, grid)
	buildGrid(bits, version, ecLevel, maskPattern, grid)
	return grid, nil
}
