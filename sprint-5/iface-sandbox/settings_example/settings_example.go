package settings_example

/*
На этом примере я показывал, как настроить запуск в GoLand
*/

func GetSum(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return a[0] + GetSum(a[1:])
}
