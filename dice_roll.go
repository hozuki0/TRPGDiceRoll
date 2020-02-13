package diceroll

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// DiceRoll は 条件付きでない構文をパースし結果を返します
// ex) 1d6 -> (1~6) / 1d6 + 2 -> (3 ~ 8)
func DiceRoll(msg string) (int, error) {
	diceSyntax := extractDiceSyntax(msg)
	if diceSyntax == "" {
		return -1, errors.New("not found dice syntax")
	}
	ret := evaluteDiceSyntax(diceSyntax)
	operator := extractOperatorFomular(msg)
	fmt.Println(operator)
	if operator != "" {
		operatorFunc := evaluteOperatorFomular(operator)
		if operatorFunc != nil {
			ret = operatorFunc(ret)
		}
	}
	return ret, nil
}

// DiceRollWithCondition は 条件付き構文をパースし結果を返します
// ex) 1d6 > 3 -> [true or false] (1 ~ 6)
func DiceRollWithCondition(msg string) (bool, int, error) {
	return false, 0, nil
}

// number d number / number D number という部分を文字列中から抜き出す
func extractDiceSyntax(msg string) string {
	regex := regexp.MustCompile("\\d+[dD]\\d+")
	ret := regex.FindStringSubmatch(strings.TrimSpace(msg))
	if len(ret) == 0 {
		return ""
	}
	return strings.ToLower(ret[0])
}

func extractOperatorFomular(msg string) string {
	regex := regexp.MustCompile("[+\\-*\\/]\\d+")
	ret := regex.FindStringSubmatch(strings.TrimSpace(msg))
	if len(ret) == 0 {
		return ""
	}
	return ret[0]
}

func evaluteDiceSyntax(msg string) int {
	re := regexp.MustCompile("\\d+")
	result := re.FindAllStringSubmatch(msg, 2)
	count, _ := strconv.Atoi(result[0][0])
	face, _ := strconv.Atoi(result[1][0])
	ret := 0
	for i := 0; i < count; i++ {
		ret += rand.Intn(face) + 1
	}
	return ret
}

func evaluteOperatorFomular(msg string) func(n int) int {
	re := regexp.MustCompile("\\d+")
	result := re.FindStringSubmatch(msg)
	value, _ := strconv.Atoi(result[0])
	if msg[0] == '+' {
		return func(n int) int {
			return n + value
		}
	} else if msg[0] == '-' {
		return func(n int) int {
			return n - value
		}
	} else if msg[0] == '*' {
		return func(n int) int {
			return n * value
		}
	} else if msg[0] == '/' {
		return func(n int) int {
			if n == 0 {
				return 0
			}
			return n / value
		}
	}
	return nil
}
