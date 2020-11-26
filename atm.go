package Atm

func Atm(n int) []int {
	billType := []int{1000, 500, 100}
	v := []int{}

	for i := 0; i < len(billType); i++ {
		if n/billType[i] >= 1 || n >= 100 {
			for j := 0; j < n/billType[i]; j++ {
				v = append(v, billType[i])
			}
			n %= billType[i]
		} else {
			println("This machine have only 1000, 500 and 100 Thb bill")
			return v
		}
	}
	return v
}
