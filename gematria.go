// Package gematria implements encoding of hebrew string into its gematria value.
//
// The exact mapping between letters and numbers is described in the
// documentation for the Value() function.
package gematria

import (
	"errors"
)

var runeValues = map[rune]int{
	rune(1488): 1,
	rune(1489): 2,
	rune(1490): 3,
	rune(1491): 4,
	rune(1492): 5,
	rune(1493): 6,
	rune(1494): 7,
	rune(1495): 8,
	rune(1496): 9,
	rune(1497): 10,
	rune(1499): 20,
	rune(1500): 30,
	rune(1502): 40,
	rune(1504): 50,
	rune(1505): 60,
	rune(1506): 70,
	rune(1508): 80,
	rune(1510): 90,
	rune(1511): 100,
	rune(1512): 200,
	rune(1513): 300,
	rune(1514): 400,

	rune(1478): 50, // Nun Hafukha

	rune(1519): 30, // Yod Triangle
	rune(1520): 12, // Double Vav
	rune(1521): 16, // Vav Yod
	rune(1522): 20, // Double Yod

	rune(1498): 20, // Final Kaf
	rune(1501): 40, // Final Mem
	rune(1503): 50, // Final Nun
	rune(1507): 80, // Final Pe
	rune(1509): 90, // Final Tsadi

	rune(64285): 10,
	rune(64287): 20, // Double Yod
	rune(64288): 70,
	rune(64289): 1,
	rune(64290): 4,
	rune(64291): 5,
	rune(64292): 20,
	rune(64293): 30,
	rune(64294): 40, // Final Mem
	rune(64295): 200,
	rune(64296): 400,
	rune(64298): 300,
	rune(64299): 300,
	rune(64300): 300,
	rune(64301): 300,
	rune(64302): 1,
	rune(64303): 1,
	rune(64304): 1,
	rune(64305): 2,
	rune(64306): 3,
	rune(64307): 4,
	rune(64308): 5,
	rune(64309): 6,
	rune(64310): 7,
	rune(64312): 9,
	rune(64313): 10,
	rune(64314): 20, // Final Kaf
	rune(64315): 20,
	rune(64316): 30,
	rune(64318): 40,
	rune(64320): 50,
	rune(64321): 60,
	rune(64323): 80, // Final Pe
	rune(64324): 80,
	rune(64326): 90,
	rune(64327): 100,
	rune(64328): 200,
	rune(64329): 300,
	rune(64330): 400,
	rune(64331): 6,
	rune(64332): 2,
	rune(64333): 20,
	rune(64334): 80,
	rune(64335): 31, // Alef Lamed
}

var valueRunes = map[int]string{
	0:   "",
	1:   string(rune(1488)),
	2:   string(rune(1489)),
	3:   string(rune(1490)),
	4:   string(rune(1491)),
	5:   string(rune(1492)),
	6:   string(rune(1493)),
	7:   string(rune(1494)),
	8:   string(rune(1495)),
	9:   string(rune(1496)),
	10:  string(rune(1497)),
	20:  string(rune(1499)),
	30:  string(rune(1500)),
	40:  string(rune(1502)),
	50:  string(rune(1504)),
	60:  string(rune(1505)),
	70:  string(rune(1506)),
	80:  string(rune(1508)),
	90:  string(rune(1510)),
	100: string(rune(1511)),
	200: string(rune(1512)),
	300: string(rune(1513)),
	400: string(rune(1514)),
	500: string(rune(1514)) + string(rune(1511)),
	600: string(rune(1514)) + string(rune(1512)),
	700: string(rune(1514)) + string(rune(1513)),
	800: string(rune(1514)) + string(rune(1514)),
	900: string(rune(1514)) + string(rune(1514)) + string(rune(1511)),
	1000: string(rune(1514)) + string(rune(1514)) + string(rune(1512)),
	1100: string(rune(1514)) + string(rune(1514)) + string(rune(1513)),
}

