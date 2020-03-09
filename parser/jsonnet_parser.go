// Code generated from Jsonnet.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Jsonnet

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 58, 388,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 42, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 49, 10, 3, 12, 3, 14, 3, 52, 11, 3, 5, 3, 54, 10, 3, 3, 3,
	5, 3, 57, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 63, 10, 3, 3, 3, 6, 3, 66,
	10, 3, 13, 3, 14, 3, 67, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 87, 10, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 94, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 112, 10, 3, 12, 3, 14, 3, 115, 11, 3, 3, 3, 3, 3, 3, 3, 5, 3, 120,
	10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 163, 10, 3, 3, 3, 3, 3, 5, 3,
	167, 10, 3, 3, 3, 3, 3, 5, 3, 171, 10, 3, 5, 3, 173, 10, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 179, 10, 3, 3, 3, 3, 3, 5, 3, 183, 10, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 188, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 194, 10, 3, 12, 3,
	14, 3, 197, 11, 3, 3, 4, 3, 4, 3, 4, 7, 4, 202, 10, 4, 12, 4, 14, 4, 205,
	11, 4, 3, 4, 5, 4, 208, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 213, 10, 4, 12,
	4, 14, 4, 216, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 225,
	10, 4, 12, 4, 14, 4, 228, 11, 4, 3, 4, 5, 4, 231, 10, 4, 3, 4, 6, 4, 234,
	10, 4, 13, 4, 14, 4, 235, 5, 4, 238, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 243,
	10, 5, 3, 6, 3, 6, 5, 6, 247, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 255, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 261, 10, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 269, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 7, 9, 279, 10, 9, 12, 9, 14, 9, 282, 11, 9, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 293, 10, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 299, 10, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 5, 13, 307, 10, 13, 3, 13, 3, 13, 3, 13, 5, 13, 312,
	10, 13, 3, 14, 3, 14, 3, 14, 7, 14, 317, 10, 14, 12, 14, 14, 14, 320, 11,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 326, 10, 14, 12, 14, 14, 14, 329,
	11, 14, 3, 14, 5, 14, 332, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 7, 14, 341, 10, 14, 12, 14, 14, 14, 344, 11, 14, 3, 14, 5, 14,
	347, 10, 14, 5, 14, 349, 10, 14, 3, 15, 3, 15, 3, 15, 7, 15, 354, 10, 15,
	12, 15, 14, 15, 357, 11, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 363, 10,
	15, 12, 15, 14, 15, 366, 11, 15, 3, 15, 5, 15, 369, 10, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 378, 10, 15, 12, 15, 14, 15,
	381, 11, 15, 3, 15, 5, 15, 384, 10, 15, 5, 15, 386, 10, 15, 3, 15, 2, 3,
	4, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 9, 7, 2,
	14, 14, 18, 18, 25, 26, 30, 30, 52, 53, 5, 2, 33, 34, 40, 40, 48, 48, 3,
	2, 35, 37, 3, 2, 33, 34, 3, 2, 46, 47, 3, 2, 41, 45, 3, 2, 31, 32, 2, 452,
	2, 30, 3, 2, 2, 2, 4, 119, 3, 2, 2, 2, 6, 237, 3, 2, 2, 2, 8, 242, 3, 2,
	2, 2, 10, 260, 3, 2, 2, 2, 12, 268, 3, 2, 2, 2, 14, 270, 3, 2, 2, 2, 16,
	273, 3, 2, 2, 2, 18, 283, 3, 2, 2, 2, 20, 292, 3, 2, 2, 2, 22, 294, 3,
	2, 2, 2, 24, 311, 3, 2, 2, 2, 26, 348, 3, 2, 2, 2, 28, 385, 3, 2, 2, 2,
	30, 31, 5, 4, 3, 2, 31, 32, 7, 2, 2, 3, 32, 3, 3, 2, 2, 2, 33, 34, 8, 3,
	1, 2, 34, 120, 9, 2, 2, 2, 35, 36, 7, 3, 2, 2, 36, 37, 5, 4, 3, 2, 37,
	38, 7, 4, 2, 2, 38, 120, 3, 2, 2, 2, 39, 41, 7, 5, 2, 2, 40, 42, 5, 6,
	4, 2, 41, 40, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 120,
	7, 6, 2, 2, 44, 53, 7, 7, 2, 2, 45, 50, 5, 4, 3, 2, 46, 47, 7, 8, 2, 2,
	47, 49, 5, 4, 3, 2, 48, 46, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3,
	2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53,
	45, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 57, 7, 8, 2,
	2, 56, 55, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 120,
	7, 9, 2, 2, 59, 60, 7, 7, 2, 2, 60, 62, 5, 4, 3, 2, 61, 63, 7, 8, 2, 2,
	62, 61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 66, 5,
	16, 9, 2, 65, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67,
	68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 7, 9, 2, 2, 70, 120, 3, 2,
	2, 2, 71, 72, 7, 27, 2, 2, 72, 73, 11, 2, 2, 2, 73, 120, 7, 54, 2, 2, 74,
	75, 7, 27, 2, 2, 75, 76, 7, 7, 2, 2, 76, 77, 5, 4, 3, 2, 77, 78, 7, 9,
	2, 2, 78, 120, 3, 2, 2, 2, 79, 120, 7, 54, 2, 2, 80, 81, 7, 21, 2, 2, 81,
	82, 5, 4, 3, 2, 82, 83, 7, 29, 2, 2, 83, 86, 5, 4, 3, 2, 84, 85, 7, 16,
	2, 2, 85, 87, 5, 4, 3, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 120,
	3, 2, 2, 2, 88, 89, 9, 3, 2, 2, 89, 120, 5, 4, 3, 21, 90, 91, 7, 20, 2,
	2, 91, 93, 7, 3, 2, 2, 92, 94, 5, 28, 15, 2, 93, 92, 3, 2, 2, 2, 93, 94,
	3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 7, 4, 2, 2, 96, 120, 5, 4, 3, 9,
	97, 98, 5, 22, 12, 2, 98, 99, 7, 12, 2, 2, 99, 100, 5, 4, 3, 8, 100, 120,
	3, 2, 2, 2, 101, 102, 7, 22, 2, 2, 102, 120, 7, 52, 2, 2, 103, 104, 7,
	23, 2, 2, 104, 120, 7, 52, 2, 2, 105, 106, 7, 17, 2, 2, 106, 120, 5, 4,
	3, 5, 107, 108, 7, 24, 2, 2, 108, 113, 5, 24, 13, 2, 109, 110, 7, 8, 2,
	2, 110, 112, 5, 24, 13, 2, 111, 109, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2,
	113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115,
	113, 3, 2, 2, 2, 116, 117, 7, 12, 2, 2, 117, 118, 5, 4, 3, 3, 118, 120,
	3, 2, 2, 2, 119, 33, 3, 2, 2, 2, 119, 35, 3, 2, 2, 2, 119, 39, 3, 2, 2,
	2, 119, 44, 3, 2, 2, 2, 119, 59, 3, 2, 2, 2, 119, 71, 3, 2, 2, 2, 119,
	74, 3, 2, 2, 2, 119, 79, 3, 2, 2, 2, 119, 80, 3, 2, 2, 2, 119, 88, 3, 2,
	2, 2, 119, 90, 3, 2, 2, 2, 119, 97, 3, 2, 2, 2, 119, 101, 3, 2, 2, 2, 119,
	103, 3, 2, 2, 2, 119, 105, 3, 2, 2, 2, 119, 107, 3, 2, 2, 2, 120, 195,
	3, 2, 2, 2, 121, 122, 12, 20, 2, 2, 122, 123, 9, 4, 2, 2, 123, 194, 5,
	4, 3, 21, 124, 125, 12, 19, 2, 2, 125, 126, 9, 5, 2, 2, 126, 194, 5, 4,
	3, 20, 127, 128, 12, 18, 2, 2, 128, 129, 9, 6, 2, 2, 129, 194, 5, 4, 3,
	19, 130, 131, 12, 17, 2, 2, 131, 132, 9, 7, 2, 2, 132, 194, 5, 4, 3, 18,
	133, 134, 12, 16, 2, 2, 134, 135, 9, 8, 2, 2, 135, 194, 5, 4, 3, 17, 136,
	137, 12, 15, 2, 2, 137, 138, 7, 49, 2, 2, 138, 194, 5, 4, 3, 16, 139, 140,
	12, 14, 2, 2, 140, 141, 7, 50, 2, 2, 141, 194, 5, 4, 3, 15, 142, 143, 12,
	13, 2, 2, 143, 144, 7, 51, 2, 2, 144, 194, 5, 4, 3, 14, 145, 146, 12, 12,
	2, 2, 146, 147, 7, 38, 2, 2, 147, 194, 5, 4, 3, 13, 148, 149, 12, 11, 2,
	2, 149, 150, 7, 39, 2, 2, 150, 194, 5, 4, 3, 12, 151, 152, 12, 29, 2, 2,
	152, 153, 7, 10, 2, 2, 153, 194, 7, 54, 2, 2, 154, 155, 12, 28, 2, 2, 155,
	156, 7, 7, 2, 2, 156, 157, 5, 4, 3, 2, 157, 158, 7, 9, 2, 2, 158, 194,
	3, 2, 2, 2, 159, 160, 12, 27, 2, 2, 160, 162, 7, 7, 2, 2, 161, 163, 5,
	4, 3, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 164, 3, 2, 2,
	2, 164, 166, 7, 11, 2, 2, 165, 167, 5, 4, 3, 2, 166, 165, 3, 2, 2, 2, 166,
	167, 3, 2, 2, 2, 167, 172, 3, 2, 2, 2, 168, 170, 7, 11, 2, 2, 169, 171,
	5, 4, 3, 2, 170, 169, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2,
	2, 2, 172, 168, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2,
	174, 194, 7, 9, 2, 2, 175, 176, 12, 24, 2, 2, 176, 178, 7, 3, 2, 2, 177,
	179, 5, 26, 14, 2, 178, 177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 180,
	3, 2, 2, 2, 180, 182, 7, 4, 2, 2, 181, 183, 7, 28, 2, 2, 182, 181, 3, 2,
	2, 2, 182, 183, 3, 2, 2, 2, 183, 194, 3, 2, 2, 2, 184, 185, 12, 10, 2,
	2, 185, 187, 7, 5, 2, 2, 186, 188, 5, 6, 4, 2, 187, 186, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 194, 7, 6, 2, 2, 190, 191,
	12, 4, 2, 2, 191, 192, 7, 45, 2, 2, 192, 194, 7, 27, 2, 2, 193, 121, 3,
	2, 2, 2, 193, 124, 3, 2, 2, 2, 193, 127, 3, 2, 2, 2, 193, 130, 3, 2, 2,
	2, 193, 133, 3, 2, 2, 2, 193, 136, 3, 2, 2, 2, 193, 139, 3, 2, 2, 2, 193,
	142, 3, 2, 2, 2, 193, 145, 3, 2, 2, 2, 193, 148, 3, 2, 2, 2, 193, 151,
	3, 2, 2, 2, 193, 154, 3, 2, 2, 2, 193, 159, 3, 2, 2, 2, 193, 175, 3, 2,
	2, 2, 193, 184, 3, 2, 2, 2, 193, 190, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2,
	195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 5, 3, 2, 2, 2, 197, 195,
	3, 2, 2, 2, 198, 203, 5, 8, 5, 2, 199, 200, 7, 8, 2, 2, 200, 202, 5, 8,
	5, 2, 201, 199, 3, 2, 2, 2, 202, 205, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2,
	203, 204, 3, 2, 2, 2, 204, 207, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 206,
	208, 7, 8, 2, 2, 207, 206, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 238,
	3, 2, 2, 2, 209, 210, 5, 14, 8, 2, 210, 211, 7, 8, 2, 2, 211, 213, 3, 2,
	2, 2, 212, 209, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2,
	214, 215, 3, 2, 2, 2, 215, 217, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217,
	218, 7, 7, 2, 2, 218, 219, 5, 4, 3, 2, 219, 220, 7, 9, 2, 2, 220, 221,
	7, 11, 2, 2, 221, 226, 5, 4, 3, 2, 222, 223, 7, 8, 2, 2, 223, 225, 5, 14,
	8, 2, 224, 222, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2,
	226, 227, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229,
	231, 7, 8, 2, 2, 230, 229, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 233,
	3, 2, 2, 2, 232, 234, 5, 16, 9, 2, 233, 232, 3, 2, 2, 2, 234, 235, 3, 2,
	2, 2, 235, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 238, 3, 2, 2, 2,
	237, 198, 3, 2, 2, 2, 237, 214, 3, 2, 2, 2, 238, 7, 3, 2, 2, 2, 239, 243,
	5, 14, 8, 2, 240, 243, 5, 22, 12, 2, 241, 243, 5, 10, 6, 2, 242, 239, 3,
	2, 2, 2, 242, 240, 3, 2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 9, 3, 2, 2, 2,
	244, 246, 5, 20, 11, 2, 245, 247, 7, 33, 2, 2, 246, 245, 3, 2, 2, 2, 246,
	247, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 249, 5, 12, 7, 2, 249, 250,
	5, 4, 3, 2, 250, 261, 3, 2, 2, 2, 251, 252, 5, 20, 11, 2, 252, 254, 7,
	3, 2, 2, 253, 255, 5, 28, 15, 2, 254, 253, 3, 2, 2, 2, 254, 255, 3, 2,
	2, 2, 255, 256, 3, 2, 2, 2, 256, 257, 7, 4, 2, 2, 257, 258, 5, 12, 7, 2,
	258, 259, 5, 4, 3, 2, 259, 261, 3, 2, 2, 2, 260, 244, 3, 2, 2, 2, 260,
	251, 3, 2, 2, 2, 261, 11, 3, 2, 2, 2, 262, 269, 7, 11, 2, 2, 263, 264,
	7, 11, 2, 2, 264, 269, 7, 11, 2, 2, 265, 266, 7, 11, 2, 2, 266, 267, 7,
	11, 2, 2, 267, 269, 7, 11, 2, 2, 268, 262, 3, 2, 2, 2, 268, 263, 3, 2,
	2, 2, 268, 265, 3, 2, 2, 2, 269, 13, 3, 2, 2, 2, 270, 271, 7, 24, 2, 2,
	271, 272, 5, 24, 13, 2, 272, 15, 3, 2, 2, 2, 273, 274, 7, 19, 2, 2, 274,
	275, 7, 54, 2, 2, 275, 276, 7, 45, 2, 2, 276, 280, 5, 4, 3, 2, 277, 279,
	5, 18, 10, 2, 278, 277, 3, 2, 2, 2, 279, 282, 3, 2, 2, 2, 280, 278, 3,
	2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 17, 3, 2, 2, 2, 282, 280, 3, 2, 2,
	2, 283, 284, 7, 21, 2, 2, 284, 285, 5, 4, 3, 2, 285, 19, 3, 2, 2, 2, 286,
	293, 7, 54, 2, 2, 287, 293, 7, 52, 2, 2, 288, 289, 7, 7, 2, 2, 289, 290,
	5, 4, 3, 2, 290, 291, 7, 9, 2, 2, 291, 293, 3, 2, 2, 2, 292, 286, 3, 2,
	2, 2, 292, 287, 3, 2, 2, 2, 292, 288, 3, 2, 2, 2, 293, 21, 3, 2, 2, 2,
	294, 295, 7, 15, 2, 2, 295, 298, 5, 4, 3, 2, 296, 297, 7, 11, 2, 2, 297,
	299, 5, 4, 3, 2, 298, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 23, 3,
	2, 2, 2, 300, 301, 7, 54, 2, 2, 301, 302, 7, 13, 2, 2, 302, 312, 5, 4,
	3, 2, 303, 304, 7, 54, 2, 2, 304, 306, 7, 3, 2, 2, 305, 307, 5, 28, 15,
	2, 306, 305, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308,
	309, 7, 4, 2, 2, 309, 310, 7, 13, 2, 2, 310, 312, 5, 4, 3, 2, 311, 300,
	3, 2, 2, 2, 311, 303, 3, 2, 2, 2, 312, 25, 3, 2, 2, 2, 313, 318, 5, 4,
	3, 2, 314, 315, 7, 8, 2, 2, 315, 317, 5, 4, 3, 2, 316, 314, 3, 2, 2, 2,
	317, 320, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319,
	327, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 321, 322, 7, 8, 2, 2, 322, 323,
	7, 54, 2, 2, 323, 324, 7, 13, 2, 2, 324, 326, 5, 4, 3, 2, 325, 321, 3,
	2, 2, 2, 326, 329, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 328, 3, 2, 2,
	2, 328, 331, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 330, 332, 7, 8, 2, 2, 331,
	330, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 349, 3, 2, 2, 2, 333, 334,
	7, 54, 2, 2, 334, 335, 7, 13, 2, 2, 335, 342, 5, 4, 3, 2, 336, 337, 7,
	8, 2, 2, 337, 338, 7, 54, 2, 2, 338, 339, 7, 13, 2, 2, 339, 341, 5, 4,
	3, 2, 340, 336, 3, 2, 2, 2, 341, 344, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2,
	342, 343, 3, 2, 2, 2, 343, 346, 3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 345,
	347, 7, 8, 2, 2, 346, 345, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 349,
	3, 2, 2, 2, 348, 313, 3, 2, 2, 2, 348, 333, 3, 2, 2, 2, 349, 27, 3, 2,
	2, 2, 350, 355, 7, 54, 2, 2, 351, 352, 7, 8, 2, 2, 352, 354, 7, 54, 2,
	2, 353, 351, 3, 2, 2, 2, 354, 357, 3, 2, 2, 2, 355, 353, 3, 2, 2, 2, 355,
	356, 3, 2, 2, 2, 356, 364, 3, 2, 2, 2, 357, 355, 3, 2, 2, 2, 358, 359,
	7, 8, 2, 2, 359, 360, 7, 54, 2, 2, 360, 361, 7, 13, 2, 2, 361, 363, 5,
	4, 3, 2, 362, 358, 3, 2, 2, 2, 363, 366, 3, 2, 2, 2, 364, 362, 3, 2, 2,
	2, 364, 365, 3, 2, 2, 2, 365, 368, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 367,
	369, 7, 8, 2, 2, 368, 367, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 386,
	3, 2, 2, 2, 370, 371, 7, 54, 2, 2, 371, 372, 7, 13, 2, 2, 372, 379, 5,
	4, 3, 2, 373, 374, 7, 8, 2, 2, 374, 375, 7, 54, 2, 2, 375, 376, 7, 13,
	2, 2, 376, 378, 5, 4, 3, 2, 377, 373, 3, 2, 2, 2, 378, 381, 3, 2, 2, 2,
	379, 377, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 383, 3, 2, 2, 2, 381,
	379, 3, 2, 2, 2, 382, 384, 7, 8, 2, 2, 383, 382, 3, 2, 2, 2, 383, 384,
	3, 2, 2, 2, 384, 386, 3, 2, 2, 2, 385, 350, 3, 2, 2, 2, 385, 370, 3, 2,
	2, 2, 386, 29, 3, 2, 2, 2, 50, 41, 50, 53, 56, 62, 67, 86, 93, 113, 119,
	162, 166, 170, 172, 178, 182, 187, 193, 195, 203, 207, 214, 226, 230, 235,
	237, 242, 246, 254, 260, 268, 280, 292, 298, 306, 311, 318, 327, 331, 342,
	346, 348, 355, 364, 368, 379, 383, 385,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'{'", "'}'", "'['", "','", "']'", "'.'", "':'", "';'",
	"'='", "'$'", "'assert'", "'else'", "'error'", "'false'", "'for'", "'function'",
	"'if'", "'import'", "'importstr'", "'local'", "'null'", "'self'", "'super'",
	"'tailstrict'", "'then'", "'true'", "'=='", "'!='", "'+'", "'-'", "'*'",
	"'/'", "'%'", "'&&'", "'||'", "'!'", "'>'", "'>='", "'<'", "'<='", "'in'",
	"'<<'", "'>>'", "'~'", "'&'", "'^'", "'|'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "DOLLAR", "ASSERT", "ELSE",
	"ERROR", "FALSE", "FOR", "FUNCTION", "IF", "IMPORT", "IMPORTSTR", "LOCAL",
	"NULL", "SELF", "SUPER", "TAILSTRICT", "THEN", "TRUE", "EQUALS", "NOTEQUALS",
	"PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULUS", "AND", "OR", "NOT", "GT",
	"GE", "LT", "LE", "IN", "SHIFTLEFT", "SHIFTRIGHT", "BITNOT", "BITAND",
	"BITXOR", "BITOR", "STRING", "NUMBER", "ID", "Whitespace", "Newline", "BlockComment",
	"LineComment",
}

