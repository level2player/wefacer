package convert

func Baidu_gender_convert(gender string) string {
	switch gender {
	case "female":
		return "女"
	case "male":
		return "男"
	}
	return "人妖"
}
func Baidu_expression_convert(expressionNo int) string {
	switch expressionNo {
	case 0:
		return "不笑"
	case 1:
		return "微笑"
	case 2:
		return "大笑"
	}
	return "似笑非笑"
}

func Baidu_glasses_convert(glasses int) string {
	switch glasses {
	case 0:
		return "无眼镜"
	case 1:
		return "普通眼镜"
	case 2:
		return "墨镜"
	}
	return "无眼镜"
}
