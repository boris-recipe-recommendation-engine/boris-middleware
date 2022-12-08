package paths

import (
	"boris-middleware/schemas"
	"fmt"

	"database/sql"

	"strings"

	"github.com/gofiber/fiber/v2"
)

type RecipeRequest struct{
	Strict bool `json:"isStrict"`
	MethodType string `json:"CookingMethod"`
	Params schemas.CookingMethodTable
}

type recipeName struct{
	Name string  `json:"name,omitempty"`
}

func CookingMethodTable(db *sql.DB) (func(c *fiber.Ctx) error){
	return (
		func(ctx *fiber.Ctx) error {
			method_req := RecipeRequest{}
			if err := ctx.BodyParser(&method_req); err != nil{
				return err
			}

			var conj string
			if method_req.Strict{
				conj = "AND"
			}else{
				conj = "OR"
			}

			query := strings.Join(
				[]string{
					fmt.Sprintf("beef = %t", method_req.Params.Beef),
					fmt.Sprintf("black_pepper = %t", method_req.Params.BlackPepper),
					fmt.Sprintf("butter = %t", method_req.Params.Butter),
					fmt.Sprintf("chicken = %t", method_req.Params.Chicken),
					fmt.Sprintf("eggs = %t", method_req.Params.Eggs),
					fmt.Sprintf("flour = %t", method_req.Params.Flour),
					fmt.Sprintf("milk = %t", method_req.Params.Milk),
					fmt.Sprintf("oil = %t", method_req.Params.Oil),
					fmt.Sprintf("paprika = %t", method_req.Params.Paprika),
					fmt.Sprintf("parsley = %t", method_req.Params.Parsley),
					fmt.Sprintf("pork = %t", method_req.Params.Pork),
					fmt.Sprintf("rice = %t", method_req.Params.Rice),
					fmt.Sprintf("salt = %t", method_req.Params.Salt),
					fmt.Sprintf("star_anise = %t", method_req.Params.Star_anise),
					fmt.Sprintf("sugar = %t", method_req.Params.Sugar),
					fmt.Sprintf("tofu = %t", method_req.Params.Tofu),
					fmt.Sprintf("vanilla = %t", method_req.Params.Vanilla),
					fmt.Sprintf("water = %t", method_req.Params.Water),
				},
				conj,
			)

			rows, err := db.Query("SELECT recipe_name FROM ? WHERE (?)", method_req.MethodType, query); 
			if err != nil{
				return err
			}

			rows.Next()

			recipe_name := recipeName{}
			rows.Scan(&recipe_name.Name)
			
			ctx.SendFile(fmt.Sprintf("./recipes/%s.txt",recipe_name.Name))

			return nil
	})
}