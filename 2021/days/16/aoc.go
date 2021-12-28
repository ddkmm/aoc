package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

var totalPackets int = 0
var totalVersion int = 0

type Trans struct {
	Message      []Packet
	versionCount int
}

type Packet struct {
	Version    []rune
	Type       []rune
	Data       []rune
	SubPacket  []Packet
	DataLength int
	LengthID   int
	Padding    int
}

func (p *Packet) getVersion() int {
	if i, err := strconv.ParseInt(string(p.Version), 2, 0); err == nil {
		return int(i)
	} else {
		return -1
	}
}

func (p *Packet) getType() int {
	if i, err := strconv.ParseInt(string(p.Type), 2, 0); err == nil {
		return int(i)
	} else {
		return -1
	}
}

func (p *Packet) setVersion(text []rune) (int, []rune) {
	for i := 0; i < 3; i++ {
		p.Version = append(p.Version, text[i])
	}
	totalVersion += getInt(p.Version)
	totalPackets++
	fmt.Printf("Packet %d, total version %d\n", totalPackets, totalVersion)
	return p.getVersion(), text[3:]
}

func getOp(typeID int) (op string) {
	switch typeID {
	case 0:
		op = "sum"
	case 1:
		op = "product"
	case 2:
		op = "minimum"
	case 3:
		op = "maximum"
	case 4:
		op = "literal"
	case 5:
		op = "greater than"
	case 6:
		op = "less than"
	case 7:
		op = "equal to"
	}
	fmt.Println(op)
	return op
}

func (p *Packet) setType(text []rune) (int, []rune) {
	for i := 0; i < 3; i++ {
		p.Type = append(p.Type, text[i])
	}
	fmt.Printf("Type %d: ", getInt(p.Type))
	getOp(getInt(p.Type))
	return p.getType(), text[3:]
}

func (p *Packet) processLiteral(text []rune) []rune {
	if text[0] == '1' {
		// multi-frame payload
		for i := 1; i < 5; i++ {
			p.Data = append(p.Data, text[i])
		}
		p.DataLength += 5
		text = p.processLiteral(text[5:])
	} else {
		// last frame
		// Part 5-bit payload
		for i := 1; i < 5; i++ {
			p.Data = append(p.Data, text[i])
		}
		p.DataLength += 5
		fmt.Printf("\tData value is %d\n", getInt(p.Data))
		return text[5:]
	}
	// Nothing more to do
	//fmt.Println(p)
	return text
}

func (p *Packet) processOpType0(text []rune) []rune {
	if p.LengthID != 0 {
		return nil
	}

	// extract 15 bits for length
	var length []rune
	for i := 1; i < 16; i++ {
		length = append(length, text[i])
	}
	text = text[16:]

	// Process the rest of the bits until opLength
	count := 0
	for len(text) > 10 {
		fmt.Printf("Sub packet %d\n", count+1)
		var sp Packet
		var t int
		_, text = sp.setVersion(text)
		count += 3
		t, text = sp.setType(text)
		count += 3
		if t == 4 {
			text = sp.processLiteral(text)
		} else {
			text = sp.processOperator(text)
		}
		count++
		p.SubPacket = append(p.SubPacket, sp)
	}
	return text
}

func getInt(text []rune) int {
	val, _ := strconv.ParseInt(string(text), 2, 0)
	return int(val)
}

func (p *Packet) processOpType1(text []rune) []rune {
	if p.LengthID != 1 {
		return nil
	}

	// 11 bits for number of subpackets
	var length []rune
	for i := 1; i < 12; i++ {
		length = append(length, text[i])
	}
	spCount := getInt(length)

	text = text[12:]
	count := 0
	//	for count < int(spCount) {
	for len(text) > 10 {
		fmt.Printf("Sub packet %d of %d\n", count+1, spCount)
		var sp Packet
		var t int
		_, text = sp.setVersion(text)
		t, text = sp.setType(text)
		if t == 4 {
			text = sp.processLiteral(text)
		} else {
			text = sp.processOperator(text)
		}
		p.SubPacket = append(p.SubPacket, sp)
		count++
	}
	return text
}

func (p *Packet) processOperator(text []rune) []rune {
	if len(text) == 0 {
		return nil
	}
	if text[0] == '0' {
		p.LengthID = 0
		text = p.processOpType0(text)
	} else {
		p.LengthID = 1
		text = p.processOpType1(text)
	}
	return text
}

