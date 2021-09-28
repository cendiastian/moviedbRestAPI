package responses

import (
	"project/ca/business/subscription"
	"time"
)

type PlanResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Expired   string    `json:"expired"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain subscription.SubcriptionPlan) PlanResponse {
	return PlanResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Expired:   domain.Expired,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(domain []subscription.SubcriptionPlan) (response []PlanResponse) {
	for _, user := range domain {
		response = append(response, FromDomain(user))
	}
	return
}
