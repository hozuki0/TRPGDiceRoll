package diceroll

import (
	"regexp"
	"strings"
)

// DiceRoll は 条件付きでない構文をパースし結果を返します
// ex) 1d6 -> (1~6) / 1d6 + 2 -> (3 ~ 8)
func DiceRoll(msg string) (int, error) {
	return 0, nil
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
