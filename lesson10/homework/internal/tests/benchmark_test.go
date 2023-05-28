package tests

import (
	"context"
	"homework10/internal/adapters/adrepo"
	"homework10/internal/adapters/userrepo"
	"homework10/internal/app"
	"testing"
)

func BenchmarkRepoApp_FindAd(b *testing.B) {
	ctx := context.Background()

	adRepo := adrepo.New()
	userRepo := userrepo.New()
	repo := app.NewApp(adRepo, userRepo)

	ad, err := repo.CreateAd(ctx, "benchmark title", "benchmark text", 1)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := repo.FindAd(ctx, ad.ID)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
