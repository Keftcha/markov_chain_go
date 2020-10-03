package jsondatabase

import (
	"testing"
)

func equalSlices(slc1 []string, slc2 []string) bool {
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

func equalMap(m1, m2 map[[2]string][]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		// Check if the key is in m2 and if value are the same
		if _, ok := m2[k]; !ok || !equalSlices(m1[k], m2[k]) {
			return false
		}
	}
	return true
}

func TestAddingAnEntry(t *testing.T) {
	data := make(map[[2]string][]string)
	got := add(data, [2]string{"a", "b"}, "c")

	expected := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c"},
	}

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestAddingMultipleEntry(t *testing.T) {
	data := make(map[[2]string][]string)
	data = add(data, [2]string{"a", "b"}, "c")
	got := add(data, [2]string{"d", "e"}, "f")

	expected := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c"},
		[2]string{"d", "e"}: []string{"f"},
	}

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestAddingMultipleEntryWithSameKey(t *testing.T) {
	data := make(map[[2]string][]string)
	data = add(data, [2]string{"a", "b"}, "c")
	got := add(data, [2]string{"a", "b"}, "f")

	expected := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c", "f"},
	}

	if !equalMap(expected, got) {
		t.Errorf(
			"expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestAddingAnEntryWithTheElemAlreadyInTheKeySubset(t *testing.T) {
	data := make(map[[2]string][]string)
	data = add(data, [2]string{"1", "2"}, "¤")
	got := add(data, [2]string{"1", "2"}, "¤")

	expected := map[[2]string][]string{
		[2]string{"1", "2"}: []string{"¤"},
	}

	if !equalMap(expected, got) {
		t.Errorf(
			"expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestGetRandomEntryFromSubset(t *testing.T) {
	got := random([]string{"¤", "§"})

	if got != "¤" && got != "§" {
		t.Errorf(
			"The value got is %q, instad of \"¤\" or \"§\"",
			got,
		)
	}
}

func TestGettingWhenThereIsNoData(t *testing.T) {
	data := make(map[[2]string][]string)

	expected := make([]string, 0)

	got := get(data, [2]string{"", ""})

	if !equalSlices(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestGettingWhenTheKeyIsntPresent(t *testing.T) {
	data := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c"},
	}

	expected := make([]string, 0)

	got := get(data, [2]string{"", ""})

	if !equalSlices(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestGetting(t *testing.T) {
	data := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c"},
	}

	expected := []string{"c"}

	got := get(data, [2]string{"a", "b"})

	if !equalSlices(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestSettingWhenTheDataIsEmpty(t *testing.T) {
	data := make(map[[2]string][]string)
	got := set(data, [2]string{"a", "b"}, []string{"c", "d", "e"})

	expected := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c", "d", "e"},
	}

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestSettingOnAnAlreadyExistingKey(t *testing.T) {
	data := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c", "d", "e"},
	}

	got := set(data, [2]string{"a", "b"}, []string{"f", "g", "h"})

	expected := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"f", "g", "h"},
	}

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestSettingOnAnNonExistingKey(t *testing.T) {
	data := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c", "d", "e"},
	}

	got := set(data, [2]string{"f", "g"}, []string{"h", "i", "j"})

	expected := map[[2]string][]string{
		[2]string{"a", "b"}: []string{"c", "d", "e"},
		[2]string{"f", "g"}: []string{"h", "i", "j"},
	}

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}
