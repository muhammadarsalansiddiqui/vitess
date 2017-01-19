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
const UPDATE = 57350
const DELETE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const OFFSET = 57359
const FOR = 57360
const ALL = 57361
const DISTINCT = 57362
const AS = 57363
const EXISTS = 57364
const ASC = 57365
const DESC = 57366
const INTO = 57367
const DUPLICATE = 57368
const KEY = 57369
const DEFAULT = 57370
const SET = 57371
const LOCK = 57372
const VALUES = 57373
const LAST_INSERT_ID = 57374
const NEXT = 57375
const VALUE = 57376
const JOIN = 57377
const STRAIGHT_JOIN = 57378
const LEFT = 57379
const RIGHT = 57380
const INNER = 57381
const OUTER = 57382
const CROSS = 57383
const NATURAL = 57384
const USE = 57385
const FORCE = 57386
const ON = 57387
const ID = 57388
const HEX = 57389
const STRING = 57390
const INTEGRAL = 57391
const FLOAT = 57392
const HEXNUM = 57393
const VALUE_ARG = 57394
const LIST_ARG = 57395
const COMMENT = 57396
const NULL = 57397
const TRUE = 57398
const FALSE = 57399
const OR = 57400
const AND = 57401
const NOT = 57402
const BETWEEN = 57403
const CASE = 57404
const WHEN = 57405
const THEN = 57406
const ELSE = 57407
const END = 57408
const LE = 57409
const GE = 57410
const NE = 57411
const NULL_SAFE_EQUAL = 57412
const IS = 57413
const LIKE = 57414
const REGEXP = 57415
const IN = 57416
const SHIFT_LEFT = 57417
const SHIFT_RIGHT = 57418
const MOD = 57419
const UNARY = 57420
const INTERVAL = 57421
const JSON_EXTRACT_OP = 57422
const JSON_UNQUOTE_EXTRACT_OP = 57423
const CREATE = 57424
const ALTER = 57425
const DROP = 57426
const RENAME = 57427
const ANALYZE = 57428
const TABLE = 57429
const INDEX = 57430
const VIEW = 57431
const TO = 57432
const IGNORE = 57433
const IF = 57434
const UNIQUE = 57435
const USING = 57436
const SHOW = 57437
const DESCRIBE = 57438
const EXPLAIN = 57439
const CURRENT_TIMESTAMP = 57440
const DATABASE = 57441
const UNUSED = 57442

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
	-1, 111,
	95, 251,
	-2, 250,
}

const yyNprod = 255
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1058

