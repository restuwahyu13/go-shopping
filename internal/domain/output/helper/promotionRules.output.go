package hopt

type PromotionRules struct {
	Key         string `json:"key,omitempty"`
	Value       any    `json:"value,omitempty"`
	Action      string `json:"action,omitempty"`
	Type        string `json:"type,omitempty"`
	Operator    string `json:"operator,omitempty"`
	OptionOne   string `json:"option_one,omitempty"`
	OptionTwo   string `json:"option_two,omitempty"`
	OptionThree string `json:"option_three,omitempty"`
	OptionFour  string `json:"option_four,omitempty"`
}
