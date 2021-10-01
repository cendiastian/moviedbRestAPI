package premium

import (
	"project/business/premium"
	"time"
)

type Premium struct {
	UserId    int
	Type      bool
	Expired   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Premium) ToDomain() premium.Premium {
	return premium.Premium{
		UserId:    user.UserId,
		Type:      user.Type,
		Expired:   user.Expired,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain premium.Premium) Premium {
	return Premium{
		UserId:    domain.UserId,
		Type:      domain.Type,
		Expired:   domain.Expired,
		UpdatedAt: domain.UpdatedAt,
	}
}
