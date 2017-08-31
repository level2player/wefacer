package convert

//Get face++表情
func Get_faceadd_expression(anger_value float64, disgust_value float64, fear_value float64, happiness_value float64, neutral_value float64, sadness_value float64, surprise_value float64) string {
	expressionlist := map[string]float64{"生气": anger_value,
		"讨厌": disgust_value,
		"恐惧": fear_value,
		"开心": happiness_value,
		"平静": neutral_value,
		"伤心": sadness_value,
		"惊讶": surprise_value}
	var max_key string
	var max_value float64
	for key, value := range expressionlist {
		if value > max_value {
			max_value = value
			max_key = key
		}
	}
	return max_key
}
