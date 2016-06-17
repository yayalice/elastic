// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import "testing"

func TestAnalyze(t *testing.T) {
	var err error

	client := setupTestClientAndCreateIndex(t)

	res, err := client.Analyze().Index(testIndexName).Analyzer("standard").Body("a test").Do()
	if err != nil {
		t.Fatal(err)
	}
	if !isEqual(res.Tokens, []Token{
		Token{
			Token:       "a",
			StartOffset: 0,
			EndOffset:   1,
			Position:    0,
			Type:        "<ALPHANUM>",
		},
		Token{
			Token:       "test",
			StartOffset: 2,
			EndOffset:   6,
			Position:    2,
			Type:        "<ALPHANUM>",
		},
	}) {
		t.Error("unexpected tokens")
	}
}

func isEqual(a, b []Token) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