var ruleNames = []string{
	"jsonnet", "expr", "objinside", "member", "field", "visibility", "objlocal",
	"forspec", "ifspec", "fieldname", "assertion", "bind", "args", "params",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JsonnetParser struct {
	*antlr.BaseParser
}

func NewJsonnetParser(input antlr.TokenStream) *JsonnetParser {
	this := new(JsonnetParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Jsonnet.g4"

	return this
}

// JsonnetParser tokens.
const (
	JsonnetParserEOF          = antlr.TokenEOF
	JsonnetParserT__0         = 1
	JsonnetParserT__1         = 2
	JsonnetParserT__2         = 3
	JsonnetParserT__3         = 4
	JsonnetParserT__4         = 5
	JsonnetParserT__5         = 6
	JsonnetParserT__6         = 7
	JsonnetParserT__7         = 8
	JsonnetParserT__8         = 9
	JsonnetParserT__9         = 10
	JsonnetParserT__10        = 11
	JsonnetParserDOLLAR       = 12
	JsonnetParserASSERT       = 13
	JsonnetParserELSE         = 14
	JsonnetParserERROR        = 15
	JsonnetParserFALSE        = 16
	JsonnetParserFOR          = 17
	JsonnetParserFUNCTION     = 18
	JsonnetParserIF           = 19
	JsonnetParserIMPORT       = 20
	JsonnetParserIMPORTSTR    = 21
	JsonnetParserLOCAL        = 22
	JsonnetParserNULL         = 23
	JsonnetParserSELF         = 24
	JsonnetParserSUPER        = 25
	JsonnetParserTAILSTRICT   = 26
	JsonnetParserTHEN         = 27
	JsonnetParserTRUE         = 28
	JsonnetParserEQUALS       = 29
	JsonnetParserNOTEQUALS    = 30
	JsonnetParserPLUS         = 31
	JsonnetParserMINUS        = 32
	JsonnetParserMULTIPLY     = 33
	JsonnetParserDIVIDE       = 34
	JsonnetParserMODULUS      = 35
	JsonnetParserAND          = 36
	JsonnetParserOR           = 37
	JsonnetParserNOT          = 38
	JsonnetParserGT           = 39
	JsonnetParserGE           = 40
	JsonnetParserLT           = 41
	JsonnetParserLE           = 42
	JsonnetParserIN           = 43
	JsonnetParserSHIFTLEFT    = 44
	JsonnetParserSHIFTRIGHT   = 45
	JsonnetParserBITNOT       = 46
	JsonnetParserBITAND       = 47
	JsonnetParserBITXOR       = 48
	JsonnetParserBITOR        = 49
	JsonnetParserSTRING       = 50
	JsonnetParserNUMBER       = 51
	JsonnetParserID           = 52
	JsonnetParserWhitespace   = 53
	JsonnetParserNewline      = 54
	JsonnetParserBlockComment = 55
	JsonnetParserLineComment  = 56
)

// JsonnetParser rules.
const (
	JsonnetParserRULE_jsonnet    = 0
	JsonnetParserRULE_expr       = 1
	JsonnetParserRULE_objinside  = 2
	JsonnetParserRULE_member     = 3
	JsonnetParserRULE_field      = 4
	JsonnetParserRULE_visibility = 5
	JsonnetParserRULE_objlocal   = 6
	JsonnetParserRULE_forspec    = 7
	JsonnetParserRULE_ifspec     = 8
	JsonnetParserRULE_fieldname  = 9
	JsonnetParserRULE_assertion  = 10
	JsonnetParserRULE_bind       = 11
	JsonnetParserRULE_args       = 12
	JsonnetParserRULE_params     = 13
)

// IJsonnetContext is an interface to support dynamic dispatch.
type IJsonnetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonnetContext differentiates from other interfaces.
	IsJsonnetContext()
}

type JsonnetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonnetContext() *JsonnetContext {
	var p = new(JsonnetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_jsonnet
	return p
}

func (*JsonnetContext) IsJsonnetContext() {}

func NewJsonnetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonnetContext {
	var p = new(JsonnetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_jsonnet

	return p
}

func (s *JsonnetContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonnetContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *JsonnetContext) EOF() antlr.TerminalNode {
	return s.GetToken(JsonnetParserEOF, 0)
}

func (s *JsonnetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonnetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonnetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterJsonnet(s)
	}
}

func (s *JsonnetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitJsonnet(s)
	}
}

func (p *JsonnetParser) Jsonnet() (localctx IJsonnetContext) {
	localctx = NewJsonnetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonnetParserRULE_jsonnet)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.expr(0)
	}
	{
		p.SetState(29)
		p.Match(JsonnetParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ImportContext struct {
	*ExprContext
}

func NewImportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportContext {
	var p = new(ImportContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserIMPORT, 0)
}

func (s *ImportContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSTRING, 0)
}

func (s *ImportContext) IMPORTSTR() antlr.TerminalNode {
	return s.GetToken(JsonnetParserIMPORTSTR, 0)
}

func (s *ImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterImport(s)
	}
}

