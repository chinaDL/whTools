package errs

import "fmt"

// InvalidFileError returns an invalid file error
var InvalidFileError = func(file string) error {
	return fmt.Errorf("invalid file %q, please make sure the file exists", file)
}

// DecodeSrcError returns a decode src error
var DecodeSrcError = func(mode string) error {
	return fmt.Errorf("invalid src, the src can't be decoded by %s", mode)
}

// InvalidSrcError returns an invalid src error
var InvalidSrcError = func(size int) error {
	return fmt.Errorf("invalid src size %d, the src must be multiple of 16", size)
}

// OverflowKeyError returns an overflow key error
var OverflowKeyError = func(size int) error {
	return fmt.Errorf("invalid key size %d, the key at least 1 byte and at most 256 bytes", size)
}

// InvalidIVError returns an invalid IV error
var InvalidIVError = func(length int, size int) error {
	return fmt.Errorf("invalid iv size %d, the iv size must be %d", length, size)
}

// InvalidModeOrPaddingError returns an invalid encryption mode or padding error
var InvalidModeOrPaddingError = func(mode, padding string) error {
	return fmt.Errorf("invalid encryption mode %q or padding %q", mode, padding)
}
