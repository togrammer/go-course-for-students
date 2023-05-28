package adrepo

import (
	"context"
	"errors"
	"homework10/internal/app"
	"testing"
	_ "time"
)

func TestRepoInit_AddAd(t *testing.T) {
	repo := New()
	ctx := context.Background()

	ad := repo.AddAd(ctx, "title", "text", 1)

	if ad.ID != 0 {
		t.Errorf("Expected ID to be 0, but got %d", ad.ID)
	}

	if ad.Title != "title" {
		t.Errorf("Expected Title to be 'title', but got '%s'", ad.Title)
	}

	if ad.Text != "text" {
		t.Errorf("Expected Text to be 'text', but got '%s'", ad.Text)
	}

	if ad.AuthorID != 1 {
		t.Errorf("Expected AuthorID to be 1, but got %d", ad.AuthorID)
	}

	if ad.Published != false {
		t.Errorf("Expected Published to be false, but got %v", ad.Published)
	}

	if ad.Created.IsZero() {
		t.Errorf("Expected Created to be non-zero, but got zero")
	}

	if ad.Updated.IsZero() {
		t.Errorf("Expected Updated to be non-zero, but got zero")
	}
}

func TestRepoInit_ChangeStatus(t *testing.T) {
	repo := New()
	ctx := context.Background()

	ad := repo.AddAd(ctx, "title", "text", 1)

	repo.ChangeStatus(ctx, ad.ID, 1, true)
	upd, _ := repo.FindAd(ctx, ad.ID)
	if upd.Published != true {
		t.Errorf("Expected Published to be true, but got %v", ad.Published)
	}
}

func TestRepoInit_UpdateAd(t *testing.T) {
	repo := New()
	ctx := context.Background()

	ad := repo.AddAd(ctx, "title", "text", 1)

	repo.UpdateAd(ctx, ad.ID, 1, "new title", "new text")

	ad, _ = repo.FindAd(ctx, ad.ID)

	if ad.Title != "new title" {
		t.Errorf("Expected Title to be 'new title', but got '%s'", ad.Title)
	}

	if ad.Text != "new text" {
		t.Errorf("Expected Text to be 'new text', but got '%s'", ad.Text)
	}
}

func TestRepoInit_FindAd(t *testing.T) {
	repo := New()
	ctx := context.Background()

	ad := repo.AddAd(ctx, "title", "text", 1)

	foundAd, err := repo.FindAd(ctx, ad.ID)

	if err != nil {
		t.Errorf("Expected err to be nil, but got %v", err)
	}

	if foundAd.ID != ad.ID {
		t.Errorf("Expected ID to be %d, but got %d", ad.ID, foundAd.ID)
	}

	if foundAd.Title != ad.Title {
		t.Errorf("Expected Title to be '%s', but got '%s'", ad.Title, foundAd.Title)
	}

	if foundAd.Text != ad.Text {
		t.Errorf("Expected Text to be '%s', but got '%s'", ad.Text, foundAd.Text)
	}

	if foundAd.AuthorID != ad.AuthorID {
		t.Errorf("Expected AuthorID to be %d, but got %d", ad.AuthorID, foundAd.AuthorID)
	}

	if foundAd.Published != ad.Published {
		t.Errorf("Expected Published to be %v, but got %v", ad.Published, foundAd.Published)
	}

	if !foundAd.Created.Equal(ad.Created) {
		t.Errorf("Expected Created to be %v, but got %v", ad.Created, foundAd.Created)
	}

	if !foundAd.Updated.Equal(ad.Updated) {
		t.Errorf("Expected Updated to be %v, but got %v", ad.Updated, foundAd.Updated)
	}
}

func TestRepoInit_FindAd_WrongId(t *testing.T) {
	repo := New()
	ctx := context.Background()

	_, err := repo.FindAd(ctx, 0)

	if !errors.Is(err, app.ErrWrongAdId) {
		t.Errorf("Expected err to be %v, but got %v", app.ErrWrongAdId, err)
	}
}

func TestRepoInit_ListAds(t *testing.T) {
	repo := New()
	ctx := context.Background()

	ad1 := repo.AddAd(ctx, "title1", "text1", 1)
	ad2 := repo.AddAd(ctx, "title2", "text2", 2)

	ads := repo.ListAds(ctx)

	if len(ads) != 2 {
		t.Errorf("Expected len(ads) to be 2, but got %d", len(ads))
	}

	if ads[0].ID != ad1.ID {
		t.Errorf("Expected ID to be %d, but got %d", ad1.ID, ads[0].ID)
	}

	if ads[0].Title != ad1.Title {
		t.Errorf("Expected Title to be '%s', but got '%s'", ad1.Title, ads[0].Title)
	}

	if ads[0].Text != ad1.Text {
		t.Errorf("Expected Text to be '%s', but got '%s'", ad1.Text, ads[0].Text)
	}

	if ads[0].AuthorID != ad1.AuthorID {
		t.Errorf("Expected AuthorID to be %d, but got %d", ad1.AuthorID, ads[0].AuthorID)
	}

	if ads[0].Published != ad1.Published {
		t.Errorf("Expected Published to be %v, but got %v", ad1.Published, ads[0].Published)
	}

	if !ads[0].Created.Equal(ad1.Created) {
		t.Errorf("Expected Created to be %v, but got %v", ad1.Created, ads[0].Created)
	}

	if !ads[0].Updated.Equal(ad1.Updated) {
		t.Errorf("Expected Updated to be %v, but got %v", ad1.Updated, ads[0].Updated)
	}

	if ads[1].ID != ad2.ID {
		t.Errorf("Expected ID to be %d, but got %d", ad2.ID, ads[1].ID)
	}

	if ads[1].Title != ad2.Title {
		t.Errorf("Expected Title to be '%s', but got '%s'", ad2.Title, ads[1].Title)
	}

	if ads[1].Text != ad2.Text {
		t.Errorf("Expected Text to be '%s', but got '%s'", ad2.Text, ads[1].Text)
	}

	if ads[1].AuthorID != ad2.AuthorID {
		t.Errorf("Expected AuthorID to be %d, but got %d", ad2.AuthorID, ads[1].AuthorID)
	}

	if ads[1].Published != ad2.Published {
		t.Errorf("Expected Published to be %v, but got %v", ad2.Published, ads[1].Published)
	}

	if !ads[1].Created.Equal(ad2.Created) {
		t.Errorf("Expected Created to be %v, but got %v", ad2.Created, ads[1].Created)
	}

	if !ads[1].Updated.Equal(ad2.Updated) {
		t.Errorf("Expected Updated to be %v, but got %v", ad2.Updated, ads[1].Updated)
	}
}
