package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

func getInt(text []rune) int {
	val, _ := strconv.ParseInt(string(text), 2, 0)
	return int(val)
}

type IntStack []int64

func (is *IntStack) isEmpty() bool {
	return len(*is) == 0
}

func (is *IntStack) push(a int64) {
	*is = append(*is, a)
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

type Trans struct {
	Message      []Packet
	versionCount int
	streamPtr    int
	stackDepth   int
}

func (t *Trans) process(text []rune) {
	var p Packet
	p.Tx = t
	p.process(text)
	p.Tx.Message = append(p.Tx.Message, p)
}

type Packet struct {
	Version    []rune
	Type       []rune
	Data       []rune
	SubPacket  []Packet
	DataLength int
	LengthID   int
	Tx         *Trans
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

func (p *Packet) setVersion(text []rune) {
	for i := 0; i < 3; i++ {
		p.Version = append(p.Version, text[p.Tx.streamPtr])
		p.Tx.streamPtr++
	}
	p.Tx.versionCount += getInt(p.Version)
}

func (p *Packet) setType(text []rune) {
	for i := 0; i < 3; i++ {
		p.Type = append(p.Type, text[p.Tx.streamPtr])
		p.Tx.streamPtr++
	}
}

func (p *Packet) process(text []rune) {
	if p.Tx.streamPtr >= len(text) {
		fmt.Printf("Total version count is %d\n", p.Tx.versionCount)
	} else if getInt(text[p.Tx.streamPtr:]) != 0 {
		fmt.Printf("Stream position (%d/%d) ", p.Tx.streamPtr, len(text))
		p.setVersion(text)
		fmt.Printf("new packet version %d, ", p.getVersion())
		p.setType(text)
		fmt.Printf("type %d, depth %d\n", p.getType(), p.Tx.stackDepth)
		if p.getType() == 4 {
			p.processLiteral(text)
		} else {
			p.processOperator(text)
		}
	}
}

func (p *Packet) processLiteral(text []rune) {
	r := text[p.Tx.streamPtr]
	p.Tx.streamPtr++

	if r == '1' {
		// multi-frame payload
		for i := 0; i < 4; i++ {
			p.Data = append(p.Data, text[p.Tx.streamPtr])
			p.Tx.streamPtr++
			p.DataLength++
		}
		p.processLiteral(text)
	} else {
		// last frame
		// Part 5-bit payload
		for i := 0; i < 4; i++ {
			p.Data = append(p.Data, text[p.Tx.streamPtr])
			p.Tx.streamPtr++
			p.DataLength++
		}
		fmt.Printf("\tData value is %d\n", getInt(p.Data))
	}
}

func (p *Packet) processOperator(text []rune) {
	r := text[p.Tx.streamPtr]
	p.Tx.streamPtr++

	if r == '0' {
		p.LengthID = 0
		p.processOpType0(text)
	} else {
		p.LengthID = 1
		p.processOpType1(text)
	}
}

func (p *Packet) processOpType0(text []rune) {
	// extract 15 bits for length
	var length []rune
	for i := 0; i < 15; i++ {
		length = append(length, text[p.Tx.streamPtr])
		p.Tx.streamPtr++
	}
	opLength := getInt(length)

	// Process the rest of the bits until opLength
	count := 0
	p.Tx.stackDepth++
	currentPtr := p.Tx.streamPtr
	for p.Tx.streamPtr-currentPtr < opLength {
		fmt.Printf("Subpacket %d at depth %d\n", count+1, p.Tx.stackDepth)
		var sp Packet
		sp.Tx = p.Tx
		sp.process(text)
		count++
		p.SubPacket = append(p.SubPacket, sp)
	}
	p.Tx.stackDepth--
}

func (p *Packet) processOpType1(text []rune) {
	// 11 bits for number of subpackets
	var length []rune
	for i := 0; i < 11; i++ {
		length = append(length, text[p.Tx.streamPtr])
		p.Tx.streamPtr++
	}
	spCount := getInt(length)

	count := 0
	fmt.Printf("%d Subpackets to process\n", spCount)
	p.Tx.stackDepth++
	for count < int(spCount) {
		//	for len(text) > 10 {
		fmt.Printf("Subpacket %d at depth %d\n", count+1, p.Tx.stackDepth)
		var sp Packet
		sp.Tx = p.Tx
		sp.process(text)
		p.SubPacket = append(p.SubPacket, sp)
		count++
	}
	p.Tx.stackDepth--
}

func (p *Packet) eval() (ret int64) {
	pType := p.getType()
	//	l := len(p.SubPacket)
	switch pType {
	case 0:
		// sum
		//fmt.Printf("+ with %d subpackets.\n", l)
		var tmpStack IntStack
		var acc int64 = 0
		for _, sp := range p.SubPacket {
			tmpStack.push(sp.eval())
		}
		for !tmpStack.isEmpty() {
			i, _ := tmpStack.pop()
			acc += i
		}
		ret = acc
	case 1:
		// product
		//	fmt.Printf("* with %d subpackets.\n", l)
		var tmpStack IntStack
		var acc int64 = 1
		p.Tx.stackDepth++
		for _, sp := range p.SubPacket {
			tmpStack.push(sp.eval())
		}
		for !tmpStack.isEmpty() {
			i, _ := tmpStack.pop()
			acc *= i
		}
		ret = acc
	case 2:
		// minimum
		// fmt.Printf("min with %d subpackets.\n", l)
		var tmpStack IntStack
		p.Tx.stackDepth++
		for _, sp := range p.SubPacket {
			tmpStack.push(sp.eval())
		}
		var acc int64 = tmpStack[0]
		for !tmpStack.isEmpty() {
			i, _ := tmpStack.pop()
			if i < acc {
				acc = i
			}
		}
		ret = acc
	case 3:
		// maximum
		// fmt.Printf("max with %d subpackets.\n", l)
		var tmpStack IntStack
		p.Tx.stackDepth++
		for _, sp := range p.SubPacket {
			tmpStack.push(sp.eval())
		}
		var acc int64 = tmpStack[0]
		for !tmpStack.isEmpty() {
			i, _ := tmpStack.pop()
			if i > acc {
				acc = i
			}
		}
		ret = acc
	case 4:
		// literal
		val := int64(getInt(p.Data))
		//	fmt.Printf("Literal value %d returning\n", val)
		ret = val
	case 5:
		// greater than
		//		fmt.Printf("> with %d subpackets:\n", l)
		var tmpStack IntStack
		p.Tx.stackDepth++
		for _, sp := range p.SubPacket {
			tmpStack.push(sp.eval())
		}
		t1, _ := tmpStack.pop()
		t2, _ := tmpStack.pop()
		if t2 > t1 {
			ret = 1
		} else {
			ret = 0
		}
	case 6:
		// less than
		//fmt.Printf("< with %d subpackets:\n", l)
		var tmpStack IntStack
		p.Tx.stackDepth++
		for _, sp := range p.SubPacket {
			tmpStack.push(sp.eval())
		}
		t1, _ := tmpStack.pop()
		t2, _ := tmpStack.pop()
		if t2 < t1 {
			ret = 1
		} else {
			ret = 0
		}
	case 7:
		// equal to
		//		fmt.Printf("== with %d subpackets:\n", l)
		var tmpStack IntStack
		p.Tx.stackDepth++
		for _, sp := range p.SubPacket {
			tmpStack.push(sp.eval())
		}
		t1, _ := tmpStack.pop()
		t2, _ := tmpStack.pop()
		if t1 == t2 {
			ret = 1
		} else {
			ret = 0
		}
	default:
		ret = -1
	}
	return
}

func part1(text []rune) Trans {
	defer utils.TimeTrack(time.Now(), "Part 1")
	var packets []Packet
	transmission := Trans{packets, 0, 0, 0}
	transmission.process(text)
	fmt.Printf("Total version is %d\n", transmission.versionCount)

	return transmission
}

func part2(t Trans) {
	defer utils.TimeTrack(time.Now(), "Part 2")
	fmt.Printf("Evaluated packet is ")
	fmt.Println(t.Message[0].eval())
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
