package main

import (
	"testing"
)

func TestFindPat1(t *testing.T) {
	data := []byte{0xc2,0x93,0x70,0x16,0x2d,0x08,0xa2,0xf1,0x3a,0x5c,0xf9,0xde,0xbc,0xee,0xfc,0x90,0x63,0x19,0xed,0x5d,0xda,0x57,0x4b,0xa0,0x22,0x2b,0x1e,0xf7,0xb1,0x66,0xf6,0x2b,0x29,0x43,0x13,0xbd,0x3f,0xe0,0x19,0x63,0x4c,0x00,0x92,0x6d,0xe6,0x43,0x4c,0x34,0xec,0xb0,0xa2,0x74,0x31,0xae,0xb3,0xc1,0x09,0x19,0x5d,0xab,0x0d,0x53,0x77,0x3b,0xd9,0xfe,0x53,0x79,0x89,0xdd,0x78,0x63,0xce,0xe8,0x89,0x16,0xb8,0x7b,0xf0,0x2c,0xab,0x9b,0x89,0xdc,0x94,0x63,0xe2,0x73,0xfe,0x9f,0x3a,0x6a,0x97,0xa8,0xf2,0xed,0x0c,0xde,0x09,0x16,0xe4,0xa9,0x8f,0x91,0x6d,0x00,0x5e,0x1a,0x64,0x5b,0xa5,0xf9,0x7b,0x36,0xe0,0xc1,0xaa,0x6e,0x04,0x94,0xf7,0x04,0x23,0x7b,0xf3,0xcc,0x2f,0x77,0x92,0x07,0x90,0xa5,0x83,0xd7,0xef,0xf3,0xd7,0x2f,0xef,0x4a,0x56,0xfb,0x17,0x77,0x2b,0xa9,0xbb,0xd6,0xd8,0x4c,0xf3,0x8b,0xf1,0x86,0xa1,0x52,0xb1,0xf5,0xf3,0x95,0x4a,0x2f,0xb2,0xc6,0xc7,0x98,0x56,0x75,0xa1,0x71,0x9e,0x96,0x02,0xb0,0x61,0x94,0xe6,0x77,0xf9,0x86,0x3c,0x93,0x56,0x26,0xf3,0xeb,0x2a,0xae,0x43,0x36,0x38,0x07,0xc2,0x09,0x39,0x95,0xcd,0x66,0x00,0xc8,0x5c,0x60,0x0b,0x7f,0x0f,0x98,0xb2,0x45,0xf3,0xe2,0xe4,0xc4,0x05,0xd2,0xe3,0xc1,0xf6,0x7d,0x99,0xd7,0xcf,0x0f,0xd0,0x7d,0x14,0x14,0xac,0x85,0xf8,0xc0,0x5d,0x40,0x19,0x28,0xf3,0x77,0xf1,0x33,0x23,0x3f,0xed,0xab,0x94,0x00,0x00,0x0d,0xf0,0x1a,0x04,0xad,0x9c,0xe0,0xa2,0xca,0x52,0x67,0x41,0xf0,0xd5,0x88,0x6c,0x8c,0x79,0xf6,0x3e,0x3d,0x37,0x00,0x59,0x34,0x38,0x83,0xf6,0x8c,0x8e,0x08,0xbd,0x5e,0xd0,0x45,0x2f,0x16,0x9e,0xf5,0x07,0xf2,0x3e,0x8f,0x85,0x8f,0x78,0xec,0x0a,0xd2,0x73,0x6b,0xc3,0xa2,0xfc,0x22,0xf4,0x8a,0xdc,0x7b,0xa6,0xe8,0xd1,0xd0,0x4f,0x5b,0xae,0x10,0x85,0x96,0xee,0x6b,0xaf,0xa8,0x79,0x7c,0x05,0x64,0xd6,0x83,0xa2,0xe5,0x66,0x40,0x40,0x65,0xed,0xb6,0xeb,0x0c,0x42,0x68,0x80,0x8d,0x4d,0x45,0xe3,0xca,0x49,0x50,0xb6,0xfc,0x12,0xe8,0x6d,0xad,0x3a,0x06,0x9d,0x96,0x04,0xac,0x14,0xc3,0x94,0xa2,0x9e,0x35,0x73,0x96,0xfb,0xb2,0xd9,0x4f,0xa6,0xb9,0x4b,0x49,0xa6,0xe7,0xc4,0x35,0xc4,0xac,0xa6,0x93,0x6a,0x4b,0xb9,0x4a,0xac,0xb0,0xd0,0xa5,0xd6,0x86,0xd6,0x71,0xf7,0x15,0x80,0x17,0x21,0x00,0x83,0xd3,0x69,0xf4,0x06,0x0f,0xcb,0x69,0x3f,0x37,0x88,0xd5,0x4b,0x42,0x0b,0xd0,0x78,0x76,0xdd,0x44,0x77,0xf8,0x45,0x19,0x96,0x18,0xad,0x65,0xbd,0xb7,0x88,0x03,0x8b,0x60,0xea,0xec,0x69,0x0d,0x0d,0x3f,0x6b,0x7f,0xf1,0x27,0x1f,0x53,0x5a,0xe8,0x41,0xb6,0x18,0xb1,0x4d,0x49,0x63,0xc3,0xb0,0x1c,0x06,0x59,0xcd,0x51,0xe6,0xcb,0x81,0x81,0x63,0xbd,0x3e,0xfe,0x69,0x02,0xda,0x2d,0x7c,0x04,0xde,0xf3,0xa2,0xa3,0x5a,0x48,0xbe,0x42,0x4c,0xce,0xb3,0x3a,0xf8,0x36,0xef,0xeb,0xaa,0xa9,0xfc,0xd6,0x9a,0x99,0xb6,0x89,0x20,0xd5,0x74,0x65,0x94,0xc9,0x53,0x4f,0xb9,0x5a,0x7b,0xfb,0xe1,0x54,0xfe,0x2d,0xf9,0x08,0xa8,0xa6,0x6b,0x7e,0xc5,0x64,0x95,0x6d,0x54,0xcd,0x72,0x05,0x68,0x7b,0xca,0x33,0x64,0x3d,0x14,0x37,0x0f,0x82,0xf4,0xdc,0x27,0xcf,0xf7,0xcf,0x27,0xe4,0xd1,0xae,0xbc,0xd7,0xe2,0x05,0x66,0x54,0x57,0xf5,0x3f,0x40,0xb1,0xc2,0xd8,0xc6,0x83,0x26,0xc8,0x34,0xfd,0xc2,0x16,0x62,0x38,0x1d,0xdf,0x20,0xe5,0x0d,0xd5,0x81,0x4d,0x21,0x75,0xfd,0xeb,0xd5,0x4a,0x9c,0x7c,0x27,0xfb,0xac,0xd1,0x01,0x59,0xb8,0x9a,0x20,0x6f,0xfb,0x7b,0x34,0x03,0x4c,0x3e,0x65,0x54,0x73,0x22,0xc1,0x7d,0x39,0x48,0x3f,0x14,0x81,0x7e,0x6f,0x85,0x07,0x94,0x5a,0xd4,0xc1,0x35,0x90,0x71,0x70,0x7a,0xc6,0x8f,0x5a,0xbf,0xfc,0xed,0x8b,0xd6,0x16,0x39,0x72,0xfd,0x30,0x80,0x79,0xea,0x2a,0x1b,0xdc,0xc3,0x2d,0x28,0x20,0xe9,0x14,0xe4,0x20,0x92,0x53,0x95,0xd2,0x92,0x7c,0x49,0xca,0x58,0x45,0x16,0x67,0x45,0x1a,0x6c,0x56,0x9f,0xed,0xec,0xfa,0xc1,0x03,0xd1,0x93,0xc1,0x74,0x31,0x8e,0x0b,0x37,0xfb,0x54,0x40,0x4b,0x92,0xe2,0x78,0x32,0x9a,0x17,0xc5,0x46,0xf3,0x24,0xa6,0xf5,0x94,0x39,0x49,0x6f,0xd4,0x39,0x76,0x8a,0x1c,0x21,0xa4,0xd8,0x70,0x65,0xae,0x64,0xa4,0x19,0xac,0x0e,0x77,0x07,0xdc,0x2a,0x79,0xc0,0x67,0xba,0x6f,0x66,0xee,0xeb,0x0e,0x51,0x04,0xf9,0xdc,0x7c,0xbe,0x04,0x69,0x6d,0x01,0x4c,0xb8,0xc1,0xbc,0x9d,0x28,0x12,0x01,0x9e,0x39,0xf6,0xda,0xa9,0x1b,0x8c,0xc0,0x75,0x8f,0x77,0xf6,0x51,0x30,0x1a,0x53,0x73,0xe9,0xb7,0xfa,0x86,0xad,0xd6,0x54,0xdd,0x8e,0x52,0x40,0x0b,0x0f,0xbc,0x71,0x29,0x0b,0xbd,0xfe,0x69,0xf4,0xd5,0x4f,0x69,0xea,0xa0,0xfe,0xcb,0xe1,0x54,0xdf,0x59,0x2e,0xfc,0x60,0x18,0x7b,0x45,0xb9,0x0a,0x47,0x13,0x3e,0xab,0x54,0xa2,0xf1,0xb4,0x42,0x21,0xa5,0xd9,0x05,0x8d,0xe5,0x43,0x09,0x95,0x79,0xe9,0x6a,0x71,0xba,0x90,0x6f,0x85,0xe2,0xe1,0x56,0x42,0x0b,0x59,0x2b,0x94,0x4c,0xf3,0x2f,0xd6,0x9a,0xb4,0x0f,0xd0,0x4a,0xbf,0x52,0x64,0x4a,0x58,0x38,0x47,0xa0,0x98,0x35,0x32,0x81,0xf0,0x31,0x0e,0xc4,0x3c,0xd4,0xf5,0x3a,0x6a,0xcb,0xb5,0x3e,0x25,0x94,0x43,0x25,0xa3,0x94,0xae,0x1d,0x05,0x8b,0x78,0x18,0xec,0xa6,0xd1,0xc5,0x5f,0x40,0x84,0xb3,0x95,0xfa,0x09,0x25,0x4d,0x96,0x17,0x24,0xe1,0x7a,0x85,0xb3,0x5b,0x8f,0x40,0xe0,0x39,0xb3,0x31,0x3b,0xef,0x96,0xcb,0x56,0xdc,0xa6,0x0b,0x34,0x8f,0x55,0x3d,0x94,0x81,0x5a,0xf6,0x42,0x33,0xdc,0x5b,0xe2,0x49,0x1c,0x8a,0xa1,0x12}

	actual, err := FindPat(data)
	if actual != 0 || err == nil {
		t.Errorf("Expect : (0, nil). But got : (%d, %s)", actual, err)
	}
}

