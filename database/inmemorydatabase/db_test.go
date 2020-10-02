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
	inMemDb.Add([2]string{"a", "b"}, "c")

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
	inMemDb.Add([2]string{"d", "e"}, "f")

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
	inMemDb.Add([2]string{"a", "b"}, "f")

	expected := make(map[[2]string][]string)
	expected[[2]string{"a", "b"}] = []string{"c", "f"}

	got := inMemDb.data

	if !equalMap(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestGetRandomEntryFromSubset(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"1", "2"}, "¤")
	inMemDb.Add([2]string{"1", "2"}, "§")

	got := inMemDb.Random([2]string{"1", "2"})

	if got != "¤" && got != "§" {
		t.Errorf(
			"The value got is %q, instad of \"¤\" or \"§\"",
			got,
		)
	}
}

func TestGettingWhenThereIsNoData(t *testing.T) {
	inMemDb := New()

	expected := make([]string, 0)

	got := inMemDb.Get([2]string{"", ""})

	if !equalSlices(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestGettingWhenTheKeyIsntPresent(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"a", "b"}, "c")

	expected := make([]string, 0)

	got := inMemDb.Get([2]string{"", ""})

	if !equalSlices(expected, got) {
		t.Errorf(
			"Expected: %q, but got: %q",
			expected,
			got,
		)
	}
}

func TestGetting(t *testing.T) {
	inMemDb := New()
	inMemDb.Add([2]string{"a", "b"}, "c")

	expected := []string{"c"}

	got := inMemDb.Get([2]string{"a", "b"})

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
	inMemDb.Set([2]string{"a", "b"}, []string{"c", "d", "e"})
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
	inMemDb.Set([2]string{"a", "b"}, []string{"f", "g", "h"})
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
	inMemDb.Set([2]string{"f", "g"}, []string{"h", "i", "j"})
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
