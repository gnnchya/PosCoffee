package implement

import "context"

type changePassword struct {
	Password 		string				`bson:"password" json:"password"`
}

func (impl *implementation)ChangePassword(ctx context.Context, uid string, password string) error{
	filter := impl.filter.MakeUIDFilters(uid)
	input := &changePassword{Password: password}
	err := impl.repo.Update(ctx, filter, input)
	if err != nil{
		return err
	}
	return nil
}
