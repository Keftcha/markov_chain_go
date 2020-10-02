package markovchaingo

import (
	"testing"
)

func equals(slc1 []string, slc2 []string) bool {
	if len(slc1) != len(slc2) {
		return false
	}

	for idx := 0; idx < len(slc1); idx++ {
		if slc1[idx] != slc2[idx] {
			return false
		}
	}
	return true
}

func TestSplitAnEmptyText(t *testing.T) {
	text := ""
	gotSplitedText := splitMessage(text)

	expectedSplitedText := []string{"", "", "\x03"}

	if !equals(expectedSplitedText, gotSplitedText) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expectedSplitedText,
			gotSplitedText,
		)
	}
}

func TestSplitSomeText(t *testing.T) {
	text := "A Markov chain is a stochastic model describing a sequence of possible events in which the probability of each event depends only on the state attained in the previous event."
	gotSplitedText := splitMessage(text)

	expectedSplitedText := []string{
		"", "", "A", "Markov", "chain", "is", "a", "stochastic",
		"model", "describing", "a", "sequence", "of",
		"possible", "events", "in", "which", "the",
		"probability", "of", "each", "event", "depends",
		"only", "on", "the", "state", "attained", "in", "the",
		"previous", "event.", "\x03",
	}

	if !equals(expectedSplitedText, gotSplitedText) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expectedSplitedText,
			gotSplitedText,
		)
	}
}

func TestSplitTextWithMultipleSpaces(t *testing.T) {
	text := "This  is     a simple   text"
	gotSplitedText := splitMessage(text)

	expectedSplitedText := []string{
		"", "", "This", "", "is", "", "", "", "", "a",
		"simple", "", "", "text", "\x03",
	}

	if !equals(expectedSplitedText, gotSplitedText) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expectedSplitedText,
			gotSplitedText,
		)
	}
}

func TestSplitTextThatHaveSpecialCaracters(t *testing.T) {
	text := "a quote \" and a back quote ` and a slash / and a back-slash \\ and this → ¤ ı 🐺"
	gotSplitedText := splitMessage(text)

	expectedSplitedText := []string{
		"", "", "a", "quote", "\"", "and", "a", "back",
		"quote", "`", "and", "a", "slash", "/", "and",
		"a", "back-slash", "\\", "and", "this", "→",
		"¤", "ı", "🐺", "\x03",
	}

	if !equals(expectedSplitedText, gotSplitedText) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expectedSplitedText,
			gotSplitedText,
		)
	}
}
