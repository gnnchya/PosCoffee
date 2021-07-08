package implement

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	"reflect"
)

func (impl *implementation) Update(ctx context.Context, input *userin.UpdateInput) (err error) {
	var (
		hashPassword []byte
	)
	err = impl.validator.Validate(input)
	if err != nil {
		return err
	}

	filters := impl.filter.MakeIdFilters(input.ID)
	users := &domain.UserStruct{}
	err = impl.repo.Read(ctx, filters, users)
	if err != nil {
		return err
	}

	//val := reflect.ValueOf(input.PassCode)
	//if !val.IsZero() {
	//	hashPassCode, err = bcrypt.GenerateFromPassword([]byte(input.PassCode), bcrypt.DefaultCost)
	//	if err != nil {
	//		return util.RepoUpdateErr(err)
	//	}
	//}
	val := reflect.ValueOf(hashPassword)
	if val.IsZero() {
		hashPassword = []byte(users.Password)
	}

	//val = reflect.ValueOf(hashPassCode)
	//if val.IsZero() {
	//	hashPassCode = []byte(users.PassCode)
	//}

	//val = reflect.ValueOf(input.MetaData.Email)
	//if !val.IsZero() {
	//	vldEquipment := map[string]bool{}
	//	exist := impl.isExist(ctx, input.ID, input.Email, Email, input.IdentifyType, vldEquipment)
	//	if exist {
	//		return util.RepoUpdateErr(errors.New(EmailExist))
	//	}
	//	users.Email = input.Email
	//}

	//val = reflect.ValueOf(input.MobileNumber)
	//if !val.IsZero() {
	//	vldMobileNumber := map[string]bool{}
	//	exist := impl.isExist(ctx, input.ID, input.MobileNumber, MobileNumber, input.IdentifyType, vldMobileNumber)
	//	if exist {
	//		return util.RepoUpdateErr(errors.New(MobileNumberExist))
	//	}
	//	users.MobileNumber = input.MobileNumber
	//}

	//val = reflect.ValueOf(input.IdentifyNumber)
	//if !val.IsZero() {
	//	encodeIdentifyNumber, err := util.Encrypt([]byte(input.IdentifyNumber), impl.cryptPassPhrase)
	//	if err != nil {
	//		return util.RepoUpdateErr(err)
	//	}
	//	vldIdentifyNumber := map[string]bool{}
	//	exist := impl.isExist(ctx, input.ID, encodeIdentifyNumber, IdentifyNumber, input.IdentifyType, vldIdentifyNumber)
	//	if exist {
	//		return util.RepoUpdateErr(errors.New(IdentifyNumberExist))
	//	}
	//	users.IdentifyNumber = encodeIdentifyNumber
	//}
	//
	//val = reflect.ValueOf(input.RoleID)
	//if val.IsZero() {
	//	input.RoleID = users.RoleID
	//}

	//val = reflect.ValueOf(input.MemberGroup)
	//if val.IsZero() {
	//	memberGroup := &usersin.MemberGroup{
	//		Name:    users.MemberGroup.Name,
	//		Members: users.MemberGroup.Members,
	//	}
	//
	//	input.MemberGroup = memberGroup
	//}
	//
	//val = reflect.ValueOf(input.MetaData)
	//if val.IsZero() {
	//	metaDate := &usersin.MetaData{
	//		Gender:    users.MetaData.Gender,
	//		BirthDate: users.MetaData.BirthDate,
	//		YearOnly:  users.MetaData.YearOnly,
	//		Weight:    users.MetaData.Weight,
	//		Height:    users.MetaData.Height,
	//	}
	//
	//	input.MetaData = metaDate
	//}

	updateInput := userin.UpdateInput{
		ID:             input.ID,
		UID:			input.UID,
		Username: 		input.Username,
		Password:       input.Password,
		MetaData:       input.MetaData,
		RoleID:         input.RoleID,
	}
	update := updateInput.ToDomain()
	update.CreatedAt = users.CreatedAt
	update.ID = ""
	err = impl.repo.Update(ctx, filters, update)
	if err != nil {
		return err
	}

	return nil
}