func (t *Trans) process(text []rune) {
	var p Packet
	if len(text) == 0 {
		fmt.Printf("Total version count is %d\n", t.versionCount)
	} else if getInt(text) != 0 {
		num, text := p.setVersion(text)
		t.versionCount += num
		num, text = p.setType(text)
		temp := 0
		if num == 4 {
			text = p.processLiteral(text)
			t.Message = append(t.Message, p)
			t.versionCount += temp
			t.process(text)
		} else {
			text = p.processOperator(text)
			t.Message = append(t.Message, p)
			t.versionCount += temp
			t.process(text)
		}
	}
}

func part1(text []rune) Trans {
	defer utils.TimeTrack(time.Now(), "Part 1")
	var packets []Packet
	transmission := Trans{packets, 0}
	transmission.process(text)

	return transmission
}

func (p *Packet) print() {
	if p.getType() != 4 {
		getOp(p.getType())
	} else {
		fmt.Println(getInt(p.Data))
	}
	for _, p := range p.SubPacket {
		p.print()
	}
}

func (p *Packet) makeStack(ops *[]string) {
	if p.getType() != 4 {
		*ops = append(*ops, getOp(p.getType()))
	} else {
		*ops = append(*ops, strconv.FormatInt(int64(getInt(p.Data)), 10))
	}
	for _, p := range p.SubPacket {
		p.makeStack(ops)
	}
}

type Stack []string
type IntStack []int64

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (is *IntStack) isEmpty() bool {
	return len(*is) == 0
}

func (s *Stack) push(a string) {
	*s = append(*s, a)
}

func (is *IntStack) push(a int64) {
	*is = append(*is, a)
}

func (s *Stack) pop() (string, bool) {
	var element string
	if s.isEmpty() {
		return "", false
	} else {
		element = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return element, true
	}
}

func (is *IntStack) pop() (int64, bool) {
	var element int64
	if is.isEmpty() {
		return 0, false
	} else {
		element = (*is)[len(*is)-1]
		*is = (*is)[:len(*is)-1]
		return element, true
	}
}

func processStack(s Stack) {
	// Make a int stack to hold values
	// have a temp accumulator
	// When an operand is on top of the stack, pop the values and
	// put the result in the acc
	// then push the acc back onto the stack as a string
	var val IntStack
	var acc int64 = 0
	for !s.isEmpty() {
		temp, res := s.pop()
		if res {
			switch temp {
			case "sum":
				acc = 0
				for !val.isEmpty() {
					tmp, _ := val.pop()
					acc += tmp
				}
				s.push(strconv.FormatInt(acc, 10))
			case "product":
				acc = 1
				for !val.isEmpty() {
					tmp, _ := val.pop()
					acc *= tmp
				}
				s.push(strconv.FormatInt(acc, 10))
			case "minimum":
				acc = val[0]
				for !val.isEmpty() {
					tmp, _ := val.pop()
					if acc > tmp {
						acc = tmp
					}
				}
				s.push(strconv.FormatInt(acc, 10))
			case "maximum":
				acc = val[0]
				for !val.isEmpty() {
					tmp, _ := val.pop()
					if acc < tmp {
						acc = tmp
					}
				}
				s.push(strconv.FormatInt(acc, 10))
			case "greater than":
				t1, _ := val.pop()
				t2, _ := val.pop()
				if t1 > t2 {
					acc = 1
				} else {
					acc = 0
				}
				s.push(strconv.FormatInt(acc, 10))
			case "less than":
				t1, _ := val.pop()
				t2, _ := val.pop()
				if t1 < t2 {
					acc = 1
				} else {
					acc = 0
				}
				s.push(strconv.FormatInt(acc, 10))
			case "equal to":
				t1, _ := val.pop()
				t2, _ := val.pop()
				if t1 == t2 {
					acc = 1
				} else {
					acc = 0
				}
				s.push(strconv.FormatInt(acc, 10))
			default:
				i, _ := strconv.ParseInt(temp, 10, 64)
				val.push(i)
			}
		}
	}
	fmt.Println(val)

}

func part2(t Trans) {
	var ops []string
	for _, p := range t.Message {
		p.makeStack(&ops)
	}
	fmt.Println(ops)
	s := Stack(ops)
	processStack(s)
}

func main() {
	text := utils.ReadInput(1)
	runes := []rune(text[0])
	var out []string
	for _, i := range runes {
		num, _ := strconv.ParseUint(string(i), 16, 4)
		binStr := fmt.Sprintf("%04s", strconv.FormatUint(num, 2))
		out = append(out, binStr)
	}
	realInput := fmt.Sprintf(strings.Join(out, ""))
	fmt.Println(realInput)
	tran := part1([]rune(realInput))
	part2(tran)
}
