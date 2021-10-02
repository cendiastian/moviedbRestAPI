package premium

import (
	"context"
	"time"
)

type Premium struct {
	UserId    int
	Type      bool
	Expired   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type Usecase interface {
// 	Detail(ctx context.Context, user int) (Premium, error)
// 	// Update(ctx context.Context, Premium Premium) error
// 	Save(ctx context.Context, Premium Premium) (Premium, error)
// }

type Repository interface {
	Detail(ctx context.Context, user int) (Premium, error)
	// Update(ctx context.Context, Premium Premium) error
	Save(ctx context.Context, Premium Premium) (Premium, error)
}
