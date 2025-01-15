package output

import (
	"github.com/fatih/color"
)

func PrintErorr(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки %d", intValue)
		return
	}
	str, ok := value.(string)
	if ok {
		color.Red("Код ошибки %d", str)
		return
	}
	erroValue, ok := value.(error)
	if ok {
		color.Red(erroValue.Error())
		return
	}
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки")
	}
}
