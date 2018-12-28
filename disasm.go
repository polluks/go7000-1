package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	mnemonicsMatrix := [...]string{
		// 0x00 - 0x0F
		// 0x00, 0x01, 0x02, 0x03,
		"NOP", "ILL", "OUTL BUS,A", "ADD A,#",
		// 0x04, 0x05, 0x06, 0x07,
		"JMP $0", "EN I", "ILL", "DEC A",
		// 0x08, 0x09, 0x0A, 0x0B,
		"INS A,BUS", "IN A,P1", "IN A,P2", "ILL",
		// 0x0C, 0x0D, 0x0E, 0x0F,
		"MOVD A,P4", "MOVD A,P5", "MOVD A,P6", "MOVD A,P7",

		// 0x10 - 0x1F
		// 0x10, 0x11, 0x12, 0x13,
		"INC @R0", "INC @R1", "JB0 $", "ADDC A,#",
		// 0x14, 0x15, 0x16, 0x17,
		"CALL $", "DIS I", "JTF $", "INC A",
		// 0x18, 0x19, 0x1A, 0x1B,
		"INC R0", "INC R1", "INC R2", "INC R3",
		// 0x1C, 0x1D, 0x1E, 0x1F,
		"INC R4", "INC R5", "INC R6", "INC R7",

		// 0x20 - 0x2F
		// 0x20, 0x21, 0x22, 0x23,
		"XCH A,@R0", "XCH A,@R1", "ILL", "MOV A,#",
		// 0x24, 0x25, 0x26, 0x27,
		"JMP $1", "EN TCNTI", "JNT0 $", "CLR A",
		// 0x28, 0x29, 0x2A, 0x2B,
		"XCH A,R0", "XCH A,R1", "XCH A,R2", "XCH A,R3",
		// 0x2C, 0x2D, 0x2E, 0x2F,
		"XCH A,R4", "XCH A,R5", "XCH A,R6", "XCH A,R7",

		// 0x30 - 0x3F
		// 0x30, 0x31, 0x32, 0x33,
		"XCHD A,@R0", "XCHD A,@R1", "JB1 $", "ILL",
		// 0x34, 0x35, 0x36, 0x37,
		"CALL $", "DIS TCNTI", "JT0 $", "CPL A",
		// 0x38, 0x39, 0x3A, 0x3B,
		"ILL", "OUTL P1,A", "OUTL P2,A", "ILL",
		// 0x3C, 0x3D, 0x3E, 0x3F,
		"MOVD P4,A", "MOVD P5,A", "MOVD P6,A", "MOVD P7,A",

		// 0x40 - 0x4F
		// 0x40, 0x41, 0x42, 0x43,
		"ORL A,@R0", "ORL A,@R1", "MOV A,T", "ORL A,#",
		// 0x44, 0x45, 0x46, 0x47,
		"JMP $2", "STRT CNT", "JNT1 $", "SWAP",
		// 0x48, 0x49, 0x4A, 0x4B,
		"ORL A,R0", "ORL A,R1", "ORL A,R2", "ORL A,R3",
		// 0x4C, 0x4D, 0x4E, 0x4F,
		"ORL A,R4", "ORL A,R5", "ORL A,R6", "ORL A,R7",

		// 0x50 - 0x5F
		// 0x50, 0x51, 0x52, 0x53,
		"ANL A,@R0", "ANL A,@R1", "JB2 $", "ANL A,#",
		// 0x54, 0x55, 0x56, 0x57,
		"CALL $", "STRT T", "JT1 $", "ILL",
		// 0x58, 0x59, 0x5A, 0x5B,
		"ANL A,R0", "ANL A,R1", "ANL A,R2", "ANL A,R3",
		// 0x5C, 0x5D, 0x5E, 0x5F,
		"ANL A,R4", "ANL A,R5", "ANL A,R6", "ANL A,R7",

		// 0x60 - 0x6F
		// 0x60, 0x61, 0x62, 0x63,
		"ADD A,@R0", "ADD A,@R1", "MOV T,A", "ILL",
		// 0x64, 0x65, 0x66, 0x67,
		"JMP $3", "STOP TCNT", "ILL", "RRC A",
		// 0x68, 0x69, 0x6A, 0x6B,
		"ADD A,R0", "ADD A,R1", "ADD A,R2", "ADD A,R3",
		// 0x6C, 0x6D, 0x6E, 0x6F,
		"ADD A,R4", "ADD A,R5", "ADD A,R6", "ADD A,R7",

		// 0x70 - 0x7F
		// 0x70, 0x71, 0x72, 0x73,
		"ADDC A,@R0", "ADDC A,@R1", "JB3 $", "ILL",
		// 0x74, 0x75, 0x76, 0x77,
		"CALL $", "ENT0 CLK", "JF1 $ ", "RR A",
		// 0x78, 0x79, 0x7A, 0x7B,
		"ADDC A,R0", "ADDC A,R1", "ADDC A,R2", "ADDC A,R3",
		// 0x7C, 0x7D, 0x7E, 0x7F,
		"ADDC A,R4", "ADDC A,R5", "ADDC A,R6", "ADDC A,R7",

		// 80-8F
		// 0x80, 0x81, 0x82, 0x83,
		"MOVX A,@R0", "MOVX A,@R1", "ILL", "RET",
		// 0x84, 0x85, 0x86, 0x87,
		"JMP $4", "CLR F0", "JNI $", "ILL",
		// 0x88, 0x89, 0x8A, 0x8B,
		"ORL BUS,#", "ORL P1,#", "ORL P2,#", "ILL",
		// 0x8C, 0x8D, 0x8E, 0x8F,
		"ORLD P4,A", "ORLD P5,A", "ORLD P6,A", "ORLD P7,A",

		// 0x90 - 0x9F
		// 0x90, 0x91, 0x92, 0x93,
		"MOVX @R0,A", "MOVX @R1,A", "JB4 $", "RETR",
		// 0x94, 0x95, 0x96, 0x97,
		"CALL $", "CPL F0", "JNZ $", "CLR C",
		// 0x98, 0x99, 0x9A, 0x9B,
		"ANL BUS,#", "ANL P1,#", "ANL P2,#", "ILL",
		// 0x9C, 0x9D, 0x9E, 0x9F,
		"ANLD P4,A", "ANLD P5,A", "ANLD P6,A", "ANLD P7,A",

		// 0xA0 - 0xAF
		// 0xA0, 0xA1, 0xA2, 0xA3,
		"MOV @R0,A", "MOV @R1,A", "ILL", "MOVP A,@A",
		// 0xA4, 0xA5, 0xA6, 0xA7,
		"JMP $5", "CLR F1", "ILL", "CPL C",
		// 0xA8, 0xA9, 0xAA, 0xAB,
		"MOV R0,A", "MOV R1,A", "MOV R2,A", "MOV R3,A",
		// 0xAC, 0xAD, 0xAE, 0xAF,
		"MOV R4,A", "MOV R5,A", "MOV R6,A", "MOV R7,A",

		// 0xB0 - 0xBF
		// 0xB0, 0xB1, 0xB2, 0xB3,
		"MOV @R0,#", "MOV @R1,#", "JB5 $", "JMPP @A",
		// 0xB4, 0xB5, 0xB6, 0xB7,
		"CALL $", "CPL F1", "JF0 $", "ILL",
		// 0xB8, 0xB9, 0xBA, 0xBB,
		"MOV R0,#", "MOV R1,#", "MOV R2,#", "MOV R3,#",
		// 0xBC, 0xBD, 0xBE, 0xBF,
		"MOV R4,#", "MOV R5,#", "MOV R6,#", "MOV R7,#",

		// 0xC0 - 0xCF
		// 0xC0, 0xC1, 0xC2, 0xC3,
		"ILL", "ILL", "ILL", "ILL",
		// 0xC4, 0xC5, 0xC6, 0xC7,
		"JMP $6", "SEL RB0", "JZ $", "MOV A,PSW",
		// 0xC8, 0xC9, 0xCA, 0xCB,
		"DEC R0", "DEC R1", "DEC R2", "DEC R3",
		// 0xCC, 0xCD, 0xCE, 0xCF,
		"DEC R4", "DEC R5", "DEC R6", "DEC R7",

		// 0xD0 - 0xDF
		// 0xD0, 0xxD1, 0xD2, 0xD3,
		"XRL A,@R0", "XRL A,@R1", "JB6 $", "XRL A,#",
		// 0xD4,0xD5,0xD6,0xD7,
		"CALL $", "SEL RB1", "ILL", "MOV PSW,A",
		// 0xD8, 0xD9, 0xDA, 0xDB,
		"XRL A,R0", "XRL A,R1", "XRL A,R2", "XRL A,R3",
		// 0xDC, 0xDD, 0xDE, 0xDF,
		"XRL A,R4", "XRL A,R5", "XRL A,R6", "XRL A,R7",

		// 0xE0 - 0xEF
		// 0xE0, 0xE1, 0xE2, 0xE3,
		"ILL", "ILL", "ILL", "MOVP3 A,@A",
		// 0xE4, 0xE5, 0xE6, 0xE7,
		"JMP $7", "SEL MB0", "JNC $", "RL A",
		// 0xE8, 0xE9, 0xEA, 0xEB,
		"DJNZ R0,$", "DJNZ R1,$", "DJNZ R2,$", "DJNZ R3,$",
		// 0xEC, 0xED, 0xEE, 0xEF,
		"DJNZ R4,$", "DJNZ R5,$", "DJNZ R6,$", "DJNZ R7,$",

		// 0xF0 - 0xFF
		// 0xF0, 0xF1, 0xF2, 0xF3,
		"MOV A,@R0", "MOV A,@R1", "JB7 $", "ILL",
		// 0xF4, 0xF5, 0xF6, 0xF7,
		"CALL $", "SEL MB1", "JC $", "RLC A",
		// 0xF8, 0xF9, 0xFA, 0xFB,
		"MOV A,R0", "MOV A,R1", "MOV A,R2", "MOV A,R3",
		// 0xFC, 0xFD, 0xFE, 0xFF,
		"MOV A,R4", "MOV A,R5", "MOV A,R6", "MOV A,R7",
	}

	opCodeLength := [...]uint8{
		//  00,01,02,03,04,05,06,07,08,09,0A,0B,0C,0D,0E,0F
		0, 3, 0, 1, 1, 0, 3, 0, 0, 0, 0, 3, 0, 0, 0, 0,
		//  10,11,12,13,14,15,16,17,18,19,1A,1B,1C,1D,1E,1F
		0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  20,21,22,23,24,25,26,27,28,29,2A,2B,2C,2D,2E,2F
		0, 0, 3, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  30,31,32,33,34,35,36,37,38,39,3A,3B,3C,3D,3E,3F
		0, 0, 1, 0, 1, 0, 1, 0, 3, 0, 0, 0, 2, 0, 0, 0,
		//  40,41,42,43,44,45,46,47,48,49,4A,4B,4C,4D,4E,4F
		0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  50,51,52,53,54,55,56,57,58,59,5A,5B,5C,5D,5E,5F
		0, 0, 1, 1, 1, 0, 1, 3, 0, 0, 0, 0, 0, 0, 0, 0,
		//  60,61,62,63,64,65,66,67,68,69,6A,6B,6C,6D,6E,6F
		0, 0, 0, 3, 1, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  70,71,72,73,74,75,76,77,78,79,7A,7B,7C,7D,7E,7F
		0, 0, 1, 3, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  80,81,82,83,84,85,86,87,88,89,8A,8B,8C,8D,8E,8F
		0, 0, 0, 0, 1, 0, 1, 3, 1, 1, 1, 3, 0, 0, 0, 0,
		//  90,91,92,93,94,95,96,97,98,99,9A,9B,9C,9D,9E,9F
		0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 3, 0, 0, 0, 0,
		//  A0,A1,A2,A3,A4,A5,A6,A7,A8,A9,AA,AB,AC,AD,AE,AF
		0, 0, 3, 0, 1, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  B0,B1,B2,B3,B4,B5,B6,B7,B8,B9,BA,BB,BC,BD,BE,BF
		1, 1, 1, 0, 1, 0, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
		//  C0,C1,C2,C3,C4,C5,C6,C7,C8,C9,CA,CB,CC,CD,CE,CF
		3, 3, 3, 3, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  D0,D1,D2,D3,D4,D5,D6,D7,D8,D9,DA,DB,DC,DD,DE,DF
		0, 0, 1, 1, 1, 1, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  E0,E1,E2,E3,E4,E5,E6,E7,E8,E9,EA,EB,EC,ED,EE,EF
		3, 3, 3, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		//  F0,F1,F2,F3,F4,F5,F6,F7,F8,F9,FA,FB,FC,FD,FE,FF
		0, 0, 1, 3, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	//fmt.Println(mnemonicsMatrix)
	//fmt.Println(opCodeLength)

	binary := readInputFile()

	for index := 0; index <= len(binary)-1; index++ {
		opCode := binary[index]

		address := fmt.Sprintf("%04x", index)
		opCodeHex := fmt.Sprintf("%02x", opCode)

		fmt.Print(address, " : ")
		fmt.Print(opCodeHex, " ")

		//fmt.Print(mnemonic)

		data := ""
		data1 := "           "
		switch {
		case opCodeLength[opCode] == 0:

		case opCodeLength[opCode] == 1:
			data1 = fmt.Sprintf("%02X         ", binary[index+1])
			data = fmt.Sprintf("%02X", binary[index+1])
			index += 1
		case opCodeLength[opCode] == 3:
			data1 = fmt.Sprintf("%02X %02X %02X   ",
				binary[index+1], binary[index+2], binary[index+3])
			data = fmt.Sprintf("%02X %02X %02X",
				binary[index+1], binary[index+2], binary[index+3])
			index += 3
		default:
			fmt.Println("Error with opCodeLength")
		}

		fmt.Print(data1)
		mnemonic := fmt.Sprintf("%s", mnemonicsMatrix[opCode])
		fmt.Print(strings.ToLower(mnemonic))
		fmt.Println(data)
	}

	for index := 0; index <= len(mnemonicsMatrix)-1; index++ {
		fmt.Printf("{\"op\":\"%02x\", \"mn\":\"%s\", \"ol\":%d},\n", index, mnemonicsMatrix[index], opCodeLength[index])
	}

	type opCodesMatrix struct {
		Op string // OpCode
		Mn string // Mnemonic
		Ol int    // OpCodeLength
	}
	birdJson := `[
				{"op":"00", "mn":"SEL MB1", "ol":1},
				{"op":"01", "mn":"MOV A,R7", "ol":1},
				{"op":"fd", "mn":"MOV A,R5", "ol":0},
				{"op":"fe", "mn":"MOV A,R6", "ol":0},
				{"op":"ff", "mn":"MOV A,R7", "ol":0}
				]`
	var opCodes []opCodesMatrix
	json.Unmarshal([]byte(birdJson), &opCodes)
	fmt.Printf("OpCodeMatrix : %+v\n", opCodes)
	myOpCode := opCodes[1]
	fmt.Println(myOpCode.Op, myOpCode.Mn, myOpCode.Ol)
}

func readInputFile() []byte {
	binaryFileName := os.Args[1]
	fmt.Println(binaryFileName)

	binary, err := ioutil.ReadFile(binaryFileName)
	check(err)

	return binary
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
