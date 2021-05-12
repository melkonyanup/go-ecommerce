package flagged_review

import (
	"errors"

	"github.com/kkamara/go-ecommerce/config"
	"github.com/kkamara/go-ecommerce/models/helper/time"
	"github.com/kkamara/go-ecommerce/models/product/product_review"
	"github.com/kkamara/go-ecommerce/schemas"
	"syreclabs.com/go/faker"
)

func Create(newFlaggedReview *schemas.ProductFlaggedReview) (flaggedReview *schemas.ProductFlaggedReview, err error) {
	db, err := config.OpenDB()
	if nil != err {
		return
	}
	now := time.Now()
	newFlaggedReview.CreatedAt = now
	newFlaggedReview.UpdatedAt = now
	res := db.Create(&newFlaggedReview)
	if res.RowsAffected < 1 {
		err = errors.New("error creating resource")
		return
	}
	flaggedReview = newFlaggedReview
	if err = res.Error; err != nil {
		return
	}
	return
}

func GetProductReviewFlagCount(productReviewId uint64) (flaggedCount uint64, err error) {
	db, err := config.OpenDB()
	if err != nil {
		return
	}
	res := db.Model(&schemas.ProductFlaggedReview{}).Select(
		"count(product_flagged_reviews.id)",
	).Joins(
		"left join product_reviews on product_flagged_reviews.product_reviews_id = product_reviews.id",
	).Where(
		"product_reviews.id = ?", productReviewId,
	).Find(&flaggedCount)
	err = res.Error
	return
}

func Random() (flaggedReview *schemas.ProductFlaggedReview, err error) {
	db, err := config.OpenDB()
	if err != nil {
		return
	}
	var (
		count int64
		pr    *schemas.ProductReview
	)
	res := db.Where(
		"deleted_at = ?", "",
	).Order("RANDOM()").Limit(1).Find(&flaggedReview).Count(&count)
	if err = res.Error; err != nil {
		return
	}
	if count == 0 {
		pr, err = product_review.Random()
		if err != nil {
			return
		}
		newFlaggedReview := &schemas.ProductFlaggedReview{
			ProductReviewsId: pr.Id,
			FlaggedFromIp:    faker.Internet().IpV4Address(),
		}

		flaggedReview, err = Create(newFlaggedReview)
		if err != nil {
			return
		}
	}
	return
}

func Seed() (err error) {
	for count := 0; count < 30; count++ {
		var pr *schemas.ProductReview
		pr, err = product_review.Random()
		if err != nil {
			return
		}
		flaggedReview := &schemas.ProductFlaggedReview{
			ProductReviewsId: pr.Id,
			FlaggedFromIp:    faker.Internet().IpV4Address(),
		}

		_, err = Create(flaggedReview)
		if err != nil {
			return
		}
	}
	return
}