func (s *ImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitImport(s)
	}
}

type ParensContext struct {
	*ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitParens(s)
	}
}

type VarContext struct {
	*ExprContext
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ID() antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, 0)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitVar(s)
	}
}

type ApplyContext struct {
	*ExprContext
}

func NewApplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplyContext {
	var p = new(ApplyContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ApplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplyContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ApplyContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *ApplyContext) TAILSTRICT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserTAILSTRICT, 0)
}

func (s *ApplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterApply(s)
	}
}

func (s *ApplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitApply(s)
	}
}

type BinaryExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewBinaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExprContext {
	var p = new(BinaryExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BinaryExprContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BinaryExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinaryExprContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(JsonnetParserMULTIPLY, 0)
}

func (s *BinaryExprContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(JsonnetParserDIVIDE, 0)
}

func (s *BinaryExprContext) MODULUS() antlr.TerminalNode {
	return s.GetToken(JsonnetParserMODULUS, 0)
}

func (s *BinaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(JsonnetParserPLUS, 0)
}

func (s *BinaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JsonnetParserMINUS, 0)
}

func (s *BinaryExprContext) SHIFTLEFT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSHIFTLEFT, 0)
}

func (s *BinaryExprContext) SHIFTRIGHT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSHIFTRIGHT, 0)
}

func (s *BinaryExprContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserGT, 0)
}

func (s *BinaryExprContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonnetParserGE, 0)
}

func (s *BinaryExprContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserLT, 0)
}

func (s *BinaryExprContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonnetParserLE, 0)
}

func (s *BinaryExprContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonnetParserIN, 0)
}

func (s *BinaryExprContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(JsonnetParserEQUALS, 0)
}

func (s *BinaryExprContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(JsonnetParserNOTEQUALS, 0)
}

func (s *BinaryExprContext) BITAND() antlr.TerminalNode {
	return s.GetToken(JsonnetParserBITAND, 0)
}

func (s *BinaryExprContext) BITXOR() antlr.TerminalNode {
	return s.GetToken(JsonnetParserBITXOR, 0)
}

func (s *BinaryExprContext) BITOR() antlr.TerminalNode {
	return s.GetToken(JsonnetParserBITOR, 0)
}

func (s *BinaryExprContext) AND() antlr.TerminalNode {
	return s.GetToken(JsonnetParserAND, 0)
}

func (s *BinaryExprContext) OR() antlr.TerminalNode {
	return s.GetToken(JsonnetParserOR, 0)
}

func (s *BinaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterBinaryExpr(s)
	}
}

func (s *BinaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitBinaryExpr(s)
	}
}

type ApplyBraceContext struct {
	*ExprContext
}

func NewApplyBraceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplyBraceContext {
	var p = new(ApplyBraceContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ApplyBraceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplyBraceContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ApplyBraceContext) Objinside() IObjinsideContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjinsideContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjinsideContext)
}

