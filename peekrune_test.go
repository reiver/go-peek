package peek_test

import (
	"testing"

	"io"
	"strings"

	"sourcecode.social/reiver/go-utf8"

	"github.com/reiver/go-peek"
)

func TestPeekRune_nilRuneScanner(t *testing.T) {

	var runescanner io.RuneScanner = nil

	r, size, err := peek.PeekRune(runescanner)

	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}

	{
		expected := "peek: nil rune-scanner"
		actual   := err.Error()

		if expected != actual {
			t.Errorf("The actual 'error message' is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}

	{
		expected := 0
		actual   := size

		if expected != actual {
			t.Errorf("The actual 'size' is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		expected := rune(0)
		actual   := r

		if expected != actual {
			t.Errorf("The actual 'rune' is not what was expected.")
			t.Logf("EXPECTED: %q (%U)", expected, expected)
			t.Logf("ACTUAL:   %q (%U)", actual, actual)
			return
		}
	}
}

func TestPeekRune_emptyRuneScanner(t *testing.T) {

	var runescanner io.RuneScanner = utf8.NewRuneScanner(strings.NewReader(""))

	r, size, err := peek.PeekRune(runescanner)

	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}

	{
		expected := io.EOF
		actual   := err

		if expected != actual {
			t.Errorf("The actual 'error' is not what was expected.")
			t.Logf("EXPECTED: (%T) %s", expected, expected)
			t.Logf("ACTUAL:   (%T) %s", actual, actual)
			return
		}
	}

	{
		expected := 0
		actual   := size

		if expected != actual {
			t.Errorf("The actual 'size' is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		expected := rune(0)
		actual   := r

		if expected != actual {
			t.Errorf("The actual 'rune' is not what was expected.")
			t.Logf("EXPECTED: %q (%U)", expected, expected)
			t.Logf("ACTUAL:   %q (%U)", actual, actual)
			return
		}
	}
}

func TestPeekRune_invalid(t *testing.T) {

	var invalid string = "\xc3\x28"

	var runescanner io.RuneScanner = utf8.NewRuneScanner(strings.NewReader(invalid))

	r, size, err := peek.PeekRune(runescanner)

	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}

	{
		expected := "peek: problem reading rune: Invalid UTF-8"
		actual   := err.Error()

		if expected != actual {
			t.Errorf("The actual 'error message' is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}

	{
		expected := 0
		actual   := size

		if expected != actual {
			t.Errorf("The actual 'size' is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		expected := rune(0)
		actual   := r

		if expected != actual {
			t.Errorf("The actual 'rune' is not what was expected.")
			t.Logf("EXPECTED: %q (%U)", expected, expected)
			t.Logf("ACTUAL:   %q (%U)", actual, actual)
			return
		}
	}
}

func TestPeekRune_slightlySmiling(t *testing.T) {

	const value rune = '🙂'

	var runescanner io.RuneScanner = utf8.NewRuneScanner(strings.NewReader(string(value)))

	r, size, err := peek.PeekRune(runescanner)

	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}

	{
		expected := utf8.RuneLength(value)
		actual   := size

		if expected != actual {
			t.Errorf("The actual 'size' is not what was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			return
		}
	}

	{
		expected := value
		actual   := r

		if expected != actual {
			t.Errorf("The actual 'rune' is not what was expected.")
			t.Logf("EXPECTED: %q (%U)", expected, expected)
			t.Logf("ACTUAL:   %q (%U)", actual, actual)
			return
		}
	}
}
