package main

import "io"

type limitedReader struct {
	r        io.Reader
	max, num int64
}

func (l *limitedReader) Read(p []byte) (n int, err error) {
	if l.num >= l.max {
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

	if l.num+int64(n) <= l.max {
		copy(p, buf[:n])
		l.num += int64(n)
		return n, nil
	} else {
		end := l.num + int64(n) - l.max - 1
		copy(p, buf[:end])
		l.num += int64(n)
		return int(end), io.EOF
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{
		r:   r,
		max: n,
	}
}
