package main
import (
  "math"
  "fmt"
)

func sumDigitDecimals(digit1, digit2, cO int8) (int8, byte) {
  res := digit1 + digit2 + cO
  var cORes byte

  if res >= 10 {
    cORes = byte(res/10)
    res = int8(math.Mod(float64(res), 10))
  }

  return res, cORes
}

func addStrings(num1 string, num2 string) string {
  var carryOver byte

  num1Ptr := len(num1)-1
  num2Ptr := len(num2)-1
  maxLen := int(math.Max(float64(len(num1)), float64(len(num2))))
  restArr := make([]byte, maxLen+1)
  resCurrIndex := len(restArr)-1

  for num1Ptr >= 0 || num2Ptr >= 0 {

    var digit int8
    if num1Ptr < 0 {
      digit, carryOver = sumDigitDecimals(int8(num2[num2Ptr]-'0'), int8(carryOver), 0)
    } else if num2Ptr < 0 {
      digit, carryOver = sumDigitDecimals(int8(num1[num1Ptr]-'0'), int8(carryOver), 0)
    } else {
      digit, carryOver = sumDigitDecimals(int8(num1[num1Ptr]-'0'), int8(num2[num2Ptr]-'0'), int8(carryOver))
    }

    restArr[resCurrIndex] = byte(digit)+'0'
    num1Ptr--
    num2Ptr--
    resCurrIndex--
  }

  if carryOver != 0 {
    restArr[0] = carryOver+'0'
  }

  if restArr[0] == '0' && num1[0] != '0' && num2[0] != '0' {
    return string(restArr[1:])
  }

  return string(restArr)
}

func main() {
  fmt.Printf("added strings %s\n", addStrings("0", "0"))
}