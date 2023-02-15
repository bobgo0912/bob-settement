package modle

var (
	Patterns [][]int
	P0       = []int{0, 6, 12, 18, 24}
	P1       = []int{4, 8, 12, 16, 20}
	P2       = []int{0, 4, 12, 20, 24}
	P3L1     = []int{0, 5, 10, 15, 20, 1, 6, 11, 16, 21} //12
	P4L2     = []int{1, 6, 11, 16, 21, 2, 7, 12, 17, 22} //23
	P5L3     = []int{2, 7, 12, 17, 22, 3, 8, 13, 18, 23} //34
	P6L4     = []int{3, 8, 13, 18, 23, 4, 9, 14, 19, 24} //45
	P7L5     = []int{0, 5, 10, 15, 20, 2, 7, 12, 17, 22} //13
	P8L6     = []int{1, 6, 11, 16, 21, 3, 8, 13, 18, 23} //24
	P9L7     = []int{2, 7, 12, 17, 22, 4, 9, 14, 19, 24} //35
	P10L8    = []int{0, 5, 10, 15, 20, 3, 8, 13, 18, 23} //14
	P11L9    = []int{1, 6, 11, 16, 21, 4, 9, 14, 19, 24} //25
	P12L10   = []int{0, 5, 10, 15, 20, 4, 9, 14, 19, 24} //15
	P13E     = []int{0, 5, 10, 15, 20, 11, 12, 13, 14}
	P14F     = []int{0, 6, 12, 18, 24, 4, 8, 16, 20}                       //
	P15G1    = []int{0, 5, 10, 15, 20, 1, 6, 11, 16, 21, 2, 7, 12, 17, 22} // 123
	P16G2    = []int{1, 6, 11, 16, 21, 2, 7, 12, 17, 22, 3, 8, 13, 18, 23} //234
	P17G3    = []int{2, 7, 12, 17, 22, 3, 8, 13, 18, 23, 4, 9, 14, 19, 24} //345
	P18G4    = []int{0, 5, 10, 15, 20, 1, 6, 11, 16, 21, 3, 8, 13, 18, 23} //124
	P19G5    = []int{0, 5, 10, 15, 20, 1, 6, 11, 16, 21, 4, 9, 14, 19, 24} //125
	P20G6    = []int{0, 5, 10, 15, 20, 2, 7, 12, 17, 22, 3, 8, 13, 18, 23} //134
	P21G7    = []int{0, 5, 10, 15, 20, 3, 8, 13, 18, 23, 4, 9, 14, 19, 24} //145
	P22G8    = []int{1, 6, 11, 16, 21, 2, 7, 12, 17, 22, 4, 9, 14, 19, 24} //235
	P23G9    = []int{1, 6, 11, 16, 21, 3, 8, 13, 18, 23, 4, 9, 14, 19, 24} //245
	P24G10   = []int{0, 5, 10, 15, 20, 2, 7, 12, 17, 22, 4, 9, 14, 19, 24} //135
	P25H     = []int{0, 1, 2, 3, 5, 10, 15, 20, 4, 9, 14, 19, 24, 21, 22, 23, 12}
	P26M1    = []int{0, 5, 10, 15, 20, 1, 6, 11, 16, 21, 2, 7, 12, 17, 22, 3, 8, 13, 18, 23} //1234
	P27M2    = []int{0, 5, 10, 15, 20, 1, 6, 11, 16, 21, 2, 7, 12, 17, 22, 4, 9, 14, 19, 24} //1235
	P28M3    = []int{0, 5, 10, 15, 20, 1, 6, 11, 16, 21, 3, 8, 13, 18, 23, 4, 9, 14, 19, 24} //1245
	P29M4    = []int{0, 5, 10, 15, 20, 2, 7, 12, 17, 22, 3, 8, 13, 18, 23, 4, 9, 14, 19, 24} // 1345                                                     //1345                                                        //1345
	P30M5    = []int{1, 6, 11, 16, 21, 2, 7, 12, 17, 22, 3, 8, 13, 18, 23, 4, 9, 14, 19, 24} //2345
	P31N     = []int{0, 5, 10, 15, 20, 4, 9, 14, 19, 24, 1, 2, 3, 21, 22, 23, 6, 8, 16, 18, 12}
)

func InitPatterns() {
	Patterns = append(Patterns, P0)
	Patterns = append(Patterns, P1)
	Patterns = append(Patterns, P2)
	Patterns = append(Patterns, P3L1)
	Patterns = append(Patterns, P4L2)
	Patterns = append(Patterns, P5L3)
	Patterns = append(Patterns, P6L4)
	Patterns = append(Patterns, P7L5)
	Patterns = append(Patterns, P8L6)
	Patterns = append(Patterns, P9L7)
	Patterns = append(Patterns, P10L8)
	Patterns = append(Patterns, P11L9)
	Patterns = append(Patterns, P12L10)
	Patterns = append(Patterns, P13E)
	Patterns = append(Patterns, P14F)
	Patterns = append(Patterns, P15G1)
	Patterns = append(Patterns, P16G2)
	Patterns = append(Patterns, P17G3)
	Patterns = append(Patterns, P18G4)
	Patterns = append(Patterns, P19G5)
	Patterns = append(Patterns, P20G6)
	Patterns = append(Patterns, P21G7)
	Patterns = append(Patterns, P22G8)
	Patterns = append(Patterns, P23G9)
	Patterns = append(Patterns, P24G10)
	Patterns = append(Patterns, P25H)
	Patterns = append(Patterns, P26M1)
	Patterns = append(Patterns, P27M2)
	Patterns = append(Patterns, P28M3)
	Patterns = append(Patterns, P29M4)
	Patterns = append(Patterns, P30M5)
	Patterns = append(Patterns, P31N)

}
