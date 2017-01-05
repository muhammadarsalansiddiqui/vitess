//line ./go/vt/sqlparser/sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line ./go/vt/sqlparser/sql.y:6
func setParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func setAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func incNesting(yylex interface{}) bool {
	yylex.(*Tokenizer).nesting++
	if yylex.(*Tokenizer).nesting == 200 {
		return true
	}
	return false
}

func decNesting(yylex interface{}) {
	yylex.(*Tokenizer).nesting--
}

func forceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

//line ./go/vt/sqlparser/sql.y:34
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	deleteExprs DeleteExprs
	deleteExpr  DeleteExpr
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	boolVal     BoolVal
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	valTuple    ValTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
	colIdent    ColIdent
	colIdents   []ColIdent
	tableIdent  TableIdent
}

const LEX_ERROR = 57346
const UNION = 57347
const SELECT = 57348
const INSERT = 57349
const UPDATE = 57350
const DELETE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const FOR = 57359
const ALL = 57360
const DISTINCT = 57361
const AS = 57362
const EXISTS = 57363
const ASC = 57364
const DESC = 57365
const INTO = 57366
const DUPLICATE = 57367
const KEY = 57368
const DEFAULT = 57369
const SET = 57370
const LOCK = 57371
const VALUES = 57372
const LAST_INSERT_ID = 57373
const NEXT = 57374
const VALUE = 57375
const JOIN = 57376
const STRAIGHT_JOIN = 57377
const LEFT = 57378
const RIGHT = 57379
const INNER = 57380
const OUTER = 57381
const CROSS = 57382
const NATURAL = 57383
const USE = 57384
const FORCE = 57385
const ON = 57386
const ID = 57387
const HEX = 57388
const STRING = 57389
const INTEGRAL = 57390
const FLOAT = 57391
const HEXNUM = 57392
const VALUE_ARG = 57393
const LIST_ARG = 57394
const COMMENT = 57395
const NULL = 57396
const TRUE = 57397
const FALSE = 57398
const OR = 57399
const AND = 57400
const NOT = 57401
const BETWEEN = 57402
const CASE = 57403
const WHEN = 57404
const THEN = 57405
const ELSE = 57406
const END = 57407
const LE = 57408
const GE = 57409
const NE = 57410
const NULL_SAFE_EQUAL = 57411
const IS = 57412
const LIKE = 57413
const REGEXP = 57414
const IN = 57415
const SHIFT_LEFT = 57416
const SHIFT_RIGHT = 57417
const MOD = 57418
const UNARY = 57419
const INTERVAL = 57420
const JSON_EXTRACT_OP = 57421
const JSON_UNQUOTE_EXTRACT_OP = 57422
const CREATE = 57423
const ALTER = 57424
const DROP = 57425
const RENAME = 57426
const ANALYZE = 57427
const TABLE = 57428
const INDEX = 57429
const VIEW = 57430
const TO = 57431
const IGNORE = 57432
const IF = 57433
const UNIQUE = 57434
const USING = 57435
const SHOW = 57436
const DESCRIBE = 57437
const EXPLAIN = 57438
const CURRENT_TIMESTAMP = 57439
const DATABASE = 57440
const UNUSED = 57441

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNION",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"ASC",
	"DESC",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"VALUES",
	"LAST_INSERT_ID",
	"NEXT",
	"VALUE",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"'('",
	"','",
	"')'",
	"ID",
	"HEX",
	"STRING",
	"INTEGRAL",
	"FLOAT",
	"HEXNUM",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"NULL",
	"TRUE",
	"FALSE",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"REGEXP",
	"IN",
	"'|'",
	"'&'",
	"SHIFT_LEFT",
	"SHIFT_RIGHT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"MOD",
	"'^'",
	"'~'",
	"UNARY",
	"INTERVAL",
	"'.'",
	"JSON_EXTRACT_OP",
	"JSON_UNQUOTE_EXTRACT_OP",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"CURRENT_TIMESTAMP",
	"DATABASE",
	"UNUSED",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 93,
	46, 50,
	109, 50,
	-2, 95,
	-1, 117,
	94, 250,
	-2, 249,
}

const yyNprod = 254
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1046

