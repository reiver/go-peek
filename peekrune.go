package peek

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

// PeekRune returns the next rune (from the underlying io.RuneScanner) without advancing the reader.
// Although technically, it does temporarily advance the io.RuneScanner, and then unadvances it.
// PeekRune uses io.RuneScanner's ReadRune and UnreadRune to peek.
func PeekRune(runescanner io.RuneScanner) (r rune, size int, err error) {
	if nil == runescanner {
		return 0, 0, errNilRuneScanner
	}

	r, size, err = runescanner.ReadRune()
	if io.EOF == err {
		return r, size, err
	}
	if nil != err {
		return r, size, erorr.Errorf("peek: problem reading rune: %w", err)
	}

	err = runescanner.UnreadRune()
	if nil != err {
		return r, size, erorr.Errorf("peek: problem unreading rune: %w", err)
	}

	return r, size, nil
}
