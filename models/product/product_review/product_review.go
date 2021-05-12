package product_review

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/kkamara/go-ecommerce/config"
	"github.com/kkamara/go-ecommerce/models/helper/number"
	"github.com/kkamara/go-ecommerce/models/helper/time"
	"github.com/kkamara/go-ecommerce/models/product"
	"github.com/kkamara/go-ecommerce/models/user"
	"github.com/kkamara/go-ecommerce/schemas"
	"syreclabs.com/go/faker"
)

func Create(newReview *schemas.ProductReview) (company *schemas.ProductReview, err error) {
	db, err := config.OpenDB()
	if nil != err {
		return
	}
	now := time.Now()
	newReview.CreatedAt = now
	newReview.UpdatedAt = now
	res := db.Create(&newReview)
	if res.RowsAffected < 1 {
		err = errors.New("error creating resource")
		return
	}
	company = newReview
	if err = res.Error; err != nil {
		return
	}
	return
}

func GetProductReviews(productId uint64) (productReviews []*schemas.ProductReview, err error) {
	db, err := config.OpenDB()
	if nil != err {
		return
	}
	res := db.Where(
		"product_id = ?", productId,
	).Where(
		"deleted_at = ?", "",
	).Find(&productReviews)
	err = res.Error
	return
}

func GetAggRating(productId uint64) string {
	defaultValue := "None"
	db, err := config.OpenDB()
	if nil != err {
		return defaultValue
	}
	var count int64
	var review float64
	res := db.Model(&schemas.ProductReview{}).Select("avg(score) as review").Where(
		"product_id = ?",
		productId,
	).Group("product_id").Distinct().Scan(&review).Count(&count)
	if err = res.Error; err != nil {
		return defaultValue
	}
	if count == 0 {
		return defaultValue
	}
	return fmt.Sprintf("%.2f", review)
}

func Random() (productReview *schemas.ProductReview, err error) {
	db, err := config.OpenDB()
	if err != nil {
		return
	}
	var (
		count int64
		u     *schemas.User
		p     *schemas.Product
	)
	res := db.Where(
		"deleted_at = ?", "",
	).Order("RANDOM()").Limit(1).Find(&productReview).Count(&count)
	if err = res.Error; err != nil {
		return
	}
	if count == 0 {
		u, err = user.Random("")
		if err != nil {
			return
		}
		p, err = product.Random()
		if err != nil {
			return
		}
		newProductReview := &schemas.ProductReview{
			UserId:    u.Id,
			ProductId: p.Id,
			Score:     uint8(rand.Intn(6)),
		}
		if rand.Intn(2) != 1 {
			newProductReview.Content = strings.Join(faker.Lorem().Words(number.RandInt(50, 150)), "\n\n")
		}
		if rand.Intn(2) != 1 {
			u, err = user.Random("")
			if err != nil {
				return
			}
			newProductReview.FlaggedReviewDecidedBy = u.Id
			newProductReview.FlaggedReviewDecisionReason = faker.Lorem().Paragraph(rand.Intn(6))
		}

		productReview, err = Create(newProductReview)
		if err != nil {
			return
		}
	}
	return
}

func Seed() (err error) {
	for count := 0; count < 100; count++ {
		var (
			u *schemas.User
			p *schemas.Product
		)
		u, err = user.Random("")
		if err != nil {
			return
		}
		p, err = product.Random()
		if err != nil {
			return
		}
		productReview := &schemas.ProductReview{
			UserId:    u.Id,
			ProductId: p.Id,
			Score:     uint8(rand.Intn(6)),
		}
		if rand.Intn(3) != 2 {
			productReview.Content = strings.Join(faker.Lorem().Sentences(number.RandInt(3, 10)), "\n\n")
		}
		if rand.Intn(2) != 1 {
			u, err = user.Random("")
			if err != nil {
				return
			}
			productReview.FlaggedReviewDecidedBy = u.Id
			productReview.FlaggedReviewDecisionReason = faker.Lorem().Paragraph(rand.Intn(6))
		}

		_, err = Create(productReview)
		if err != nil {
			return
		}
	}
	return
}
