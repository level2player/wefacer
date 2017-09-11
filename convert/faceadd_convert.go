package convert

import (
	"log"
	"strconv"
)

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

func Get_faceadd_eyestatus(NormalGlassEyeOpen float64, NoGlassEyeClose float64, Occlusion float64, NoGlassEyeOpen float64, NormalGlassEyeClose float64, DarkGlasses float64) string {
	expressionlist := map[string]float64{"佩戴普通眼镜且睁眼": NormalGlassEyeOpen,
		"不戴眼镜且闭眼":   NoGlassEyeClose,
		"眼睛被遮挡":     Occlusion,
		"不戴眼镜且睁眼":   NoGlassEyeOpen,
		"佩戴普通眼镜且闭眼": NormalGlassEyeClose,
		"佩戴墨镜":      DarkGlasses,
	}
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

func Get_faceadd_beauty(Gender string, Male_score float64, Female_score float64) string {
	if Male_score == 0 && Female_score == 0 {
		return "等待开放此功能"
	}
	switch Gender {
	case "Male":
		log.Println(Male_score)
		return strconv.Itoa(int(Male_score))
	case "Female":
		log.Println(Female_score)
		return strconv.Itoa(int(Female_score))
	default:
		return "无法检测颜值"
	}
}
func Get_faceadd_gender(gender string) string {
	switch gender {
	case "Female":
		return "女"
	case "Male":
		return "男"
	}
	return "无法识别"
}
func Get_faceadd_ethnicity(ethnicity string) string {
	switch ethnicity {
	case "Asian":
		return "亚洲人"
	case "White":
		return "白人"
	case "Black":
		return "黑人"
	}
	return "无法识别"
}
