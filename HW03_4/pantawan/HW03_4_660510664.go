// Pantawan Sosamer
// 660510664
// HW03_4
// 204203 Sec 001

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func roundToEven(x string, bPlace uint8) string {

	// หาจุด
	var ansX string
	var count int = 0
	var toCheck int = 0
	var findI, lenc int
	var j uint8
	count = strings.Count(x, ".")
	if count > 0 {
		findI = strings.Index(x, ".")
	} else {
		// เติมจุดให้ตัวที่ไม่มีจุด
		x = x + "."
		findI = strings.Index(x, ".")
	}

	// เติมให้มีทศนิยมเท่า bPlace
	var coppy string
	coppy = ""
	coppy = x[findI+1:]

	// fmt.Println(coppy)
	if bPlace == 0 {
		xx := x[:findI]
		yy := x[findI:]
		if len(coppy) == 0 { //ดักอะไร
			return xx
		} else if x[findI+1] == 48 {
			// fmt.Println("a")
			return xx
		} else if string(xx[len(xx)-1]) == "1" {
			// fmt.Println("b",string(xx[len(xx)-1]))
			// fmt.Println("cc")
			toIII := make([]int64, len(xx))
			for i := 0; i < len(xx); i++ {
				toIII[i] = scanToInt(string(xx[i]))
			}
			// fmt.Println(xx)
			toIII[len(toIII)-1]++ //last index add 1
			// fmt.Println(toIII)
			l := len(toIII) - 1

			//edit binary as correct
			for i := 0; i < len(toIII); i++ {
				if toIII[l] != 2 {
					break
				} else if toIII[l] == 2 {
					toIII[l] = 0
					if i != len(toIII)-1 {
						toIII[l-1]++
					}
				}
				l--
			}

			//convert string to num
			var z string = ""
			for i := 0; i < len(toIII); i++ {
				z += strconv.Itoa(int(toIII[i]))
			}
			ansX = z

			//
			var bb bool = true
			for i := 0; i < len(ansX); i++ {
				if string(ansX[i]) == "." {
					continue
				} else if string(ansX[i]) == "1" {
					bb = false
				}
			}

			//if first is not '1'
			if bb {
				ansX = "1" + ansX
			}
			return ansX

		} else {
			yy = x[findI+1:]
			// fmt.Println(yy)

			// check half
			for p := 1; p < len(yy); p++ {
				// fmt.Println("fuc")
				if yy[p] == 49 { // yy[p]=='1'
					toCheck++
				}
			}
			// fmt.Println(toCheck)
			// fmt.Println("yy")
			if toCheck > 0 { // found '1' in yy
				// fmt.Println("yy")
				// fmt.Println(xx)
				toIII := make([]int64, len(xx))
				for i := 0; i < len(xx); i++ {
					toIII[i] = scanToInt(string(xx[i]))
				}
				toIII[len(toIII)-1]++ // add 1 in last index
				l := len(toIII) - 1

				//edit binary as correct
				for i := 0; i < len(toIII); i++ {
					if toIII[l] != 2 {
						break
					} else if toIII[l] == 2 {
						toIII[l] = 0
						if i != len(toIII)-1 {
							toIII[l-1]++
						}
					}
					l--
				}

				// fmt.Println(xx)
				// fmt.Println(toIII)
				var z string = ""
				for i := 0; i < len(toIII); i++ {
					z += strconv.Itoa(int(toIII[i]))
				}
				ansX = z

				var bb bool = true
				for i := 0; i < len(ansX); i++ {
					if string(ansX[i]) == "." {
						continue
					} else if string(ansX[i]) == "1" {
						bb = false
					}
				}
				if bb == true {
					ansX = "1" + ansX
				}
				return ansX

			} else if (toCheck == 0) && (xx[len(xx)-1] == 49) {
				toIII := make([]int64, len(xx))
				for i := 0; i < len(xx); i++ {
					toIII[i] = scanToInt(string(xx[i]))
				}
				toIII[len(toIII)-1]++
				l := len(toIII) - 1
				for i := 0; i < len(toIII); i++ {
					if toIII[l] != 2 {
						break
					} else if toIII[l] == 2 {
						toIII[l] = 0
						if i != len(toIII)-1 {
							toIII[l-1]++
						}
					}
					l--
				}

				// fmt.Println(xx)
				// fmt.Println(toIII)
				var z string = ""
				for i := 0; i < len(toIII); i++ {
					z += strconv.Itoa(int(toIII[i]))
				}
				ansX = z

				var bb bool = true
				for i := 0; i < len(ansX); i++ {
					if string(ansX[i]) == "." {
						continue
					} else if string(ansX[i]) == "1" {
						bb = false
					}
				}
				if bb == true {
					ansX = "1" + ansX
				}
				return ansX
				// return x[:findI]
			} else {
				return x[:findI]
			}
		}
	} else if len(coppy) == 0 {
		for j = 0; j < bPlace; j++ {
			x += "0"
		}
		ansX = x
		return x
	} else if uint8(len(coppy)) < bPlace {
		for j = 0; j < bPlace-uint8(len(coppy)); j++ {
			x += "0"
		}
		ansX = x
		return ansX
	} else if uint8(len(coppy)) == bPlace {
		return x
	}

	coppy = x[findI+1:]
	// fmt.Println("cppy",coppy)
	if coppy[bPlace] == '0' {
		// fmt.Println("xx")
		// if coppy[bPlace] == '0' {
		ansX = x[:findI+1] + string(coppy[:bPlace]) //ถ้าตัวหลังตำแหน่งต้องการปัดเป็น 0 ปัดข้างหลังทิ้ง
		// fmt.Println("xxอ")
		return ansX
		// }
	} else {
		// fmt.Println("777")
		coppyB := string(coppy[bPlace:]) //เช็คหลังตำแหน่งที่ต้องการปัดเป็นต้นไป
		// fmt.Println("cb",coppy)
		lenc = len(coppyB)
		if lenc == 1 {
			if coppy[bPlace-1] == 48 {
				ansX = x[:findI+1] + string(coppy[:bPlace]) //ถ้าเหลือตัวเดียว ปัดข้างหลังทิ้ง
				// fmt.Println("lenc 1",ansX)
			} else {
				toCheck++
			}
		} else {
			// เช็คว่าหลังตัวที่ต้องการปัด มากกว่าครึ่งหรือไม่
			// fmt.Println("coppy")
			// fmt.Println("lenc",lenc)
			toCheck = 0
			for p := 1; p < lenc; p++ {

				if coppyB[p] == 49 {
					// fmt.Println("coppyB[p]",coppyB[p])
					// toCheck = 0
					toCheck++
				}
			}
		}
	}

	// fmt.Println("toCheck",toCheck)
	// fmt.Println(ansX)
	// fmt.Println(coppy)

	if toCheck == 0 && coppy[bPlace-1] != 49 {
		// if toCheck == 0 {
		x = x[:findI+1] + string(coppy[:bPlace]) //ถ้าไม่มากกว่าครึ่ง ปัดข้างหลังทิ้ง
		ansX = x
		// fmt.Println("bb")
		// } else if ansX[len(ansX) - 1] == 48{
		// fmt.Println(ansX)
		// 	return ansX
	} else {
		// fmt.Println("cc")
		var k int = 0
		x = x[:findI+1] + string(coppy[:bPlace])
		toI := make([]int64, len(x))
		toII := make([]int64, len(x)-1)

		// fmt.Println(x)

		for i := 0; i < len(x); i++ {
			if string(x[i]) == "." {
				toI[i] = 5
			} else {
				toI[i] = scanToInt(string(x[i]))
			}
		}

		for i := 0; i < len(x); i++ {
			if string(x[i]) == "." {
				continue
			} else {
				toII[k] = scanToInt(string(x[i]))
			}
			k++
		}

		// fmt.Println(toI)
		// fmt.Println(toII)

		toII[len(toII)-1]++
		lenT := len(toII) - 1
		for i := 0; i < len(toII); i++ {
			// fmt.Println(toII)
			if toII[lenT] != 2 {
				break
			} else if toII[lenT] == 2 {
				toII[lenT] = 0
				if i != len(toII)-1 {
					toII[lenT-1]++
				}
			}
			lenT--
		}

		k = 0
		for i := 0; i < len(toI); i++ {
			if toI[i] == 5 {
				continue
			} else {
				toI[i] = toII[k]
			}
			k++
		}

		// fmt.Println(toI)
		var keep string = ""
		for i := 0; i < len(toI); i++ {
			keep += strconv.Itoa(int(toI[i]))
		}

		ansX = strings.Replace(keep, "5", ".", 1)
		// fmt.Println(ansX)

	}

	var b bool = true
	for i := 0; i < len(ansX); i++ {
		if string(ansX[i]) == "." {
			continue
		} else if string(ansX[i]) == "1" {
			b = false
		}
	}
	if b == true {
		ansX = "1" + ansX
	}

	return ansX
}

func scanToInt(s string) int64 {
	var n int64
	fmt.Sscan(s, &n)
	return n
}
