package search

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		testCase   string
		dictionary Dictionary
		key        string
		value      string
		want       string
		wantErr    error
	}{
		{
			"Add item not in dictionary",
			Dictionary{},
			"test",
			"this is just a test",
			"this is just a test",
			nil,
		},
		{
			"search key not found, error returned",
			Dictionary{"test": "this is just a test"},
			"test",
			"",
			"this is just a test",
			ErrWordExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			tt.dictionary.Add(tt.key, tt.value)
			got, err := tt.dictionary.Search(tt.key)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("received error but did not expect to, %s", err.Error())
			}
			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		testCase   string
		dictionary Dictionary
		key        string
		wantErr    error
	}{
		{
			"Delete item not in dictionary",
			Dictionary{},
			"test",
			ErrNotFound,
		},
		{
			"Delete item in dictionary",
			Dictionary{"test": "this is just a test"},
			"test",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			err := tt.dictionary.Delete(tt.key)
			got, _ := tt.dictionary.Search(tt.key)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("received error but did not expect to, %s", err.Error())
			}
			if got != "" {
				t.Errorf("got %q", got)
			}
		})
	}
}

func TestSearch(t *testing.T) {

	tests := []struct {
		testCase   string
		dictionary Dictionary
		search     string
		want       string
		wantErr    error
	}{
		{
			"search key found",
			Dictionary{"test": "this is just a test"},
			"test",
			"this is just a test",
			nil,
		},
		{
			"search key not found, error returned",
			Dictionary{"test": "this is just a test"},
			"bad-key",
			"",
			ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			got, err := tt.dictionary.Search(tt.search)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("received error but did not expect to, %s", err.Error())
			}
			assertStrings(t, got, tt.want, tt.search)
		})
	}
	// assertStrings(t, got, want)
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		testCase   string
		dictionary Dictionary
		key        string
		value      string
		want       string
		wantErr    error
	}{
		{
			"Update item not in dictionary",
			Dictionary{},
			"test",
			"this is just a test",
			"",
			ErrNotFound,
		},
		{
			"Update item in dictionary",
			Dictionary{"test": "this is just a test"},
			"test",
			"",
			"",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			err := tt.dictionary.Update(tt.key, tt.value)
			got, _ := tt.dictionary.Search(tt.key)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("received error but did not expect to, %s", err.Error())
			}
			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}

func assertStrings(t testing.TB, got, want, search string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, search)
	}
}
