package main

import "io"

type limitedReader struct {
	r io.Reader
	n int64 // bytes remaining
}

func (l *limitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}

	buf := make([]byte, len(p))

	n, err = l.r.Read(buf)
	if err == io.EOF {
		return
	}
	if err != nil {
		return
	}

	if l.n-int64(n) >= 0 {
		copy(p, buf[:n])
		l.n -= int64(n)
		return n, nil
	} else {
		end := int64(n) - l.n - 1
		copy(p, buf[:end])
		l.n -= end
		return int(end), io.EOF
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{
		r: r,
		n: n,
	}
}

// ioRead realize limit  Reader method as in standard library.
func (l *limitedReader) ioRead(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.n {
		p = p[:l.n]
	}

	n, err = l.r.Read(p)
	l.n -= int64(n)
	return
}
