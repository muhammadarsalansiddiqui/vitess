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
	yys              int
	empty            struct{}
	statement        Statement
	selStmt          SelectStatement
	byt              byte
	bytes            []byte
	bytes2           [][]byte
	str              string
	selectExprs      SelectExprs
	selectExpr       SelectExpr
	columns          Columns
	colName          *ColName
	tableExprs       TableExprs
	tableExpr        TableExpr
	tableName        *TableName
	indexHints       *IndexHints
	expr             Expr
	boolExpr         BoolExpr
	boolVal          BoolVal
	valExpr          ValExpr
	colTuple         ColTuple
	valExprs         ValExprs
	values           Values
	valTuple         ValTuple
	subquery         *Subquery
	caseExpr         *CaseExpr
	whens            []*When
	when             *When
	orderBy          OrderBy
	order            *Order
	limit            *Limit
	insRows          InsertRows
	updateExprs      UpdateExprs
	updateExpr       *UpdateExpr
	onDupUpdateExprs OnDupUpdateExprs
	onDupUpdateExpr  *OnDupUpdateExpr
	colIdent         ColIdent
	colIdents        []ColIdent
	tableIdent       TableIdent
}

const LEX_ERROR = 57346
const UNION = 57347
const SELECT = 57348
const INSERT = 57349
const REPLACE = 57350
const UPDATE = 57351
const DELETE = 57352
const FROM = 57353
const WHERE = 57354
const GROUP = 57355
const HAVING = 57356
const ORDER = 57357
const BY = 57358
const LIMIT = 57359
const OFFSET = 57360
const FOR = 57361
const ALL = 57362
const DISTINCT = 57363
const AS = 57364
const EXISTS = 57365
const ASC = 57366
const DESC = 57367
const INTO = 57368
const DUPLICATE = 57369
const KEY = 57370
const DEFAULT = 57371
const SET = 57372
const LOCK = 57373
const VALUES = 57374
const LAST_INSERT_ID = 57375
const NEXT = 57376
const VALUE = 57377
const JOIN = 57378
const STRAIGHT_JOIN = 57379
const LEFT = 57380
const RIGHT = 57381
const INNER = 57382
const OUTER = 57383
const CROSS = 57384
const NATURAL = 57385
const USE = 57386
const FORCE = 57387
const ON = 57388
const ID = 57389
const HEX = 57390
const STRING = 57391
const INTEGRAL = 57392
const FLOAT = 57393
const HEXNUM = 57394
const VALUE_ARG = 57395
const LIST_ARG = 57396
const COMMENT = 57397
const NULL = 57398
const TRUE = 57399
const FALSE = 57400
const OR = 57401
const AND = 57402
const NOT = 57403
const BETWEEN = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const LE = 57410
const GE = 57411
const NE = 57412
const NULL_SAFE_EQUAL = 57413
const IS = 57414
const LIKE = 57415
const REGEXP = 57416
const IN = 57417
const SHIFT_LEFT = 57418
const SHIFT_RIGHT = 57419
const MOD = 57420
const UNARY = 57421
const INTERVAL = 57422
const JSON_EXTRACT_OP = 57423
const JSON_UNQUOTE_EXTRACT_OP = 57424
const CREATE = 57425
const ALTER = 57426
const DROP = 57427
const RENAME = 57428
const ANALYZE = 57429
const TABLE = 57430
const INDEX = 57431
const VIEW = 57432
const TO = 57433
const IGNORE = 57434
const IF = 57435
const UNIQUE = 57436
const USING = 57437
const SHOW = 57438
const DESCRIBE = 57439
const EXPLAIN = 57440
const CURRENT_TIMESTAMP = 57441
const DATABASE = 57442
const UNUSED = 57443

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNION",
	"SELECT",
	"INSERT",
	"REPLACE",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"OFFSET",
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
	-1, 120,
	96, 254,
	-2, 253,
}

const yyNprod = 258
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1090

