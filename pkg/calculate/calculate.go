package calculate

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	elements []interface{} //elements
}

func NewStack() *Stack {
	return &Stack{make([]interface{}, 0, 100)}
}

func (s *Stack) Push(value ...interface{}) {
	s.elements = append(s.elements, value...)
}

func (s *Stack) Top() (value interface{}) {
	if s.Size() > 0 {
		return s.elements[s.Size()-1]
	}
	return nil
}

func (s *Stack) Pop() (value interface{}, err error) {
	if s.Size() > 0 {
		value = s.elements[s.Size()-1]
		s.elements = s.elements[:s.Size()-1]
		return
	}
	return nil, errors.New("s is empty") //read empty s
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Empty() bool {
	if s.elements == nil || s.Size() == 0 {
		return true
	}
	return false
}

type Calculator struct {
	stDat *Stack
	stSym *Stack
}

func NewCalculator() *Calculator {
	return &Calculator{NewStack(), NewStack()}
}

func math(num1, num2 float64, sym byte) (result float64) {
	switch sym {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		result = num1 / num2
	}
	return
}

func (c *Calculator) Step(num float64, sym byte) error {
	if sym == '\n' {
		if c.stSym.Size() == 0 {
			c.stDat.Push(num)
		} else {
			length := c.stSym.Size()
			for i := 0; i < length; i++ {
				symbol, _ := c.stSym.Pop()
				ifc, _ := c.stDat.Pop()
				num1 := ifc.(float64)
				num = math(num1, num, symbol.(byte))
			}
			c.stDat.Push(num)
		}
	} else {
		if c.stSym.Size() != 0 {
			tSymbol := c.stSym.Top().(byte)
			pc := priorityCompare(tSymbol, sym)
			switch {
			case pc < 0:
				c.stDat.Push(num)
				c.stSym.Push(sym)
			case pc >= 0:
				// 遇到运算符优先级相等的时候需要先把栈内数据计算
				length := c.stSym.Size()
				for i := 0; i < length; i++ {
					symbol, _ := c.stSym.Pop()
					ifc, _ := c.stDat.Pop()
					num1 := ifc.(float64)
					num = math(num1, num, symbol.(byte))
				}
				c.stDat.Push(num)
				c.stSym.Push(sym)
			}
		} else {
			c.stDat.Push(num)
			c.stSym.Push(sym)
		}
	}
	return nil
}

func (c *Calculator) Result() (result float64, err error) {
	if c.stSym.Empty() && c.stDat.Size() == 1 {
		result = c.stDat.Top().(float64)
	} else {
		//计算式未结束或者计算式有误
		err = errors.New("计算式未结束或者计算式有误")
	}
	return
}

// 比较运算符优先级，return >0,则s1高于s2; =0,则s1、s2相同; <0,则s1低于s2
func priorityCompare(s1, s2 byte) int {
	level := func(sym byte) int {
		lvl := 0
		switch sym {
		case '+', '-':
			lvl = 1
		case '*', '/':
			lvl = 2
		}
		return lvl
	}

	return level(s1) - level(s2)
}

func Calculate(str []byte) (result float64, err error) {
	cal := NewCalculator()
	sNum, num := make([]byte, 0, 100), 0.0

	for idx := 0; idx < len(str); idx++ {
		c := str[idx]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sNum = append(sNum, c)
		case '+', '-', '*', '/':
			if (len(sNum) > 0 && num != 0.0) || (len(sNum) == 0 && num == 0.0) {
				err = errors.New("计算式有误")
				return
			}

			if len(sNum) > 0 {
				strNum := string(sNum)
				intNum, _ := strconv.Atoi(strNum)
				num = float64(intNum)
			}
			cal.Step(num, c)
			sNum, num = make([]byte, 0, 100), 0.0
		default:
			err = errors.New(fmt.Sprintf("无效符号： %s", string(c)))
			return
		}
	}

	// 扫描结束
	if (len(sNum) > 0 && num != 0.0) || (len(sNum) == 0 && num == 0.0) {
		err = errors.New("计算式有误")
		return
	}

	if len(sNum) > 0 {
		strNum := string(sNum)
		intNum, _ := strconv.Atoi(strNum)
		num = float64(intNum)
	}
	cal.Step(num, '\n')
	return cal.Result()
}
