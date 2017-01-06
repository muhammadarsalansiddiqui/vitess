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
	-2, 96,
	-1, 117,
	94, 251,
	-2, 250,
}

const yyNprod = 255
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1064

var yyAct = [...]int{

	129, 228, 448, 65, 231, 345, 47, 336, 250, 392,
	299, 403, 322, 268, 258, 162, 123, 292, 163, 110,
	121, 230, 3, 165, 95, 132, 194, 178, 362, 364,
	77, 70, 48, 49, 184, 67, 157, 111, 72, 41,
	282, 74, 161, 40, 411, 41, 115, 182, 105, 50,
	14, 15, 16, 17, 410, 84, 58, 409, 35, 73,
	37, 71, 46, 124, 38, 43, 44, 45, 186, 42,
	66, 412, 18, 63, 233, 234, 374, 275, 276, 102,
	109, 104, 198, 116, 160, 91, 96, 159, 90, 211,
	67, 368, 363, 155, 68, 117, 60, 464, 67, 201,
	172, 214, 215, 216, 217, 211, 175, 98, 423, 200,
	199, 199, 154, 100, 196, 425, 189, 337, 181, 183,
	180, 62, 232, 306, 87, 201, 201, 235, 236, 237,
	238, 68, 417, 287, 169, 185, 134, 304, 305, 303,
	60, 19, 20, 22, 21, 23, 302, 337, 244, 390,
	375, 376, 377, 117, 24, 25, 26, 61, 166, 212,
	213, 214, 215, 216, 217, 211, 227, 229, 171, 168,
	60, 323, 271, 245, 200, 199, 277, 101, 68, 279,
	197, 455, 323, 92, 200, 199, 280, 78, 68, 97,
	201, 85, 249, 116, 86, 60, 196, 241, 255, 289,
	201, 158, 253, 134, 298, 252, 60, 307, 308, 309,
	274, 311, 312, 313, 314, 315, 316, 317, 318, 319,
	320, 321, 286, 166, 191, 323, 310, 283, 297, 284,
	325, 166, 293, 295, 296, 285, 256, 294, 28, 329,
	14, 116, 116, 324, 326, 430, 290, 291, 67, 343,
	427, 341, 330, 333, 327, 328, 166, 192, 344, 331,
	334, 136, 135, 137, 138, 139, 140, 301, 340, 141,
	325, 323, 349, 94, 351, 174, 350, 289, 352, 134,
	323, 360, 60, 365, 393, 269, 97, 367, 195, 256,
	348, 134, 76, 191, 117, 136, 135, 137, 138, 139,
	140, 158, 277, 141, 134, 378, 380, 381, 382, 95,
	146, 405, 82, 60, 379, 271, 68, 393, 197, 278,
	166, 166, 166, 166, 399, 323, 408, 384, 256, 323,
	125, 126, 108, 407, 116, 145, 97, 127, 79, 128,
	357, 247, 354, 389, 385, 358, 353, 387, 401, 404,
	400, 397, 386, 142, 39, 398, 395, 120, 134, 143,
	144, 55, 260, 263, 264, 265, 261, 301, 262, 266,
	167, 355, 406, 14, 54, 418, 356, 413, 359, 153,
	264, 265, 459, 152, 391, 421, 57, 89, 444, 426,
	88, 419, 424, 176, 460, 120, 120, 339, 277, 422,
	428, 107, 273, 51, 52, 239, 240, 151, 346, 242,
	416, 347, 251, 418, 150, 415, 373, 158, 277, 463,
	453, 441, 442, 439, 14, 28, 30, 443, 1, 272,
	248, 446, 404, 267, 120, 167, 193, 449, 449, 449,
	450, 451, 447, 167, 445, 452, 177, 36, 281, 179,
	69, 149, 434, 435, 67, 342, 462, 461, 246, 458,
	431, 454, 440, 456, 457, 465, 466, 120, 167, 402,
	130, 414, 372, 388, 173, 120, 120, 204, 207, 300,
	243, 335, 131, 122, 220, 221, 222, 223, 224, 225,
	226, 208, 205, 206, 203, 210, 209, 218, 219, 212,
	213, 214, 215, 216, 217, 211, 394, 59, 64, 83,
	338, 202, 118, 361, 259, 120, 120, 75, 257, 164,
	190, 80, 209, 218, 219, 212, 213, 214, 215, 216,
	217, 211, 167, 167, 167, 167, 93, 113, 81, 260,
	263, 264, 265, 261, 99, 262, 266, 53, 103, 27,
	56, 106, 13, 12, 11, 10, 114, 29, 133, 9,
	8, 7, 6, 59, 5, 156, 4, 2, 0, 59,
	64, 170, 0, 31, 32, 33, 34, 0, 0, 300,
	187, 0, 134, 188, 323, 117, 136, 135, 137, 138,
	139, 140, 0, 0, 141, 147, 148, 0, 0, 119,
	0, 146, 0, 0, 0, 0, 0, 0, 120, 0,
	0, 0, 0, 120, 0, 396, 0, 0, 0, 0,
	0, 125, 126, 112, 59, 0, 145, 0, 127, 0,
	128, 0, 0, 0, 59, 254, 0, 0, 0, 0,
	270, 0, 59, 0, 142, 0, 0, 0, 0, 134,
	143, 144, 117, 136, 135, 137, 138, 139, 140, 0,
	0, 141, 147, 148, 0, 0, 114, 59, 146, 0,
	0, 0, 0, 288, 0, 0, 0, 0, 0, 0,
	0, 120, 120, 0, 0, 436, 437, 438, 125, 126,
	0, 120, 0, 145, 0, 127, 0, 128, 0, 0,
	0, 0, 0, 396, 0, 0, 0, 0, 0, 0,
	0, 142, 0, 0, 114, 114, 0, 143, 144, 218,
	219, 212, 213, 214, 215, 216, 217, 211, 432, 433,
	0, 59, 59, 59, 59, 0, 332, 0, 133, 0,
	0, 0, 0, 0, 270, 0, 0, 366, 0, 0,
	0, 369, 0, 370, 0, 0, 371, 0, 0, 0,
	0, 0, 134, 0, 323, 117, 136, 135, 137, 138,
	139, 140, 0, 0, 141, 147, 148, 0, 0, 119,
	0, 146, 0, 0, 0, 133, 210, 209, 218, 219,
	212, 213, 214, 215, 216, 217, 211, 0, 0, 0,
	0, 125, 126, 112, 0, 0, 145, 114, 127, 134,
	128, 0, 117, 136, 135, 137, 138, 139, 140, 0,
	0, 141, 147, 148, 142, 429, 119, 0, 146, 0,
	143, 144, 0, 0, 0, 0, 0, 14, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 125, 126,
	112, 0, 133, 145, 0, 127, 0, 128, 0, 210,
	209, 218, 219, 212, 213, 214, 215, 216, 217, 211,
	0, 142, 0, 0, 0, 0, 134, 143, 144, 117,
	136, 135, 137, 138, 139, 140, 0, 0, 141, 147,
	148, 0, 0, 119, 0, 146, 0, 0, 0, 133,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	14, 0, 0, 0, 0, 125, 126, 0, 0, 0,
	145, 0, 127, 134, 128, 0, 117, 136, 135, 137,
	138, 139, 140, 0, 0, 141, 147, 148, 142, 0,
	119, 0, 146, 0, 143, 144, 0, 0, 0, 134,
	0, 0, 117, 136, 135, 137, 138, 139, 140, 0,
	0, 141, 125, 126, 0, 0, 0, 145, 146, 127,
	0, 128, 0, 0, 0, 0, 0, 0, 0, 420,
	0, 0, 0, 0, 0, 142, 0, 0, 125, 126,
	0, 143, 144, 145, 0, 127, 0, 128, 210, 209,
	218, 219, 212, 213, 214, 215, 216, 217, 211, 383,
	68, 142, 0, 0, 0, 0, 0, 143, 144, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 210, 209,
	218, 219, 212, 213, 214, 215, 216, 217, 211, 0,
	0, 0, 210, 209, 218, 219, 212, 213, 214, 215,
	216, 217, 211, 210, 209, 218, 219, 212, 213, 214,
	215, 216, 217, 211,
}
var yyPact = [...]int{

	44, -1000, -1000, 420, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -44,
	-61, -33, -37, -40, -1000, -1000, -1000, 418, 385, 342,
	-1000, -67, 92, 147, 83, -76, -42, 83, -1000, -43,
	83, -1000, 92, -77, 139, -77, 92, -1000, -1000, -1000,
	-1000, -1000, -1000, 277, 140, -1000, 68, 366, 359, -6,
	-1000, 92, 263, -1000, -8, 143, -1000, 38, -1000, 92,
	51, 129, -1000, 92, -1000, -57, 92, 380, 288, 83,
	-1000, 764, -1000, 397, -1000, 353, 349, -1000, 92, 83,
	92, 406, -22, -10, 158, 92, 48, 83, 246, -1000,
	372, -82, -1000, 20, -1000, 92, -1000, -1000, 92, -1000,
	247, -1000, -1000, 268, -12, 114, 415, -1000, -1000, 878,
	831, -1000, -21, -1000, -1000, 246, 246, 246, 246, 259,
	259, -1000, -1000, 259, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 246, -1000, -1000, 92,
	-1000, -1000, -1000, -1000, 313, 290, -1000, 398, 878, 158,
	48, 190, 505, -1000, -1000, 265, 382, 234, -1000, -1000,
	-17, -1000, 973, -16, 904, -1000, -1000, 275, 83, -1000,
	-65, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	406, 764, 158, -1000, -1000, 130, -1000, -1000, 47, 878,
	878, 175, 604, 91, 60, 246, 246, 246, 175, 246,
	246, 246, 246, 246, 246, 246, 246, 246, 246, 246,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 23, 415, 124,
	233, 224, 415, 212, 212, -1000, -1000, -1000, 962, 717,
	537, -1000, 418, 52, 973, -1000, 367, 83, 83, 398,
	392, 396, 114, 190, -17, -1000, 158, 158, 158, 158,
	-1000, 312, 308, -1000, 337, 306, 344, -14, -1000, 122,
	-1000, -1000, 92, -1000, 282, 5, 105, 973, 92, -1000,
	-1000, 92, -1000, 404, -1000, 243, -1000, -1000, -18, -1000,
	23, 50, -1000, -1000, 93, -1000, -1000, -1000, 973, -1000,
	904, -1000, -1000, 91, 246, 246, 246, 973, 973, 948,
	-1000, 637, 441, -1000, 15, 15, -1, -1, -1, -1,
	75, 75, -1000, -1000, -1000, 246, -1000, -1000, -1000, -1000,
	-1000, 178, 764, -1000, 178, 82, -1000, 878, 273, 259,
	420, 240, 278, -1000, 392, -1000, 246, 246, -1000, 505,
	267, 328, -1000, -1000, -1000, -1000, 299, -1000, 292, -1000,
	-1000, -1000, -46, -49, -59, -1000, -1000, -1000, -1000, -23,
	-1000, -1000, 402, 395, 46, -1000, -1000, -1000, 224, -1000,
	973, 973, 918, 246, 973, -1000, 178, -1000, 40, -1000,
	246, 49, -1000, 364, 204, -1000, 246, -1000, -1000, 83,
	-1000, 779, 199, -1000, 706, 878, 878, -1000, -1000, 259,
	259, 259, 83, -1000, 398, 878, 246, -1000, -1000, -1000,
	246, 973, -1000, -1000, 973, 246, 362, 259, -1000, 246,
	246, -1000, -1000, -1000, 114, 114, 83, 83, 83, 392,
	114, 184, 973, 973, 412, -1000, 973, -1000, 135, -1000,
	135, 135, 365, 83, -1000, 83, -1000, -1000, -1000, 411,
	18, 143, -1000, -1000, 83, 83, -1000,
}
var yyPgo = [...]int{

	0, 567, 21, 566, 564, 562, 561, 560, 559, 555,
	554, 553, 552, 557, 550, 549, 547, 538, 121, 73,
	19, 37, 537, 520, 42, 15, 18, 519, 518, 14,
	514, 23, 513, 2, 36, 46, 512, 25, 511, 510,
	20, 1, 509, 17, 10, 4, 506, 16, 63, 483,
	482, 481, 7, 480, 473, 472, 471, 470, 8, 469,
	11, 460, 5, 459, 458, 455, 9, 3, 70, 451,
	354, 292, 450, 449, 448, 447, 446, 0, 26, 436,
	474, 13, 433, 429, 6, 428, 426, 275, 12,
}
var yyR1 = [...]int{

	0, 85, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 3, 3, 4, 5,
	5, 5, 6, 7, 7, 7, 8, 8, 8, 9,
	10, 10, 10, 11, 12, 12, 12, 86, 13, 14,
	14, 15, 15, 15, 16, 16, 17, 17, 18, 18,
	19, 19, 19, 20, 20, 21, 21, 21, 21, 22,
	22, 79, 79, 79, 78, 78, 23, 23, 24, 24,
	25, 25, 26, 26, 26, 27, 27, 27, 27, 83,
	83, 82, 82, 82, 81, 81, 28, 28, 28, 28,
	29, 29, 29, 29, 30, 30, 31, 31, 32, 32,
	32, 32, 33, 33, 34, 34, 35, 35, 35, 35,
	35, 35, 37, 37, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 43, 43, 43,
	43, 43, 43, 38, 38, 38, 38, 38, 38, 38,
	44, 44, 44, 48, 45, 45, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 57, 57, 57, 57, 50, 53, 53,
	51, 51, 52, 54, 54, 49, 49, 49, 40, 40,
	40, 40, 40, 40, 40, 42, 42, 42, 55, 55,
	56, 56, 58, 58, 59, 59, 60, 61, 61, 61,
	62, 62, 62, 63, 63, 63, 64, 64, 65, 65,
	66, 66, 39, 39, 46, 46, 47, 67, 67, 68,
	69, 69, 71, 71, 72, 72, 70, 70, 73, 73,
	73, 73, 73, 73, 74, 74, 75, 75, 76, 76,
	77, 80, 87, 88, 84,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 12, 6, 3, 8, 8, 8, 7,
	6, 7, 3, 5, 8, 4, 6, 7, 4, 5,
	4, 5, 5, 3, 2, 2, 2, 0, 2, 0,
	2, 1, 2, 2, 0, 1, 0, 1, 1, 3,
	1, 3, 5, 1, 3, 1, 2, 3, 5, 1,
	1, 0, 1, 2, 1, 1, 0, 2, 1, 3,
	1, 1, 3, 3, 3, 3, 5, 5, 3, 0,
	1, 0, 1, 2, 1, 1, 1, 2, 2, 1,
	2, 3, 2, 3, 2, 2, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 1, 3, 3, 2,
	3, 3, 1, 1, 1, 3, 3, 3, 4, 3,
	4, 3, 4, 5, 6, 3, 2, 1, 2, 1,
	2, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 1, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 3, 3, 4, 5,
	3, 4, 1, 1, 1, 1, 1, 5, 0, 1,
	1, 2, 4, 0, 2, 1, 3, 5, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 0, 2, 4, 0, 3, 1, 3,
	0, 5, 2, 1, 1, 3, 3, 1, 3, 3,
	1, 1, 0, 2, 0, 3, 0, 1, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 1, 1, 1, 0,
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
	-80, -68, -41, -80, -87, -84, 21, -76, 109, -73,
	100, 98, 27, 99, 14, 115, 48, -80, -80, -84,
	-23, 46, 10, -79, -78, 20, -77, 50, 94, 61,
	60, 76, -38, 79, 62, 77, 78, 63, 76, 81,
	80, 90, 84, 85, 86, 87, 88, 89, 82, 83,
	69, 70, 71, 72, 73, 74, 75, -35, -41, -35,
	-2, -45, -41, 95, 96, -41, -41, -41, -41, -87,
	-87, -48, -87, -53, -41, -31, -64, 28, -87, -34,
	-58, 14, -35, -24, -80, -34, 46, -28, -29, -30,
	34, 38, 40, 35, 36, 37, 41, -82, -81, 20,
	-80, 50, -83, 20, -24, 94, 94, -41, 44, -77,
	-84, -74, 105, -34, -21, -24, -78, 86, -80, -77,
	-35, -35, -43, 57, 62, 58, 59, -37, -41, -44,
	-87, -48, 55, 79, 77, 78, 63, -41, -41, -41,
	-43, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -88, 47, -88, 46, -88, -40, -40, -77,
	-88, -20, 19, -88, -20, -51, -52, 65, -39, 30,
	-2, -67, -65, -77, -58, -62, 16, 15, -34, -25,
	-26, -25, -26, 34, 34, 34, 39, 34, 39, 34,
	-29, -32, 42, 106, 43, -81, -80, -88, 86, -80,
	-80, -80, -55, 12, 94, 57, 58, 59, -45, -44,
	-41, -41, -41, 61, -41, -88, -20, -88, -54, -52,
	67, -35, -66, 44, -46, -47, -87, -66, -88, 46,
	-62, -41, -59, -60, -41, 44, 44, 34, 34, 103,
	103, 103, 94, -84, -56, 13, 15, 86, -77, -88,
	61, -41, -88, 68, -41, 66, 25, 46, -77, 46,
	46, -61, 22, 23, -35, -35, -87, -87, -87, -58,
	-35, -45, -41, -41, 26, -47, -41, -60, -33, -77,
	-33, -33, -62, 8, -88, 46, -88, -88, -63, 17,
	29, -67, -77, 8, 79, -77, -77,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 37, 37, 37, 37, 37, 246,
	236, 0, 0, 0, 254, 254, 254, 0, 41, 44,
	39, 236, 0, 0, 0, 234, 0, 0, 247, 0,
	0, 237, 0, 232, 0, 232, 0, 34, 35, 36,
	15, 42, 43, 46, 0, 45, 38, 0, 0, 96,
	251, 0, 0, 48, 50, 22, 227, 0, 250, 0,
	0, 0, 254, 0, 254, 0, 0, 0, 0, 0,
	33, 0, 47, 0, 195, 0, 0, 40, 0, 0,
	0, 104, 0, -2, 0, 0, 0, 0, 0, 254,
	0, 248, 25, 0, 28, 0, 30, 233, 0, 254,
	66, 53, 55, 61, 0, 59, 60, -2, 106, 0,
	0, 146, 147, 148, 149, 0, 0, 0, 0, 185,
	0, 172, 114, 0, 252, 188, 189, 190, 191, 192,
	193, 194, 173, 174, 175, 176, 178, 112, 113, 0,
	230, 231, 196, 197, 216, 104, 97, 202, 0, 0,
	0, 104, 68, 70, 71, 81, 79, 0, 49, 51,
	0, 228, 229, 0, 0, 23, 235, 0, 0, 254,
	244, 238, 239, 240, 241, 242, 243, 29, 31, 32,
	104, 0, 0, 56, 62, 0, 64, 65, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	133, 134, 135, 136, 137, 138, 139, 109, 0, 0,
	0, 0, 144, 0, 0, 163, 164, 165, 0, 0,
	0, 126, 0, 0, 179, 14, 0, 0, 0, 202,
	210, 0, 105, 104, 97, 20, 0, 0, 0, 0,
	86, 0, 0, 89, 0, 0, 0, 98, 82, 0,
	84, 85, 0, 80, 0, 0, 0, 144, 0, 249,
	26, 0, 245, 198, 54, 67, 63, 57, 0, 186,
	107, 108, 111, 127, 0, 129, 131, 115, 116, 117,
	0, 141, 142, 0, 0, 0, 0, 119, 121, 0,
	125, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 110, 253, 143, 0, 226, 161, 162, 166,
	167, 0, 0, 170, 0, 183, 180, 0, 220, 0,
	223, 220, 0, 218, 210, 19, 0, 0, 21, 69,
	75, 0, 78, 87, 88, 90, 0, 92, 0, 94,
	95, 72, 0, 0, 0, 83, 73, 74, 52, 0,
	254, 27, 200, 0, 0, 128, 130, 132, 0, 118,
	120, 122, 0, 0, 145, 168, 0, 171, 0, 181,
	0, 0, 16, 0, 222, 224, 0, 17, 217, 0,
	18, 211, 203, 204, 207, 0, 0, 91, 93, 0,
	0, 0, 0, 24, 202, 0, 0, 58, 187, 140,
	0, 123, 169, 177, 184, 0, 0, 0, 219, 0,
	0, 206, 208, 209, 76, 77, 0, 0, 0, 210,
	201, 199, 124, 182, 0, 225, 212, 205, 0, 102,
	0, 0, 213, 0, 99, 0, 100, 101, 13, 0,
	0, 221, 103, 214, 0, 0, 215,
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
			yyVAL.statement = &DeleteMulti{Comments: Comments(yyDollar[2].bytes2), DeleteExprs: yyDollar[3].deleteExprs, From: yyDollar[5].tableExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), Using: false}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:251
		{
			yyVAL.statement = &DeleteMulti{Comments: Comments(yyDollar[2].bytes2), DeleteExprs: yyDollar[4].deleteExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), Using: true}
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
			yyVAL.deleteExpr = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:407
		{
			yyVAL.deleteExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:411
		{
			yyVAL.deleteExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:417
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:421
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:427
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:431
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:435
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:439
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:445
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:449
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:454
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:458
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:462
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:469
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:474
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:478
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:484
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:488
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:498
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:502
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:506
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:519
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:523
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:527
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:531
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:536
		{
			yyVAL.empty = struct{}{}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:538
		{
			yyVAL.empty = struct{}{}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:541
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:545
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:549
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:556
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:570
		{
			yyVAL.str = JoinStr
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:574
		{
			yyVAL.str = StraightJoinStr
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:580
		{
			yyVAL.str = LeftJoinStr
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:584
		{
			yyVAL.str = LeftJoinStr
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:588
		{
			yyVAL.str = RightJoinStr
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:592
		{
			yyVAL.str = RightJoinStr
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:598
		{
			yyVAL.str = NaturalJoinStr
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:602
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:612
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:616
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:621
		{
			yyVAL.indexHints = nil
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:625
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:629
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:633
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:639
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:643
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:648
		{
			yyVAL.boolExpr = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:652
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:659
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:663
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:667
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:671
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:675
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:681
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:685
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:691
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:695
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:699
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:703
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:707
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:711
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:715
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:719
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:723
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:727
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:731
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:735
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:739
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:745
		{
			yyVAL.str = IsNullStr
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:749
		{
			yyVAL.str = IsNotNullStr
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:753
		{
			yyVAL.str = IsTrueStr
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:757
		{
			yyVAL.str = IsNotTrueStr
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:761
		{
			yyVAL.str = IsFalseStr
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:765
		{
			yyVAL.str = IsNotFalseStr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:771
		{
			yyVAL.str = EqualStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:775
		{
			yyVAL.str = LessThanStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:779
		{
			yyVAL.str = GreaterThanStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:783
		{
			yyVAL.str = LessEqualStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:787
		{
			yyVAL.str = GreaterEqualStr
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:791
		{
			yyVAL.str = NotEqualStr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:795
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:801
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:805
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:809
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:815
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:821
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:825
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:831
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:835
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:839
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:843
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:847
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:851
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:855
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:859
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:863
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:867
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:871
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
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
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:883
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:887
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:891
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:895
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:899
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:907
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
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:921
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:925
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:933
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 168:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:937
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 169:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:941
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:945
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:949
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:953
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:959
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:963
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:967
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:971
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 177:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:977
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:982
		{
			yyVAL.valExpr = nil
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:986
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:992
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:996
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1002
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1007
		{
			yyVAL.valExpr = nil
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1011
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1017
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1021
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 187:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1025
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1031
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1035
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1039
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1043
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1047
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1051
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1055
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1061
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1070
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1074
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1079
		{
			yyVAL.valExprs = nil
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1083
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1088
		{
			yyVAL.boolExpr = nil
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1092
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1097
		{
			yyVAL.orderBy = nil
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1101
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1107
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1111
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1117
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1122
		{
			yyVAL.str = AscScr
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1126
		{
			yyVAL.str = AscScr
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1130
		{
			yyVAL.str = DescScr
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1135
		{
			yyVAL.limit = nil
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1139
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 212:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1143
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1148
		{
			yyVAL.str = ""
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1152
		{
			yyVAL.str = ForUpdateStr
		}
	case 215:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1156
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
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1169
		{
			yyVAL.columns = nil
		}
	case 217:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1173
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1179
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 219:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1183
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1188
		{
			yyVAL.updateExprs = nil
		}
	case 221:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1192
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1198
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1202
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1208
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1212
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1218
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1224
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1228
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1234
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1243
		{
			yyVAL.byt = 0
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1245
		{
			yyVAL.byt = 1
		}
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1248
		{
			yyVAL.empty = struct{}{}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1250
		{
			yyVAL.empty = struct{}{}
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1253
		{
			yyVAL.str = ""
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1255
		{
			yyVAL.str = IgnoreStr
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1259
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1261
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1265
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1267
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1269
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1272
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1274
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1277
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1279
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1282
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1284
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1288
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1294
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1300
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1309
		{
			decNesting(yylex)
		}
	case 254:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1314
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
