package templatesutils

import (
	"fmt"
	"forum/utils/declension"
	"time"
)

func timeAgo(e interface{}) string {
	var t time.Time
	switch v := e.(type) {
	case time.Time:
		t = v
	case *time.Time:
		t = *v
	default:
		return ""
	}

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return ""
	}

	t = t.In(location)

	seconds := int(time.Since(t).Seconds())
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	if years > 0 {
		return fmt.Sprintf("%d %s назад", years, declension.Declension(years, "год", "года", "лет"))
	} else if months > 0 {
		return fmt.Sprintf("%d %s назад", months, declension.Declension(months, "месяц", "месяца", "месяцев"))
	} else if weeks > 0 {
		return fmt.Sprintf("%d %s назад", weeks, declension.Declension(weeks, "неделя", "недели", "недель"))
	} else if days > 0 {
		if days == 1 {
			return "вчера"
		}
		if days == 2 {
			return "позавчера"
		}
		return fmt.Sprintf("%d %s назад", days, declension.Declension(days, "день", "дня", "дней"))
	} else if hours > 0 {
		return fmt.Sprintf("%d %s назад", hours, declension.Declension(hours, "час", "часа", "часов"))
	} else if minutes > 0 {
		return fmt.Sprintf("%d %s назад", minutes, declension.Declension(minutes, "минута", "минуты", "минут"))
	} else if seconds > 0 {
		return fmt.Sprintf("%d %s назад", seconds, declension.Declension(seconds, "секунда", "секунды", "секунд"))
	}
	return "только что"
}
