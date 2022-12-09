package paths

import (
	"boris-middleware/schemas"
	"fmt"
	"os"
	"strings"

	"database/sql"

	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type LaxRecipeRequest struct{
	MethodType string `json:"CookingMethod"`
	Params []string `json:"Params"`
}

func CookingMethodLax(ctx *fiber.Ctx, db *sql.DB) error{

	method_req := LaxRecipeRequest{}
	
	if err := json.Unmarshal(ctx.Body(), &method_req); err != nil{
		return ctx.SendString(err.Error())
	}
	query := strings.Join(method_req.Params, " = 1, ")+" = 1"

	rows, err := db.Query(fmt.Sprintf("SELECT recipe_name FROM %s WHERE (%s)", method_req.MethodType, query)); 
	if err != nil{
		return ctx.SendString(err.Error())
	}

	var dat []byte

	var res_arr = []string{}

	for rows.Next(){
		recipe_name := schemas.RecipeName{}
		rows.Scan(&recipe_name.Name)

		dat, err = os.ReadFile(fmt.Sprintf("./recipes/%s.json",recipe_name.Name))
		if err != nil{
			return ctx.SendString(err.Error())
		}
		res_arr = append(res_arr, string(dat))
	}
	rep := strings.Join(res_arr, "~~")
	ctx.SendString(rep)

	return nil
}