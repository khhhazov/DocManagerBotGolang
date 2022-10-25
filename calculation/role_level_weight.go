package calculation

func CalcRoleWeight(buttonData string) int {
	switch buttonData {
	case "Министерство образования":
		return 1
	case "Ректор университета":
		return 2
	case "Президент университета":
		return 2
	case "Проректор":
		return 3
	case "Директор":
		return 4
	case "Заведущий кафедры":
		return 5
	case "Преподаватель":
		return 6
	case "Староста учебной группы":
		return 7
	case "Студент":
		return 7
	default:
		return 0
	}
}
