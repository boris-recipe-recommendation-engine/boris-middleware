package paths

import (
	"boris-middleware/schemas"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type StrictRecipeRequest struct{
	MethodType string `json:"CookingMethod"`
	Params schemas.CookingMethodTable `json:"Params"`
}

func CookingMethodStrict(ctx *fiber.Ctx, db *sql.DB) error{
	method_req := StrictRecipeRequest{}
	if err := json.Unmarshal(ctx.Body(), &method_req); err != nil{
		return ctx.SendString(err.Error())
	}

	query := fmt.Sprintf(
		`beef = %t, black_pepper = %t, butter = %t, chicken = %t, eggs = %t, flour = %t, milk = %t, oil = %t, paprika = %t, parsley = %t, pork = %t, rice = %t, salt = %t, star_anise = %t, sugar = %t, tofu = %t, vanilla = %t, water = %t,
		corn_starch = %t, soy_sauce = %t, cooking_wine = %t, ginger = %t, scallion = %t, vinegar = %t, cabbage = %t, mushroom = %t, chicken_powder = %t, yeast = %t, tomato_sauce = %t, tomato_paste = %t, tomato = %t, basil = %t, oregano = %t, garlic = %t, onion_powder = %t, pepperoni = %t, cheese = %t, pepperPowder = %t`,
		method_req.Params.Beef, method_req.Params.BlackPepper, method_req.Params.Butter, method_req.Params.Chicken, method_req.Params.Eggs, method_req.Params.Flour, method_req.Params.Milk, method_req.Params.Oil, method_req.Params.Paprika, method_req.Params.Parsley, method_req.Params.Pork, method_req.Params.Rice, method_req.Params.Salt, method_req.Params.Star_anise, method_req.Params.Sugar, method_req.Params.Tofu, method_req.Params.Vanilla, method_req.Params.Water,
		method_req.Params.CornStarch, method_req.Params.SoySauce, method_req.Params.CookingWine, method_req.Params.Ginger, method_req.Params.Scallion, method_req.Params.Vinegar, method_req.Params.Cabbage, method_req.Params.Mushroom, method_req.Params.ChickenPowder, method_req.Params.Yeast, method_req.Params.Tomato_sauce, method_req.Params.Tomato_paste, method_req.Params.Tomato, method_req.Params.Basil, method_req.Params.Oregano, method_req.Params.Garlic, method_req.Params.OnionPowder, method_req.Params.Pepperoni, method_req.Params.Cheese, method_req.Params.PepperPowder)

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
	ctx.SendString(strings.Join(res_arr, "~~"))

	return nil
}