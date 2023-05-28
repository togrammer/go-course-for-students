package tests

import (
	"context"
	"homework10/internal/adapters/adrepo"
	"strconv"
	"testing"
)

func FuzzMapRepo(f *testing.F) {
	testcases := []int64{1, 1000}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, n int64) {
		if n <= 0 {
			t.Skip("n must be positive")
		}
		testRepo := adrepo.New()
		ctx := context.Background()
		for i := int64(0); i < n; i += 1 {
			_ = testRepo.AddAd(ctx, strconv.Itoa(int(i)), "some text", n)
		}
		got := testRepo.AddAd(ctx, strconv.Itoa(int(n)), "some text", 1)
		expect := n

		if got.ID != expect {
			t.Errorf("Expect: %d, but got: %d", expect, got.ID)
		}
	})
}
