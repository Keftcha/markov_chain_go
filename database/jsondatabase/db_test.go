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

func equalMap(m1, m2 map[string][]string) bool {
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
	data := make(map[string][]string)
	got, err := add(data, [2]string{"a", "b"}, "c")

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := map[string][]string{
		"[\"a\",\"b\"]": []string{"c"},
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
	data := make(map[string][]string)
	data, _ = add(data, [2]string{"a", "b"}, "c")
	got, err := add(data, [2]string{"d", "e"}, "f")

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := map[string][]string{
		"[\"a\",\"b\"]": []string{"c"},
		"[\"d\",\"e\"]": []string{"f"},
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
	data := make(map[string][]string)
	data, _ = add(data, [2]string{"a", "b"}, "c")
	got, err := add(data, [2]string{"a", "b"}, "f")

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := map[string][]string{
		"[\"a\",\"b\"]": []string{"c", "f"},
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
	data := make(map[string][]string)
	data, _ = add(data, [2]string{"1", "2"}, "¤")
	got, err := add(data, [2]string{"1", "2"}, "¤")

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := map[string][]string{
		"[\"1\",\"2\"]": []string{"¤"},
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
	got, err := random([]string{"¤", "§"})

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	if got != "¤" && got != "§" {
		t.Errorf(
			"The value got is %q, instad of \"¤\" or \"§\"",
			got,
		)
	}
}

func TestGetRandomEntryFromEmptySubset(t *testing.T) {
	got, err := random(make([]string, 0))

	if err.Error() != "The key haven't any words in his subset" || got != "" {
		t.Errorf(
			"Error got is %q, instead of %q",
			err.Error(),
			"The key haven't any words in his subset",
		)
	}
}

func TestGettingWhenThereIsNoData(t *testing.T) {
	data := make(map[string][]string)

	got, err := get(data, [2]string{"", ""})

	if err.Error() != "Key not found" || got != nil {
		t.Errorf(
			"Error got is %q, instead of %q",
			err.Error(),
			"Key not found",
		)
	}
}

func TestGettingWhenTheKeyIsntPresent(t *testing.T) {
	data := map[string][]string{
		"[\"a\",\"b\"]": []string{"c"},
	}

	got, err := get(data, [2]string{"", ""})

	if err.Error() != "Key not found" || got != nil {
		t.Errorf(
			"Error got is %q, instead of %q",
			err.Error(),
			"Key not found",
		)
	}
}

func TestGetting(t *testing.T) {
	data := map[string][]string{
		"[\"a\",\"b\"]": []string{"c"},
	}

	expected := []string{"c"}

	got, err := get(data, [2]string{"a", "b"})

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	if !equalSlices(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestSettingWhenTheDataIsEmpty(t *testing.T) {
	data := make(map[string][]string)
	got, err := set(data, [2]string{"a", "b"}, []string{"c", "d", "e"})

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := map[string][]string{
		"[\"a\",\"b\"]": []string{"c", "d", "e"},
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
	data := map[string][]string{
		"[\"a\",\"b\"]": []string{"c", "d", "e"},
	}

	got, err := set(data, [2]string{"a", "b"}, []string{"f", "g", "h"})

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := map[string][]string{
		"[\"a\",\"b\"]": []string{"f", "g", "h"},
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
	data := map[string][]string{
		"[\"a\",\"b\"]": []string{"c", "d", "e"},
	}

	got, err := set(data, [2]string{"f", "g"}, []string{"h", "i", "j"})

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := map[string][]string{
		"[\"a\",\"b\"]": []string{"c", "d", "e"},
		"[\"f\",\"g\"]": []string{"h", "i", "j"},
	}

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}
