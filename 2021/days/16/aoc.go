package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dkim/aoc/2021/utils"
)

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

func (p *Packet) getPacketLength() int {
	sum := 0
	sum += len(p.Version)
	sum += len(p.Type)
	sum += len(p.Data)
	return sum
}

func (p *Packet) getVersion() int {
	if i, err := strconv.ParseInt(string(p.Version), 2, 0); err == nil {
		fmt.Printf("Packet version %s: %d\n", string(p.Version), i)
		return int(i)
	} else {
		return -1
	}
}

func (p *Packet) getType() int {
	if i, err := strconv.ParseInt(string(p.Type), 2, 0); err == nil {
		/*
			fmt.Printf("Packet type %s: %d ", string(p.Type), i)
			if i == 4 {
				fmt.Printf("Literal\n")
			} else {
				fmt.Printf("Operator\n")
			}
		*/
		return int(i)
	} else {
		return -1
	}
}

func (p *Packet) setVersion(text []rune) (int, []rune) {
	for i := 0; i < 3; i++ {
		p.Version = append(p.Version, text[i])
	}
	return p.getVersion(), text[3:]
}

func (p *Packet) setType(text []rune) (int, []rune) {
	for i := 0; i < 3; i++ {
		p.Type = append(p.Type, text[i])
	}
	return p.getType(), text[3:]
}

func (p *Packet) processLiteral(text []rune, pad bool) ([]rune, int) {
	ver := 0
	temp := 0
	if text[0] == '1' {
		for i := 1; i < 5; i++ {
			p.Data = append(p.Data, text[i])
		}
		p.DataLength += 5

		text, temp = p.processLiteral(text[5:], pad)
		ver += temp
	} else {
		// Part 5-bit payload
		for i := 1; i < 5; i++ {
			p.Data = append(p.Data, text[i])
		}
		p.DataLength += 5
		// Identify padding
		if pad {
			p.Padding = p.DataLength % 4
		}
		return text[5+p.Padding:], ver
	}
	// Nothing more to do
	fmt.Println(p)
	return text, ver
}

func (p *Packet) processOpType0(text []rune) ([]rune, int) {
	ver := 0
	if p.LengthID != 0 {
		return nil, ver
	}

	// extract 15 bits for length
	var length []rune
	for i := 1; i < 16; i++ {
		length = append(length, text[i])
	}
	opLength, _ := strconv.ParseInt(string(length), 2, 0)
	text = text[16:]

	// Process the rest of the bits until opLength
	count := 0
	for count < int(opLength) {
		var sp Packet
		var t int
		var temp int
		temp, text = sp.setVersion(text)
		ver += temp
		count += 3
		t, text = sp.setType(text)
		count += 3
		if t == 4 {
			text, temp = sp.processLiteral(text, false)
			ver += temp
		} else {
			text, temp = sp.processOperator(text)
			ver += temp
		}
		count += sp.DataLength
		p.SubPacket = append(p.SubPacket, sp)
	}
	return text, ver

}

func (p *Packet) processOpType1(text []rune) ([]rune, int) {
	ver := 0
	if p.LengthID != 1 {
		return nil, ver
	}

	// 11 bits for number of subpackets
	var length []rune
	for i := 1; i < 12; i++ {
		length = append(length, text[i])
	}

	spCount, _ := strconv.ParseInt(string(length), 2, 0)
	text = text[12:]
	count := 0
	for count < int(spCount) {
		var sp Packet
		var t int
		var temp int
		temp, text = sp.setVersion(text)
		ver += temp
		t, text = sp.setType(text)
		if t == 4 {
			text, temp = sp.processLiteral(text, false)
			ver += temp
		} else {
			text, temp = sp.processOperator(text)
			ver += temp
		}
		fmt.Printf("Sub packet\n")
		p.SubPacket = append(p.SubPacket, sp)
		count++
	}
	return text, ver
}

func (p *Packet) processOperator(text []rune) ([]rune, int) {
	ver := 0
	temp := 0
	if text[0] == '0' {
		p.LengthID = 0
		text, temp = p.processOpType0(text)
		ver += temp
	} else {
		p.LengthID = 1
		text, temp = p.processOpType1(text)
		ver += temp
	}
	return text, ver
}

func (t *Trans) process(text []rune) {
	var p Packet
	done := true
	if len(text) == 0 {
		fmt.Printf("Total version count is %d\n", t.versionCount)
	} else {
		for _, r := range text {
			if r == '1' {
				done = false
				break
			}
		}
	}
	if !done {
		num, text := p.setVersion(text)
		t.versionCount += num
		num, text = p.setType(text)
		temp := 0
		if num == 4 {
			//fmt.Println("Literal packet")
			text, temp = p.processLiteral(text, true)
			t.Message = append(t.Message, p)
			t.versionCount += temp
			t.process(text)
		} else {
			//fmt.Println("Operator packet")
			text, temp = p.processOperator(text)
			t.Message = append(t.Message, p)
			t.versionCount += temp
			t.process(text)
		}
	}
	fmt.Println(t.versionCount)
}

func part1(text []rune) {
	defer utils.TimeTrack(time.Now(), "Part 1")
	var packets []Packet
	transmission := Trans{packets, 0}
	transmission.process(text)

	/*
		version := 0
			fmt.Println("Counting")
			for i, p := range transmission.Message {
				fmt.Printf("Packet %d - ", i)
				t := p.getType()
				if t == 4 {
					version += p.getVersion()
				} else {
					version += p.getVersion()
					for _, sp := range p.SubPacket {
						version += sp.getVersion()
					}
				}
			}
			fmt.Printf("Version total is %d\n", version)
	*/
}

func main() {
	text := utils.ReadInput(0)
	runes := []rune(text[0])
	var out []string
	for _, i := range runes {
		num, _ := strconv.ParseUint(string(i), 16, 4)
		binStr := fmt.Sprintf("%04s", strconv.FormatUint(num, 2))
		out = append(out, binStr)
	}
	realInput := fmt.Sprintf(strings.Join(out, ""))
	fmt.Println(realInput)
	part1([]rune(realInput))

}