func TestFindPat2(t *testing.T) {
	data := []byte{0x86,0x47,0x40,0x00,0xc6,0x53,0x72,0xa2,0xf2,0xc3,0x4c,0x0d,0x87,0x50,0xc4,0xe4,0x2e,0xdd,0x36,0x5a,0xc3,0x65,0xb7,0xd5,0x94,0x7e,0x9e,0xc1,0xcd,0xa6,0xde,0xec,0x80,0xa7,0xd0,0x61,0x5b,0xaa,0x98,0x0c,0x4d,0x71,0x9d,0xca,0xd5,0x24,0x16,0x3d,0xbf,0x47,0x5f,0x9e,0x3d,0x13,0xe7,0x4d,0x62,0xc8,0x05,0x7e,0x34,0xc4,0x80,0x5d,0x78,0x5c,0x90,0x61,0x1d,0x97,0x7e,0x58,0x5d,0x71,0xc2,0x8f,0x0e,0x5f,0xeb,0x00,0x9a,0xa5,0xfb,0x09,0xa5,0x58,0x64,0x6b,0xeb,0xe9,0x89,0x0e,0x3f,0x25,0x56,0x64,0x59,0xb2,0x78,0x2c,0x9d,0x97,0x41,0xd9,0xc0,0x21,0x72,0xac,0x38,0x5b,0x75,0x2c,0xe2,0xe0,0xdb,0x0c,0x1a,0xef,0x21,0x82,0xf5,0x4a,0x1a,0x14,0x66,0x97,0xbf,0xa3,0xdb,0x4e,0x95,0xa3,0x8a,0xe7,0x19,0x81,0x0a,0x77,0xef,0xf7,0xaf,0x1c,0x7f,0x1a,0x4e,0xfa,0x51,0x1b,0xdf,0xcd,0x88,0xd4,0x57,0x25,0x2a,0xbd,0xdc,0x60,0xa6,0x9b,0xad,0x79,0x1a,0xe8,0x25,0x8d,0x9b,0x44,0xf2,0xeb,0xe3,0x28,0xd8,0x71,0x8d,0x00,0xc2,0xbb,0xa2,0x5e,0x47,0x67,0x7c,0xba,0x31,0xac,0x91,0x61,0xdc,0x47,0x69,0xeb,0xaa,0x8e,0xdc,0x68,0xf0,0x8f,0x64,0xd6,0x92,0xde,0x94,0x7a,0x90,0x82,0xa3,0xf5,0x12,0x31,0xce,0xac,0x34,0x90,0x6c,0x5d,0xa7,0xa2,0xec,0x45,0xfc,0xd5,0x98,0xde,0x91,0x46,0xdf,0x59,0xca,0x74,0xd3,0xcc,0x43,0x6a,0x84,0xb2,0xbd,0xca,0x9e,0xc6,0xdb,0x30,0xf1,0xb2,0x37,0x75,0x9c,0xdf,0x6e,0xae,0xae,0x1f,0x4a,0xf0,0x42,0x9d,0xa0,0xb4,0xb3,0x5b,0x25,0xfe,0x94,0xb1,0xd7,0x56,0x1f,0xb2,0x30,0x53,0xd9,0x6b,0xa8,0x5f,0x3d,0x67,0xfb,0x4d,0x77,0xb3,0xfd,0xbd,0xf4,0x93,0xa0,0x64,0x3b,0x77,0x81,0x51,0x1a,0xcc,0x3a,0xf4,0xc2,0x19,0x3a,0x23,0x29,0xc1,0x39,0xfd,0x3d,0xd0,0x42,0x41,0xe1,0xf8,0x76,0x72,0x04,0x1b,0x30,0x31,0x18,0xa4,0x0e,0xb9,0x2f,0xaf,0x9b,0x55,0x89,0xeb,0x41,0x90,0xa3,0xda,0x48,0x08,0xa8,0x0d,0x5e,0x57,0xc6,0x49,0xe0,0x2a,0x98,0x73,0x6d,0xde,0x2e,0x16,0x08,0x9d,0x42,0xd6,0x0b,0x9f,0x81,0x77,0x31,0x28,0xfb,0xcc,0x1b,0x63,0x5f,0xbf,0xb3,0xac,0x82,0x67,0x32,0x99,0x81,0x16,0xf1,0xe0,0xd0,0x88,0x20,0xe4,0x6e,0x37,0x03,0x47,0x13,0xdf,0xf0,0x05,0xcf,0xeb,0xe9,0x01,0x85,0x97,0x48,0xe7,0x23,0x58,0x71,0x47,0x3a,0x00,0x57,0x28,0xcf,0x76,0xfd,0xfa,0x7f,0x2c,0x2f,0xac,0x23,0x22,0x14,0xd7,0xd1,0xc0,0x91,0xdf,0x6c,0xa0,0x16,0x01,0x24,0xb2,0xd9,0x80,0x2d,0xb8,0x7a,0xfa,0xef,0x57,0x32,0x8c,0x47,0x8a,0x8d,0x57,0x6b,0x9b,0x4b,0x5b,0xa4,0x89,0xbd,0x93,0x8f,0xd6,0x62,0xfd,0x74,0x32,0x66,0xe6,0xca,0x86,0xba,0x70,0x01,0xa8,0xb8,0xa7,0x97,0xe0,0xd4,0xef,0xf9,0x59,0xfe,0xb1,0x6b,0x15,0x9f,0x59,0x99,0xb9,0xe3,0x32,0x69,0x7d,0x0d,0x49,0x84,0x06,0x99,0xa4,0x39,0x8a,0x40,0xd6,0xc6,0x62,0x3e,0x39,0x5f,0xb1,0x5b,0x0c,0xda,0x35,0x90,0x26,0x0a,0x64,0xa4,0x0a,0x5a,0x40,0x2a,0x5b,0xe8,0xd4,0x13,0x47,0xf8,0xc3,0x5a,0x6f,0xf4,0x1e,0xe3,0xb0,0xa1,0xab,0x5c,0x2e,0xa0,0x78,0x53,0x55,0xca,0xf2,0xa1,0xb4,0x2a,0x9a,0x8f,0xc5,0xfe,0x1f,0xdf,0x1e,0xf7,0x7c,0xf7,0x3e,0x9d,0x3a,0x8e,0xca,0x6d,0xbf,0x5c,0x3f,0x9a,0x37,0x5b,0x09,0x12,0xd8,0x85,0x3e,0x6d,0xdd,0x19,0xc9,0xd3,0x36,0xa2,0x47,0x64,0xc3,0x99,0x83,0x84,0x60,0xef,0x40,0x8c,0x8c,0x48,0x40,0x7c,0x11,0x3e,0xf1,0x06,0x54,0x75,0xb7,0x46,0x5a,0xbc,0x59,0xe6,0x0e,0x1d,0xdf,0xfa,0xd7,0x7a,0x79,0x66,0x0d,0x3a,0xa7,0xc9,0x85,0x58,0x77,0x1a,0x48,0x6e,0x06,0xc5,0x55,0x47,0xce,0x54,0xbf,0x19,0x87,0x87,0xf0,0x86,0x9d,0x1c,0x01,0x50,0xad,0xde,0xad,0xc3,0x76,0x8e,0xae,0x63,0x0a,0x1b,0x8f,0xf9,0xb1,0xfa,0x4b,0x65,0x66,0xf7,0x67,0x69,0x8d,0x45,0xe6,0x22,0xe8,0xf6,0xed,0xe3,0xeb,0xb5,0x6b,0xfb,0xaa,0xd5,0x30,0x52,0xdb,0xb8,0x98,0x99,0xcf,0x1d,0x5c,0xdc,0x2f,0x1d,0x8c,0x6c,0xd6,0x3e,0x01,0xf2,0x1c,0x27,0x53,0xea,0xc3,0x53,0x5b,0xba,0x57,0x58,0xea,0x38,0xef,0x55,0x0c,0xfb,0x2a,0xcb,0xdd,0x8b,0x75,0xb8,0xc2,0x04,0xbf,0xce,0xd1,0xc2,0x07,0xba,0x79,0xc5,0xb9,0x40,0x64,0x78,0x00,0xd4,0xce,0x57,0x9d,0x26,0xdc,0xca,0xc4,0x7f,0x4a,0x1f,0x43,0xa6,0x58,0xe3,0x88,0x40,0x39,0x55,0x2a,0x12,0x96,0x0b,0xef,0x96,0xda,0xc1,0xe0,0x64,0x6b,0xe4,0xb7,0x44,0x20,0x08,0x9e,0x23,0x12,0xa0,0x47,0xd6,0xfd,0xb0,0x02,0x1d,0xee,0xbf,0x8f,0xfe,0x1d,0x11,0xa4,0xa1,0xa5,0x5f,0x1c,0x50,0x8a,0xfe,0x0c,0x7d,0x09,0xef,0x63,0xfb,0xc5,0x18,0xa9,0xb5,0xa9,0x01,0x1d,0x03,0xb7,0x13,0xcb,0xd1,0x25,0x80,0x9f,0x27,0x2f,0x40,0xb7,0x8c,0xf4,0xae,0xd3,0x18,0xc5,0xcb,0x6d,0xb8,0x92,0xd2,0x5e,0x75,0x6f,0xae,0xb2,0xb6,0x52,0x14,0xa8,0xde,0xbe,0x14,0x62,0xd2,0xfc,0x98,0xaa,0x1a,0xc7,0xdc,0xf1,0x9f,0xd3,0x3c,0x78,0x76,0xcd,0xcf,0xf8,0xe4,0xbd,0xd3,0x0e,0xdd,0xb9,0xe1,0xa1,0x48,0x05,0x32,0x3b,0xda,0xe6,0x22,0xa2,0xfa,0x5a,0x96,0x9f,0x38,0x23,0x6d,0x60,0x18,0x19,0x38,0xf7,0xbf,0x4b,0x0c,0xae,0x1c,0x72,0x8f,0xbf,0x06,0x0d,0x5f,0xe9,0x0e,0xd4,0xa1,0xdf,0x1f,0x69,0xd6,0x24,0xb7,0x78,0xa9,0xcf,0xa0,0x01,0x0c,0x79,0x65,0x58,0xf8,0x7c,0x2a,0x45,0xf2,0x0a,0xce,0x80,0x32,0xab,0xef,0x65,0x35,0x7e,0x18,0x60,0xd4,0x80,0xbe,0xdc,0x93,0xd2,0x2b,0x29,0x4f,0x82,0x8a,0x80,0x30,0xc4,0xe7,0x26,0xa2,0xec,0xd3,0x26,0xf5,0x84,0xdc,0x21,0x6c,0x13,0xd9,0x90}

	actual, err := FindPat(data)
	if actual != 1 || err != nil {
		t.Errorf("Expect : (1, nil). But got : (%d, %s)", actual, err)
	}
}