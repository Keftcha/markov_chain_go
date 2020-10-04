package inmemorydatabase

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

func TestNewFunctionReturnAnMemoryDatabaseEmpty(t *testing.T) {
	expectedData := make(map[[2]string][]string)
	gotData := New().data

	if !equalMap(expectedData, gotData) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expectedData,
			gotData,
		)
	}
}

func TestAddingAnEntry(t *testing.T) {
	inMemDb := New()
	err := inMemDb.Add([2]string{"a", "b"}, "c")

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := make(map[[2]string][]string)
	expected[[2]string{"a", "b"}] = []string{"c"}

	got := inMemDb.data

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestAddingMultipleEntry(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"a", "b"}, "c")
	err := inMemDb.Add([2]string{"d", "e"}, "f")

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := make(map[[2]string][]string)
	expected[[2]string{"a", "b"}] = []string{"c"}
	expected[[2]string{"d", "e"}] = []string{"f"}

	got := inMemDb.data

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestAddingMultipleEntryWithSameKey(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"a", "b"}, "c")
	err := inMemDb.Add([2]string{"a", "b"}, "f")

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := make(map[[2]string][]string)
	expected[[2]string{"a", "b"}] = []string{"c", "f"}

	got := inMemDb.data

	if !equalMap(expected, got) {
		t.Errorf(
			"expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestAddingAnEntryWithTheElemAlreadyInTheKeySubset(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"1", "2"}, "¤")
	err := inMemDb.Add([2]string{"1", "2"}, "¤")

	if err != nil {
		t.Errorf(
			"Got an error: %q",
			err,
		)
	}

	expected := map[[2]string][]string{
		[2]string{"1", "2"}: []string{"¤"},
	}

	got := inMemDb.data

	if !equalMap(expected, got) {
		t.Errorf(
			"expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestGetRandomEntryFromSubset(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"1", "2"}, "¤")
	inMemDb.Add([2]string{"1", "2"}, "§")

	got, err := inMemDb.Random([2]string{"1", "2"})

	if err != nil {
		t.Errorf(
			"An error occured: %s",
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
	inMemDb := New()
	inMemDb.data = map[[2]string][]string{
		[2]string{"1", "2"}: make([]string, 0),
	}

	got, err := inMemDb.Random([2]string{"1", "2"})

	if err.Error() != "The key haven't any words in his subset" || got != "" {
		t.Errorf(
			"Error got is %q, instead of %q",
			err.Error(),
			"The key haven't any words in his subset",
		)
	}
}

func TestGetRandomEntryFromSubsetOfMissingKey(t *testing.T) {
	inMemDb := New()

	got, err := inMemDb.Random([2]string{"1", "2"})

	if err.Error() != "Key not found" || got != "" {
		t.Errorf(
			"Error got is %q, instead of %q",
			err.Error(),
			"Key not found",
		)
	}
}

func TestGettingWhenThereIsNoData(t *testing.T) {
	inMemDb := New()

	got, err := inMemDb.Get([2]string{"", ""})

	if err.Error() != "Key not found" || got != nil {
		t.Errorf(
			"Error got is %q, instead of %q",
			err.Error(),
			"Key not found",
		)
	}
}

func TestGettingWhenTheKeyIsntPresent(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"a", "b"}, "c")

	got, err := inMemDb.Get([2]string{"", ""})

	if err.Error() != "Key not found" || got != nil {
		t.Errorf(
			"Error got is %q, instead of %q",
			err.Error(),
			"Key not found",
		)
	}
}

func TestGetting(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"a", "b"}, "c")

	expected := []string{"c"}

	got, err := inMemDb.Get([2]string{"a", "b"})

	if err != nil {
		t.Errorf(
			"An error occured: %s",
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
	inMemDb := New()
	err := inMemDb.Set([2]string{"a", "b"}, []string{"c", "d", "e"})

	if err != nil {
		t.Errorf(
			"An error occured: %s",
			err,
		)
	}

	got := inMemDb.data

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
	inMemDb := New()
	inMemDb.Set([2]string{"a", "b"}, []string{"c", "d", "e"})
	err := inMemDb.Set([2]string{"a", "b"}, []string{"f", "g", "h"})

	if err != nil {
		t.Errorf(
			"An error occured: %s",
			err,
		)
	}

	got := inMemDb.data

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
	inMemDb := New()
	inMemDb.Set([2]string{"a", "b"}, []string{"c", "d", "e"})
	err := inMemDb.Set([2]string{"f", "g"}, []string{"h", "i", "j"})

	if err != nil {
		t.Errorf(
			"An error occured: %s",
			err,
		)
	}

	got := inMemDb.data

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
