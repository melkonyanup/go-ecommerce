package engine

import (
	"github.com/gofiber/template/html"
	"github.com/kkamara/go-ecommerce/engine/app"
	"github.com/kkamara/go-ecommerce/engine/cart"
	companyEngine "github.com/kkamara/go-ecommerce/engine/company"
	"github.com/kkamara/go-ecommerce/engine/pagination"
	"github.com/kkamara/go-ecommerce/engine/product"
	"github.com/kkamara/go-ecommerce/engine/product/flagged_review"
	"github.com/kkamara/go-ecommerce/engine/product/product_review"
	"github.com/kkamara/go-ecommerce/engine/user"
)

func GetEngine() (engine *html.Engine) {
	engine = html.New("./views", ".html")
	engine.AddFunc("appname", app.AppName)
	engine.AddFunc("htmlSafe", app.HtmlSafe)
	engine.AddFunc("newlinestobr", app.NewLinesToBR)
	engine.AddFunc("matchstring", app.MatchString)
	engine.AddFunc("timesince", app.TimeSince)
	engine.AddFunc("cartCount", cart.CartCount)
	engine.AddFunc("cartPrice", cart.CartPrice)
	engine.AddFunc("role", user.Role)
	engine.AddFunc("userauth", user.Userauth)
	engine.AddFunc("userSlug", user.UserSlug)
	engine.AddFunc("isAuthUser", user.IsAuthUser)
	engine.AddFunc("companySlug", companyEngine.CompanySlug)
	engine.AddFunc("formattedCost", app.FormattedCost)
	engine.AddFunc("formattedFloat64", app.FormattedFloat64)
	engine.AddFunc("doesUserOwnProduct", product.DoesUserOwnProduct)
	engine.AddFunc("productReview", product_review.ProductReview)
	engine.AddFunc("isFlaggedFiveTimes", flagged_review.IsFlaggedFiveTimes)
	engine.AddFunc("companyName", companyEngine.CompanyName)
	engine.AddFunc("shouldDisableLessPage", pagination.ShouldDisableLessPage)
	engine.AddFunc("shouldDisablePrevPage", pagination.ShouldDisablePrevPage)
	engine.AddFunc("shouldDisableNextPage", pagination.ShouldDisableNextPage)
	engine.AddFunc("shouldDisableMorePage", pagination.ShouldDisableMorePage)
	engine.AddFunc("getPrevPage", pagination.GetPrevPage)
	engine.AddFunc("getNextPage", pagination.GetNextPage)
	engine.AddFunc("getLessPage", pagination.GetLessPage)
	engine.AddFunc("getMorePage", pagination.GetMorePage)
	return
}
