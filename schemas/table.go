package schemas

type CookingMethodTable struct {
	Beef          bool `json:"beef,omitempty"`
	BlackPepper   bool `json:"black_pepper,omitempty"`
	Butter        bool `json:"butter,omitempty"`
	Chicken       bool `json:"chicken,omitempty"`
	Eggs          bool `json:"eggs,omitempty"`
	Flour         bool `json:"flour,omitempty"`
	Milk          bool `json:"milk,omitempty"`
	Oil           bool `json:"oil,omitempty"`
	Paprika       bool `json:"paprika,omitempty"`
	Parsley       bool `json:"parsley,omitempty"`
	Pork          bool `json:"pork,omitempty"`
	Rice          bool `json:"rice,omitempty"`
	Salt          bool `json:"salt,omitempty"`
	Star_anise    bool `json:"star_anise,omitempty"`
	Sugar         bool `json:"sugar,omitempty"`
	Tofu          bool `json:"tofu,omitempty"`
	Vanilla       bool `json:"vanilla,omitempty"`
	Water         bool `json:"water,omitempty"`
	CornStarch    bool `json:"corn_starch,omitempty"`
	SoySauce      bool `json:"soy_sauce,omitempty"`
	CookingWine   bool `json:"cooking_wine,omitempty"`
	Ginger        bool `json:"ginger,omitempty"`
	Scallion      bool `json:"scallion,omitempty"`
	Vinegar       bool `json:"vinegar,omitempty"`
	Cabbage       bool `json:"cabbage,omitempty"`
	Mushroom      bool `json:"mushroom,omitempty"`
	ChickenPowder bool `json:"chicken_powder,omitempty"`
	Yeast         bool `json:"yeast,omitempty"`
	Tomato_sauce  bool `json:"tomato_sauce,omitempty"`
	Tomato_paste  bool `json:"tomato_paste,omitempty"`
	Tomato        bool `json:"tomato,omitempty"`
	Basil         bool `json:"basil,omitempty"`
	Oregano       bool `json:"oregano,omitempty"`
	Garlic        bool `json:"garlic,omitempty"`
	OnionPowder   bool `json:"onion_powder,omitempty"`
	Pepperoni     bool `json:"pepperoni,omitempty"`
	Cheese        bool `json:"cheese,omitempty"`
	PepperPowder  bool `json:"pepperPowder,omitempty"`
}

type RecipeName struct {
	Name string `json:"name,omitempty"`
}