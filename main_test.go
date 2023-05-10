package main

import (
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

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_genHuman(t *testing.T) {
	t.Parallel()

	r := rand.New(rand.NewSource(69))
	normal := genHuman(r, 5, " ", false, false)

	t.Run("separator", func(t *testing.T) {
		t.Parallel()

		r := rand.New(rand.NewSource(69))
		res := genHuman(r, 5, "-", false, false)
		if strings.ReplaceAll(res, "-", " ") != normal {
			t.Errorf("genHuman() = %v, not validated", res)
		}
	})

	t.Run("capitalize", func(t *testing.T) {
		t.Parallel()

		r := rand.New(rand.NewSource(69))
		res := genHuman(r, 5, " ", true, false)
		if cases.Title(language.English).String(normal) != res {
			t.Errorf("%q not validated", res)
		}
	})

	t.Run("no_repeat", func(t *testing.T) {
		t.Parallel()

		r := rand.New(rand.NewSource(69))
		res := genHuman(r, 50, " ", false, false)
		if hasDuplicate(strings.Split(res, " ")) {
			t.Errorf("%q should not have duplicates", res)
		}
	})
}

func Test_genRandom(t *testing.T) {
	t.Parallel()
	r := rand.New(rand.NewSource(69))
	t.Run("gen_lowercase", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < N; i++ {
			res := genRandom(r, i%len(LowerLetters), false, false, false, true)
			if res != strings.ToLower(res) {
				t.Errorf("%q is not lowercase", res)
			}
		}
	})

	t.Run("gen_uppercase", func(t *testing.T) {
		t.Parallel()

		res := genRandom(r, N, true, false, false, true)
		if res == strings.ToLower(res) {
			t.Errorf("%q does not include uppercase", res)
		}
	})

	t.Run("gen_symbol", func(t *testing.T) {
		t.Parallel()

		res := genRandom(r, N, false, false, true, true)
		if !strings.ContainsAny(res, Symbols) {
			t.Errorf("%q does not include any symbols", res)
		}
	})

	t.Run("gen_digit", func(t *testing.T) {
		t.Parallel()

		res := genRandom(r, N, false, true, false, true)
		if !strings.ContainsAny(res, Digits) {
			t.Errorf("%q does not include any digits", res)
		}
	})

	t.Run("gen_no_repeat", func(t *testing.T) {
		t.Parallel()

		res := genRandom(r, len(LowerLetters+UpperLetters+Digits+Symbols), true, true, true, false)
		if hasDuplicate(strings.Split(res, "")) {
			t.Errorf("%q should not have duplicate", res)
		}
	})
}

func Test_genPIN(t *testing.T) {
	type args struct {
		r           *rand.Rand
		length      int
		allowRepeat bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genPIN(tt.args.r, tt.args.length, tt.args.allowRepeat); got != tt.want {
				t.Errorf("genPIN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randElement(t *testing.T) {
	type args struct {
		r *rand.Rand
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randElement(tt.args.r, tt.args.s); got != tt.want {
				t.Errorf("randElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randInsert(t *testing.T) {
	type args struct {
		r *rand.Rand
		s string
		e string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randInsert(tt.args.r, tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("randInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
