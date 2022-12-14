package encoding

import (
	"bytes"
	"io"
	"strconv"
)

//代码参考于: https://github.com/Dasio/base45

const (
	baseSize         = 45
	chunkSize        = 2
	encodedChunkSize = 3
)

var decodeMap = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 36, 255, 255, 255, 37, 38, 255, 255, 255, 255, 39, 40, 255, 41, 42, 43,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 44, 255, 255, 255, 255, 255, 255, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// Encode encodes src using the encoding enc, writing
// EncodedLen(len(src)) bytes to dst.
//
// The encoding pads the output to a multiple of 3 bytes,
// so Encode is not appropriate for use on individual blocks
// of a large data stream. Use NewEncoder() instead.
func b45encode(dst, src []byte) {
	encode := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"

	if len(src) == 0 {
		return
	}

	di, si := 0, 0
	n := (len(src) / chunkSize) * chunkSize
	for si < n {
		val := uint(src[si+0])<<8 | uint(src[si+1])
		dst[di+0] = encode[val%baseSize]
		dst[di+1] = encode[(val/baseSize)%baseSize]
		dst[di+2] = encode[(val/(baseSize*baseSize))%baseSize]

		si += chunkSize
		di += encodedChunkSize
	}

	if len(src)-si == 0 {
		return
	}

	val := uint(src[si])

	dst[di+0] = encode[val%baseSize]
	dst[di+1] = encode[(val/baseSize)%baseSize]
}

// EncodedLen returns the length in bytes of the base45 encoding
// of an input buffer of length n.
func encodedLen(n int) int {
	res := n / chunkSize * encodedChunkSize
	if n%chunkSize != 0 {
		res += 2
	}
	return res
}

type b45encoder struct {
	err  error
	w    io.Writer
	buf  [chunkSize]byte // buffered data waiting to be encoded
	nbuf int             // number of bytes in buf
	out  [1024]byte      // output buffer
}

func (e *b45encoder) Write(p []byte) (n int, err error) {
	if e.err != nil {
		return 0, e.err
	}

	// Leading fringe.
	if e.nbuf > 0 {
		var i int
		for i = 0; i < len(p) && e.nbuf < chunkSize; i++ {
			e.buf[e.nbuf] = p[i]
			e.nbuf++
		}
		n += i
		p = p[i:]
		if e.nbuf < chunkSize {
			return
		}
		b45encode(e.out[:], e.buf[:])
		if _, e.err = e.w.Write(e.out[:encodedChunkSize]); e.err != nil {
			return n, e.err
		}
		e.nbuf = 0
	}

	// Large interior chunks.
	for len(p) >= chunkSize {
		nn := len(e.out) / encodedChunkSize * chunkSize
		if nn > len(p) {
			nn = len(p)
			nn -= nn % chunkSize
		}
		b45encode(e.out[:], p[:nn])
		if _, e.err = e.w.Write(e.out[0 : nn/chunkSize*encodedChunkSize]); e.err != nil {
			return n, e.err
		}
		n += nn
		p = p[nn:]
	}

	// Trailing fringe.
	for i := 0; i < len(p); i++ {
		e.buf[i] = p[i]
	}
	e.nbuf = len(p)
	n += len(p)
	return
}

// Close flushes any pending output from the encoder.
// It is an error to call Write after calling Close.
func (e *b45encoder) Close() error {
	// If there's anything left in the buffer, flush it out
	if e.err == nil && e.nbuf > 0 {
		b45encode(e.out[:], e.buf[:e.nbuf])
		_, e.err = e.w.Write(e.out[:encodedLen(e.nbuf)])
		e.nbuf = 0
	}
	return e.err
}