func (s *ApplyBraceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterApplyBrace(s)
	}
}

func (s *ApplyBraceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitApplyBrace(s)
	}
}

type IndexContext struct {
	*ExprContext
}

func NewIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexContext {
	var p = new(IndexContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexContext) ID() antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, 0)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitIndex(s)
	}
}

type UnaryExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryExprContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(JsonnetParserPLUS, 0)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JsonnetParserMINUS, 0)
}

func (s *UnaryExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserNOT, 0)
}

func (s *UnaryExprContext) BITNOT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserBITNOT, 0)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

type IndexExprContext struct {
	*ExprContext
}

func NewIndexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExprContext {
	var p = new(IndexExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IndexExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterIndexExpr(s)
	}
}

func (s *IndexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitIndexExpr(s)
	}
}

type ArrayContext struct {
	*ExprContext
	_expr IExprContext
	elems []IExprContext
}

func NewArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayContext {
	var p = new(ArrayContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ArrayContext) Get_expr() IExprContext { return s._expr }

func (s *ArrayContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ArrayContext) GetElems() []IExprContext { return s.elems }

func (s *ArrayContext) SetElems(v []IExprContext) { s.elems = v }

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArrayContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitArray(s)
	}
}

type FunctionContext struct {
	*ExprContext
}

func NewFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionContext {
	var p = new(FunctionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(JsonnetParserFUNCTION, 0)
}

func (s *FunctionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionContext) Params() IParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitFunction(s)
	}
}

type ErrorExprContext struct {
	*ExprContext
}

func NewErrorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ErrorExprContext {
	var p = new(ErrorExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ErrorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorExprContext) ERROR() antlr.TerminalNode {
	return s.GetToken(JsonnetParserERROR, 0)
}

func (s *ErrorExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ErrorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterErrorExpr(s)
	}
}

func (s *ErrorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitErrorExpr(s)
	}
}

type InSuperContext struct {
	*ExprContext
}

func NewInSuperContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InSuperContext {
	var p = new(InSuperContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InSuperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InSuperContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InSuperContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonnetParserIN, 0)
}

func (s *InSuperContext) SUPER() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSUPER, 0)
}

func (s *InSuperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterInSuper(s)
	}
}

func (s *InSuperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitInSuper(s)
	}
}

type ArrayCompContext struct {
	*ExprContext
}

func NewArrayCompContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayCompContext {
	var p = new(ArrayCompContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ArrayCompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayCompContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayCompContext) AllForspec() []IForspecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForspecContext)(nil)).Elem())
	var tst = make([]IForspecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForspecContext)
		}
	}

	return tst
}

func (s *ArrayCompContext) Forspec(i int) IForspecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForspecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForspecContext)
}

func (s *ArrayCompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterArrayComp(s)
	}
}

func (s *ArrayCompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitArrayComp(s)
	}
}

type AssertContext struct {
	*ExprContext
}

func NewAssertContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssertContext {
	var p = new(AssertContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertContext) Assertion() IAssertionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssertionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *AssertContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterAssert(s)
	}
}

func (s *AssertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitAssert(s)
	}
}

type IndexSuperExprContext struct {
	*ExprContext
}

func NewIndexSuperExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexSuperExprContext {
	var p = new(IndexSuperExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IndexSuperExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexSuperExprContext) SUPER() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSUPER, 0)
}

func (s *IndexSuperExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexSuperExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterIndexSuperExpr(s)
	}
}

func (s *IndexSuperExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitIndexSuperExpr(s)
	}
}

type SliceContext struct {
	*ExprContext
	startIdx IExprContext
	endIdx   IExprContext
	step     IExprContext
}

func NewSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceContext {
	var p = new(SliceContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SliceContext) GetStartIdx() IExprContext { return s.startIdx }

func (s *SliceContext) GetEndIdx() IExprContext { return s.endIdx }

func (s *SliceContext) GetStep() IExprContext { return s.step }

func (s *SliceContext) SetStartIdx(v IExprContext) { s.startIdx = v }

func (s *SliceContext) SetEndIdx(v IExprContext) { s.endIdx = v }

func (s *SliceContext) SetStep(v IExprContext) { s.step = v }

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SliceContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterSlice(s)
	}
}

func (s *SliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitSlice(s)
	}
}

type ValueContext struct {
	*ExprContext
	value antlr.Token
}

func NewValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueContext {
	var p = new(ValueContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ValueContext) GetValue() antlr.Token { return s.value }

func (s *ValueContext) SetValue(v antlr.Token) { s.value = v }

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(JsonnetParserNULL, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(JsonnetParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(JsonnetParserFALSE, 0)
}

func (s *ValueContext) SELF() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSELF, 0)
}

func (s *ValueContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(JsonnetParserDOLLAR, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JsonnetParserNUMBER, 0)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitValue(s)
	}
}

type IndexSuperContext struct {
	*ExprContext
}

func NewIndexSuperContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexSuperContext {
	var p = new(IndexSuperContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IndexSuperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexSuperContext) SUPER() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSUPER, 0)
}

func (s *IndexSuperContext) ID() antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, 0)
}

func (s *IndexSuperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterIndexSuper(s)
	}
}

func (s *IndexSuperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitIndexSuper(s)
	}
}

type ObjectContext struct {
	*ExprContext
}

func NewObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectContext {
	var p = new(ObjectContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) Objinside() IObjinsideContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjinsideContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjinsideContext)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitObject(s)
	}
}

type IfThenElseContext struct {
	*ExprContext
}

func NewIfThenElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfThenElseContext {
	var p = new(IfThenElseContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IfThenElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfThenElseContext) IF() antlr.TerminalNode {
	return s.GetToken(JsonnetParserIF, 0)
}

func (s *IfThenElseContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IfThenElseContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfThenElseContext) THEN() antlr.TerminalNode {
	return s.GetToken(JsonnetParserTHEN, 0)
}

func (s *IfThenElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(JsonnetParserELSE, 0)
}

func (s *IfThenElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterIfThenElse(s)
	}
}

func (s *IfThenElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitIfThenElse(s)
	}
}

type LocalBindContext struct {
	*ExprContext
	_bind IBindContext
	binds []IBindContext
}

func NewLocalBindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalBindContext {
	var p = new(LocalBindContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LocalBindContext) Get_bind() IBindContext { return s._bind }

func (s *LocalBindContext) Set_bind(v IBindContext) { s._bind = v }

func (s *LocalBindContext) GetBinds() []IBindContext { return s.binds }

func (s *LocalBindContext) SetBinds(v []IBindContext) { s.binds = v }

func (s *LocalBindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalBindContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(JsonnetParserLOCAL, 0)
}

func (s *LocalBindContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LocalBindContext) AllBind() []IBindContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBindContext)(nil)).Elem())
	var tst = make([]IBindContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBindContext)
		}
	}

	return tst
}

func (s *LocalBindContext) Bind(i int) IBindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBindContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBindContext)
}

func (s *LocalBindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterLocalBind(s)
	}
}

func (s *LocalBindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitLocalBind(s)
	}
}