var yyAct = [...]int{

	129, 227, 399, 65, 444, 342, 123, 296, 230, 249,
	388, 266, 319, 256, 163, 333, 162, 229, 3, 95,
	110, 121, 289, 132, 165, 177, 47, 193, 359, 361,
	77, 183, 111, 70, 161, 67, 124, 157, 72, 41,
	40, 74, 41, 279, 181, 50, 14, 15, 16, 17,
	105, 407, 48, 49, 35, 84, 37, 58, 406, 405,
	38, 71, 115, 73, 46, 185, 42, 66, 18, 43,
	44, 45, 232, 233, 408, 63, 370, 273, 197, 160,
	109, 96, 159, 116, 90, 210, 91, 68, 169, 460,
	67, 200, 360, 155, 98, 428, 429, 117, 67, 102,
	171, 104, 209, 208, 217, 218, 211, 212, 213, 214,
	215, 216, 210, 154, 195, 180, 182, 179, 198, 419,
	334, 100, 231, 87, 68, 413, 174, 234, 235, 236,
	237, 166, 184, 200, 60, 284, 188, 19, 20, 22,
	21, 23, 213, 214, 215, 216, 210, 334, 243, 386,
	24, 25, 26, 209, 208, 217, 218, 211, 212, 213,
	214, 215, 216, 210, 60, 170, 134, 62, 117, 97,
	240, 168, 169, 101, 244, 274, 299, 78, 276, 371,
	372, 373, 226, 228, 211, 212, 213, 214, 215, 216,
	210, 322, 116, 248, 252, 195, 166, 134, 286, 253,
	60, 426, 272, 295, 166, 277, 304, 305, 306, 423,
	308, 309, 310, 311, 312, 313, 314, 315, 316, 317,
	318, 251, 283, 281, 254, 294, 282, 280, 166, 92,
	307, 451, 320, 199, 198, 60, 267, 269, 326, 298,
	116, 116, 321, 323, 68, 173, 196, 67, 340, 200,
	338, 327, 330, 303, 324, 325, 190, 320, 341, 328,
	331, 287, 288, 337, 60, 401, 269, 301, 302, 300,
	347, 346, 349, 348, 286, 425, 322, 320, 357, 362,
	199, 198, 290, 292, 293, 364, 421, 291, 61, 76,
	345, 166, 166, 166, 166, 68, 200, 28, 85, 274,
	194, 86, 134, 376, 377, 378, 374, 320, 375, 209,
	208, 217, 218, 211, 212, 213, 214, 215, 216, 210,
	199, 198, 395, 320, 380, 14, 60, 120, 68, 389,
	196, 116, 254, 320, 416, 79, 200, 298, 158, 320,
	167, 381, 158, 391, 383, 397, 400, 396, 385, 393,
	382, 275, 394, 209, 208, 217, 218, 211, 212, 213,
	214, 215, 216, 210, 134, 120, 120, 60, 389, 246,
	97, 414, 191, 254, 94, 238, 239, 97, 108, 241,
	354, 417, 82, 404, 352, 355, 134, 415, 420, 353,
	403, 351, 350, 409, 274, 418, 424, 387, 55, 356,
	247, 262, 263, 153, 120, 167, 152, 14, 190, 414,
	95, 54, 455, 167, 274, 39, 89, 440, 438, 422,
	435, 437, 88, 439, 456, 175, 107, 442, 400, 443,
	441, 336, 271, 445, 445, 445, 120, 167, 446, 447,
	343, 448, 151, 172, 120, 120, 412, 57, 297, 150,
	67, 250, 458, 457, 51, 52, 344, 450, 411, 452,
	453, 461, 462, 29, 430, 431, 369, 136, 135, 137,
	138, 139, 140, 158, 436, 141, 59, 64, 459, 31,
	32, 33, 34, 449, 120, 120, 75, 14, 28, 30,
	80, 217, 218, 211, 212, 213, 214, 215, 216, 210,
	167, 167, 167, 167, 1, 93, 270, 265, 258, 261,
	262, 263, 259, 99, 260, 264, 192, 103, 402, 176,
	106, 36, 278, 178, 69, 114, 149, 133, 339, 245,
	454, 427, 59, 398, 156, 130, 410, 368, 59, 64,
	384, 258, 261, 262, 263, 259, 297, 260, 264, 186,
	242, 134, 187, 320, 117, 136, 135, 137, 138, 139,
	140, 332, 131, 141, 147, 148, 122, 390, 119, 83,
	146, 335, 201, 118, 358, 120, 257, 255, 164, 189,
	120, 113, 392, 81, 53, 27, 56, 13, 12, 11,
	125, 126, 112, 59, 10, 145, 9, 127, 8, 128,
	7, 6, 5, 59, 156, 4, 2, 0, 0, 268,
	0, 59, 0, 142, 0, 0, 0, 0, 0, 143,
	144, 0, 0, 0, 134, 0, 0, 117, 136, 135,
	137, 138, 139, 140, 114, 59, 141, 147, 148, 0,
	0, 285, 0, 146, 0, 0, 0, 120, 120, 0,
	0, 432, 433, 434, 0, 0, 0, 120, 0, 0,
	0, 0, 0, 125, 126, 0, 0, 0, 145, 392,
	127, 0, 128, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 114, 114, 0, 0, 142, 0, 0, 0,
	0, 0, 143, 144, 0, 329, 0, 133, 59, 59,
	59, 59, 0, 0, 0, 0, 0, 0, 0, 68,
	0, 268, 0, 0, 363, 0, 0, 365, 0, 366,
	0, 134, 367, 320, 117, 136, 135, 137, 138, 139,
	140, 0, 0, 141, 147, 148, 0, 0, 119, 0,
	146, 209, 208, 217, 218, 211, 212, 213, 214, 215,
	216, 210, 0, 0, 0, 0, 0, 0, 0, 0,
	125, 126, 112, 0, 133, 145, 0, 127, 0, 128,
	0, 0, 0, 114, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 142, 0, 0, 0, 0, 134, 143,
	144, 117, 136, 135, 137, 138, 139, 140, 0, 0,
	141, 147, 148, 0, 0, 119, 0, 146, 0, 0,
	0, 0, 0, 0, 0, 0, 14, 0, 379, 0,
	0, 0, 0, 0, 0, 0, 0, 125, 126, 112,
	0, 133, 145, 0, 127, 0, 128, 209, 208, 217,
	218, 211, 212, 213, 214, 215, 216, 210, 0, 0,
	142, 0, 0, 0, 0, 134, 143, 144, 117, 136,
	135, 137, 138, 139, 140, 0, 0, 141, 147, 148,
	0, 0, 119, 0, 146, 0, 0, 0, 133, 208,
	217, 218, 211, 212, 213, 214, 215, 216, 210, 14,
	0, 0, 0, 0, 125, 126, 0, 0, 0, 145,
	0, 127, 134, 128, 0, 117, 136, 135, 137, 138,
	139, 140, 0, 0, 141, 147, 148, 142, 0, 119,
	0, 146, 0, 143, 144, 0, 0, 0, 134, 0,
	0, 117, 136, 135, 137, 138, 139, 140, 0, 0,
	141, 125, 126, 0, 0, 0, 145, 146, 127, 0,
	128, 0, 0, 0, 134, 0, 0, 117, 136, 135,
	137, 138, 139, 140, 142, 0, 141, 125, 126, 0,
	143, 144, 145, 146, 127, 0, 128, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	142, 0, 0, 125, 126, 0, 143, 144, 145, 0,
	127, 0, 128, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 142, 203, 206, 0,
	0, 0, 143, 144, 219, 220, 221, 222, 223, 224,
	225, 207, 204, 205, 202, 209, 208, 217, 218, 211,
	212, 213, 214, 215, 216, 210,
}
var yyPact = [...]int{

	40, -1000, -1000, 483, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -48,
	-64, -36, -33, -38, -1000, -1000, -1000, 481, 436, 379,
	-1000, -67, 116, 278, 76, -74, -42, 76, -1000, -39,
	76, -1000, 116, -77, 129, -77, 116, -1000, -1000, -1000,
	-1000, -1000, -1000, 347, 247, -1000, 67, 398, 388, -10,
	-1000, 116, 364, -1000, -13, 123, -1000, 25, -1000, 116,
	59, 125, -1000, 116, -1000, -55, 116, 405, 334, 76,
	-1000, 743, -1000, 432, -1000, 376, 373, -1000, 116, 76,
	116, 462, -27, -15, 152, 116, 2, 76, 909, -1000,
	404, -84, -1000, 17, -1000, 116, -1000, -1000, 116, -1000,
	362, -1000, -1000, 280, -16, 173, 955, -1000, -1000, 857,
	810, -1000, -23, -1000, -1000, 909, 909, 909, 909, 257,
	257, -1000, -1000, 257, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 909, -1000, -1000, 116,
	-1000, -1000, -1000, -1000, 341, 331, -1000, 437, 857, 152,
	86, 327, 507, -1000, -1000, 216, 412, 319, -1000, -1000,
	-1000, 22, -17, 883, -1000, -1000, 307, 76, -1000, -62,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 462,
	743, 152, -1000, -1000, 196, -1000, -1000, 49, 857, 857,
	225, 579, 121, 190, 909, 909, 909, 225, 909, 909,
	909, 909, 909, 909, 909, 909, 909, 909, 909, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 15, 955, 260, 292,
	230, 955, 418, 418, -1000, -1000, -1000, 661, 676, 506,
	-1000, 481, 55, 22, -1000, 401, 76, 76, 437, 424,
	441, 173, 327, -1000, 152, 152, 152, 152, -1000, 358,
	357, -1000, 350, 346, 365, -14, -1000, 187, -1000, -1000,
	116, -1000, 286, 120, 22, 116, -1000, -1000, 116, -1000,
	454, -1000, 178, -1000, -1000, -18, -1000, 15, 57, -1000,
	-1000, 122, -1000, -1000, -1000, 22, -1000, 883, -1000, -1000,
	121, 909, 909, 909, 22, 22, 757, -1000, 409, 798,
	-1000, 56, 56, -5, -5, -5, -5, 100, 100, -1000,
	-1000, -1000, 909, -1000, -1000, -1000, -1000, -1000, 210, 743,
	-1000, 210, 82, -1000, 857, 285, 257, 483, 324, 276,
	-1000, 424, -1000, 909, 909, -1000, 507, 221, 474, -1000,
	-1000, -1000, -1000, 356, -1000, 349, -1000, -1000, -1000, -44,
	-45, -52, -1000, -1000, -1000, -20, -1000, -1000, 445, 431,
	39, -1000, -1000, -1000, 230, -1000, 22, 22, 273, 909,
	22, -1000, 210, -1000, 51, -1000, 909, 220, -1000, 394,
	163, -1000, 909, -1000, -1000, 76, -1000, 229, 155, -1000,
	73, 857, 857, -1000, -1000, 257, 257, 257, 76, -1000,
	437, 857, 909, -1000, -1000, -1000, 909, 22, -1000, -1000,
	22, 909, 391, 257, -1000, 909, 909, -1000, -1000, -1000,
	173, 173, 76, 76, 76, 424, 173, 145, 22, 22,
	475, -1000, 22, -1000, 185, -1000, 185, 185, 395, 76,
	-1000, 76, -1000, -1000, -1000, 470, 10, 123, -1000, -1000,
	76, 76, -1000,
}
var yyPgo = [...]int{

	0, 606, 17, 605, 602, 601, 600, 598, 596, 594,
	589, 588, 587, 463, 586, 585, 584, 583, 167, 75,
	20, 32, 581, 579, 34, 16, 14, 578, 577, 13,
	576, 24, 574, 4, 37, 62, 573, 23, 572, 571,
	21, 1, 569, 22, 7, 8, 567, 6, 36, 566,
	562, 561, 15, 550, 540, 537, 536, 535, 9, 533,
	2, 531, 5, 530, 529, 528, 10, 3, 67, 526,
	415, 289, 524, 523, 522, 521, 519, 0, 27, 516,
	443, 11, 507, 506, 26, 504, 489, 245, 12,
}
var yyR1 = [...]int{

	0, 85, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 3, 3, 4, 5,
	5, 5, 6, 7, 7, 7, 8, 8, 8, 9,
	10, 10, 10, 11, 12, 12, 12, 86, 13, 14,
	14, 15, 15, 15, 16, 16, 17, 17, 18, 18,
	19, 19, 20, 20, 21, 21, 21, 21, 22, 22,
	79, 79, 79, 78, 78, 23, 23, 24, 24, 25,
	25, 26, 26, 26, 27, 27, 27, 27, 83, 83,
	82, 82, 82, 81, 81, 28, 28, 28, 28, 29,
	29, 29, 29, 30, 30, 31, 31, 32, 32, 32,
	32, 33, 33, 34, 34, 35, 35, 35, 35, 35,
	35, 37, 37, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 43, 43, 43, 43,
	43, 43, 38, 38, 38, 38, 38, 38, 38, 44,
	44, 44, 48, 45, 45, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 57, 57, 57, 57, 50, 53, 53, 51,
	51, 52, 54, 54, 49, 49, 49, 40, 40, 40,
	40, 40, 40, 40, 42, 42, 42, 55, 55, 56,
	56, 58, 58, 59, 59, 60, 61, 61, 61, 62,
	62, 62, 63, 63, 63, 64, 64, 65, 65, 66,
	66, 39, 39, 46, 46, 47, 67, 67, 68, 69,
	69, 71, 71, 72, 72, 70, 70, 73, 73, 73,
	73, 73, 73, 74, 74, 75, 75, 76, 76, 77,
	80, 87, 88, 84,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 12, 6, 3, 8, 8, 8, 7,
	6, 7, 3, 5, 8, 4, 6, 7, 4, 5,
	4, 5, 5, 3, 2, 2, 2, 0, 2, 0,
	2, 1, 2, 2, 0, 1, 0, 1, 1, 3,
	1, 3, 1, 3, 1, 2, 3, 5, 1, 1,
	0, 1, 2, 1, 1, 0, 2, 1, 3, 1,
	1, 3, 3, 3, 3, 5, 5, 3, 0, 1,
	0, 1, 2, 1, 1, 1, 2, 2, 1, 2,
	3, 2, 3, 2, 2, 1, 3, 0, 5, 5,
	5, 1, 3, 0, 2, 1, 3, 3, 2, 3,
	3, 1, 1, 1, 3, 3, 3, 4, 3, 4,
	3, 4, 5, 6, 3, 2, 1, 2, 1, 2,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 1, 3, 1, 3, 1, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 3, 3, 4, 5, 3,
	4, 1, 1, 1, 1, 1, 5, 0, 1, 1,
	2, 4, 0, 2, 1, 3, 5, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 0, 3, 0,
	2, 0, 3, 1, 3, 2, 0, 1, 1, 0,
	2, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 2, 1, 1, 3, 3, 1, 3, 3, 1,
	1, 0, 2, 0, 3, 0, 1, 1, 1, 1,
	1, 1, 1, 0, 1, 0, 1, 0, 2, 1,
	1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -85, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, 6, 7, 8, 9, 28, 97,
	98, 100, 99, 101, 110, 111, 112, -15, 5, -13,
	-86, -13, -13, -13, -13, 102, -75, 104, 108, -70,
	104, 106, 102, 102, 103, 104, 102, -84, -84, -84,
	-2, 18, 19, -16, 32, 19, -14, -70, -31, -80,
	48, 10, -18, -19, -80, -67, -68, -77, 48, -72,
	107, 103, -77, 102, -77, -80, -71, 107, 48, -71,
	-80, -17, 35, -42, -77, 51, 54, 56, 24, 28,
	94, -31, -18, -80, 10, 46, 94, 46, 69, -80,
	62, 48, -84, -80, -84, 105, -80, 21, 44, -77,
	-20, -21, 86, -22, -80, -35, -41, 48, -36, 62,
	-87, -40, -49, -47, -48, 84, 85, 91, 93, -77,
	-57, -50, -37, 21, 45, 50, 49, 51, 52, 53,
	54, 57, 107, 113, 114, 89, 64, 58, 59, -69,
	17, 10, 30, 30, -31, -67, -80, -34, 11, 109,
	94, -24, -25, -26, -27, -31, -48, -87, -19, 86,
	-68, -41, -80, -87, -84, 21, -76, 109, -73, 100,
	98, 27, 99, 14, 115, 48, -80, -80, -84, -23,
	46, 10, -79, -78, 20, -77, 50, 94, 61, 60,
	76, -38, 79, 62, 77, 78, 63, 76, 81, 80,
	90, 84, 85, 86, 87, 88, 89, 82, 83, 69,
	70, 71, 72, 73, 74, 75, -35, -41, -35, -2,
	-45, -41, 95, 96, -41, -41, -41, -41, -87, -87,
	-48, -87, -53, -41, -31, -64, 28, -87, -34, -58,
	14, -35, -24, -34, 46, -28, -29, -30, 34, 38,
	40, 35, 36, 37, 41, -82, -81, 20, -80, 50,
	-83, 20, -24, 94, -41, 44, -77, -84, -74, 105,
	-34, -21, -24, -78, 86, -80, -77, -35, -35, -43,
	57, 62, 58, 59, -37, -41, -44, -87, -48, 55,
	79, 77, 78, 63, -41, -41, -41, -43, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -88,
	47, -88, 46, -88, -40, -40, -77, -88, -20, 19,
	-88, -20, -51, -52, 65, -39, 30, -2, -67, -65,
	-77, -58, -62, 16, 15, -34, -25, -26, -25, -26,
	34, 34, 34, 39, 34, 39, 34, -29, -32, 42,
	106, 43, -81, -80, -88, -80, -80, -80, -55, 12,
	94, 57, 58, 59, -45, -44, -41, -41, -41, 61,
	-41, -88, -20, -88, -54, -52, 67, -35, -66, 44,
	-46, -47, -87, -66, -88, 46, -62, -41, -59, -60,
	-41, 44, 44, 34, 34, 103, 103, 103, 94, -84,
	-56, 13, 15, 86, -77, -88, 61, -41, -88, 68,
	-41, 66, 25, 46, -77, 46, 46, -61, 22, 23,
	-35, -35, -87, -87, -87, -58, -35, -45, -41, -41,
	26, -47, -41, -60, -33, -77, -33, -33, -62, 8,
	-88, 46, -88, -88, -63, 17, 29, -67, -77, 8,
	79, -77, -77,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 37, 37, 37, 37, 37, 245,
	235, 0, 0, 0, 253, 253, 253, 0, 41, 44,
	39, 235, 0, 0, 0, 233, 0, 0, 246, 0,
	0, 236, 0, 231, 0, 231, 0, 34, 35, 36,
	15, 42, 43, 46, 0, 45, 38, 0, 0, 95,
	250, 0, 0, 48, 50, 22, 226, 0, 249, 0,
	0, 0, 253, 0, 253, 0, 0, 0, 0, 0,
	33, 0, 47, 0, 194, 0, 0, 40, 0, 0,
	0, 103, 0, -2, 0, 0, 0, 0, 0, 253,
	0, 247, 25, 0, 28, 0, 30, 232, 0, 253,
	65, 52, 54, 60, 0, 58, 59, -2, 105, 0,
	0, 145, 146, 147, 148, 0, 0, 0, 0, 184,
	0, 171, 113, 0, 251, 187, 188, 189, 190, 191,
	192, 193, 172, 173, 174, 175, 177, 111, 112, 0,
	229, 230, 195, 196, 215, 103, 96, 201, 0, 0,
	0, 103, 67, 69, 70, 80, 78, 0, 49, 51,
	227, 228, 0, 0, 23, 234, 0, 0, 253, 243,
	237, 238, 239, 240, 241, 242, 29, 31, 32, 103,
	0, 0, 55, 61, 0, 63, 64, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 132,
	133, 134, 135, 136, 137, 138, 108, 0, 0, 0,
	0, 143, 0, 0, 162, 163, 164, 0, 0, 0,
	125, 0, 0, 178, 14, 0, 0, 0, 201, 209,
	0, 104, 103, 20, 0, 0, 0, 0, 85, 0,
	0, 88, 0, 0, 0, 97, 81, 0, 83, 84,
	0, 79, 0, 0, 143, 0, 248, 26, 0, 244,
	197, 53, 66, 62, 56, 0, 185, 106, 107, 110,
	126, 0, 128, 130, 114, 115, 116, 0, 140, 141,
	0, 0, 0, 0, 118, 120, 0, 124, 149, 150,
	151, 152, 153, 154, 155, 156, 157, 158, 159, 109,
	252, 142, 0, 225, 160, 161, 165, 166, 0, 0,
	169, 0, 182, 179, 0, 219, 0, 222, 219, 0,
	217, 209, 19, 0, 0, 21, 68, 74, 0, 77,
	86, 87, 89, 0, 91, 0, 93, 94, 71, 0,
	0, 0, 82, 72, 73, 0, 253, 27, 199, 0,
	0, 127, 129, 131, 0, 117, 119, 121, 0, 0,
	144, 167, 0, 170, 0, 180, 0, 0, 16, 0,
	221, 223, 0, 17, 216, 0, 18, 210, 202, 203,
	206, 0, 0, 90, 92, 0, 0, 0, 0, 24,
	201, 0, 0, 57, 186, 139, 0, 122, 168, 176,
	183, 0, 0, 0, 218, 0, 0, 205, 207, 208,
	75, 76, 0, 0, 0, 209, 200, 198, 123, 181,
	0, 224, 211, 204, 0, 101, 0, 0, 212, 0,
	98, 0, 99, 100, 13, 0, 0, 220, 102, 213,
	0, 0, 214,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 88, 81, 3,
	45, 47, 86, 84, 46, 85, 94, 87, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	70, 69, 71, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 90, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 80, 3, 91,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 72, 73, 74, 75, 76, 77,
	78, 79, 82, 83, 89, 92, 93, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:185
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:191
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:207
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:211
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{Expr: yyDollar[4].valExpr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:215
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:221
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:225
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:237
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:243
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 20:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:247
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), DeleteExprs: yyDollar[3].deleteExprs, From: yyDollar[5].tableExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:251
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), DeleteExprs: yyDollar[4].deleteExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr)}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:257
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:263
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableIdent}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:267
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:272
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: NewTableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:278
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[4].tableIdent}
		}
	case 27:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:282
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:287
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: NewTableIdent(yyDollar[3].colIdent.Lowered()), NewName: NewTableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:293
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:299
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:307
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:312
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: NewTableIdent(yyDollar[4].colIdent.Lowered()), IfExists: exists}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:322
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[3].tableIdent}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:328
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:332
		{
			yyVAL.statement = &Other{}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:336
		{
			yyVAL.statement = &Other{}
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:341
		{
			setAllowComments(yylex, true)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:345
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:351
		{
			yyVAL.bytes2 = nil
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:355
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:361
		{
			yyVAL.str = UnionStr
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:365
		{
			yyVAL.str = UnionAllStr
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:369
		{
			yyVAL.str = UnionDistinctStr
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:374
		{
			yyVAL.str = ""
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:378
		{
			yyVAL.str = DistinctStr
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:383
		{
			yyVAL.str = ""
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:387
		{
			yyVAL.str = StraightJoinHint
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:393
		{
			yyVAL.deleteExprs = DeleteExprs{yyDollar[1].deleteExpr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:397
		{
			yyVAL.deleteExprs = append(yyVAL.deleteExprs, yyDollar[3].deleteExpr)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:403
		{
			yyVAL.deleteExpr = &NonStarExpr{Expr: yyDollar[1].tableIdent}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:407
		{
			yyVAL.deleteExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:413
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:417
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:423
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:427
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:431
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:435
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:441
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:445
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:450
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:454
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:458
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:465
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:470
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:474
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:480
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:484
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:494
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:498
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:502
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:515
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:519
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:523
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:527
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 78:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:532
		{
			yyVAL.empty = struct{}{}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:534
		{
			yyVAL.empty = struct{}{}
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:537
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:541
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:545
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:552
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:558
		{
			yyVAL.str = JoinStr
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:562
		{
			yyVAL.str = JoinStr
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:566
		{
			yyVAL.str = JoinStr
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:570
		{
			yyVAL.str = StraightJoinStr
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:576
		{
			yyVAL.str = LeftJoinStr
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:580
		{
			yyVAL.str = LeftJoinStr
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:584
		{
			yyVAL.str = RightJoinStr
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:588
		{
			yyVAL.str = RightJoinStr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:594
		{
			yyVAL.str = NaturalJoinStr
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:598
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:608
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:612
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:617
		{
			yyVAL.indexHints = nil
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:621
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:625
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:629
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:635
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:639
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:644
		{
			yyVAL.boolExpr = nil
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:648
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:655
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:659
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:663
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:667
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:671
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:677
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:681
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:687
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:691
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:695
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:699
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:703
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:707
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:711
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:715
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:719
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:723
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:727
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:731
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:735
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:741
		{
			yyVAL.str = IsNullStr
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:745
		{
			yyVAL.str = IsNotNullStr
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:749
		{
			yyVAL.str = IsTrueStr
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:753
		{
			yyVAL.str = IsNotTrueStr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:757
		{
			yyVAL.str = IsFalseStr
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:761
		{
			yyVAL.str = IsNotFalseStr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:767
		{
			yyVAL.str = EqualStr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:771
		{
			yyVAL.str = LessThanStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:775
		{
			yyVAL.str = GreaterThanStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:779
		{
			yyVAL.str = LessEqualStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:783
		{
			yyVAL.str = GreaterEqualStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:787
		{
			yyVAL.str = NotEqualStr
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:791
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:797
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:801
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:805
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:811
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:817
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:821
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:827
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:831
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:835
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:839
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:843
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:847
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:851
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:855
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:859
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:863
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:867
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:871
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:875
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:879
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:883
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:887
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:891
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:895
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:903
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				// Handle double negative
				if num.Val[0] == '-' {
					num.Val = num.Val[1:]
					yyVAL.valExpr = num
				} else {
					yyVAL.valExpr = NewIntVal(append([]byte("-"), num.Val...))
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UMinusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:917
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:921
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:929
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:933
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 168:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:937
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:941
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 170:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:945
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:949
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:955
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:959
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:963
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:967
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 176:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:973
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:978
		{
			yyVAL.valExpr = nil
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:982
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:988
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:992
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 181:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:998
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1003
		{
			yyVAL.valExpr = nil
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1007
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1013
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1017
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 186:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1021
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1027
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1031
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1035
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1039
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1043
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1047
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1051
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1057
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1066
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1070
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1075
		{
			yyVAL.valExprs = nil
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1079
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1084
		{
			yyVAL.boolExpr = nil
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1088
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1093
		{
			yyVAL.orderBy = nil
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1097
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1103
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1107
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1113
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1118
		{
			yyVAL.str = AscScr
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1122
		{
			yyVAL.str = AscScr
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1126
		{
			yyVAL.str = DescScr
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1131
		{
			yyVAL.limit = nil
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1135
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 211:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1139
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1144
		{
			yyVAL.str = ""
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1148
		{
			yyVAL.str = ForUpdateStr
		}
	case 214:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1152
		{
			if yyDollar[3].colIdent.Lowered() != "share" {
				yylex.Error("expecting share")
				return 1
			}
			if yyDollar[4].colIdent.Lowered() != "mode" {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = ShareModeStr
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1165
		{
			yyVAL.columns = nil
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1169
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1175
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 218:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1179
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1184
		{
			yyVAL.updateExprs = nil
		}
	case 220:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1188
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1194
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1198
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1204
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1208
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1214
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1220
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1224
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1230
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 231:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1239
		{
			yyVAL.byt = 0
		}
	case 232:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1241
		{
			yyVAL.byt = 1
		}
	case 233:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1244
		{
			yyVAL.empty = struct{}{}
		}
	case 234:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1246
		{
			yyVAL.empty = struct{}{}
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1249
		{
			yyVAL.str = ""
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1251
		{
			yyVAL.str = IgnoreStr
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1255
		{
			yyVAL.empty = struct{}{}
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1257
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1259
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1261
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1265
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1268
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1270
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1273
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1275
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1278
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1280
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1284
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1290
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1296
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1305
		{
			decNesting(yylex)
		}
	case 253:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1310
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
