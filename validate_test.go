package main

import (
	"testing"
)

func TestValidateChirp(t *testing.T) {
	tests := []struct {
		body         string
		expectedBody string
	}{
		{"I hear Mastodon is better than Chirpy. kerfuffle I need to migrate", "I hear Mastodon is better than Chirpy. **** I need to migrate"},
		{"I hear Mastodon is better than Chirpy. sharbert I need to migrate", "I hear Mastodon is better than Chirpy. **** I need to migrate"},
		{"I hear Mastodon is better than Chirpy. fornax I need to migrate", "I hear Mastodon is better than Chirpy. **** I need to migrate"},
		{"I hear Mastodon is better than Chirpy. I need to migrate", "I hear Mastodon is better than Chirpy. I need to migrate"},
	}

	for _, test := range tests {
		cleaned_body := replaceBadWords(test.body, badWords)
		if cleaned_body != test.expectedBody {
			t.Errorf("Expected '%s', got '%s'", test.expectedBody, cleaned_body)
		}
	}
}