// NewEncoder returns a new base45 stream encoder.
// Base45 encodings operate in 3-byte blocks; when finished
// writing, the caller must Close the returned encoder to flush any
// partially written blocks.
func newB45Encoder(w io.Writer) io.WriteCloser {
	return &b45encoder{w: w}
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base45-encoded data
func b45decodedLen(n int) int {
	res := n / encodedChunkSize * chunkSize
	if n%encodedChunkSize != 0 {
		res++
	}
	return res
}

type corruptInputError int64

func (e corruptInputError) Error() string {
	return "illegal base45 data at input byte " + strconv.FormatInt(int64(e), 10)
}

func decodeTriplet(dst, src []byte, si int) (nsi, n int, err error) {
	// Decode triplet using the  alphabet
	var dbuf [3]byte
	dlen := 3

	for j := 0; j < len(dbuf); j++ {
		if len(src) == si {
			switch {
			case j == 0:
				return si, 0, nil
			case j == 1:
				return si, 0, corruptInputError(si - j)
			}
			dlen = j
			break
		}
		in := src[si]
		si++

		out := decodeMap[in]
		if out == 0xFF {
			return si, 0, corruptInputError(si - j)
		}
		dbuf[j] = out
	}

	val := int(dbuf[0]) + baseSize*int(dbuf[1]) + baseSize*baseSize*int(dbuf[2])
	if val > 0xFFFF {
		err = corruptInputError(si)
	}
	switch dlen {
	case 3:
		dst[0] = byte(val / 256)
		dst[1] = byte(val % 256)
	case 2:
		dst[0] = byte(val % 256)
	}

	return si, dlen - 1, err
}

func b45decode(dst, src []byte) (n int, err error) {
	if len(src) == 0 {
		return 0, nil
	}

	si := 0

	for len(src)-si >= encodedChunkSize && len(dst)-n >= encodedChunkSize {
		var ninc int
		si, ninc, err = decodeTriplet(dst[n:], src, si)
		n += ninc
		if err != nil {
			return n, err
		}
	}
	for si < len(src) {
		var ninc int
		si, ninc, err = decodeTriplet(dst[n:], src, si)
		n += ninc
		if err != nil {
			return n, err
		}
	}
	return n, err
}

type b45decoder struct {
	err     error
	readErr error // error from r.Read
	r       io.Reader
	buf     [1024]byte // leftover input
	nbuf    int
	out     []byte // leftover decoded output
	outbuf  [1024 / encodedChunkSize * chunkSize]byte
}

// NewDecoder constructs a new base45 stream decoder.
func newB45Decoder(r io.Reader) io.Reader {
	return &b45decoder{r: r}
}

func (d *b45decoder) Read(p []byte) (n int, err error) {
	// Use leftover decoded output from last read.
	if len(d.out) > 0 {
		n = copy(p, d.out)
		d.out = d.out[n:]
		return n, nil
	}

	if d.err != nil {
		return 0, d.err
	}

	// Refill buffer.
	for d.nbuf < encodedChunkSize && d.readErr == nil {
		nn := len(p) / chunkSize * encodedChunkSize
		if nn < encodedChunkSize {
			nn = encodedChunkSize
		}
		// Try to read one more segment.
		nn += encodedChunkSize
		if nn > len(d.buf) {
			nn = len(d.buf)
		}
		nn, d.readErr = d.r.Read(d.buf[d.nbuf:nn])
		d.nbuf += nn
	}

	if d.nbuf < encodedChunkSize {
		if d.nbuf > 0 {
			// Decode final fragment.
			var nw int
			nw, d.err = b45decode(d.outbuf[:], d.buf[:d.nbuf])
			d.nbuf = 0
			d.out = d.outbuf[:nw]
			n = copy(p, d.out)
			d.out = d.out[n:]
			if n > 0 || len(p) == 0 && len(d.out) > 0 {
				return n, nil
			}
			if d.err != nil {
				return 0, d.err
			}
		}
		d.err = d.readErr
		if d.err == io.EOF && d.nbuf > 0 {
			d.err = io.ErrUnexpectedEOF
		}
		return 0, d.err
	}

	// Decode chunk into p, or d.out and then p if p is too small.
	nr := d.nbuf
	nw := b45decodedLen(d.nbuf)
	if nr%encodedChunkSize == 1 {
		nr--
	}
	if nw > len(p) {
		nw, d.err = b45decode(d.outbuf[:], d.buf[:nr])
		d.out = d.outbuf[:nw]
		n = copy(p, d.out)
		d.out = d.out[n:]
	} else {
		n, d.err = b45decode(p, d.buf[:nr])
	}
	d.nbuf -= nr
	copy(d.buf[:d.nbuf], d.buf[nr:])
	return n, d.err
}

// ByBase45 encodes by base64.
// 通过 base45 编码
func (e Encode) ByBase45() Encode {
	if len(e.src) == 0 {
		return e
	}
	b := &bytes.Buffer{}
	encode := newB45Encoder(b)
	_, _ = encode.Write(e.src)
	_ = encode.Close()
	e.dst = b.Bytes()
	return e
}

// ByBase45 decodes by base45.
// 通过 base45 解码
func (d Decode) ByBase45() Decode {
	if len(d.src) == 0 {
		return d
	}
	decoder := newB45Decoder(bytes.NewReader(d.src))
	dBuf := make([]byte, b45decodedLen(len(d.src)))
	count, err := decoder.Read(dBuf)
	d.dst, d.Err = dBuf[:count], err
	return d
}
