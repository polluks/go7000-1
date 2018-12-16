package main

import "fmt"

func main() {
	opCodeMatrix := [...]string{
		// 0x00 - 0x0F
		// 0x00, 0x01, 0x02, 0x03,
		"NOP", "ILL", "OUTLBUS,A", "ADDA,#",
		// 0x04, 0x05, 0x06, 0x07,
		"JMP$0", "ENI", "ILL", "DECA",
		// 0x08, 0x09, 0x0A, 0x0B,
		"INSA,BUS", "INA,P1", "INA,P2", "ILL",
		// 0x0C, 0x0D, 0x0E, 0x0F,
		"MOVDA,P4", "MOVDA,P5", "MOVDA,P6", "MOVDA,P7",

		// 0x10 - 0x1F
		// 0x10, 0x11, 0x12, 0x13,
		"INC@R0", "INC@R1", "JB0$", "ADDCA,#",
		// 0x14, 0x15, 0x16, 0x17,
		"CALL$", "DISI", "JTF$", "INCA",
		// 0x18, 0x19, 0x1A, 0x1B,
		"INCR0", "INCR1", "INCR2", "INCR3",
		// 0x1C, 0x1D, 0x1E, 0x1F,
		"INCR4", "INCR5", "INCR6", "INCR7",

		// 0x20 - 0x2F
		// 0x20, 0x21, 0x22, 0x23,
		"XCHA,@R0", "XCHA,@R1", "ILL", "MOVA,#",
		// 0x24, 0x25, 0x26, 0x27,
		"JMP$1", "ENTCNTI", "JNT0$", "CLRA",
		// 0x28, 0x29, 0x2A, 0x2B,
		"XCHA,R0", "XCHA,R1", "XCHA,R2", "XCHA,R3",
		// 0x2C, 0x2D, 0x2E, 0x2F,
		"XCHA,R4", "XCHA,R5", "XCHA,R6", "XCHA,R7",

		// 0x30 - 0x3F
		// 0x30, 0x31, 0x32, 0x33,
		"XCHDA,@R0", "XCHDA,@R1", "JB1$", "ILL",
		// 0x34, 0x35, 0x36, 0x37,
		"CALL$", "DISTCNTI", "JT0$", "CPLA",
		// 0x38, 0x39, 0x3A, 0x3B,
		"ILL", "OUTLP1,A", "OUTLP2,A", "ILL",
		// 0x3C, 0x3D, 0x3E, 0x3F,
		"MOVDP4,A", "MOVDP5,A", "MOVDP6,A", "MOVDP7,A",

		// 0x40 - 0x4F
		// 0x40, 0x41, 0x42, 0x43,
		"ORLA,@R0", "ORLA,@R1", "MOVA,T", "ORLA,#",
		// 0x44, 0x45, 0x46, 0x47,
		"JMP$2", "STRTCNT", "JNT1$", "SWAP",
		// 0x48, 0x49, 0x4A, 0x4B,
		"ORLA,R0", "ORLA,R1", "ORLA,R2", "ORLA,R3",
		// 0x4C, 0x4D, 0x4E, 0x4F,
		"ORLA,R4", "ORLA,R5", "ORLA,R6", "ORLA,R7",

		// 0x50 - 0x5F
		// 0x50, 0x51, 0x52, 0x53,
		"ANLA,@R0", "ANLA,@R1", "JB2$", "ANLA,#",
		// 0x54, 0x55, 0x56, 0x57,
		"CALL$", "STRTT", "JT1$", "ILL",
		// 0x58, 0x59, 0x5A, 0x5B,
		"ANLA,R0", "ANLA,R1", "ANLA,R2", "ANLA,R3",
		// 0x5C, 0x5D, 0x5E, 0x5F,
		"ANLA,R4", "ANLA,R5", "ANLA,R6", "ANLA,R7",

		// 0x60 - 0x6F
		// 0x60, 0x61, 0x62, 0x63,
		"ADDA,@R0", "ADDA,@R1", "MOVT,A", "ILL",
		// 0x64, 0x65, 0x66, 0x67,
		"JMP$3", "STOPTCNT", "ILL", "RRCA",
		// 0x68, 0x69, 0x6A, 0x6B,
		"ADDA,R0", "ADDA,R1", "ADDA,R2", "ADDA,R3",
		// 0x6C, 0x6D, 0x6E, 0x6F,
		"ADDA,R4", "ADDA,R5", "ADDA,R6", "ADDA,R7",

		// 0x70 - 0x7F
		// 0x70, 0x71, 0x72, 0x73,
		"ADDCA,@R0", "ADDCA,@R1", "JB3$", "ILL",
		// 0x74, 0x75, 0x76, 0x77,
		"CALL$", "ENT0CLK", "JF1$", "RRA",
		// 0x78, 0x79, 0x7A, 0x7B,
		"ADDCA,R0", "ADDCA,R1", "ADDCA,R2", "ADDCA,R3",
		// 0x7C, 0x7D, 0x7E, 0x7F,
		"ADDCA,R4", "ADDCA,R5", "ADDCA,R6", "ADDCA,R7",

		// 80-8F
		// 0x80, 0x81, 0x82, 0x83,
		"MOVXA,@R0", "MOVXA,@R1", "ILL", "RET",
		// 0x84, 0x85, 0x86, 0x87,
		"JMP$4", "CLRF0", "JNI$", "ILL",
		// 0x88, 0x89, 0x8A, 0x8B,
		"ORLBUS,#", "ORLP1,#", "ORLP2,#", "ILL",
		// 0x8C, 0x8D, 0x8E, 0x8F,
		"ORLDP4,A", "ORLDP5,A", "ORLDP6,A", "ORLDP7,A",

		// 0x90 - 0x9F
		// 0x90, 0x91, 0x92, 0x93,
		"MOVX@R0,A", "MOVX@R1,A", "JB4$", "RETR",
		// 0x94, 0x95, 0x96, 0x97,
		"CALL$", "CPLF0", "JNZ$", "CLRC",
		// 0x98, 0x99, 0x9A, 0x9B,
		"ANLBUS,#", "ANLP1,#", "ANLP2,#", "ILL",
		// 0x9C, 0x9D, 0x9E, 0x9F,
		"ANLDP4,A", "ANLDP5,A", "ANLDP6,A", "ANLDP7,A",

		// 0xA0 - 0xAF
		// 0xA0, 0xA1, 0xA2, 0xA3,
		"MOV@R0,A", "MOV@R1,A", "ILL", "MOVPA,@A",
		// 0xA4, 0xA5, 0xA6, 0xA7,
		"JMP$5", "CLRF1", "ILL", "CPLC",
		// 0xA8, 0xA9, 0xAA, 0xAB,
		"MOVR0,A", "MOVR1,A", "MOVR2,A", "MOVR3,A",
		// 0xAC, 0xAD, 0xAE, 0xAF,
		"MOVR4,A", "MOVR5,A", "MOVR6,A", "MOVR7,A",

		// 0xB0 - 0xBF
		// 0xB0, 0xB1, 0xB2, 0xB3,
		"MOV@R0,#", "MOV@R1,#", "JB5$", "JMPP@A",
		// 0xB4, 0xB5, 0xB6, 0xB7,
		"CALL$", "CPLF1", "JF0$", "ILL",
		// 0xB8, 0xB9, 0xBA, 0xBB,
		"MOVR0,#", "MOVR1,#", "MOVR2,#", "MOVR3,#",
		// 0xBC, 0xBD, 0xBE, 0xBF,
		"MOVR4,#", "MOVR5,#", "MOVR6,#", "MOVR7,#",

		// 0xC0 - 0xCF
		// 0xC0, 0xC1, 0xC2, 0xC3,
		"ILL", "ILL", "ILL", "ILL",
		// 0xC4, 0xC5, 0xC6, 0xC7,
		"JMP$6", "SELRB0", "JZ$", "MOVA,PSW",
		// 0xC8, 0xC9, 0xCA, 0xCB,
		"DECR0", "DECR1", "DECR2", "DECR3",
		// 0xCC, 0xCD, 0xCE, 0xCF,
		"DECR4", "DECR5", "DECR6", "DECR7",

		// 0xD0 - 0xDF
		// 0xD0, 0xxD1, 0xD2, 0xD3,
		"XRLA,@R0", "XRLA,@R1", "JB6$", "XRLA,#",
		// 0xD4,0xD5,0xD6,0xD7,
		"CALL$", "SELRB1", "ILL", "MOVPSW,A",
		// 0xD8, 0xD9, 0xDA, 0xDB,
		"XRLA,R0", "XRLA,R1", "XRLA,R2", "XRLA,R3",
		// 0xDC, 0xDD, 0xDE, 0xDF,
		"XRLA,R4", "XRLA,R5", "XRLA,R6", "XRLA,R7",

		// 0xE0 - 0xEF
		// 0xE0, 0xE1, 0xE2, 0xE3,
		"ILL", "ILL", "ILL", "MOVP3A,@A",
		// 0xE4, 0xE5, 0xE6, 0xE7,
		"JMP$7", "SELMB0", "JNC$", "RLA",
		// 0xE8, 0xE9, 0xEA, 0xEB,
		"DJNZR0,$", "DJNZR1,$", "DJNZR2,$", "DJNZR3,$",
		// 0xEC, 0xED, 0xEE, 0xEF,
		"DJNZR4,$", "DJNZR5,$", "DJNZR6,$", "DJNZR7,$",

		// 0xF0 - 0xFF
		// 0xF0, 0xF1, 0xF2, 0xF3,
		"MOVA,@R0", "MOVA,@R1", "JB7$", "ILL",
		// 0xF4, 0xF5, 0xF6, 0xF7,
		"CALL$", "SELMB1", "JC$", "RLCA",
		// 0xF8, 0xF9, 0xFA, 0xFB,
		"MOVA,R0", "MOVA,R1", "MOVA,R2", "MOVA,R3",
		// 0xFC, 0xFD, 0xFE, 0xFF,
		"MOVA,R4", "MOVA,R5", "MOVA,R6", "MOVA,R7",
	}

	opCodeLength := [...]int8{
		//	00,01,02,03,04,05,06,07,08,09,0A,0B,0C,0D,0E,0F
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
		0, 0, 3, 0, 1, 0, 1, 3, 1, 1, 1, 3, 0, 0, 0, 0,
		//  90,91,92,93,94,95,96,97,98,99,9A,9B,9C,9D,9E,9F
		0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 3, 0, 0, 0, 0,
		//  A0,A1,A2,A3,A4,A5,A6,A7,A8,A9,AA,AB,AC,AD,AE,AF
		0, 0, 3, 0, 1, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  B0,B1,B2,B3,B4,B5,B6,B7,B8,B9,BA,BB,BC,BD,BE,BF
		1, 1, 1, 0, 1, 0, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
		//  C0,C1,C2,C3,C4,C5,C6,C7,C8,C9,CA,CB,CC,CD,CE,CF
		3, 3, 3, 3, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  D0,D1,D2,D3,D4,D5,D6,D7,D8,D9,DA,DB,DC,DD,DE,DF
		0, 0, 1, 1, 1, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		//  E0,E1,E2,E3,E4,E5,E6,E7,E8,E9,EA,EB,EC,ED,EE,EF
		3, 3, 3, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		//  F0,F1,F2,F3,F4,F5,F6,F7,F8,F9,FA,FB,FC,FD,FE,FF
		0, 0, 1, 3, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	fmt.Println(opCodeMatrix)
	fmt.Println(opCodeLength)
}