func (p *JsonnetParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *JsonnetParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, JsonnetParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(32)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ValueContext).value = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonnetParserDOLLAR)|(1<<JsonnetParserFALSE)|(1<<JsonnetParserNULL)|(1<<JsonnetParserSELF)|(1<<JsonnetParserTRUE))) != 0) || _la == JsonnetParserSTRING || _la == JsonnetParserNUMBER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ValueContext).value = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.Match(JsonnetParserT__0)
		}
		{
			p.SetState(34)
			p.expr(0)
		}
		{
			p.SetState(35)
			p.Match(JsonnetParserT__1)
		}

	case 3:
		localctx = NewObjectContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.Match(JsonnetParserT__2)
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonnetParserT__4)|(1<<JsonnetParserASSERT)|(1<<JsonnetParserLOCAL))) != 0) || _la == JsonnetParserSTRING || _la == JsonnetParserID {
			{
				p.SetState(38)
				p.Objinside()
			}

		}
		{
			p.SetState(41)
			p.Match(JsonnetParserT__3)
		}

	case 4:
		localctx = NewArrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(42)
			p.Match(JsonnetParserT__4)
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonnetParserT__0)|(1<<JsonnetParserT__2)|(1<<JsonnetParserT__4)|(1<<JsonnetParserDOLLAR)|(1<<JsonnetParserASSERT)|(1<<JsonnetParserERROR)|(1<<JsonnetParserFALSE)|(1<<JsonnetParserFUNCTION)|(1<<JsonnetParserIF)|(1<<JsonnetParserIMPORT)|(1<<JsonnetParserIMPORTSTR)|(1<<JsonnetParserLOCAL)|(1<<JsonnetParserNULL)|(1<<JsonnetParserSELF)|(1<<JsonnetParserSUPER)|(1<<JsonnetParserTRUE)|(1<<JsonnetParserPLUS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(JsonnetParserMINUS-32))|(1<<(JsonnetParserNOT-32))|(1<<(JsonnetParserBITNOT-32))|(1<<(JsonnetParserSTRING-32))|(1<<(JsonnetParserNUMBER-32))|(1<<(JsonnetParserID-32)))) != 0) {
			{
				p.SetState(43)

				var _x = p.expr(0)

				localctx.(*ArrayContext)._expr = _x
			}
			localctx.(*ArrayContext).elems = append(localctx.(*ArrayContext).elems, localctx.(*ArrayContext)._expr)
			p.SetState(48)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(44)
						p.Match(JsonnetParserT__5)
					}
					{
						p.SetState(45)

						var _x = p.expr(0)

						localctx.(*ArrayContext)._expr = _x
					}
					localctx.(*ArrayContext).elems = append(localctx.(*ArrayContext).elems, localctx.(*ArrayContext)._expr)

				}
				p.SetState(50)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserT__5 {
			{
				p.SetState(53)
				p.Match(JsonnetParserT__5)
			}

		}
		{
			p.SetState(56)
			p.Match(JsonnetParserT__6)
		}

	case 5:
		localctx = NewArrayCompContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(57)
			p.Match(JsonnetParserT__4)
		}
		{
			p.SetState(58)
			p.expr(0)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserT__5 {
			{
				p.SetState(59)
				p.Match(JsonnetParserT__5)
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == JsonnetParserFOR {
			{
				p.SetState(62)
				p.Forspec()
			}

			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(67)
			p.Match(JsonnetParserT__6)
		}

	case 6:
		localctx = NewIndexSuperContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(JsonnetParserSUPER)
		}
		p.SetState(70)
		p.MatchWildcard()

		{
			p.SetState(71)
			p.Match(JsonnetParserID)
		}

	case 7:
		localctx = NewIndexSuperExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(72)
			p.Match(JsonnetParserSUPER)
		}
		{
			p.SetState(73)
			p.Match(JsonnetParserT__4)
		}
		{
			p.SetState(74)
			p.expr(0)
		}
		{
			p.SetState(75)
			p.Match(JsonnetParserT__6)
		}

	case 8:
		localctx = NewVarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(77)
			p.Match(JsonnetParserID)
		}

	case 9:
		localctx = NewIfThenElseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(78)
			p.Match(JsonnetParserIF)
		}
		{
			p.SetState(79)
			p.expr(0)
		}
		{
			p.SetState(80)
			p.Match(JsonnetParserTHEN)
		}
		{
			p.SetState(81)
			p.expr(0)
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(82)
				p.Match(JsonnetParserELSE)
			}
			{
				p.SetState(83)
				p.expr(0)
			}

		}

	case 10:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(86)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(JsonnetParserPLUS-31))|(1<<(JsonnetParserMINUS-31))|(1<<(JsonnetParserNOT-31))|(1<<(JsonnetParserBITNOT-31)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(87)
			p.expr(19)
		}

	case 11:
		localctx = NewFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(JsonnetParserFUNCTION)
		}
		{
			p.SetState(89)
			p.Match(JsonnetParserT__0)
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserID {
			{
				p.SetState(90)
				p.Params()
			}

		}
		{
			p.SetState(93)
			p.Match(JsonnetParserT__1)
		}
		{
			p.SetState(94)
			p.expr(7)
		}

	case 12:
		localctx = NewAssertContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			p.Assertion()
		}
		{
			p.SetState(96)
			p.Match(JsonnetParserT__9)
		}
		{
			p.SetState(97)
			p.expr(6)
		}

	case 13:
		localctx = NewImportContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(JsonnetParserIMPORT)
		}
		{
			p.SetState(100)
			p.Match(JsonnetParserSTRING)
		}

	case 14:
		localctx = NewImportContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(101)
			p.Match(JsonnetParserIMPORTSTR)
		}
		{
			p.SetState(102)
			p.Match(JsonnetParserSTRING)
		}

	case 15:
		localctx = NewErrorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(103)
			p.Match(JsonnetParserERROR)
		}
		{
			p.SetState(104)
			p.expr(3)
		}

	case 16:
		localctx = NewLocalBindContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(105)
			p.Match(JsonnetParserLOCAL)
		}
		{
			p.SetState(106)

			var _x = p.Bind()

			localctx.(*LocalBindContext)._bind = _x
		}
		localctx.(*LocalBindContext).binds = append(localctx.(*LocalBindContext).binds, localctx.(*LocalBindContext)._bind)
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonnetParserT__5 {
			{
				p.SetState(107)
				p.Match(JsonnetParserT__5)
			}
			{
				p.SetState(108)

				var _x = p.Bind()

				localctx.(*LocalBindContext)._bind = _x
			}
			localctx.(*LocalBindContext).binds = append(localctx.(*LocalBindContext).binds, localctx.(*LocalBindContext)._bind)

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(114)
			p.Match(JsonnetParserT__9)
		}
		{
			p.SetState(115)
			p.expr(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(191)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(120)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(JsonnetParserMULTIPLY-33))|(1<<(JsonnetParserDIVIDE-33))|(1<<(JsonnetParserMODULUS-33)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(121)
					p.expr(19)
				}

			case 2:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(123)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonnetParserPLUS || _la == JsonnetParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(124)
					p.expr(18)
				}

			case 3:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(126)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonnetParserSHIFTLEFT || _la == JsonnetParserSHIFTRIGHT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(127)
					p.expr(17)
				}

			case 4:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(129)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(JsonnetParserGT-39))|(1<<(JsonnetParserGE-39))|(1<<(JsonnetParserLT-39))|(1<<(JsonnetParserLE-39))|(1<<(JsonnetParserIN-39)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(130)
					p.expr(16)
				}

			case 5:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(132)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonnetParserEQUALS || _la == JsonnetParserNOTEQUALS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(133)
					p.expr(15)
				}

			case 6:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(135)

					var _m = p.Match(JsonnetParserBITAND)

					localctx.(*BinaryExprContext).op = _m
				}
				{
					p.SetState(136)
					p.expr(14)
				}

			case 7:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(138)

					var _m = p.Match(JsonnetParserBITXOR)

					localctx.(*BinaryExprContext).op = _m
				}
				{
					p.SetState(139)
					p.expr(13)
				}

			case 8:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(141)

					var _m = p.Match(JsonnetParserBITOR)

					localctx.(*BinaryExprContext).op = _m
				}
				{
					p.SetState(142)
					p.expr(12)
				}

			case 9:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(144)

					var _m = p.Match(JsonnetParserAND)

					localctx.(*BinaryExprContext).op = _m
				}
				{
					p.SetState(145)
					p.expr(11)
				}

			case 10:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(147)

					var _m = p.Match(JsonnetParserOR)

					localctx.(*BinaryExprContext).op = _m
				}
				{
					p.SetState(148)
					p.expr(10)
				}

			case 11:
				localctx = NewIndexContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
				}
				{
					p.SetState(150)
					p.Match(JsonnetParserT__7)
				}
				{
					p.SetState(151)
					p.Match(JsonnetParserID)
				}

			case 12:
				localctx = NewIndexExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
				}
				{
					p.SetState(153)
					p.Match(JsonnetParserT__4)
				}
				{
					p.SetState(154)
					p.expr(0)
				}
				{
					p.SetState(155)
					p.Match(JsonnetParserT__6)
				}

			case 13:
				localctx = NewSliceContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(157)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				{
					p.SetState(158)
					p.Match(JsonnetParserT__4)
				}
				p.SetState(160)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonnetParserT__0)|(1<<JsonnetParserT__2)|(1<<JsonnetParserT__4)|(1<<JsonnetParserDOLLAR)|(1<<JsonnetParserASSERT)|(1<<JsonnetParserERROR)|(1<<JsonnetParserFALSE)|(1<<JsonnetParserFUNCTION)|(1<<JsonnetParserIF)|(1<<JsonnetParserIMPORT)|(1<<JsonnetParserIMPORTSTR)|(1<<JsonnetParserLOCAL)|(1<<JsonnetParserNULL)|(1<<JsonnetParserSELF)|(1<<JsonnetParserSUPER)|(1<<JsonnetParserTRUE)|(1<<JsonnetParserPLUS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(JsonnetParserMINUS-32))|(1<<(JsonnetParserNOT-32))|(1<<(JsonnetParserBITNOT-32))|(1<<(JsonnetParserSTRING-32))|(1<<(JsonnetParserNUMBER-32))|(1<<(JsonnetParserID-32)))) != 0) {
					{
						p.SetState(159)

						var _x = p.expr(0)

						localctx.(*SliceContext).startIdx = _x
					}

				}
				{
					p.SetState(162)
					p.Match(JsonnetParserT__8)
				}
				p.SetState(164)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonnetParserT__0)|(1<<JsonnetParserT__2)|(1<<JsonnetParserT__4)|(1<<JsonnetParserDOLLAR)|(1<<JsonnetParserASSERT)|(1<<JsonnetParserERROR)|(1<<JsonnetParserFALSE)|(1<<JsonnetParserFUNCTION)|(1<<JsonnetParserIF)|(1<<JsonnetParserIMPORT)|(1<<JsonnetParserIMPORTSTR)|(1<<JsonnetParserLOCAL)|(1<<JsonnetParserNULL)|(1<<JsonnetParserSELF)|(1<<JsonnetParserSUPER)|(1<<JsonnetParserTRUE)|(1<<JsonnetParserPLUS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(JsonnetParserMINUS-32))|(1<<(JsonnetParserNOT-32))|(1<<(JsonnetParserBITNOT-32))|(1<<(JsonnetParserSTRING-32))|(1<<(JsonnetParserNUMBER-32))|(1<<(JsonnetParserID-32)))) != 0) {
					{
						p.SetState(163)

						var _x = p.expr(0)

						localctx.(*SliceContext).endIdx = _x
					}

				}
				p.SetState(170)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonnetParserT__8 {
					{
						p.SetState(166)
						p.Match(JsonnetParserT__8)
					}
					p.SetState(168)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonnetParserT__0)|(1<<JsonnetParserT__2)|(1<<JsonnetParserT__4)|(1<<JsonnetParserDOLLAR)|(1<<JsonnetParserASSERT)|(1<<JsonnetParserERROR)|(1<<JsonnetParserFALSE)|(1<<JsonnetParserFUNCTION)|(1<<JsonnetParserIF)|(1<<JsonnetParserIMPORT)|(1<<JsonnetParserIMPORTSTR)|(1<<JsonnetParserLOCAL)|(1<<JsonnetParserNULL)|(1<<JsonnetParserSELF)|(1<<JsonnetParserSUPER)|(1<<JsonnetParserTRUE)|(1<<JsonnetParserPLUS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(JsonnetParserMINUS-32))|(1<<(JsonnetParserNOT-32))|(1<<(JsonnetParserBITNOT-32))|(1<<(JsonnetParserSTRING-32))|(1<<(JsonnetParserNUMBER-32))|(1<<(JsonnetParserID-32)))) != 0) {
						{
							p.SetState(167)

							var _x = p.expr(0)

							localctx.(*SliceContext).step = _x
						}

					}

				}
				{
					p.SetState(172)
					p.Match(JsonnetParserT__6)
				}

			case 14:
				localctx = NewApplyContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(174)
					p.Match(JsonnetParserT__0)
				}
				p.SetState(176)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonnetParserT__0)|(1<<JsonnetParserT__2)|(1<<JsonnetParserT__4)|(1<<JsonnetParserDOLLAR)|(1<<JsonnetParserASSERT)|(1<<JsonnetParserERROR)|(1<<JsonnetParserFALSE)|(1<<JsonnetParserFUNCTION)|(1<<JsonnetParserIF)|(1<<JsonnetParserIMPORT)|(1<<JsonnetParserIMPORTSTR)|(1<<JsonnetParserLOCAL)|(1<<JsonnetParserNULL)|(1<<JsonnetParserSELF)|(1<<JsonnetParserSUPER)|(1<<JsonnetParserTRUE)|(1<<JsonnetParserPLUS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(JsonnetParserMINUS-32))|(1<<(JsonnetParserNOT-32))|(1<<(JsonnetParserBITNOT-32))|(1<<(JsonnetParserSTRING-32))|(1<<(JsonnetParserNUMBER-32))|(1<<(JsonnetParserID-32)))) != 0) {
					{
						p.SetState(175)
						p.Args()
					}

				}
				{
					p.SetState(178)
					p.Match(JsonnetParserT__1)
				}
				p.SetState(180)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(179)
						p.Match(JsonnetParserTAILSTRICT)
					}

				}

			case 15:
				localctx = NewApplyBraceContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(183)
					p.Match(JsonnetParserT__2)
				}
				p.SetState(185)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonnetParserT__4)|(1<<JsonnetParserASSERT)|(1<<JsonnetParserLOCAL))) != 0) || _la == JsonnetParserSTRING || _la == JsonnetParserID {
					{
						p.SetState(184)
						p.Objinside()
					}

				}
				{
					p.SetState(187)
					p.Match(JsonnetParserT__3)
				}

			case 16:
				localctx = NewInSuperContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonnetParserRULE_expr)
				p.SetState(188)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(189)
					p.Match(JsonnetParserIN)
				}
				{
					p.SetState(190)
					p.Match(JsonnetParserSUPER)
				}

			}

		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IObjinsideContext is an interface to support dynamic dispatch.
type IObjinsideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjinsideContext differentiates from other interfaces.
	IsObjinsideContext()
}

type ObjinsideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjinsideContext() *ObjinsideContext {
	var p = new(ObjinsideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_objinside
	return p
}

func (*ObjinsideContext) IsObjinsideContext() {}

func NewObjinsideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjinsideContext {
	var p = new(ObjinsideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_objinside

	return p
}

func (s *ObjinsideContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjinsideContext) CopyFrom(ctx *ObjinsideContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ObjinsideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjinsideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MembersContext struct {
	*ObjinsideContext
	_member IMemberContext
	members []IMemberContext
}

func NewMembersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MembersContext {
	var p = new(MembersContext)

	p.ObjinsideContext = NewEmptyObjinsideContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjinsideContext))

	return p
}

func (s *MembersContext) Get_member() IMemberContext { return s._member }

func (s *MembersContext) Set_member(v IMemberContext) { s._member = v }

func (s *MembersContext) GetMembers() []IMemberContext { return s.members }

func (s *MembersContext) SetMembers(v []IMemberContext) { s.members = v }

func (s *MembersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembersContext) AllMember() []IMemberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberContext)(nil)).Elem())
	var tst = make([]IMemberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberContext)
		}
	}

	return tst
}

func (s *MembersContext) Member(i int) IMemberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *MembersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterMembers(s)
	}
}

func (s *MembersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitMembers(s)
	}
}

type ObjectCompContext struct {
	*ObjinsideContext
	key   IExprContext
	value IExprContext
}

func NewObjectCompContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectCompContext {
	var p = new(ObjectCompContext)

	p.ObjinsideContext = NewEmptyObjinsideContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjinsideContext))

	return p
}

func (s *ObjectCompContext) GetKey() IExprContext { return s.key }

func (s *ObjectCompContext) GetValue() IExprContext { return s.value }

func (s *ObjectCompContext) SetKey(v IExprContext) { s.key = v }

func (s *ObjectCompContext) SetValue(v IExprContext) { s.value = v }

func (s *ObjectCompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectCompContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ObjectCompContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ObjectCompContext) AllObjlocal() []IObjlocalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjlocalContext)(nil)).Elem())
	var tst = make([]IObjlocalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjlocalContext)
		}
	}

	return tst
}

func (s *ObjectCompContext) Objlocal(i int) IObjlocalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjlocalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjlocalContext)
}

func (s *ObjectCompContext) AllForspec() []IForspecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForspecContext)(nil)).Elem())
	var tst = make([]IForspecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForspecContext)
		}
	}

	return tst
}

func (s *ObjectCompContext) Forspec(i int) IForspecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForspecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForspecContext)
}

func (s *ObjectCompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterObjectComp(s)
	}
}

func (s *ObjectCompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitObjectComp(s)
	}
}

func (p *JsonnetParser) Objinside() (localctx IObjinsideContext) {
	localctx = NewObjinsideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonnetParserRULE_objinside)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMembersContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)

			var _x = p.Member()

			localctx.(*MembersContext)._member = _x
		}
		localctx.(*MembersContext).members = append(localctx.(*MembersContext).members, localctx.(*MembersContext)._member)
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(197)
					p.Match(JsonnetParserT__5)
				}
				{
					p.SetState(198)

					var _x = p.Member()

					localctx.(*MembersContext)._member = _x
				}
				localctx.(*MembersContext).members = append(localctx.(*MembersContext).members, localctx.(*MembersContext)._member)

			}
			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserT__5 {
			{
				p.SetState(204)
				p.Match(JsonnetParserT__5)
			}

		}

	case 2:
		localctx = NewObjectCompContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonnetParserLOCAL {
			{
				p.SetState(207)
				p.Objlocal()
			}
			{
				p.SetState(208)
				p.Match(JsonnetParserT__5)
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(215)
			p.Match(JsonnetParserT__4)
		}
		{
			p.SetState(216)

			var _x = p.expr(0)

			localctx.(*ObjectCompContext).key = _x
		}
		{
			p.SetState(217)
			p.Match(JsonnetParserT__6)
		}
		{
			p.SetState(218)
			p.Match(JsonnetParserT__8)
		}
		{
			p.SetState(219)

			var _x = p.expr(0)

			localctx.(*ObjectCompContext).value = _x
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(220)
					p.Match(JsonnetParserT__5)
				}
				{
					p.SetState(221)
					p.Objlocal()
				}

			}
			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserT__5 {
			{
				p.SetState(227)
				p.Match(JsonnetParserT__5)
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == JsonnetParserFOR {
			{
				p.SetState(230)
				p.Forspec()
			}

			p.SetState(233)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_member
	return p
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) Objlocal() IObjlocalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjlocalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjlocalContext)
}

func (s *MemberContext) Assertion() IAssertionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssertionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *MemberContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterMember(s)
	}
}

func (s *MemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitMember(s)
	}
}

func (p *JsonnetParser) Member() (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonnetParserRULE_member)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonnetParserLOCAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Objlocal()
		}

	case JsonnetParserASSERT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Assertion()
		}

	case JsonnetParserT__4, JsonnetParserSTRING, JsonnetParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.Field()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) CopyFrom(ctx *FieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionFieldContext struct {
	*FieldContext
}

func NewFunctionFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionFieldContext {
	var p = new(FunctionFieldContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *FunctionFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionFieldContext) Fieldname() IFieldnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldnameContext)
}

func (s *FunctionFieldContext) Visibility() IVisibilityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVisibilityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVisibilityContext)
}

func (s *FunctionFieldContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionFieldContext) Params() IParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunctionFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterFunctionField(s)
	}
}

func (s *FunctionFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitFunctionField(s)
	}
}

type ValueFieldContext struct {
	*FieldContext
}

func NewValueFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueFieldContext {
	var p = new(ValueFieldContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *ValueFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueFieldContext) Fieldname() IFieldnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldnameContext)
}

func (s *ValueFieldContext) Visibility() IVisibilityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVisibilityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVisibilityContext)
}

func (s *ValueFieldContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueFieldContext) PLUS() antlr.TerminalNode {
	return s.GetToken(JsonnetParserPLUS, 0)
}

func (s *ValueFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterValueField(s)
	}
}

func (s *ValueFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitValueField(s)
	}
}