var yyAct = [...]int{

	225, 348, 69, 443, 378, 449, 328, 228, 252, 266,
	338, 126, 284, 50, 124, 265, 320, 264, 135, 277,
	322, 68, 175, 158, 227, 3, 248, 191, 398, 400,
	80, 166, 73, 114, 118, 44, 113, 38, 43, 40,
	44, 51, 52, 41, 268, 46, 47, 48, 15, 16,
	17, 18, 19, 261, 53, 108, 429, 428, 427, 74,
	76, 171, 49, 45, 230, 231, 380, 211, 212, 213,
	214, 208, 20, 355, 181, 208, 255, 195, 97, 63,
	66, 92, 456, 71, 120, 119, 198, 454, 179, 105,
	101, 107, 399, 215, 216, 209, 210, 211, 212, 213,
	214, 208, 169, 168, 196, 63, 410, 96, 61, 183,
	321, 321, 99, 371, 103, 161, 172, 160, 90, 198,
	165, 404, 272, 94, 229, 65, 186, 351, 250, 232,
	233, 234, 235, 71, 127, 194, 291, 356, 357, 358,
	71, 21, 22, 24, 23, 25, 123, 15, 65, 241,
	289, 290, 288, 94, 26, 27, 28, 224, 226, 178,
	180, 177, 209, 210, 211, 212, 213, 214, 208, 91,
	120, 246, 256, 156, 104, 62, 182, 197, 196, 244,
	243, 160, 95, 412, 123, 123, 132, 81, 95, 119,
	259, 65, 287, 198, 236, 237, 242, 251, 239, 65,
	283, 453, 254, 292, 293, 294, 189, 296, 297, 298,
	299, 300, 301, 302, 303, 304, 305, 306, 282, 262,
	271, 247, 263, 95, 70, 295, 65, 75, 250, 123,
	77, 275, 276, 197, 196, 458, 250, 119, 119, 188,
	250, 197, 196, 188, 87, 311, 312, 309, 256, 198,
	123, 270, 30, 307, 308, 310, 416, 198, 123, 123,
	327, 336, 285, 314, 317, 324, 167, 309, 250, 112,
	323, 238, 100, 315, 318, 336, 250, 278, 280, 281,
	70, 163, 279, 349, 79, 70, 256, 70, 354, 325,
	361, 362, 363, 359, 71, 192, 250, 88, 123, 123,
	89, 360, 100, 193, 138, 137, 139, 140, 141, 142,
	365, 65, 143, 351, 249, 250, 100, 119, 95, 426,
	423, 157, 93, 71, 269, 194, 323, 257, 111, 425,
	376, 379, 270, 82, 375, 286, 370, 374, 95, 95,
	85, 390, 366, 393, 70, 368, 381, 386, 394, 388,
	285, 401, 385, 367, 387, 391, 372, 396, 389, 42,
	392, 395, 258, 344, 345, 408, 155, 340, 343, 344,
	345, 341, 411, 342, 346, 58, 15, 424, 123, 193,
	154, 403, 274, 123, 447, 256, 406, 247, 57, 98,
	413, 420, 422, 60, 409, 373, 448, 173, 270, 270,
	270, 270, 159, 110, 353, 269, 54, 55, 430, 329,
	384, 31, 330, 431, 253, 433, 434, 379, 421, 383,
	335, 435, 313, 286, 167, 153, 67, 436, 33, 34,
	35, 36, 37, 152, 170, 444, 326, 340, 343, 344,
	345, 341, 274, 342, 346, 123, 451, 452, 455, 432,
	15, 30, 32, 1, 352, 169, 444, 461, 437, 438,
	347, 190, 174, 39, 260, 176, 72, 151, 442, 64,
	64, 269, 269, 269, 269, 162, 457, 446, 459, 460,
	78, 417, 377, 133, 83, 123, 123, 382, 334, 439,
	440, 441, 369, 240, 468, 64, 319, 64, 134, 125,
	245, 86, 64, 199, 121, 397, 339, 102, 337, 267,
	187, 106, 418, 419, 109, 116, 84, 56, 29, 117,
	59, 14, 13, 316, 465, 136, 12, 11, 10, 9,
	8, 7, 164, 6, 415, 5, 4, 2, 0, 0,
	0, 0, 405, 184, 0, 0, 185, 0, 0, 95,
	0, 250, 120, 138, 137, 139, 140, 141, 142, 0,
	0, 143, 149, 150, 414, 0, 122, 405, 148, 0,
	207, 206, 215, 216, 209, 210, 211, 212, 213, 214,
	208, 0, 0, 0, 0, 0, 64, 0, 128, 129,
	115, 0, 0, 147, 0, 130, 0, 131, 207, 206,
	215, 216, 209, 210, 211, 212, 213, 214, 208, 0,
	0, 144, 0, 0, 462, 0, 0, 145, 146, 445,
	0, 71, 0, 117, 64, 0, 450, 450, 450, 95,
	273, 0, 120, 138, 137, 139, 140, 141, 142, 0,
	445, 143, 0, 463, 0, 464, 0, 0, 148, 0,
	466, 0, 467, 207, 206, 215, 216, 209, 210, 211,
	212, 213, 214, 208, 0, 0, 0, 0, 128, 129,
	0, 117, 117, 147, 0, 130, 0, 131, 0, 0,
	0, 0, 0, 0, 136, 0, 0, 0, 0, 0,
	331, 144, 332, 0, 0, 333, 0, 145, 146, 0,
	0, 0, 0, 350, 0, 64, 0, 0, 95, 407,
	250, 120, 138, 137, 139, 140, 141, 142, 0, 0,
	143, 149, 150, 0, 0, 122, 0, 148, 207, 206,
	215, 216, 209, 210, 211, 212, 213, 214, 208, 0,
	0, 0, 0, 0, 0, 0, 0, 128, 129, 115,
	0, 117, 147, 0, 130, 0, 131, 0, 0, 0,
	0, 136, 0, 0, 0, 0, 0, 0, 0, 0,
	144, 64, 64, 64, 64, 0, 145, 146, 0, 0,
	0, 0, 0, 0, 350, 95, 0, 402, 120, 138,
	137, 139, 140, 141, 142, 0, 0, 143, 149, 150,
	0, 0, 122, 0, 148, 0, 0, 0, 0, 0,
	0, 15, 0, 0, 0, 364, 0, 0, 0, 0,
	0, 0, 0, 0, 128, 129, 115, 0, 136, 147,
	0, 130, 0, 131, 207, 206, 215, 216, 209, 210,
	211, 212, 213, 214, 208, 0, 0, 144, 0, 0,
	0, 0, 95, 145, 146, 120, 138, 137, 139, 140,
	141, 142, 0, 0, 143, 149, 150, 0, 0, 122,
	0, 148, 0, 0, 0, 136, 207, 206, 215, 216,
	209, 210, 211, 212, 213, 214, 208, 0, 0, 0,
	0, 128, 129, 0, 0, 0, 147, 0, 130, 95,
	131, 0, 120, 138, 137, 139, 140, 141, 142, 0,
	0, 143, 149, 150, 144, 0, 122, 0, 148, 0,
	145, 146, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 15, 0, 0, 0, 0, 0, 0, 128, 129,
	0, 0, 0, 147, 0, 130, 95, 131, 0, 120,
	138, 137, 139, 140, 141, 142, 0, 0, 143, 149,
	150, 144, 0, 0, 0, 148, 0, 145, 146, 0,
	0, 0, 95, 0, 0, 120, 138, 137, 139, 140,
	141, 142, 0, 0, 143, 128, 129, 0, 0, 0,
	147, 148, 130, 0, 131, 0, 0, 0, 95, 0,
	0, 120, 138, 137, 139, 140, 141, 142, 144, 0,
	143, 128, 129, 0, 145, 146, 147, 148, 130, 0,
	131, 206, 215, 216, 209, 210, 211, 212, 213, 214,
	208, 0, 0, 0, 144, 0, 0, 128, 129, 0,
	145, 146, 147, 0, 130, 0, 131, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	144, 201, 204, 0, 0, 0, 145, 146, 217, 218,
	219, 220, 221, 222, 223, 205, 202, 203, 200, 207,
	206, 215, 216, 209, 210, 211, 212, 213, 214, 208,
}
var yyPact = [...]int{

	42, -1000, -1000, 446, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -67, -68, -41, -59, -42, -1000, -1000, -1000, 444,
	386, 354, -1000, -73, 149, 98, 415, 90, -77, -46,
	90, -1000, -44, 90, -1000, 98, -79, 137, -79, 98,
	-1000, -1000, -1000, -1000, -1000, -1000, 303, 244, -1000, 60,
	149, 292, 98, -1000, -18, -1000, 359, 98, 268, -1000,
	19, -1000, 98, 50, 124, -1000, 98, -1000, -52, 98,
	380, 282, 90, -1000, 738, -1000, 414, -1000, 348, 334,
	-1000, 291, 370, 90, 90, -1000, -1000, 98, 90, 412,
	90, 951, -1000, 374, -89, -1000, 59, -1000, 98, -1000,
	-1000, 98, -1000, 195, -1000, -1000, 273, -19, 171, 997,
	-1000, -1000, 852, 805, -1000, -33, -1000, -1000, 951, 951,
	951, 951, 271, 271, -1000, -1000, 271, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 951, -1000,
	-1000, 98, -1000, -1000, -1000, -1000, 370, 90, -1000, 271,
	446, 268, 266, -1000, -1000, 254, 399, 852, -1000, 794,
	-20, 925, -1000, -1000, 281, 90, -1000, -54, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 412, 738, 176,
	-1000, -1000, 83, -1000, -1000, 34, 852, 852, 218, 899,
	135, 71, 951, 951, 951, 218, 951, 951, 951, 951,
	951, 951, 951, 951, 951, 951, 951, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 8, 997, 179, 247, 219, 997,
	253, 253, -1000, -1000, -1000, 571, 502, 661, -1000, 444,
	43, 794, -1000, 280, 224, 241, -1000, 951, -1000, 90,
	-1000, 399, 392, 396, 171, 120, 794, 98, -1000, -1000,
	98, -1000, 407, -1000, 213, 401, -1000, -1000, 261, 382,
	141, -1000, -1000, -23, -1000, 8, 41, -1000, -1000, 78,
	-1000, -1000, -1000, 794, -1000, 925, -1000, -1000, 135, 951,
	951, 951, 794, 794, 752, -1000, 9, 938, -1000, -21,
	-21, -17, -17, -17, -17, 76, 76, -1000, -1000, 951,
	-1000, -1000, -1000, -1000, -1000, 191, 738, -1000, 191, 44,
	-1000, 852, -1000, 368, -1000, 271, -1000, 392, -1000, 951,
	951, -30, -1000, -1000, 405, 394, 176, 176, 176, 176,
	-1000, 322, 305, -1000, 319, 307, 325, -16, -1000, 75,
	-1000, -1000, 98, -1000, 227, 33, -1000, -1000, -1000, 219,
	-1000, 794, 794, 646, 951, 794, -1000, 191, -1000, 36,
	-1000, 951, 115, 362, -1000, -1000, 516, 208, -1000, 488,
	90, -1000, 399, 852, 951, 401, 274, 331, -1000, -1000,
	-1000, -1000, 293, -1000, 283, -1000, -1000, -1000, -47, -48,
	-49, -1000, -1000, -1000, -1000, -1000, -1000, 951, 794, -1000,
	-1000, 794, 951, 440, 951, 951, 951, -1000, -1000, -1000,
	392, 171, 199, 852, 852, -1000, -1000, 271, 271, 271,
	794, 794, 90, 794, 794, -1000, 365, 171, 171, 90,
	90, 90, 153, -1000, -1000, 16, -1000, 439, 1, 187,
	-1000, 187, 187, 90, 582, -1000, 90, -1000, 90, -1000,
	-1000, -1000, 271, 90, -1000, 90, -1000, 79, -1000,
}
var yyPgo = [...]int{

	0, 537, 24, 536, 535, 533, 531, 530, 529, 528,
	527, 526, 522, 521, 411, 520, 518, 517, 516, 36,
	33, 515, 510, 17, 15, 9, 509, 508, 10, 506,
	44, 108, 505, 5, 31, 34, 504, 18, 503, 23,
	14, 0, 501, 19, 12, 7, 500, 11, 134, 499,
	498, 496, 16, 493, 492, 488, 487, 483, 8, 482,
	4, 481, 6, 477, 81, 475, 20, 468, 3, 21,
	2, 467, 359, 284, 466, 465, 464, 463, 462, 186,
	27, 461, 434, 1, 460, 454, 13, 453, 452, 61,
	26,
}
var yyR1 = [...]int{

	0, 87, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 3, 3, 4,
	4, 5, 6, 7, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 11, 12, 13, 13, 13, 88, 14,
	15, 15, 16, 16, 16, 17, 17, 18, 18, 19,
	19, 20, 20, 20, 20, 21, 21, 81, 81, 81,
	80, 80, 22, 22, 23, 23, 24, 24, 25, 25,
	25, 26, 26, 26, 26, 85, 85, 84, 84, 84,
	83, 83, 27, 27, 27, 27, 28, 28, 28, 28,
	29, 29, 31, 31, 30, 30, 32, 32, 32, 32,
	33, 33, 34, 34, 35, 35, 35, 35, 35, 35,
	37, 37, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 43, 43, 43, 43, 43,
	43, 38, 38, 38, 38, 38, 38, 38, 44, 44,
	44, 48, 45, 45, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 57, 57, 57, 57, 50, 53, 53, 51, 51,
	52, 54, 54, 49, 49, 49, 40, 40, 40, 40,
	40, 40, 40, 42, 42, 42, 55, 55, 56, 56,
	58, 58, 59, 59, 60, 61, 61, 61, 62, 62,
	62, 62, 63, 63, 63, 64, 64, 65, 65, 66,
	66, 39, 39, 46, 46, 47, 67, 67, 68, 68,
	69, 69, 70, 71, 71, 73, 73, 74, 74, 72,
	72, 75, 75, 75, 75, 75, 75, 76, 76, 77,
	77, 78, 78, 79, 82, 89, 90, 86,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 12, 6, 3, 7, 7, 5,
	5, 8, 7, 3, 5, 8, 4, 6, 7, 4,
	5, 4, 5, 5, 3, 2, 2, 2, 0, 2,
	0, 2, 1, 2, 2, 0, 1, 0, 1, 1,
	3, 1, 2, 3, 5, 1, 1, 0, 1, 2,
	1, 1, 0, 2, 1, 3, 1, 1, 3, 3,
	3, 3, 5, 5, 3, 0, 1, 0, 1, 2,
	1, 1, 1, 2, 2, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 1, 3, 0, 5, 5, 5,
	1, 3, 0, 2, 1, 3, 3, 2, 3, 3,
	1, 1, 1, 3, 3, 3, 4, 3, 4, 3,
	4, 5, 6, 3, 2, 1, 2, 1, 2, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	1, 3, 1, 3, 1, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 3, 3, 4, 5, 3, 4,
	1, 1, 1, 1, 1, 5, 0, 1, 1, 2,
	4, 0, 2, 1, 3, 5, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 0, 3, 0, 2,
	0, 3, 1, 3, 2, 0, 1, 1, 0, 2,
	4, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 2, 1, 1, 3, 3, 1, 3, 1, 6,
	1, 3, 3, 1, 1, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -87, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, -13, 6, 7, 8, 9, 10,
	30, 99, 100, 102, 101, 103, 112, 113, 114, -16,
	5, -14, -88, -14, -14, -14, -14, -14, 104, -77,
	106, 110, -72, 106, 108, 104, 104, 105, 106, 104,
	-86, -86, -86, -2, 20, 21, -17, 34, 21, -15,
	-72, -31, 26, -30, -82, 50, -30, 11, -69, -70,
	-79, 50, -74, 109, 105, -79, 104, -79, -82, -73,
	109, 50, -73, -82, -18, 37, -42, -79, 53, 56,
	58, -31, -64, 30, -89, 47, -30, 96, 30, -30,
	48, 71, -82, 64, 50, -86, -82, -86, 107, -82,
	23, 46, -79, -19, -20, 88, -21, -82, -35, -41,
	50, -36, 64, -89, -40, -49, -47, -48, 86, 87,
	93, 95, -79, -57, -50, -37, 23, 52, 51, 53,
	54, 55, 56, 59, 109, 115, 116, 91, 66, 60,
	61, -71, 19, 11, 32, 32, -64, 30, -39, 32,
	-2, -69, -65, -79, -82, -69, -34, 12, -70, -41,
	-82, -89, -86, 23, -78, 111, -75, 102, 100, 29,
	101, 15, 117, 50, -82, -82, -86, -22, 48, 11,
	-81, -80, 22, -79, 52, 96, 63, 62, 78, -38,
	81, 64, 79, 80, 65, 78, 83, 82, 92, 86,
	87, 88, 89, 90, 91, 84, 85, 71, 72, 73,
	74, 75, 76, 77, -35, -41, -35, -2, -45, -41,
	97, 98, -41, -41, -41, -41, -89, -89, -48, -89,
	-53, -41, -30, -39, -69, -46, -47, -89, -90, 48,
	49, -34, -58, 15, -35, 96, -41, 46, -79, -86,
	-76, 107, -34, -20, -23, -24, -25, -26, -30, -48,
	-89, -80, 88, -82, -79, -35, -35, -43, 59, 64,
	60, 61, -37, -41, -44, -89, -48, 57, 81, 79,
	80, 65, -41, -41, -41, -43, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -90, -90, 48,
	-90, -40, -40, -79, -90, -19, 21, -90, -19, -51,
	-52, 67, -66, 46, -66, 48, -79, -58, -62, 17,
	16, -82, -82, -82, -55, 13, 48, -27, -28, -29,
	36, 40, 42, 37, 38, 39, 43, -84, -83, 22,
	-82, 52, -85, 22, -23, 96, 59, 60, 61, -45,
	-44, -41, -41, -41, 63, -41, -90, -19, -90, -54,
	-52, 69, -35, 27, -47, -62, -41, -59, -60, -41,
	96, -86, -56, 14, 16, -24, -25, -24, -25, 36,
	36, 36, 41, 36, 41, 36, -28, -32, 44, 108,
	45, -83, -82, -90, 88, -79, -90, 63, -41, -90,
	70, -41, 68, 28, 48, 18, 48, -61, 24, 25,
	-58, -35, -45, 46, 46, 36, 36, 105, 105, 105,
	-41, -41, 9, -41, -41, -60, -62, -35, -35, -89,
	-89, -89, -67, -68, -70, -79, -63, 19, 31, -33,
	-79, -33, -33, 48, 71, 9, 81, -90, 48, -90,
	-90, -68, 32, -79, -79, -89, -79, -79, -90,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 38, 38, 38, 38, 38,
	38, 249, 239, 0, 0, 0, 257, 257, 257, 0,
	42, 45, 40, 239, 0, 0, 0, 0, 237, 0,
	0, 250, 0, 0, 240, 0, 235, 0, 235, 0,
	35, 36, 37, 16, 43, 44, 47, 0, 46, 39,
	0, 215, 0, 93, 94, 254, 0, 0, 23, 230,
	0, 253, 0, 0, 0, 257, 0, 257, 0, 0,
	0, 0, 0, 34, 0, 48, 0, 193, 0, 0,
	41, 215, 0, 0, 0, 255, 92, 0, 0, 102,
	0, 0, 257, 0, 251, 26, 0, 29, 0, 31,
	236, 0, 257, 62, 49, 51, 57, 0, 55, 56,
	-2, 104, 0, 0, 144, 145, 146, 147, 0, 0,
	0, 0, 183, 0, 170, 112, 0, 186, 187, 188,
	189, 190, 191, 192, 171, 172, 173, 174, 176, 110,
	111, 0, 233, 234, 194, 195, 0, 0, 19, 0,
	222, 20, 0, 217, 95, 102, 200, 0, 231, 232,
	0, 0, 24, 238, 0, 0, 257, 247, 241, 242,
	243, 244, 245, 246, 30, 32, 33, 102, 0, 0,
	52, 58, 0, 60, 61, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 131, 132, 133,
	134, 135, 136, 137, 107, 0, 0, 0, 0, 142,
	0, 0, 161, 162, 163, 0, 0, 0, 124, 0,
	0, 177, 15, 219, 219, 221, 223, 0, 216, 0,
	256, 200, 208, 0, 103, 0, 142, 0, 252, 27,
	0, 248, 196, 50, 63, 64, 66, 67, 77, 75,
	0, 59, 53, 0, 184, 105, 106, 109, 125, 0,
	127, 129, 113, 114, 115, 0, 139, 140, 0, 0,
	0, 0, 117, 119, 0, 123, 148, 149, 150, 151,
	152, 153, 154, 155, 156, 157, 158, 108, 141, 0,
	225, 159, 160, 164, 165, 0, 0, 168, 0, 181,
	178, 0, 17, 0, 18, 0, 218, 208, 22, 0,
	0, 0, 257, 28, 198, 0, 0, 0, 0, 0,
	82, 0, 0, 85, 0, 0, 0, 96, 78, 0,
	80, 81, 0, 76, 0, 0, 126, 128, 130, 0,
	116, 118, 120, 0, 0, 143, 166, 0, 169, 0,
	179, 0, 0, 0, 224, 21, 209, 201, 202, 205,
	0, 25, 200, 0, 0, 65, 71, 0, 74, 83,
	84, 86, 0, 88, 0, 90, 91, 68, 0, 0,
	0, 79, 69, 70, 54, 185, 138, 0, 121, 167,
	175, 182, 0, 0, 0, 0, 0, 204, 206, 207,
	208, 199, 197, 0, 0, 87, 89, 0, 0, 0,
	122, 180, 0, 210, 211, 203, 212, 72, 73, 0,
	0, 0, 220, 226, 228, 0, 14, 0, 0, 0,
	100, 0, 0, 0, 0, 213, 0, 97, 0, 98,
	99, 227, 0, 0, 101, 0, 214, 0, 229,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 90, 83, 3,
	47, 49, 88, 86, 48, 87, 96, 89, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	72, 71, 73, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 92, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 82, 3, 93,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 74, 75, 76, 77,
	78, 79, 80, 81, 84, 85, 91, 94, 95, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
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
		//line ./go/vt/sqlparser/sql.y:184
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:190
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 14:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:207
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:211
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{Expr: yyDollar[4].valExpr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:215
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:221
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].onDupUpdateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:225
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].onDupUpdateExprs))
			for _, updateList := range yyDollar[6].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].onDupUpdateExprs)}
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:237
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Columns: yyDollar[4].columns, Rows: yyDollar[5].insRows, Replace: true}
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:241
		{
			cols := make(Columns, 0, len(yyDollar[5].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[5].updateExprs))
			for _, updateList := range yyDollar[5].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Columns: cols, Rows: Values{vals}, Replace: true}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:253
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 22:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:259
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:265
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:271
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableIdent}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:275
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:280
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: NewTableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:286
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[4].tableIdent}
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:290
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:295
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: NewTableIdent(yyDollar[3].colIdent.Lowered()), NewName: NewTableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:301
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:307
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:315
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:320
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: NewTableIdent(yyDollar[4].colIdent.Lowered()), IfExists: exists}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:330
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[3].tableIdent}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:336
		{
			yyVAL.statement = &Other{}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:340
		{
			yyVAL.statement = &Other{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:344
		{
			yyVAL.statement = &Other{}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:349
		{
			setAllowComments(yylex, true)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:353
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:359
		{
			yyVAL.bytes2 = nil
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:363
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:369
		{
			yyVAL.str = UnionStr
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:373
		{
			yyVAL.str = UnionAllStr
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:377
		{
			yyVAL.str = UnionDistinctStr
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:382
		{
			yyVAL.str = ""
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:386
		{
			yyVAL.str = DistinctStr
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:391
		{
			yyVAL.str = ""
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:395
		{
			yyVAL.str = StraightJoinHint
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:401
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:405
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:411
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:415
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:419
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:423
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:429
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:433
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:438
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:442
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:446
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:453
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:458
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:462
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:468
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:472
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:482
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:486
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:490
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:503
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:507
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:511
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:515
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:520
		{
			yyVAL.empty = struct{}{}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:522
		{
			yyVAL.empty = struct{}{}
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:525
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:529
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:533
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:540
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:546
		{
			yyVAL.str = JoinStr
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:550
		{
			yyVAL.str = JoinStr
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:554
		{
			yyVAL.str = JoinStr
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:558
		{
			yyVAL.str = StraightJoinStr
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:564
		{
			yyVAL.str = LeftJoinStr
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:568
		{
			yyVAL.str = LeftJoinStr
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:572
		{
			yyVAL.str = RightJoinStr
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:576
		{
			yyVAL.str = RightJoinStr
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:582
		{
			yyVAL.str = NaturalJoinStr
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:586
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:596
		{
			yyVAL.tableName = yyDollar[2].tableName
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:600
		{
			yyVAL.tableName = yyDollar[1].tableName
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:606
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:610
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:615
		{
			yyVAL.indexHints = nil
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:619
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:623
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:627
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:633
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:637
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:642
		{
			yyVAL.boolExpr = nil
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:646
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:653
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:657
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:661
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:665
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:669
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:675
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:679
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:685
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:689
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:693
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:697
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:701
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:705
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:709
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:713
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:717
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:721
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:725
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:729
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:733
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:739
		{
			yyVAL.str = IsNullStr
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:743
		{
			yyVAL.str = IsNotNullStr
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:747
		{
			yyVAL.str = IsTrueStr
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:751
		{
			yyVAL.str = IsNotTrueStr
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:755
		{
			yyVAL.str = IsFalseStr
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:759
		{
			yyVAL.str = IsNotFalseStr
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:765
		{
			yyVAL.str = EqualStr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:769
		{
			yyVAL.str = LessThanStr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:773
		{
			yyVAL.str = GreaterThanStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:777
		{
			yyVAL.str = LessEqualStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:781
		{
			yyVAL.str = GreaterEqualStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:785
		{
			yyVAL.str = NotEqualStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:789
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:795
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:799
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:803
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:809
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:815
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:819
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:825
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:829
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:833
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:837
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:841
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:845
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:849
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:853
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:857
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:861
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:865
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:869
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:873
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:877
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:881
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:885
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:889
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:893
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:901
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
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:915
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:919
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:927
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 166:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:931
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 167:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:935
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:939
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:943
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:947
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:953
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:957
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:961
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:965
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 175:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:971
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:976
		{
			yyVAL.valExpr = nil
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:980
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:986
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:990
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 180:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:996
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1001
		{
			yyVAL.valExpr = nil
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1005
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1011
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1015
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 185:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1019
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1025
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1029
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1033
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1037
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1041
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1045
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1049
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1055
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1064
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1068
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1073
		{
			yyVAL.valExprs = nil
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1077
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1082
		{
			yyVAL.boolExpr = nil
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1086
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1091
		{
			yyVAL.orderBy = nil
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1095
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1101
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1105
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1111
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1116
		{
			yyVAL.str = AscScr
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1120
		{
			yyVAL.str = AscScr
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1124
		{
			yyVAL.str = DescScr
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1129
		{
			yyVAL.limit = nil
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1133
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 210:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1137
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 211:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1141
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1146
		{
			yyVAL.str = ""
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1150
		{
			yyVAL.str = ForUpdateStr
		}
	case 214:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1154
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
		//line ./go/vt/sqlparser/sql.y:1167
		{
			yyVAL.columns = nil
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1171
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1177
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 218:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1181
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1186
		{
			yyVAL.onDupUpdateExprs = nil
		}
	case 220:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1190
		{
			yyVAL.onDupUpdateExprs = yyDollar[5].onDupUpdateExprs
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1196
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1200
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1206
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1210
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1216
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1222
		{
			yyVAL.onDupUpdateExprs = OnDupUpdateExprs{yyDollar[1].onDupUpdateExpr}
		}
	case 227:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1226
		{
			yyVAL.onDupUpdateExprs = append(yyDollar[1].onDupUpdateExprs, yyDollar[3].onDupUpdateExpr)
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1232
		{
			yyVAL.onDupUpdateExpr = &OnDupUpdateExpr{Name: yyDollar[1].updateExpr.Name, Expr: yyDollar[1].updateExpr.Expr, UseLookup: false}
		}
	case 229:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1236
		{
			yyVAL.onDupUpdateExpr = &OnDupUpdateExpr{Name: yyDollar[1].colIdent, LookupName: yyDollar[5].colIdent, UseLookup: true}
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1242
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1246
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1252
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1261
		{
			yyVAL.byt = 0
		}
	case 236:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1263
		{
			yyVAL.byt = 1
		}
	case 237:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1266
		{
			yyVAL.empty = struct{}{}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1268
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1271
		{
			yyVAL.str = ""
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1273
		{
			yyVAL.str = IgnoreStr
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1277
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1279
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1281
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1283
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1285
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1287
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1290
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1292
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1295
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1297
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1300
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1302
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1306
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1312
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1318
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1327
		{
			decNesting(yylex)
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1332
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
