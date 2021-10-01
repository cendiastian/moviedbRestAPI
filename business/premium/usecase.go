package premium

// type PaymentUsecase struct {
// 	ConfigJWT      middlewares.ConfigJWT
// 	Repo           Repository
// 	contextTimeout time.Duration
// }

// func NewPaymentUsecase(repo Repository, timeout time.Duration) Usecase {
// 	return &PaymentUsecase{
// 		Repo:           repo,
// 		contextTimeout: timeout,
// 	}
// }

// func (uc *PaymentUsecase) Detail(c context.Context, id int) (res Premium, err error) {
// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()

// 	premium, err := uc.Repo.Detail(ctx, id)
// 	if err != nil {
// 		return Premium{}, err
// 	}

// 	today := time.Now()
// 	if premium.Expired.Before(today) == true {
// 		res = Premium{
// 			UserId:  premium.UserId,
// 			Type:    false,
// 			Expired: time.Time{},
// 		}
// 		premium, err = uc.Repo.Save(ctx, res)
// 		if err != nil {
// 			return Premium{}, err
// 		}
// 	}

// 	return premium, nil

// }

// func (uc *PaymentUsecase) Save(c context.Context, domain Premium) (err error) {

// 	if domain.Id == 0 {
// 		return errors.New("mohon isi ID")
// 	}

// 	ctx, error := context.WithTimeout(c, uc.contextTimeout)
// 	defer error()
// 	_, err = uc.Repo.Detail(ctx, domain.Id)
// 	if err != nil {
// 		return err
// 	}
// 	domain.UpdatedAt = time.Now()

// 	err = uc.Repo.Update(ctx, domain)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
