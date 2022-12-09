package paths

import (
	"boris-middleware/schemas"
	"fmt"
	"os"
	"strings"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/near/borsh-go"
)

type LaxRecipeRequest struct{
	MethodType string `json:"CookingMethod"`
	Params []string `json:"Params"`
}

func CookingMethodLax(db *sql.DB) (func(c *fiber.Ctx) error){
	return (
		func(ctx *fiber.Ctx) error {
			method_req := LaxRecipeRequest{}
			if err := ctx.BodyParser(&method_req); err != nil{
				return err
			}
			query := strings.Join(method_req.Params, "= 1, ")

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