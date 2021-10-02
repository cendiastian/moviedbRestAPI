package responses

import (
	"project/business/premium"
	"time"
)

type ProResponse struct {
	UserId    int       `json:"user_id"`
	Type      bool      `json:"type"`
	Expired   time.Time `json:"expiredAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomainPro(domain premium.Premium) ProResponse {
	return ProResponse{
		UserId:    domain.UserId,
		Type:      domain.Type,
		Expired:   domain.Expired,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