var yyAct = [...]int{

	213, 441, 435, 369, 63, 318, 117, 216, 238, 338,
	252, 328, 251, 270, 309, 62, 250, 307, 263, 293,
	215, 3, 104, 109, 35, 126, 37, 254, 163, 179,
	38, 115, 74, 67, 389, 391, 41, 105, 159, 40,
	421, 41, 43, 44, 45, 169, 247, 99, 50, 420,
	419, 68, 70, 46, 42, 371, 123, 47, 154, 167,
	58, 195, 194, 203, 204, 197, 198, 199, 200, 201,
	202, 196, 218, 219, 199, 200, 201, 202, 196, 110,
	171, 345, 448, 48, 49, 87, 186, 241, 183, 90,
	89, 64, 196, 157, 69, 446, 156, 71, 390, 92,
	308, 65, 111, 118, 152, 294, 401, 94, 84, 185,
	184, 81, 86, 294, 151, 217, 128, 114, 185, 184,
	220, 221, 222, 223, 150, 186, 273, 96, 277, 98,
	166, 168, 165, 103, 186, 184, 60, 212, 214, 395,
	258, 229, 275, 276, 274, 64, 128, 170, 64, 60,
	186, 160, 114, 114, 308, 65, 361, 346, 347, 348,
	242, 174, 224, 225, 181, 234, 227, 60, 60, 233,
	341, 230, 111, 264, 266, 267, 95, 110, 265, 240,
	197, 198, 199, 200, 201, 202, 196, 75, 269, 339,
	445, 278, 279, 280, 114, 282, 283, 284, 285, 286,
	287, 288, 289, 290, 291, 292, 64, 236, 261, 262,
	257, 237, 281, 268, 249, 114, 256, 60, 296, 341,
	244, 408, 245, 114, 114, 110, 110, 271, 65, 364,
	182, 226, 450, 294, 248, 295, 297, 181, 128, 312,
	260, 176, 294, 326, 301, 304, 317, 302, 305, 314,
	298, 299, 185, 184, 65, 315, 91, 82, 403, 415,
	83, 180, 310, 114, 114, 296, 294, 14, 186, 326,
	294, 313, 242, 344, 316, 294, 351, 352, 353, 349,
	300, 255, 130, 129, 131, 132, 133, 134, 350, 65,
	135, 182, 272, 28, 310, 256, 91, 355, 260, 330,
	333, 334, 335, 331, 110, 332, 336, 128, 155, 416,
	60, 243, 149, 271, 242, 73, 102, 79, 147, 418,
	367, 370, 356, 366, 360, 358, 357, 410, 411, 128,
	384, 417, 362, 177, 381, 385, 294, 380, 377, 376,
	379, 378, 114, 386, 91, 334, 335, 114, 387, 392,
	330, 333, 334, 335, 331, 399, 332, 336, 146, 88,
	255, 76, 402, 363, 394, 256, 256, 256, 256, 397,
	176, 405, 404, 365, 382, 14, 242, 400, 272, 383,
	372, 161, 412, 414, 55, 195, 194, 203, 204, 197,
	198, 199, 200, 201, 202, 196, 101, 54, 413, 422,
	232, 343, 396, 313, 423, 439, 39, 425, 426, 370,
	51, 52, 427, 114, 319, 375, 320, 440, 428, 203,
	204, 197, 198, 199, 200, 201, 202, 196, 396, 436,
	255, 255, 255, 255, 443, 444, 145, 158, 57, 429,
	430, 239, 374, 325, 144, 155, 61, 157, 453, 29,
	436, 447, 424, 14, 114, 114, 28, 30, 431, 432,
	433, 449, 1, 451, 452, 31, 32, 33, 34, 342,
	59, 337, 178, 162, 36, 246, 164, 66, 143, 460,
	72, 437, 434, 235, 77, 148, 438, 409, 442, 442,
	442, 368, 124, 457, 373, 59, 324, 359, 228, 59,
	407, 306, 437, 125, 93, 455, 116, 456, 97, 311,
	80, 100, 458, 303, 459, 127, 108, 231, 187, 112,
	388, 85, 329, 327, 59, 253, 175, 153, 107, 78,
	406, 53, 27, 56, 13, 12, 11, 172, 10, 128,
	173, 294, 111, 130, 129, 131, 132, 133, 134, 9,
	8, 135, 141, 142, 7, 6, 113, 5, 140, 4,
	2, 0, 127, 0, 195, 194, 203, 204, 197, 198,
	199, 200, 201, 202, 196, 0, 0, 0, 119, 120,
	106, 59, 0, 139, 0, 121, 128, 122, 294, 111,
	130, 129, 131, 132, 133, 134, 0, 0, 135, 141,
	142, 136, 0, 113, 0, 140, 0, 137, 138, 0,
	0, 0, 0, 0, 108, 59, 0, 0, 0, 0,
	0, 259, 0, 0, 0, 119, 120, 106, 0, 0,
	139, 0, 121, 0, 122, 0, 0, 127, 194, 203,
	204, 197, 198, 199, 200, 201, 202, 196, 136, 0,
	0, 0, 0, 0, 137, 138, 0, 0, 0, 0,
	0, 128, 108, 108, 111, 130, 129, 131, 132, 133,
	134, 0, 0, 135, 141, 142, 0, 0, 113, 321,
	140, 322, 0, 0, 323, 0, 0, 0, 0, 0,
	0, 0, 340, 0, 59, 0, 0, 0, 0, 14,
	119, 120, 106, 0, 0, 139, 0, 121, 0, 122,
	0, 0, 0, 0, 0, 127, 0, 14, 15, 16,
	17, 0, 0, 136, 0, 0, 0, 0, 0, 137,
	138, 0, 0, 0, 0, 0, 0, 0, 0, 128,
	18, 108, 111, 130, 129, 131, 132, 133, 134, 0,
	0, 135, 141, 142, 0, 0, 113, 0, 140, 0,
	0, 0, 127, 0, 59, 59, 59, 59, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 340, 119, 120,
	393, 0, 0, 139, 0, 121, 128, 122, 0, 111,
	130, 129, 131, 132, 133, 134, 0, 0, 135, 141,
	142, 136, 0, 113, 0, 140, 0, 137, 138, 19,
	20, 22, 21, 23, 0, 0, 0, 0, 0, 0,
	0, 0, 24, 25, 26, 119, 120, 0, 0, 0,
	139, 0, 121, 128, 122, 0, 111, 130, 129, 131,
	132, 133, 134, 0, 454, 135, 141, 142, 136, 0,
	65, 0, 140, 0, 137, 138, 0, 0, 0, 128,
	0, 0, 111, 130, 129, 131, 132, 133, 134, 0,
	14, 135, 119, 120, 0, 0, 0, 139, 140, 121,
	0, 122, 195, 194, 203, 204, 197, 198, 199, 200,
	201, 202, 196, 0, 0, 136, 0, 0, 119, 120,
	0, 137, 138, 139, 0, 121, 0, 122, 0, 0,
	128, 0, 0, 111, 130, 129, 131, 132, 133, 134,
	0, 136, 135, 0, 0, 0, 0, 137, 138, 140,
	0, 0, 0, 0, 0, 0, 128, 0, 0, 111,
	130, 129, 131, 132, 133, 134, 0, 0, 135, 119,
	120, 0, 0, 0, 139, 140, 121, 0, 122, 0,
	0, 0, 0, 0, 0, 0, 398, 0, 0, 0,
	0, 0, 136, 0, 0, 119, 120, 0, 137, 138,
	139, 0, 121, 0, 122, 195, 194, 203, 204, 197,
	198, 199, 200, 201, 202, 196, 0, 0, 136, 189,
	192, 0, 0, 0, 137, 138, 205, 206, 207, 208,
	209, 210, 211, 193, 190, 191, 188, 195, 194, 203,
	204, 197, 198, 199, 200, 201, 202, 196, 354, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 195, 194, 203,
	204, 197, 198, 199, 200, 201, 202, 196,
}
var yyPact = [...]int{

	711, -1000, -1000, 451, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -79,
	-66, -49, -61, -50, -1000, -1000, -1000, 447, 391, 364,
	-1000, -71, 118, 436, 106, -75, -53, 106, -1000, -51,
	106, -1000, 118, -76, 138, -76, 118, -1000, -1000, -1000,
	-1000, -1000, -1000, 281, 205, -1000, 51, 87, 330, -5,
	-1000, 118, 209, -1000, 29, -1000, 118, 44, 127, -1000,
	118, -1000, -59, 118, 374, 271, 106, -1000, 615, -1000,
	426, -1000, 327, 287, -1000, 283, 118, -1000, 106, 118,
	434, 106, 890, -1000, 359, -82, -1000, 31, -1000, 118,
	-1000, -1000, 118, -1000, 323, -1000, -1000, 240, -7, 48,
	936, -1000, -1000, 740, 693, -1000, -24, -1000, -1000, 890,
	890, 890, 890, 192, 192, -1000, -1000, 192, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	890, -1000, -1000, 118, -1000, -1000, -1000, -1000, 369, 106,
	106, -1000, 297, -1000, 427, 740, -1000, -20, -8, 864,
	-1000, -1000, 266, 106, -1000, -60, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 434, 615, 100, -1000, -1000,
	179, -1000, -1000, 53, 740, 740, 115, 787, 70, 64,
	890, 890, 890, 115, 890, 890, 890, 890, 890, 890,
	890, 890, 890, 890, 890, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 9, 936, 57, 288, 218, 936, 232, 232,
	-1000, -1000, -1000, 801, 493, 540, -1000, 447, 34, -20,
	-1000, 217, 192, 451, 249, 227, -1000, 427, 398, 401,
	48, 123, -20, 118, -1000, -1000, 118, -1000, 431, -1000,
	196, 315, -1000, -1000, 168, 380, 261, -1000, -1000, -14,
	-1000, 9, 73, -1000, -1000, 99, -1000, -1000, -1000, -20,
	-1000, 864, -1000, -1000, 70, 890, 890, 890, -20, -20,
	966, -1000, 336, 556, -1000, -13, -13, 1, 1, 1,
	1, 95, 95, -1000, -1000, -1000, 890, -1000, -1000, -1000,
	-1000, -1000, 194, 615, -1000, 194, 88, -1000, 740, -1000,
	337, 182, -1000, 890, -1000, -1000, 106, 398, -1000, 890,
	890, -40, -1000, -1000, 429, 400, 100, 100, 100, 100,
	-1000, 302, 299, -1000, 339, 295, 308, -9, -1000, 119,
	-1000, -1000, 118, -1000, 222, 52, -1000, -1000, -1000, 218,
	-1000, -20, -20, 904, 890, -20, -1000, 194, -1000, 37,
	-1000, 890, 191, 345, 192, -1000, -1000, 483, 174, -1000,
	304, 106, -1000, 427, 740, 890, 315, 214, 264, -1000,
	-1000, -1000, -1000, 296, -1000, 284, -1000, -1000, -1000, -54,
	-55, -64, -1000, -1000, -1000, -1000, -1000, -1000, 890, -20,
	-1000, -1000, -20, 890, 444, -1000, 890, 890, 890, -1000,
	-1000, -1000, 398, 48, 171, 740, 740, -1000, -1000, 192,
	192, 192, -20, -20, 106, -20, -20, -1000, 387, 48,
	48, 106, 106, 106, 143, -1000, -1000, 25, -1000, 443,
	2, 185, -1000, 185, 185, 106, 813, -1000, 106, -1000,
	106, -1000, -1000, -1000, 192, 106, -1000, 106, -1000, 65,
	-1000,
}
var yyPgo = [...]int{

	0, 560, 20, 559, 557, 555, 554, 550, 549, 538,
	536, 535, 534, 449, 533, 532, 531, 529, 22, 37,
	528, 526, 16, 12, 10, 525, 523, 11, 522, 27,
	521, 520, 1, 58, 23, 519, 25, 518, 517, 31,
	0, 510, 18, 13, 7, 509, 6, 103, 506, 503,
	501, 17, 498, 497, 496, 494, 492, 8, 491, 3,
	487, 5, 486, 485, 483, 14, 482, 2, 15, 4,
	478, 406, 315, 477, 476, 475, 474, 473, 56, 29,
	472, 437, 9, 471, 469, 57, 462, 457, 38, 19,
}
var yyR1 = [...]int{

	0, 86, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 3, 3, 4, 5,
	6, 7, 7, 7, 8, 8, 8, 9, 10, 10,
	10, 11, 12, 12, 12, 87, 13, 14, 14, 15,
	15, 15, 16, 16, 17, 17, 18, 18, 19, 19,
	19, 19, 20, 20, 80, 80, 80, 79, 79, 21,
	21, 22, 22, 23, 23, 24, 24, 24, 25, 25,
	25, 25, 84, 84, 83, 83, 83, 82, 82, 26,
	26, 26, 26, 27, 27, 27, 27, 28, 28, 30,
	30, 29, 29, 31, 31, 31, 31, 32, 32, 33,
	33, 34, 34, 34, 34, 34, 34, 36, 36, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 42, 42, 42, 42, 42, 42, 37, 37,
	37, 37, 37, 37, 37, 43, 43, 43, 47, 44,
	44, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 56, 56,
	56, 56, 49, 52, 52, 50, 50, 51, 53, 53,
	48, 48, 48, 39, 39, 39, 39, 39, 39, 39,
	41, 41, 41, 54, 54, 55, 55, 57, 57, 58,
	58, 59, 60, 60, 60, 61, 61, 61, 61, 62,
	62, 62, 63, 63, 64, 64, 65, 65, 38, 38,
	45, 45, 46, 66, 66, 67, 67, 68, 68, 69,
	70, 70, 72, 72, 73, 73, 71, 71, 74, 74,
	74, 74, 74, 74, 75, 75, 76, 76, 77, 77,
	78, 81, 88, 89, 85,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 12, 6, 3, 7, 7, 8, 7,
	3, 5, 8, 4, 6, 7, 4, 5, 4, 5,
	5, 3, 2, 2, 2, 0, 2, 0, 2, 1,
	2, 2, 0, 1, 0, 1, 1, 3, 1, 2,
	3, 5, 1, 1, 0, 1, 2, 1, 1, 0,
	2, 1, 3, 1, 1, 3, 3, 3, 3, 5,
	5, 3, 0, 1, 0, 1, 2, 1, 1, 1,
	2, 2, 1, 2, 3, 2, 3, 2, 2, 2,
	1, 1, 3, 0, 5, 5, 5, 1, 3, 0,
	2, 1, 3, 3, 2, 3, 3, 1, 1, 1,
	3, 3, 3, 4, 3, 4, 3, 4, 5, 6,
	3, 2, 1, 2, 1, 2, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 3, 3, 4, 5, 3, 4, 1, 1, 1,
	1, 1, 5, 0, 1, 1, 2, 4, 0, 2,
	1, 3, 5, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 2, 1,
	1, 3, 3, 1, 3, 1, 6, 1, 3, 3,
	1, 1, 0, 2, 0, 3, 0, 1, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -86, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, 6, 7, 8, 9, 29, 98,
	99, 101, 100, 102, 111, 112, 113, -15, 5, -13,
	-87, -13, -13, -13, -13, 103, -76, 105, 109, -71,
	105, 107, 103, 103, 104, 105, 103, -85, -85, -85,
	-2, 19, 20, -16, 33, 20, -14, -71, -29, -81,
	49, 10, -68, -69, -78, 49, -73, 108, 104, -78,
	103, -78, -81, -72, 108, 49, -72, -81, -17, 36,
	-41, -78, 52, 55, 57, -30, 25, -29, 29, 95,
	-29, 47, 70, -81, 63, 49, -85, -81, -85, 106,
	-81, 22, 45, -78, -18, -19, 87, -20, -81, -34,
	-40, 49, -35, 63, -88, -39, -48, -46, -47, 85,
	86, 92, 94, -78, -56, -49, -36, 22, 46, 51,
	50, 52, 53, 54, 55, 58, 108, 114, 115, 90,
	65, 59, 60, -70, 18, 10, 31, 31, -63, 29,
	-88, -29, -68, -81, -33, 11, -69, -40, -81, -88,
	-85, 22, -77, 110, -74, 101, 99, 28, 100, 14,
	116, 49, -81, -81, -85, -21, 47, 10, -80, -79,
	21, -78, 51, 95, 62, 61, 77, -37, 80, 63,
	78, 79, 64, 77, 82, 81, 91, 85, 86, 87,
	88, 89, 90, 83, 84, 70, 71, 72, 73, 74,
	75, 76, -34, -40, -34, -2, -44, -40, 96, 97,
	-40, -40, -40, -40, -88, -88, -47, -88, -52, -40,
	-29, -38, 31, -2, -68, -64, -78, -33, -57, 14,
	-34, 95, -40, 45, -78, -85, -75, 106, -33, -19,
	-22, -23, -24, -25, -29, -47, -88, -79, 87, -81,
	-78, -34, -34, -42, 58, 63, 59, 60, -36, -40,
	-43, -88, -47, 56, 80, 78, 79, 64, -40, -40,
	-40, -42, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -89, 48, -89, 47, -89, -39, -39,
	-78, -89, -18, 20, -89, -18, -50, -51, 66, -65,
	45, -45, -46, -88, -65, -89, 47, -57, -61, 16,
	15, -81, -81, -81, -54, 12, 47, -26, -27, -28,
	35, 39, 41, 36, 37, 38, 42, -83, -82, 21,
	-81, 51, -84, 21, -22, 95, 58, 59, 60, -44,
	-43, -40, -40, -40, 62, -40, -89, -18, -89, -53,
	-51, 68, -34, 26, 47, -78, -61, -40, -58, -59,
	-40, 95, -85, -55, 13, 15, -23, -24, -23, -24,
	35, 35, 35, 40, 35, 40, 35, -27, -31, 43,
	107, 44, -82, -81, -89, 87, -78, -89, 62, -40,
	-89, 69, -40, 67, 27, -46, 47, 17, 47, -60,
	23, 24, -57, -34, -44, 45, 45, 35, 35, 104,
	104, 104, -40, -40, 8, -40, -40, -59, -61, -34,
	-34, -88, -88, -88, -66, -67, -69, -78, -62, 18,
	30, -32, -78, -32, -32, 47, 70, 8, 80, -89,
	47, -89, -89, -67, 31, -78, -78, -88, -78, -78,
	-89,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 35, 35, 35, 35, 35, 246,
	236, 0, 0, 0, 254, 254, 254, 0, 39, 42,
	37, 236, 0, 0, 0, 234, 0, 0, 247, 0,
	0, 237, 0, 232, 0, 232, 0, 32, 33, 34,
	15, 40, 41, 44, 0, 43, 36, 0, 0, 91,
	251, 0, 20, 227, 0, 250, 0, 0, 0, 254,
	0, 254, 0, 0, 0, 0, 0, 31, 0, 45,
	0, 190, 0, 0, 38, 212, 0, 90, 0, 0,
	99, 0, 0, 254, 0, 248, 23, 0, 26, 0,
	28, 233, 0, 254, 59, 46, 48, 54, 0, 52,
	53, -2, 101, 0, 0, 141, 142, 143, 144, 0,
	0, 0, 0, 180, 0, 167, 109, 0, 252, 183,
	184, 185, 186, 187, 188, 189, 168, 169, 170, 171,
	173, 107, 108, 0, 230, 231, 191, 192, 0, 0,
	0, 89, 99, 92, 197, 0, 228, 229, 0, 0,
	21, 235, 0, 0, 254, 244, 238, 239, 240, 241,
	242, 243, 27, 29, 30, 99, 0, 0, 49, 55,
	0, 57, 58, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 128, 129, 130, 131, 132,
	133, 134, 104, 0, 0, 0, 0, 139, 0, 0,
	158, 159, 160, 0, 0, 0, 121, 0, 0, 174,
	14, 216, 0, 219, 216, 0, 214, 197, 205, 0,
	100, 0, 139, 0, 249, 24, 0, 245, 193, 47,
	60, 61, 63, 64, 74, 72, 0, 56, 50, 0,
	181, 102, 103, 106, 122, 0, 124, 126, 110, 111,
	112, 0, 136, 137, 0, 0, 0, 0, 114, 116,
	0, 120, 145, 146, 147, 148, 149, 150, 151, 152,
	153, 154, 155, 105, 253, 138, 0, 222, 156, 157,
	161, 162, 0, 0, 165, 0, 178, 175, 0, 16,
	0, 218, 220, 0, 17, 213, 0, 205, 19, 0,
	0, 0, 254, 25, 195, 0, 0, 0, 0, 0,
	79, 0, 0, 82, 0, 0, 0, 93, 75, 0,
	77, 78, 0, 73, 0, 0, 123, 125, 127, 0,
	113, 115, 117, 0, 0, 140, 163, 0, 166, 0,
	176, 0, 0, 0, 0, 215, 18, 206, 198, 199,
	202, 0, 22, 197, 0, 0, 62, 68, 0, 71,
	80, 81, 83, 0, 85, 0, 87, 88, 65, 0,
	0, 0, 76, 66, 67, 51, 182, 135, 0, 118,
	164, 172, 179, 0, 0, 221, 0, 0, 0, 201,
	203, 204, 205, 196, 194, 0, 0, 84, 86, 0,
	0, 0, 119, 177, 0, 207, 208, 200, 209, 69,
	70, 0, 0, 0, 217, 223, 225, 0, 13, 0,
	0, 0, 97, 0, 0, 0, 0, 210, 0, 94,
	0, 95, 96, 224, 0, 0, 98, 0, 211, 0,
	226,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 89, 82, 3,
	46, 48, 87, 85, 47, 86, 95, 88, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	71, 70, 72, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 91, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 81, 3, 92,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 73, 74, 75, 76, 77,
	78, 79, 80, 83, 84, 90, 93, 94, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116,
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
	case 13:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:206
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:210
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{Expr: yyDollar[4].valExpr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:214
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:220
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].onDupUpdateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:224
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].onDupUpdateExprs))
			for _, updateList := range yyDollar[6].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].onDupUpdateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:236
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:242
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:248
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:254
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableIdent}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:258
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:263
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: NewTableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:269
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[4].tableIdent}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:273
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:278
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: NewTableIdent(yyDollar[3].colIdent.Lowered()), NewName: NewTableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:284
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:290
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:298
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:303
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: NewTableIdent(yyDollar[4].colIdent.Lowered()), IfExists: exists}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:313
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[3].tableIdent}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:319
		{
			yyVAL.statement = &Other{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:323
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:327
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:332
		{
			setAllowComments(yylex, true)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:336
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:342
		{
			yyVAL.bytes2 = nil
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:346
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:352
		{
			yyVAL.str = UnionStr
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:356
		{
			yyVAL.str = UnionAllStr
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:360
		{
			yyVAL.str = UnionDistinctStr
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:365
		{
			yyVAL.str = ""
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:369
		{
			yyVAL.str = DistinctStr
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
			yyVAL.str = StraightJoinHint
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:384
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:388
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:394
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:398
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:402
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:406
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:412
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:416
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:421
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:425
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:429
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:436
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:441
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:445
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:451
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:455
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:465
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:469
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:473
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:486
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:490
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:494
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:498
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:503
		{
			yyVAL.empty = struct{}{}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:505
		{
			yyVAL.empty = struct{}{}
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:508
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:512
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:516
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:523
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:529
		{
			yyVAL.str = JoinStr
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:533
		{
			yyVAL.str = JoinStr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:537
		{
			yyVAL.str = JoinStr
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:541
		{
			yyVAL.str = StraightJoinStr
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:547
		{
			yyVAL.str = LeftJoinStr
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:551
		{
			yyVAL.str = LeftJoinStr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:555
		{
			yyVAL.str = RightJoinStr
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:559
		{
			yyVAL.str = RightJoinStr
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:565
		{
			yyVAL.str = NaturalJoinStr
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:569
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:579
		{
			yyVAL.tableName = yyDollar[2].tableName
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:583
		{
			yyVAL.tableName = yyDollar[1].tableName
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:589
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:593
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:598
		{
			yyVAL.indexHints = nil
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:602
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:606
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:610
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:616
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:620
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:625
		{
			yyVAL.boolExpr = nil
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:629
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:636
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:640
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:644
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:648
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:652
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:658
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:662
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:668
		{
			yyVAL.boolExpr = yyDollar[1].boolVal
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:672
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:676
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:680
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:684
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:688
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:692
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:696
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:700
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:704
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:708
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:712
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:716
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:722
		{
			yyVAL.str = IsNullStr
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:726
		{
			yyVAL.str = IsNotNullStr
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:730
		{
			yyVAL.str = IsTrueStr
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:734
		{
			yyVAL.str = IsNotTrueStr
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:738
		{
			yyVAL.str = IsFalseStr
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:742
		{
			yyVAL.str = IsNotFalseStr
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:748
		{
			yyVAL.str = EqualStr
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:752
		{
			yyVAL.str = LessThanStr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:756
		{
			yyVAL.str = GreaterThanStr
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:760
		{
			yyVAL.str = LessEqualStr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:764
		{
			yyVAL.str = GreaterEqualStr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:768
		{
			yyVAL.str = NotEqualStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:772
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:778
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:782
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:786
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:792
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:798
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:802
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:808
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:812
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:816
		{
			yyVAL.valExpr = yyDollar[1].valTuple
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:820
		{
			yyVAL.valExpr = yyDollar[1].subquery
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:824
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:828
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:832
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:836
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:840
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:844
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:848
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:852
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:856
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:860
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:864
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:868
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:872
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:876
		{
			if num, ok := yyDollar[2].valExpr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:884
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
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:898
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:902
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:910
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:914
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 164:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:918
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:922
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 166:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:926
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:930
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:936
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:940
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:944
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:948
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 172:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:954
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:959
		{
			yyVAL.valExpr = nil
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:963
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:969
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:973
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 177:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:979
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:984
		{
			yyVAL.valExpr = nil
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:988
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:994
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:998
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 182:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1002
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1008
		{
			yyVAL.valExpr = NewStrVal(yyDollar[1].bytes)
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1012
		{
			yyVAL.valExpr = NewHexVal(yyDollar[1].bytes)
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1016
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1020
		{
			yyVAL.valExpr = NewFloatVal(yyDollar[1].bytes)
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1024
		{
			yyVAL.valExpr = NewHexNum(yyDollar[1].bytes)
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1028
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1032
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1038
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.valExpr = NewIntVal([]byte("1"))
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1047
		{
			yyVAL.valExpr = NewIntVal(yyDollar[1].bytes)
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1051
		{
			yyVAL.valExpr = NewValArg(yyDollar[1].bytes)
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1056
		{
			yyVAL.valExprs = nil
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1060
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1065
		{
			yyVAL.boolExpr = nil
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1069
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1074
		{
			yyVAL.orderBy = nil
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1078
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1084
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 200:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1088
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1094
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1099
		{
			yyVAL.str = AscScr
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1103
		{
			yyVAL.str = AscScr
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1107
		{
			yyVAL.str = DescScr
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1112
		{
			yyVAL.limit = nil
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1116
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 207:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1120
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 208:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1124
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1129
		{
			yyVAL.str = ""
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1133
		{
			yyVAL.str = ForUpdateStr
		}
	case 211:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1137
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
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1150
		{
			yyVAL.columns = nil
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1154
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1160
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 215:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1164
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1169
		{
			yyVAL.onDupUpdateExprs = nil
		}
	case 217:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1173
		{
			yyVAL.onDupUpdateExprs = yyDollar[5].onDupUpdateExprs
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1179
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1183
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1189
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1193
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 222:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1199
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1205
		{
			yyVAL.onDupUpdateExprs = OnDupUpdateExprs{yyDollar[1].onDupUpdateExpr}
		}
	case 224:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1209
		{
			yyVAL.onDupUpdateExprs = append(yyDollar[1].onDupUpdateExprs, yyDollar[3].onDupUpdateExpr)
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1215
		{
			yyVAL.onDupUpdateExpr = &OnDupUpdateExpr{Name: yyDollar[1].updateExpr.Name, Expr: yyDollar[1].updateExpr.Expr, UseLookup: false}
		}
	case 226:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1219
		{
			yyVAL.onDupUpdateExpr = &OnDupUpdateExpr{Name: yyDollar[1].colIdent, LookupName: yyDollar[5].colIdent, UseLookup: true}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1225
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1229
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1235
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1244
		{
			yyVAL.byt = 0
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1246
		{
			yyVAL.byt = 1
		}
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1249
		{
			yyVAL.empty = struct{}{}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1251
		{
			yyVAL.empty = struct{}{}
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1254
		{
			yyVAL.str = ""
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1256
		{
			yyVAL.str = IgnoreStr
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1260
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1262
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1264
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1266
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1268
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1270
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1273
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1275
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1278
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1280
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1283
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1285
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1289
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1295
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1301
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1310
		{
			decNesting(yylex)
		}
	case 254:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1315
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
