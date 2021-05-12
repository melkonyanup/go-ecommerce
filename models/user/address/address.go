package address

import (
	"errors"
	"math/rand"

	"github.com/kkamara/go-ecommerce/config"
	"github.com/kkamara/go-ecommerce/models/helper/time"
	"github.com/kkamara/go-ecommerce/models/user"
	"github.com/kkamara/go-ecommerce/schemas"
	"syreclabs.com/go/faker"
)

func Create(newAddress *schemas.UserAddress) (address *schemas.UserAddress, err error) {
	db, err := config.OpenDB()
	if nil != err {
		return
	}
	now := time.Now()
	newAddress.CreatedAt = now
	newAddress.UpdatedAt = now
	res := db.Create(&newAddress)
	if res.RowsAffected < 1 {
		err = errors.New("error creating resource")
		return
	}
	address = newAddress
	if err = res.Error; err != nil {
		return
	}
	return
}

func GetAddress(userId uint64) (addressConfig *schemas.UserAddress, err error) {
	db, err := config.OpenDB()
	if nil != err {
		return
	}
	res := db.Model(&schemas.UserAddress{}).Joins(
		"left join users on user_addresses.user_id = users.id",
	).Where(
		"users.id = ?",
		userId,
	).Limit(1).Find(&addressConfig)
	err = res.Error
	return
}

func Random() (userAddress *schemas.UserAddress, err error) {
	db, err := config.OpenDB()
	if err != nil {
		return
	}
	var count int64
	res := db.Where("deleted_at = ?", "").Order("RANDOM()").Limit(1).Find(&userAddress).Count(&count)
	if err = res.Error; err != nil {
		return
	}
	if count == 0 {
		var u *schemas.User
		u, err = user.Random("")
		if err != nil {
			return
		}
		var buildingName string
		if rand.Intn(2) == 1 {
			buildingName = faker.Company().Name()
		}
		address := faker.Address()
		addr := &schemas.UserAddress{
			UserId:         u.Id,
			MobileNumber:   faker.PhoneNumber().CellPhone(),
			BuildingName:   buildingName,
			StreetAddress1: address.StreetAddress(),
			City:           address.City(),
			Country:        address.Country(),
			Postcode:       address.Postcode(),
		}
		userAddress, err = Create(addr)
		if err != nil {
			return
		}
	}
	return
}

func Seed() (err error) {
	for count := 0; count < 30; count++ {
		var u *schemas.User
		u, err = user.Random("")
		if err != nil {
			return
		}
		var buildingName string
		if rand.Intn(2) == 1 {
			buildingName = faker.Company().Name()
		}
		address := faker.Address()
		addressConfig := &schemas.UserAddress{
			UserId:         u.Id,
			MobileNumber:   faker.PhoneNumber().CellPhone(),
			BuildingName:   buildingName,
			StreetAddress1: address.StreetAddress(),
			City:           address.City(),
			Country:        address.Country(),
			Postcode:       address.Postcode(),
		}

		_, err = Create(addressConfig)
		if err != nil {
			return
		}
	}
	return
}
