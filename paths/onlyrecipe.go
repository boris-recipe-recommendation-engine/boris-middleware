package paths

import (
	"boris-middleware/schemas"
	"fmt"
	"os"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/near/borsh-go"
)

type StrictRecipeRequest struct{
	MethodType string `json:"CookingMethod"`
	Params schemas.CookingMethodTable `json:"Params"`
}

func CookingMethodStrict(db *sql.DB) (func(c *fiber.Ctx) error){
	return (
		func(ctx *fiber.Ctx) error {
			method_req := StrictRecipeRequest{}
			if err := ctx.BodyParser(&method_req); err != nil{
				return err
			}

			query := fmt.Sprintf("beef = %t, black_pepper = %t, butter = %t, chicken = %t, eggs = %t, flour = %t, milk = %t, oil = %t, paprika = %t, parsley = %t, pork = %t, rice = %t, salt = %t, star_anise = %t, sugar = %t, tofu = %t, vanilla = %t, water = %t",
			method_req.Params.Beef, method_req.Params.BlackPepper, method_req.Params.Butter, method_req.Params.Chicken, method_req.Params.Eggs, method_req.Params.Flour, method_req.Params.Milk, method_req.Params.Oil, method_req.Params.Paprika, method_req.Params.Parsley, method_req.Params.Pork, method_req.Params.Rice, method_req.Params.Salt, method_req.Params.Star_anise, method_req.Params.Sugar, method_req.Params.Tofu, method_req.Params.Vanilla, method_req.Params.Water)

			rows, err := db.Query("SELECT recipe_name FROM ? WHERE (?)", method_req.MethodType, query); 
			if err != nil{
				return err
			}

			var dat []byte

			var res_arr = []string{}

			for rows.Next(){
				recipe_name := schemas.RecipeName{}
				rows.Scan(&recipe_name.Name)
	
				dat, err = os.ReadFile(fmt.Sprintf("./recipes/%s.txt",recipe_name.Name))
				if err != nil{
					return err
				}

				res_arr = append(res_arr, string(dat))
			}

			data, err := borsh.Serialize(res_arr)
			if err != nil {
				return err
			}			
			ctx.Send(data)

			return nil
	})
}