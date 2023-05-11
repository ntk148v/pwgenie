// Copyright 2023 Kien Nguyen-Tuan <kiennt2609@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"math/rand"
	"strings"
	"testing"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const N = 1000

func hasDuplicate(s []string) bool {
	found := make(map[string]struct{}, len(s))
	for _, e := range s {
		if _, ok := found[e]; ok {
			return true
		}

		found[e] = struct{}{}
	}
	return false
}

func Test_genHuman(t *testing.T) {
	t.Parallel()

	r := rand.New(rand.NewSource(69))
	normal, _ := genHuman(r, 5, " ", false, false)

	t.Run("separator", func(t *testing.T) {
		t.Parallel()

		r := rand.New(rand.NewSource(69))
		res, err := genHuman(r, 5, "-", false, false)
		if err != nil {
			t.Error(err)
		}
		if strings.ReplaceAll(res, "-", " ") != normal {
			t.Errorf("genHuman() = %v, not validated", res)
		}
	})

	t.Run("capitalize", func(t *testing.T) {
		t.Parallel()

		r := rand.New(rand.NewSource(69))
		res, err := genHuman(r, 5, " ", true, false)
		if err != nil {
			t.Error(err)
		}

		if cases.Title(language.English).String(normal) != res {
			t.Errorf("%q not validated", res)
		}
	})

	t.Run("no_repeat", func(t *testing.T) {
		t.Parallel()

		res, err := genHuman(r, len(EFFWords), " ", false, false)
		if err != nil {
			t.Error(err)
		}
		if hasDuplicate(strings.Split(res, " ")) {
			t.Errorf("%q should not have duplicates", res)
		}
	})

	t.Run("no_repeat_failed", func(t *testing.T) {
		t.Parallel()

		_, err := genHuman(r, len(EFFWords)+1, " ", false, false)
		if err != nil {
			if !errors.Is(err, ErrTooManyCharacters) {
				t.Errorf("%q should be %q", err, ErrTooManyCharacters)
			}
		}
	})
}

func Test_genRandom(t *testing.T) {
	t.Parallel()
	r := rand.New(rand.NewSource(69))
	t.Run("gen_lowercase", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < N; i++ {
			res, err := genRandom(r, i%len(LowerLetters), false, false, false, true)
			if err != nil {
				t.Error(err)
			}
			if res != strings.ToLower(res) {
				t.Errorf("%q is not lowercase", res)
			}
		}
	})

	t.Run("gen_uppercase", func(t *testing.T) {
		t.Parallel()

		res, err := genRandom(r, N, true, false, false, true)
		if err != nil {
			t.Error(err)
		}
		if res == strings.ToLower(res) {
			t.Errorf("%q does not include uppercase", res)
		}
	})

	t.Run("gen_symbol", func(t *testing.T) {
		t.Parallel()

		res, err := genRandom(r, N, false, false, true, true)
		if err != nil {
			t.Error(err)
		}
		if !strings.ContainsAny(res, Symbols) {
			t.Errorf("%q does not include any symbols", res)
		}
	})

	t.Run("gen_digit", func(t *testing.T) {
		t.Parallel()

		res, err := genRandom(r, N, false, true, false, true)
		if err != nil {
			t.Error(err)
		}
		if !strings.ContainsAny(res, Digits) {
			t.Errorf("%q does not include any digits", res)
		}
	})

	t.Run("gen_no_repeat", func(t *testing.T) {
		t.Parallel()

		res, err := genRandom(r, len(LowerLetters+UpperLetters+Digits+Symbols), true, true, true, false)
		if err != nil {
			t.Error(err)
		}
		if hasDuplicate(strings.Split(res, "")) {
			t.Errorf("%q should not have duplicate", res)
		}
	})

	t.Run("gen_no_repeat_failed", func(t *testing.T) {
		t.Parallel()

		_, err := genRandom(r, len(LowerLetters+UpperLetters+Digits+Symbols)+1, true, true, true, false)
		if err != nil {
			if !errors.Is(err, ErrTooManyCharacters) {
				t.Errorf("%q should be %q", err, ErrTooManyCharacters)
			}
		}
	})
}

func Test_genPIN(t *testing.T) {
	t.Parallel()
	r := rand.New(rand.NewSource(69))

	t.Run("gen_no_repeat", func(t *testing.T) {
		t.Parallel()
		res, err := genPIN(r, len(Digits)-1, false)
		if err != nil {
			t.Error(err)
		}

		if hasDuplicate(strings.Split(res, "")) {
			t.Errorf("%q should not have duplicate", res)
		}
	})

	t.Run("gen_no_repeat_failed", func(t *testing.T) {
		t.Parallel()

		_, err := genPIN(r, len(Digits)+1, false)
		if err != nil {
			if !errors.Is(err, ErrTooManyCharacters) {
				t.Errorf("%q should be %q", err, ErrTooManyCharacters)
			}
		}
	})
}
