package persistence

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/models"
)

func TestListReviews(t *testing.T) {
	// TestCase
	tests := []struct {
		name    string
		want    []*models.Review
		wantErr error
	}{
		{
			name: "ok: list reviews",
			want: []*models.Review{
				{ID: 1, Text: "test message id 1"},
				{ID: 2, Text: "test message id 2"},
				{ID: 3, Text: "test message id 3"},
				{ID: 4, Text: "test message id 4"},
			},
		},
	}

	for _, tt := range tests {
		got, err := reviewRepo.ListReviews(ctx)
		assert.NoError(t, err)
		if diff := cmp.Diff(tt.wantErr, err); len(diff) != 0 {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
		if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}

	}
}

func TestGetReview(t *testing.T) {
	t.Run("ok: exist id 1", func(t *testing.T) {
		got, err := reviewRepo.GetReview(ctx, "1")
		assert.NoError(t, err)
		assert.NotEmpty(t, got)
		assert.NotNil(t, got)
		assert.Equal(t, int64(1), got.ID)
		assert.Equal(t, "test message id 1", got.Text)
	})

	t.Run("ok: exist id 2", func(t *testing.T) {
		got, err := reviewRepo.GetReview(ctx, "2")
		assert.NoError(t, err)
		assert.NotEmpty(t, got)
		assert.NotNil(t, got)
		assert.Equal(t, int64(2), got.ID)
		assert.Equal(t, "test message id 2", got.Text)
	})

	t.Run("ok: not exist id", func(t *testing.T) {
		got, err := reviewRepo.GetReview(ctx, "99999")
		assert.Error(t, err)
		assert.Empty(t, got)
		assert.NotNil(t, got)
	})

	t.Run("blank args", func(t *testing.T) {
		got, err := reviewRepo.GetReview(ctx, "")
		assert.Error(t, err)
		assert.Empty(t, got)
		assert.NotNil(t, got)
	})

}
