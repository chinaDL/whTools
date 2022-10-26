package gstr

import (
	"context"
	"math"
	"strconv"
	"strings"
)

func GenerateDictDo(charset string, n int, fn func(string, context.CancelFunc, int)) {
	ctx, cancel := context.WithCancel(context.Background())
	rCh := make(chan string, 0)
	count := 0
	go GenerateDict(charset, n, ctx, rCh, &count)
	defer cancel()
For:
	for {
		select {
		case v, ok := <-rCh:
			if !ok {

				break For
			}
			fn(v, cancel, count)
		case <-ctx.Done():
			break For
		}

	}
}

func GenerateDict(charset string, n int, context context.Context, dictCh chan string, retCount *int) {
	m := strings.Split(charset, "")
	mc := len(m)
	cIndexList := make([]string, n)
	count := int(math.Pow(float64(len(m)), float64(n)))
	//fmt.Printf("字典总数: %d\n", count)
	*retCount = count
	for i := 0; i < n; i++ {
		cIndexList[i] = "0"
	}
	r := ""
For:
	for c := 0; c < count; c++ {

		r = ""
		for _, ci := range cIndexList {
			index, _ := strconv.Atoi(ci)
			r += m[index%len(m)]
		}
		cIndexStr := strings.Join(cIndexList, ",")
		cIndex := anyToDecimal(cIndexStr, mc)
		cIndex++
		tStr := decimalToAny(cIndex, mc)
		cIndexList = listLeftPad(strings.Split(tStr, ","), "0", n)
		if dictCh != nil {
			dictCh <- r
		}
		select {
		case <-context.Done():
			break For
		default:

		}
	}
	if dictCh != nil {
		close(dictCh)
	}
}

func listLeftPad(oList []string, padChar string, count int) []string {
	ret := make([]string, 0)
	c := count - len(oList)
	for i := 0; i < c; i++ {
		ret = append(ret, padChar)
	}
	for _, v := range oList {
		ret = append(ret, v)
	}
	return ret
}

func anyToDecimal(num string, n int) int {
	var newNum float64
	newNum = 0.0
	ns := strings.Split(num, ",")
	nNum := len(ns) - 1
	for _, value := range ns {
		tValue, _ := strconv.Atoi(value)
		tmp := float64(tValue)
		if tmp != -1 {
			newNum = newNum + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(newNum)
}

func decimalToAny(num, n int) string {
	newNumStr := ""
	var remainder int
	var remainderString string
	for num != 0 {
		remainder = num % n

		remainderString = strconv.Itoa(remainder)

		newNumStr = "," + remainderString + newNumStr
		num = num / n
	}
	return newNumStr[1:]
}