// https://github.com/chaimleib/hebrew-special-numbers/blob/master/styles/default.yml
var specials = map[int]string{
	15:  valueRunes[9] + valueRunes[6],
	16:  valueRunes[9] + valueRunes[7],
	115: valueRunes[100] + valueRunes[9] + valueRunes[6],
	116: valueRunes[100] + valueRunes[9] + valueRunes[7],
	215: valueRunes[200] + valueRunes[9] + valueRunes[6],
	216: valueRunes[200] + valueRunes[9] + valueRunes[7],
	270: valueRunes[70] + valueRunes[200],
	272: valueRunes[70] + valueRunes[200] + valueRunes[2],
	274: valueRunes[70] + valueRunes[4] + valueRunes[200],
	275: valueRunes[70] + valueRunes[200] + valueRunes[5],
	298: valueRunes[200] + valueRunes[8] + valueRunes[90],
	304: valueRunes[4] + valueRunes[300],
	315: valueRunes[300] + valueRunes[9] + valueRunes[6],
	316: valueRunes[300] + valueRunes[9] + valueRunes[7],
	344: valueRunes[300] + valueRunes[4] + valueRunes[40],
	415: valueRunes[400] + valueRunes[9] + valueRunes[6],
	416: valueRunes[400] + valueRunes[9] + valueRunes[7],
	515: valueRunes[500] + valueRunes[9] + valueRunes[6],
	516: valueRunes[500] + valueRunes[9] + valueRunes[7],
	615: valueRunes[600] + valueRunes[9] + valueRunes[6],
	616: valueRunes[600] + valueRunes[9] + valueRunes[7],
	670: valueRunes[70] + valueRunes[600],
	672: valueRunes[400] + valueRunes[70] + valueRunes[200] + valueRunes[2],
	674: valueRunes[70] + valueRunes[4] + valueRunes[200] + valueRunes[400],
	698: valueRunes[600] + valueRunes[8] + valueRunes[90],
	715: valueRunes[700] + valueRunes[9] + valueRunes[6],
	716: valueRunes[700] + valueRunes[9] + valueRunes[7],
	744: valueRunes[700] + valueRunes[4] + valueRunes[40],
	815: valueRunes[800] + valueRunes[9] + valueRunes[6],
	816: valueRunes[800] + valueRunes[9] + valueRunes[7],
	915: valueRunes[900] + valueRunes[9] + valueRunes[6],
	916: valueRunes[900] + valueRunes[9] + valueRunes[7],
	1015: valueRunes[1000] + valueRunes[9] + valueRunes[6],
	1016: valueRunes[1000] + valueRunes[9] + valueRunes[7],
	1115: valueRunes[1100] + valueRunes[9] + valueRunes[6],
	1116: valueRunes[1100] + valueRunes[9] + valueRunes[7],
}

// https://en.wikipedia.org/wiki/Geresh
const geresh = "\u05F3"

// https://en.wikipedia.org/wiki/Gershayim
const gershayim = rune(1524)

// Value returns the gematria value of str, and an error if result has overflowed.
// Value is calculated using the standard encoding,
// assigning the values 1–9, 10–90, 100–400 to the 22 Hebrew letters in order.
// Final letters have the value as the non-final. Non Hebrew characters are ignored.
func Value(str string) (int, error) {
	var sum int
	for _, r := range str {

		result, ok := add(sum, runeValues[r])
		sum = result

		if !ok {
			return result, errors.New("string is too long")
		}
	}
	return sum, nil
}

// add returns the sum of its arguments and a boolean flag which is false if the result overflows an int.
func add(a, b int) (value int, ok bool) {
	result := a + b
	//Overflow if both arguments have the opposite sign of the result.
	if ((a ^ result) & (b ^ result)) < 0 {
		return result, false
	}

	return result, true
}

// Hebrew converts a number to a numeric Hebrew string
// Valid range: 1-1199
func Hebrew(i int) (string, error) {
	if _, ok := specials[i]; ok {
		return specials[i], nil
	}

	result := ""

	if i < 1 || i > 1199 {
		return result, errors.New("Number out of range (1-1199)")
	}

	ones := i % 10
	tens := (i - ones) % 100
	hundreds := (i - ones - tens)

	result += valueRunes[hundreds]
	result += valueRunes[tens]
	result += valueRunes[ones]

	return result, nil
}

// AddGeresh converts adds a geresh or gershayim to a numeric Hebrew string as returned by Hebrew()
// If the string contains one rune, a geresh is appended
// If the string contains more than one rune, gershayim is inserted before the last rune
// If the function receives a non-Hebrew string, the result is undefined
func AddGeresh(heb string) (string, error) {
	result := ""
	tmp := []rune(heb)
	if len(tmp) == 1 {
		result = heb + geresh
	} else if len(tmp) > 1 {
		last := tmp[len(tmp)-1:][0]
		result = string(append(tmp[:len(tmp)-1], gershayim, last))
	}
	return result, nil
}