func (p *JsonnetParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonnetParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Fieldname()
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserPLUS {
			{
				p.SetState(243)
				p.Match(JsonnetParserPLUS)
			}

		}
		{
			p.SetState(246)
			p.Visibility()
		}
		{
			p.SetState(247)
			p.expr(0)
		}

	case 2:
		localctx = NewFunctionFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.Fieldname()
		}
		{
			p.SetState(250)
			p.Match(JsonnetParserT__0)
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserID {
			{
				p.SetState(251)
				p.Params()
			}

		}
		{
			p.SetState(254)
			p.Match(JsonnetParserT__1)
		}
		{
			p.SetState(255)
			p.Visibility()
		}
		{
			p.SetState(256)
			p.expr(0)
		}

	}

	return localctx
}

// IVisibilityContext is an interface to support dynamic dispatch.
type IVisibilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVisibilityContext differentiates from other interfaces.
	IsVisibilityContext()
}

type VisibilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVisibilityContext() *VisibilityContext {
	var p = new(VisibilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_visibility
	return p
}

func (*VisibilityContext) IsVisibilityContext() {}

func NewVisibilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VisibilityContext {
	var p = new(VisibilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_visibility

	return p
}

func (s *VisibilityContext) GetParser() antlr.Parser { return s.parser }
func (s *VisibilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VisibilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VisibilityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterVisibility(s)
	}
}

func (s *VisibilityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitVisibility(s)
	}
}

func (p *JsonnetParser) Visibility() (localctx IVisibilityContext) {
	localctx = NewVisibilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonnetParserRULE_visibility)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.Match(JsonnetParserT__8)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Match(JsonnetParserT__8)
		}
		{
			p.SetState(262)
			p.Match(JsonnetParserT__8)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(263)
			p.Match(JsonnetParserT__8)
		}
		{
			p.SetState(264)
			p.Match(JsonnetParserT__8)
		}
		{
			p.SetState(265)
			p.Match(JsonnetParserT__8)
		}

	}

	return localctx
}

// IObjlocalContext is an interface to support dynamic dispatch.
type IObjlocalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjlocalContext differentiates from other interfaces.
	IsObjlocalContext()
}

type ObjlocalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjlocalContext() *ObjlocalContext {
	var p = new(ObjlocalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_objlocal
	return p
}

func (*ObjlocalContext) IsObjlocalContext() {}

func NewObjlocalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjlocalContext {
	var p = new(ObjlocalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_objlocal

	return p
}

func (s *ObjlocalContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjlocalContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(JsonnetParserLOCAL, 0)
}

func (s *ObjlocalContext) Bind() IBindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBindContext)
}

func (s *ObjlocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjlocalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjlocalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterObjlocal(s)
	}
}

func (s *ObjlocalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitObjlocal(s)
	}
}

func (p *JsonnetParser) Objlocal() (localctx IObjlocalContext) {
	localctx = NewObjlocalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonnetParserRULE_objlocal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(JsonnetParserLOCAL)
	}
	{
		p.SetState(269)
		p.Bind()
	}

	return localctx
}

// IForspecContext is an interface to support dynamic dispatch.
type IForspecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForspecContext differentiates from other interfaces.
	IsForspecContext()
}

type ForspecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForspecContext() *ForspecContext {
	var p = new(ForspecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_forspec
	return p
}

func (*ForspecContext) IsForspecContext() {}

func NewForspecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForspecContext {
	var p = new(ForspecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_forspec

	return p
}

func (s *ForspecContext) GetParser() antlr.Parser { return s.parser }

func (s *ForspecContext) FOR() antlr.TerminalNode {
	return s.GetToken(JsonnetParserFOR, 0)
}

func (s *ForspecContext) ID() antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, 0)
}

func (s *ForspecContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonnetParserIN, 0)
}

func (s *ForspecContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForspecContext) AllIfspec() []IIfspecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfspecContext)(nil)).Elem())
	var tst = make([]IIfspecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfspecContext)
		}
	}

	return tst
}

func (s *ForspecContext) Ifspec(i int) IIfspecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfspecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfspecContext)
}

func (s *ForspecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForspecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForspecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterForspec(s)
	}
}

func (s *ForspecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitForspec(s)
	}
}

func (p *JsonnetParser) Forspec() (localctx IForspecContext) {
	localctx = NewForspecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonnetParserRULE_forspec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(JsonnetParserFOR)
	}
	{
		p.SetState(272)
		p.Match(JsonnetParserID)
	}
	{
		p.SetState(273)
		p.Match(JsonnetParserIN)
	}
	{
		p.SetState(274)
		p.expr(0)
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonnetParserIF {
		{
			p.SetState(275)
			p.Ifspec()
		}

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIfspecContext is an interface to support dynamic dispatch.
type IIfspecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfspecContext differentiates from other interfaces.
	IsIfspecContext()
}

type IfspecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfspecContext() *IfspecContext {
	var p = new(IfspecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_ifspec
	return p
}

func (*IfspecContext) IsIfspecContext() {}

func NewIfspecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfspecContext {
	var p = new(IfspecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_ifspec

	return p
}

func (s *IfspecContext) GetParser() antlr.Parser { return s.parser }

func (s *IfspecContext) IF() antlr.TerminalNode {
	return s.GetToken(JsonnetParserIF, 0)
}

func (s *IfspecContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfspecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfspecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfspecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterIfspec(s)
	}
}

func (s *IfspecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitIfspec(s)
	}
}

func (p *JsonnetParser) Ifspec() (localctx IIfspecContext) {
	localctx = NewIfspecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonnetParserRULE_ifspec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(JsonnetParserIF)
	}
	{
		p.SetState(282)
		p.expr(0)
	}

	return localctx
}

// IFieldnameContext is an interface to support dynamic dispatch.
type IFieldnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldnameContext differentiates from other interfaces.
	IsFieldnameContext()
}

type FieldnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldnameContext() *FieldnameContext {
	var p = new(FieldnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_fieldname
	return p
}

func (*FieldnameContext) IsFieldnameContext() {}

func NewFieldnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldnameContext {
	var p = new(FieldnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_fieldname

	return p
}

func (s *FieldnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldnameContext) ID() antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, 0)
}

func (s *FieldnameContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonnetParserSTRING, 0)
}

func (s *FieldnameContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterFieldname(s)
	}
}

func (s *FieldnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitFieldname(s)
	}
}

func (p *JsonnetParser) Fieldname() (localctx IFieldnameContext) {
	localctx = NewFieldnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonnetParserRULE_fieldname)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonnetParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(JsonnetParserID)
		}

	case JsonnetParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(JsonnetParserSTRING)
		}

	case JsonnetParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(286)
			p.Match(JsonnetParserT__4)
		}
		{
			p.SetState(287)
			p.expr(0)
		}
		{
			p.SetState(288)
			p.Match(JsonnetParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssertionContext is an interface to support dynamic dispatch.
type IAssertionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCondition returns the condition rule contexts.
	GetCondition() IExprContext

	// GetMessage returns the message rule contexts.
	GetMessage() IExprContext

	// SetCondition sets the condition rule contexts.
	SetCondition(IExprContext)

	// SetMessage sets the message rule contexts.
	SetMessage(IExprContext)

	// IsAssertionContext differentiates from other interfaces.
	IsAssertionContext()
}

type AssertionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	condition IExprContext
	message   IExprContext
}

func NewEmptyAssertionContext() *AssertionContext {
	var p = new(AssertionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_assertion
	return p
}

func (*AssertionContext) IsAssertionContext() {}

func NewAssertionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssertionContext {
	var p = new(AssertionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_assertion

	return p
}

func (s *AssertionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssertionContext) GetCondition() IExprContext { return s.condition }

func (s *AssertionContext) GetMessage() IExprContext { return s.message }

func (s *AssertionContext) SetCondition(v IExprContext) { s.condition = v }

func (s *AssertionContext) SetMessage(v IExprContext) { s.message = v }

func (s *AssertionContext) ASSERT() antlr.TerminalNode {
	return s.GetToken(JsonnetParserASSERT, 0)
}

func (s *AssertionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AssertionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssertionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssertionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterAssertion(s)
	}
}

func (s *AssertionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitAssertion(s)
	}
}

func (p *JsonnetParser) Assertion() (localctx IAssertionContext) {
	localctx = NewAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonnetParserRULE_assertion)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(JsonnetParserASSERT)
	}
	{
		p.SetState(293)

		var _x = p.expr(0)

		localctx.(*AssertionContext).condition = _x
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonnetParserT__8 {
		{
			p.SetState(294)
			p.Match(JsonnetParserT__8)
		}
		{
			p.SetState(295)

			var _x = p.expr(0)

			localctx.(*AssertionContext).message = _x
		}

	}

	return localctx
}

// IBindContext is an interface to support dynamic dispatch.
type IBindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBindContext differentiates from other interfaces.
	IsBindContext()
}

type BindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindContext() *BindContext {
	var p = new(BindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_bind
	return p
}

func (*BindContext) IsBindContext() {}

func NewBindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindContext {
	var p = new(BindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_bind

	return p
}

func (s *BindContext) GetParser() antlr.Parser { return s.parser }

func (s *BindContext) CopyFrom(ctx *BindContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueBindContext struct {
	*BindContext
}

func NewValueBindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueBindContext {
	var p = new(ValueBindContext)

	p.BindContext = NewEmptyBindContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BindContext))

	return p
}

func (s *ValueBindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueBindContext) ID() antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, 0)
}

func (s *ValueBindContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueBindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterValueBind(s)
	}
}

func (s *ValueBindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitValueBind(s)
	}
}

type FunctionBindContext struct {
	*BindContext
}

func NewFunctionBindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionBindContext {
	var p = new(FunctionBindContext)

	p.BindContext = NewEmptyBindContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BindContext))

	return p
}

func (s *FunctionBindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBindContext) ID() antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, 0)
}

func (s *FunctionBindContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionBindContext) Params() IParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunctionBindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterFunctionBind(s)
	}
}

func (s *FunctionBindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitFunctionBind(s)
	}
}

func (p *JsonnetParser) Bind() (localctx IBindContext) {
	localctx = NewBindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JsonnetParserRULE_bind)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueBindContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.Match(JsonnetParserID)
		}
		{
			p.SetState(299)
			p.Match(JsonnetParserT__10)
		}
		{
			p.SetState(300)
			p.expr(0)
		}

	case 2:
		localctx = NewFunctionBindContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(301)
			p.Match(JsonnetParserID)
		}
		{
			p.SetState(302)
			p.Match(JsonnetParserT__0)
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserID {
			{
				p.SetState(303)
				p.Params()
			}

		}
		{
			p.SetState(306)
			p.Match(JsonnetParserT__1)
		}
		{
			p.SetState(307)
			p.Match(JsonnetParserT__10)
		}
		{
			p.SetState(308)
			p.expr(0)
		}

	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetNames returns the names token list.
	GetNames() []antlr.Token

	// SetNames sets the names token list.
	SetNames([]antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetPos returns the pos rule context list.
	GetPos() []IExprContext

	// GetNamed returns the named rule context list.
	GetNamed() []IExprContext

	// SetPos sets the pos rule context list.
	SetPos([]IExprContext)

	// SetNamed sets the named rule context list.
	SetNamed([]IExprContext)

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_expr  IExprContext
	pos    []IExprContext
	_ID    antlr.Token
	names  []antlr.Token
	named  []IExprContext
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) Get_ID() antlr.Token { return s._ID }

func (s *ArgsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ArgsContext) GetNames() []antlr.Token { return s.names }

func (s *ArgsContext) SetNames(v []antlr.Token) { s.names = v }

func (s *ArgsContext) Get_expr() IExprContext { return s._expr }

func (s *ArgsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ArgsContext) GetPos() []IExprContext { return s.pos }

func (s *ArgsContext) GetNamed() []IExprContext { return s.named }

func (s *ArgsContext) SetPos(v []IExprContext) { s.pos = v }

func (s *ArgsContext) SetNamed(v []IExprContext) { s.named = v }

func (s *ArgsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArgsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JsonnetParserID)
}

func (s *ArgsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *JsonnetParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonnetParserRULE_args)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)

			var _x = p.expr(0)

			localctx.(*ArgsContext)._expr = _x
		}
		localctx.(*ArgsContext).pos = append(localctx.(*ArgsContext).pos, localctx.(*ArgsContext)._expr)
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(312)
					p.Match(JsonnetParserT__5)
				}
				{
					p.SetState(313)

					var _x = p.expr(0)

					localctx.(*ArgsContext)._expr = _x
				}
				localctx.(*ArgsContext).pos = append(localctx.(*ArgsContext).pos, localctx.(*ArgsContext)._expr)

			}
			p.SetState(318)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(319)
					p.Match(JsonnetParserT__5)
				}
				{
					p.SetState(320)

					var _m = p.Match(JsonnetParserID)

					localctx.(*ArgsContext)._ID = _m
				}
				localctx.(*ArgsContext).names = append(localctx.(*ArgsContext).names, localctx.(*ArgsContext)._ID)
				{
					p.SetState(321)
					p.Match(JsonnetParserT__10)
				}
				{
					p.SetState(322)

					var _x = p.expr(0)

					localctx.(*ArgsContext)._expr = _x
				}
				localctx.(*ArgsContext).named = append(localctx.(*ArgsContext).named, localctx.(*ArgsContext)._expr)

			}
			p.SetState(327)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserT__5 {
			{
				p.SetState(328)
				p.Match(JsonnetParserT__5)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)

			var _m = p.Match(JsonnetParserID)

			localctx.(*ArgsContext)._ID = _m
		}
		localctx.(*ArgsContext).names = append(localctx.(*ArgsContext).names, localctx.(*ArgsContext)._ID)
		{
			p.SetState(332)
			p.Match(JsonnetParserT__10)
		}
		{
			p.SetState(333)

			var _x = p.expr(0)

			localctx.(*ArgsContext)._expr = _x
		}
		localctx.(*ArgsContext).named = append(localctx.(*ArgsContext).named, localctx.(*ArgsContext)._expr)
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(334)
					p.Match(JsonnetParserT__5)
				}
				{
					p.SetState(335)

					var _m = p.Match(JsonnetParserID)

					localctx.(*ArgsContext)._ID = _m
				}
				localctx.(*ArgsContext).names = append(localctx.(*ArgsContext).names, localctx.(*ArgsContext)._ID)
				{
					p.SetState(336)
					p.Match(JsonnetParserT__10)
				}
				{
					p.SetState(337)

					var _x = p.expr(0)

					localctx.(*ArgsContext)._expr = _x
				}
				localctx.(*ArgsContext).named = append(localctx.(*ArgsContext).named, localctx.(*ArgsContext)._expr)

			}
			p.SetState(342)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserT__5 {
			{
				p.SetState(343)
				p.Match(JsonnetParserT__5)
			}

		}

	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetPos returns the pos token list.
	GetPos() []antlr.Token

	// GetNames returns the names token list.
	GetNames() []antlr.Token

	// SetPos sets the pos token list.
	SetPos([]antlr.Token)

	// SetNames sets the names token list.
	SetNames([]antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetDefaults returns the defaults rule context list.
	GetDefaults() []IExprContext

	// SetDefaults sets the defaults rule context list.
	SetDefaults([]IExprContext)

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_ID      antlr.Token
	pos      []antlr.Token
	names    []antlr.Token
	_expr    IExprContext
	defaults []IExprContext
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonnetParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonnetParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) Get_ID() antlr.Token { return s._ID }

func (s *ParamsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ParamsContext) GetPos() []antlr.Token { return s.pos }

func (s *ParamsContext) GetNames() []antlr.Token { return s.names }

func (s *ParamsContext) SetPos(v []antlr.Token) { s.pos = v }

func (s *ParamsContext) SetNames(v []antlr.Token) { s.names = v }

func (s *ParamsContext) Get_expr() IExprContext { return s._expr }

func (s *ParamsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ParamsContext) GetDefaults() []IExprContext { return s.defaults }

func (s *ParamsContext) SetDefaults(v []IExprContext) { s.defaults = v }

func (s *ParamsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JsonnetParserID)
}

func (s *ParamsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JsonnetParserID, i)
}

func (s *ParamsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ParamsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonnetListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *JsonnetParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JsonnetParserRULE_params)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)

			var _m = p.Match(JsonnetParserID)

			localctx.(*ParamsContext)._ID = _m
		}
		localctx.(*ParamsContext).pos = append(localctx.(*ParamsContext).pos, localctx.(*ParamsContext)._ID)
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(349)
					p.Match(JsonnetParserT__5)
				}
				{
					p.SetState(350)

					var _m = p.Match(JsonnetParserID)

					localctx.(*ParamsContext)._ID = _m
				}
				localctx.(*ParamsContext).pos = append(localctx.(*ParamsContext).pos, localctx.(*ParamsContext)._ID)

			}
			p.SetState(355)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(356)
					p.Match(JsonnetParserT__5)
				}
				{
					p.SetState(357)

					var _m = p.Match(JsonnetParserID)

					localctx.(*ParamsContext)._ID = _m
				}
				localctx.(*ParamsContext).names = append(localctx.(*ParamsContext).names, localctx.(*ParamsContext)._ID)
				{
					p.SetState(358)
					p.Match(JsonnetParserT__10)
				}
				{
					p.SetState(359)

					var _x = p.expr(0)

					localctx.(*ParamsContext)._expr = _x
				}
				localctx.(*ParamsContext).defaults = append(localctx.(*ParamsContext).defaults, localctx.(*ParamsContext)._expr)

			}
			p.SetState(364)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserT__5 {
			{
				p.SetState(365)
				p.Match(JsonnetParserT__5)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)

			var _m = p.Match(JsonnetParserID)

			localctx.(*ParamsContext)._ID = _m
		}
		localctx.(*ParamsContext).names = append(localctx.(*ParamsContext).names, localctx.(*ParamsContext)._ID)
		{
			p.SetState(369)
			p.Match(JsonnetParserT__10)
		}
		{
			p.SetState(370)

			var _x = p.expr(0)

			localctx.(*ParamsContext)._expr = _x
		}
		localctx.(*ParamsContext).defaults = append(localctx.(*ParamsContext).defaults, localctx.(*ParamsContext)._expr)
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(371)
					p.Match(JsonnetParserT__5)
				}
				{
					p.SetState(372)

					var _m = p.Match(JsonnetParserID)

					localctx.(*ParamsContext)._ID = _m
				}
				localctx.(*ParamsContext).names = append(localctx.(*ParamsContext).names, localctx.(*ParamsContext)._ID)
				{
					p.SetState(373)
					p.Match(JsonnetParserT__10)
				}
				{
					p.SetState(374)

					var _x = p.expr(0)

					localctx.(*ParamsContext)._expr = _x
				}
				localctx.(*ParamsContext).defaults = append(localctx.(*ParamsContext).defaults, localctx.(*ParamsContext)._expr)

			}
			p.SetState(379)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonnetParserT__5 {
			{
				p.SetState(380)
				p.Match(JsonnetParserT__5)
			}

		}

	}

	return localctx
}

func (p *JsonnetParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonnetParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
