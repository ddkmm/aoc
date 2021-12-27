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

func (p *Packet) getPacketLength() int {
	sum := 0
	sum += len(p.Version)
	sum += len(p.Type)
	sum += len(p.Data)
	return sum
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
	totalVersion += getInt(p.Version)
	totalPackets++
	fmt.Printf("Packet %d, total version %d\n", totalPackets, totalVersion)
	return p.getVersion(), text[3:]
}

func (p *Packet) setType(text []rune) (int, []rune) {
	for i := 0; i < 3; i++ {
		p.Type = append(p.Type, text[i])
	}
	return p.getType(), text[3:]
}

func (p *Packet) processLiteral(text []rune) []rune {
	if len(text) < 11 {
		return nil
	}
	fmt.Printf("ProcessLiteral, %d length\n", len(text))
	if text[0] == '1' {
		for i := 1; i < 5; i++ {
			p.Data = append(p.Data, text[i])
		}
		p.DataLength += 5
		text = p.processLiteral(text[5:])
	} else {
		// Part 5-bit payload
		for i := 1; i < 5; i++ {
			p.Data = append(p.Data, text[i])
		}
		p.DataLength += 5
		return text[5:]
	}
	// Nothing more to do
	//fmt.Println(p)
	return text
}

func (p *Packet) processOpType0(text []rune) []rune {
	fmt.Printf("Processing type 0 operation: %d in text\n", len(text))
	if p.LengthID != 0 {
		return nil
	}
	if len(text) < 22 {
		return nil
	}

	// extract 15 bits for length
	var length []rune
	for i := 1; i < 16; i++ {
		length = append(length, text[i])
	}
	opLength := getInt(length)
	fmt.Printf("Op 0 has %d length\n", opLength)
	text = text[16:]
	fmt.Printf("Removed 15 bits for lenth. Now text length is %d\n", len(text))

	// Process the rest of the bits until opLength
	count := 0
	//for count < int(opLength) {
	for len(text) > 11 {
		var sp Packet
		var t int
		_, text = sp.setVersion(text)
		count += 3
		t, text = sp.setType(text)
		count += 3
		fmt.Printf("%d length remaining\n", len(text))
		if t == 4 {
			fmt.Println("\tStarting literal subpacket")
			text = sp.processLiteral(text)
			fmt.Println("\tEnding literal subpacket")
		} else {
			fmt.Println("\tStarting operator subpacket")
			text = sp.processOperator(text)
			fmt.Println("\tEnding operator subpacket")
		}
		count += sp.DataLength
		p.SubPacket = append(p.SubPacket, sp)
		fmt.Printf("Subpacket finsished. %d processed out of %d from buffer\n", count, opLength)
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
	if len(text) < 23 {
		return nil
	}
	fmt.Printf("Processing type 1 operation: %d in text\n", len(text))

	// 11 bits for number of subpackets
	var length []rune
	for i := 1; i < 12; i++ {
		length = append(length, text[i])
	}

	spCount := getInt(length)
	fmt.Printf("%d packets\n", spCount)
	text = text[12:]
	count := 0
	//	for count < int(spCount) {
	for len(text) > 11 {
		var sp Packet
		var t int
		_, text = sp.setVersion(text)
		t, text = sp.setType(text)
		if t == 4 {
			text = sp.processLiteral(text)
		} else {
			text = sp.processOperator(text)
		}
		fmt.Printf("Sub packet\n")
		p.SubPacket = append(p.SubPacket, sp)
		count++
	}
	return text
}

func (p *Packet) processOperator(text []rune) []rune {
	if len(text) == 0 {
		return nil
	}
	fmt.Printf("ProcessOperator, %d length\n", len(text))
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
			//fmt.Println("Literal packet")
			text = p.processLiteral(text)
			t.Message = append(t.Message, p)
			t.versionCount += temp
			t.process(text)
		} else {
			//fmt.Println("Operator packet")
			text = p.processOperator(text)
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
	part1([]rune(realInput))

}
