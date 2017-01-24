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
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	exprs       Exprs
	boolVal     BoolVal
	colTuple    ColTuple
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
const DIV = 57419
const MOD = 57420
const UNARY = 57421
const COLLATE = 57422
const INTERVAL = 57423
const JSON_EXTRACT_OP = 57424
const JSON_UNQUOTE_EXTRACT_OP = 57425
const CREATE = 57426
const ALTER = 57427
const DROP = 57428
const RENAME = 57429
const ANALYZE = 57430
const TABLE = 57431
const INDEX = 57432
const VIEW = 57433
const TO = 57434
const IGNORE = 57435
const IF = 57436
const UNIQUE = 57437
const USING = 57438
const SHOW = 57439
const DESCRIBE = 57440
const EXPLAIN = 57441
const CURRENT_TIMESTAMP = 57442
const DATABASE = 57443
const CURRENT_DATE = 57444
const UNIX_TIMESTAMP = 57445
const CURRENT_TIME = 57446
const LOCALTIME = 57447
const LOCALTIMESTAMP = 57448
const UTC_DATE = 57449
const UTC_TIME = 57450
const UTC_TIMESTAMP = 57451
const REPLACE = 57452
const UNUSED = 57453

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
	"'!'",
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
	"DIV",
	"'%'",
	"MOD",
	"'^'",
	"'~'",
	"UNARY",
	"COLLATE",
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
	"CURRENT_DATE",
	"UNIX_TIMESTAMP",
	"CURRENT_TIME",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"UTC_DATE",
	"UTC_TIME",
	"UTC_TIMESTAMP",
	"REPLACE",
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
	-1, 67,
	98, 263,
	-2, 262,
	-1, 129,
	46, 180,
	-2, 169,
}

const yyNprod = 267
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1267

var yyAct = [...]int{

	127, 142, 453, 336, 119, 243, 62, 141, 387, 314,
	255, 356, 325, 346, 267, 115, 47, 268, 266, 285,
	327, 118, 170, 108, 109, 270, 35, 179, 37, 242,
	3, 41, 38, 116, 77, 65, 406, 408, 71, 64,
	198, 74, 48, 49, 69, 40, 185, 41, 263, 103,
	111, 43, 44, 45, 436, 84, 435, 50, 58, 434,
	183, 210, 220, 221, 213, 214, 215, 216, 217, 218,
	219, 212, 121, 70, 222, 73, 46, 42, 232, 233,
	107, 187, 366, 90, 258, 202, 96, 93, 100, 92,
	166, 102, 65, 222, 458, 65, 64, 174, 168, 64,
	197, 212, 292, 407, 222, 215, 216, 217, 218, 219,
	212, 63, 200, 222, 176, 167, 290, 291, 289, 231,
	72, 195, 95, 67, 190, 416, 98, 413, 326, 239,
	240, 72, 87, 241, 182, 184, 181, 197, 234, 235,
	236, 237, 238, 425, 426, 67, 173, 211, 210, 220,
	221, 213, 214, 215, 216, 217, 218, 219, 212, 412,
	186, 222, 281, 158, 60, 203, 65, 253, 94, 326,
	64, 379, 251, 288, 99, 196, 195, 363, 364, 365,
	260, 196, 195, 89, 78, 247, 14, 15, 16, 17,
	230, 254, 197, 244, 250, 272, 72, 261, 197, 85,
	200, 317, 86, 174, 196, 195, 172, 60, 246, 18,
	60, 418, 359, 286, 264, 322, 309, 265, 310, 193,
	284, 197, 257, 293, 294, 295, 296, 297, 298, 299,
	300, 301, 302, 303, 304, 305, 306, 307, 283, 313,
	280, 357, 320, 276, 278, 279, 273, 274, 277, 423,
	318, 331, 321, 323, 311, 312, 192, 330, 199, 339,
	422, 382, 333, 315, 319, 335, 271, 158, 72, 60,
	201, 359, 332, 14, 272, 344, 158, 28, 287, 60,
	76, 19, 20, 22, 21, 23, 72, 328, 201, 94,
	421, 362, 460, 317, 24, 25, 26, 286, 196, 195,
	192, 317, 344, 317, 196, 195, 368, 369, 370, 367,
	334, 317, 82, 158, 165, 197, 60, 66, 322, 317,
	317, 197, 171, 430, 328, 372, 79, 259, 106, 374,
	375, 158, 433, 401, 432, 383, 399, 378, 402, 384,
	373, 400, 398, 397, 39, 271, 272, 272, 272, 272,
	59, 403, 163, 352, 353, 162, 14, 389, 94, 393,
	75, 395, 287, 394, 80, 396, 451, 339, 404, 409,
	91, 419, 411, 376, 177, 59, 57, 380, 452, 59,
	55, 249, 244, 415, 331, 381, 97, 414, 385, 388,
	420, 101, 105, 54, 104, 361, 51, 52, 429, 112,
	337, 427, 348, 351, 352, 353, 349, 59, 350, 354,
	169, 392, 431, 338, 175, 161, 391, 271, 271, 271,
	271, 188, 256, 160, 189, 343, 171, 61, 457, 437,
	417, 443, 442, 439, 14, 28, 446, 447, 448, 30,
	65, 1, 428, 244, 64, 308, 449, 454, 454, 454,
	455, 456, 135, 134, 136, 137, 138, 139, 360, 463,
	140, 464, 355, 459, 465, 461, 462, 194, 178, 438,
	36, 262, 440, 441, 388, 180, 68, 59, 159, 252,
	164, 444, 445, 450, 220, 221, 213, 214, 215, 216,
	217, 218, 219, 212, 424, 316, 222, 117, 213, 214,
	215, 216, 217, 218, 219, 212, 131, 386, 222, 129,
	112, 59, 148, 128, 390, 342, 377, 245, 324, 130,
	282, 158, 371, 317, 67, 135, 134, 136, 137, 138,
	139, 120, 329, 140, 132, 133, 275, 83, 114, 125,
	248, 157, 211, 210, 220, 221, 213, 214, 215, 216,
	217, 218, 219, 212, 204, 113, 222, 112, 112, 405,
	88, 122, 123, 110, 347, 29, 345, 146, 269, 124,
	72, 191, 126, 81, 53, 27, 56, 340, 13, 12,
	341, 31, 32, 33, 34, 11, 143, 10, 358, 9,
	59, 8, 149, 144, 150, 145, 151, 155, 156, 154,
	153, 152, 147, 211, 210, 220, 221, 213, 214, 215,
	216, 217, 218, 219, 212, 117, 7, 222, 6, 348,
	351, 352, 353, 349, 131, 350, 354, 5, 4, 2,
	148, 0, 0, 0, 112, 0, 0, 0, 0, 158,
	0, 317, 67, 135, 134, 136, 137, 138, 139, 0,
	0, 140, 132, 133, 0, 0, 114, 125, 0, 157,
	0, 0, 59, 59, 59, 59, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 358, 0, 0, 410, 122,
	123, 110, 0, 0, 0, 146, 0, 124, 0, 0,
	126, 211, 210, 220, 221, 213, 214, 215, 216, 217,
	218, 219, 212, 0, 143, 222, 0, 0, 0, 0,
	149, 144, 150, 145, 151, 155, 156, 154, 153, 152,
	147, 117, 0, 0, 0, 0, 0, 0, 0, 0,
	131, 0, 0, 0, 0, 0, 148, 0, 0, 0,
	0, 0, 0, 0, 0, 158, 0, 0, 67, 135,
	134, 136, 137, 138, 139, 0, 0, 140, 132, 133,
	0, 0, 114, 125, 0, 157, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 122, 123, 110, 0, 0,
	0, 146, 0, 124, 0, 0, 126, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	143, 14, 0, 0, 0, 0, 149, 144, 150, 145,
	151, 155, 156, 154, 153, 152, 147, 117, 0, 0,
	0, 0, 0, 0, 0, 0, 131, 0, 0, 0,
	0, 0, 148, 0, 0, 0, 0, 0, 0, 0,
	0, 158, 0, 0, 67, 135, 134, 136, 137, 138,
	139, 0, 0, 140, 132, 133, 0, 0, 114, 125,
	0, 157, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 122, 123, 0, 0, 0, 0, 146, 0, 124,
	0, 0, 126, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 143, 0, 0, 0,
	0, 0, 149, 144, 150, 145, 151, 155, 156, 154,
	153, 152, 147, 117, 0, 0, 0, 0, 0, 0,
	0, 0, 131, 0, 0, 0, 0, 0, 148, 0,
	0, 0, 0, 0, 0, 0, 0, 158, 0, 0,
	67, 135, 134, 136, 137, 138, 139, 0, 0, 140,
	132, 133, 0, 0, 114, 125, 0, 157, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 122, 123, 0,
	0, 0, 0, 146, 0, 124, 0, 0, 126, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 143, 0, 0, 0, 0, 0, 149, 144,
	150, 145, 151, 155, 156, 154, 153, 152, 147, 131,
	0, 0, 0, 0, 0, 148, 0, 0, 0, 0,
	0, 0, 0, 0, 158, 0, 0, 67, 135, 134,
	136, 137, 138, 139, 0, 0, 140, 132, 133, 0,
	0, 0, 125, 0, 157, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 122, 123, 0, 0, 0, 0,
	146, 0, 124, 0, 0, 126, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 143,
	0, 0, 0, 0, 0, 149, 144, 150, 145, 151,
	155, 156, 154, 153, 152, 147, 131, 0, 0, 0,
	0, 0, 148, 0, 0, 0, 0, 0, 0, 0,
	0, 158, 0, 0, 67, 135, 134, 136, 137, 138,
	139, 0, 0, 140, 0, 0, 0, 0, 0, 125,
	0, 157, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 122, 123, 0, 0, 0, 0, 146, 0, 124,
	0, 0, 126, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 143, 0, 0, 0,
	0, 0, 149, 144, 150, 145, 151, 155, 156, 154,
	153, 152, 147, 206, 0, 209, 0, 0, 0, 0,
	0, 223, 224, 225, 226, 227, 228, 229, 0, 207,
	208, 205, 211, 210, 220, 221, 213, 214, 215, 216,
	217, 218, 219, 212, 0, 0, 222,
}
var yyPact = [...]int{

	180, -1000, -1000, 430, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -80,
	-63, -29, -55, -30, -1000, -1000, -1000, 428, 377, 360,
	-1000, -79, 115, 417, 96, -67, -34, 82, -1000, -31,
	82, -1000, 115, -77, 135, -77, 115, -1000, -1000, -1000,
	-1000, -1000, -1000, 276, 147, -1000, 75, 158, 341, -9,
	-1000, 115, 121, -1000, 51, -1000, -12, -1000, 115, 63,
	125, -1000, -1000, 115, -1000, -60, 115, 370, 283, 82,
	-1000, 699, -1000, 405, -1000, 324, 321, -1000, 285, 115,
	-1000, 96, 115, 415, 96, 911, 96, -1000, 352, -86,
	-1000, 32, -1000, 115, -1000, -1000, 115, -1000, 209, -1000,
	-1000, 237, -13, -1000, 911, 1170, -1000, 221, -1000, -21,
	-1000, -1000, 1105, 1105, 1105, 1105, 1105, 221, 221, -1000,
	-1000, 221, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 805, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 911, -1000, 115,
	-1000, -1000, -1000, -1000, 350, 96, 82, -1000, 311, -1000,
	408, 911, -1000, 114, -1000, -14, -1000, -1000, 282, 82,
	-1000, -61, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 415, 699, 230, -1000, 911, 911, 185, -1000, 219,
	-1000, -1000, 74, 22, 1008, 117, 37, 1105, 1105, 1105,
	1105, 1105, 1105, 1105, 1105, 1105, 1105, 1105, 1105, 1105,
	1105, 1105, 167, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 428, 402, 402, -3, -3, -3, -3, 521, 475,
	593, 82, 272, 271, 114, 61, 114, -1000, 279, 221,
	430, 242, 263, -1000, 408, 384, 398, 114, 82, 115,
	-1000, -1000, 115, -1000, 413, -1000, 228, 584, -1000, -1000,
	220, 374, 267, 22, 59, -1000, -1000, 119, -1000, -1000,
	-1000, -1000, -16, -1000, 609, -1000, -1000, -1000, -1000, 117,
	1105, 1105, 1105, 609, 609, 460, 400, -22, -3, 17,
	17, 8, 8, 8, 8, 8, 412, 412, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 253, 699, -1000, -1000, 253,
	153, -1000, 911, -1000, 102, -1000, 911, -1000, 359, 214,
	-1000, 911, -1000, -1000, 82, 384, -1000, 911, 911, -1000,
	-1000, -1000, 403, 396, 230, 230, 230, 230, -1000, 308,
	307, -1000, 301, 298, 316, -7, -1000, 161, -1000, -1000,
	115, -1000, 255, -1000, -1000, -1000, 71, -1000, 609, 609,
	65, 1105, -1000, 253, -1000, -1000, 114, 55, -1000, 911,
	143, 344, 221, -1000, -1000, 243, 202, -1000, 120, -1000,
	408, 911, 911, 584, 278, 367, -1000, -1000, -1000, -1000,
	299, -1000, 297, -1000, -1000, -1000, -48, -51, -53, -1000,
	-1000, -1000, -1000, 1105, 609, -1000, -1000, 114, 911, 425,
	-1000, 911, 911, 911, -1000, -1000, -1000, 384, 114, 168,
	911, 911, -1000, -1000, 221, 221, 221, 609, 114, 96,
	114, 114, -1000, 348, 114, 114, 82, 82, 82, 121,
	-1000, 420, 13, 245, -1000, 245, 245, -1000, 82, -1000,
	82, -1000, -1000, 82, -1000, -1000,
}
var yyPgo = [...]int{

	0, 629, 29, 628, 627, 618, 616, 591, 589, 587,
	585, 579, 578, 565, 576, 575, 574, 573, 23, 24,
	50, 571, 18, 14, 17, 568, 566, 13, 564, 25,
	560, 559, 2, 22, 555, 33, 554, 540, 21, 15,
	537, 536, 19, 5, 532, 7, 531, 72, 4, 519,
	518, 12, 517, 516, 515, 514, 513, 509, 10, 507,
	8, 494, 3, 483, 480, 479, 20, 6, 111, 478,
	344, 280, 476, 475, 471, 470, 468, 0, 40, 467,
	317, 11, 462, 458, 16, 445, 441, 439, 1, 9,
}
var yyR1 = [...]int{

	0, 86, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 3, 3, 4, 5,
	6, 7, 7, 7, 8, 8, 8, 9, 10, 10,
	10, 11, 12, 12, 12, 87, 13, 14, 14, 15,
	15, 15, 16, 16, 17, 17, 18, 18, 19, 19,
	19, 19, 79, 79, 79, 78, 78, 21, 21, 22,
	22, 23, 23, 24, 24, 24, 25, 25, 25, 25,
	83, 83, 82, 82, 82, 81, 81, 26, 26, 26,
	26, 27, 27, 27, 27, 28, 28, 30, 30, 29,
	29, 31, 31, 31, 31, 32, 32, 33, 33, 20,
	20, 20, 20, 20, 20, 35, 35, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 41,
	41, 41, 41, 41, 41, 36, 36, 36, 36, 36,
	36, 36, 42, 42, 42, 47, 43, 43, 85, 85,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 57, 57, 57, 57, 57, 57, 57, 57,
	56, 56, 56, 56, 56, 56, 56, 49, 52, 52,
	50, 50, 51, 53, 53, 48, 48, 48, 38, 38,
	38, 38, 38, 38, 38, 40, 40, 40, 54, 54,
	55, 55, 58, 58, 59, 59, 60, 61, 61, 61,
	62, 62, 62, 62, 63, 63, 63, 64, 64, 65,
	65, 66, 66, 37, 37, 44, 44, 45, 46, 67,
	67, 68, 69, 69, 71, 71, 72, 72, 70, 70,
	73, 73, 73, 73, 73, 73, 74, 74, 75, 75,
	76, 76, 77, 80, 88, 89, 84,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 12, 6, 3, 7, 7, 8, 7,
	3, 5, 8, 4, 6, 7, 4, 5, 4, 5,
	5, 3, 2, 2, 2, 0, 2, 0, 2, 1,
	2, 2, 0, 1, 0, 1, 1, 3, 1, 2,
	3, 5, 0, 1, 2, 1, 1, 0, 2, 1,
	3, 1, 1, 3, 3, 3, 3, 5, 5, 3,
	0, 1, 0, 1, 2, 1, 1, 1, 2, 2,
	1, 2, 3, 2, 3, 2, 2, 2, 1, 1,
	3, 0, 5, 5, 5, 1, 3, 0, 2, 1,
	3, 3, 2, 3, 1, 1, 1, 1, 3, 3,
	3, 4, 3, 4, 3, 4, 5, 6, 2, 1,
	2, 1, 2, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 3, 1, 1,
	1, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 3, 3, 4, 5, 3, 4, 1,
	1, 4, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 5, 0, 1,
	1, 2, 4, 0, 2, 1, 3, 5, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 4, 0, 2, 4, 0, 3, 1,
	3, 0, 5, 2, 1, 1, 3, 3, 1, 1,
	3, 3, 1, 1, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -86, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, 6, 7, 8, 9, 29, 101,
	102, 104, 103, 105, 114, 115, 116, -15, 5, -13,
	-87, -13, -13, -13, -13, 106, -75, 108, 112, -70,
	108, 110, 106, 106, 107, 108, 106, -84, -84, -84,
	-2, 19, 20, -16, 33, 20, -14, -70, -29, -80,
	49, 10, -67, -68, -48, -77, -80, 49, -72, 111,
	107, -77, 49, 106, -77, -80, -71, 111, 49, -71,
	-80, -17, 36, -40, -77, 52, 55, 57, -30, 25,
	-29, 29, 98, -29, 47, 71, 98, -80, 63, 49,
	-84, -80, -84, 109, -80, 22, 45, -77, -18, -19,
	88, -20, -80, -34, 63, -39, -35, 22, -38, -48,
	-46, -47, 86, 87, 94, 64, 97, -77, -56, -57,
	-49, 31, 59, 60, 51, 50, 52, 53, 54, 55,
	58, -45, -88, 111, 118, 120, 92, 127, 37, 117,
	119, 121, 126, 125, 124, 122, 123, 66, 46, -69,
	18, 10, 31, 31, -64, 29, -88, -29, -67, -80,
	-33, 11, -68, -20, -77, -80, -84, 22, -76, 113,
	-73, 104, 102, 28, 103, 14, 128, 49, -80, -80,
	-84, -21, 47, 10, -79, 62, 61, 78, -78, 21,
	-77, 51, 98, -20, -36, 81, 63, 79, 80, 65,
	83, 82, 93, 86, 87, 88, 89, 90, 91, 92,
	84, 85, 96, 71, 72, 73, 74, 75, 76, 77,
	-47, -88, 99, 100, -39, -39, -39, -39, -39, -88,
	-88, -88, -2, -43, -20, -52, -20, -29, -37, 31,
	-2, -67, -65, -77, -33, -58, 14, -20, 98, 45,
	-77, -84, -74, 109, -33, -19, -22, -23, -24, -25,
	-29, -47, -88, -20, -20, -41, 58, 63, 59, 60,
	-78, 88, -80, -35, -39, -42, -45, -47, 56, 81,
	79, 80, 65, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -85, 49,
	51, -38, -38, -77, -89, -18, 20, 48, -89, -18,
	-77, -89, 47, -89, -50, -51, 67, -66, 45, -44,
	-45, -88, -66, -89, 47, -58, -62, 16, 15, -77,
	-80, -80, -54, 12, 47, -26, -27, -28, 35, 39,
	41, 36, 37, 38, 42, -82, -81, 21, -80, 51,
	-83, 21, -22, 58, 59, 60, 98, -42, -39, -39,
	-39, 62, -89, -18, -89, -89, -20, -53, -51, 69,
	-20, 26, 47, -77, -62, -20, -59, -60, -20, -84,
	-55, 13, 15, -23, -24, -23, -24, 35, 35, 35,
	40, 35, 40, 35, -27, -31, 43, 110, 44, -81,
	-80, -89, 88, 62, -39, -89, 70, -20, 68, 27,
	-45, 47, 17, 47, -61, 23, 24, -58, -20, -43,
	45, 45, 35, 35, 107, 107, 107, -39, -20, 8,
	-20, -20, -60, -62, -20, -20, -88, -88, -88, -67,
	-63, 18, 30, -32, -77, -32, -32, 8, 81, -89,
	47, -89, -89, -77, -77, -77,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 35, 35, 35, 35, 35, 258,
	248, 0, 0, 0, 266, 266, 266, 0, 39, 42,
	37, 248, 0, 0, 0, 246, 0, 0, 259, 0,
	0, 249, 0, 244, 0, 244, 0, 32, 33, 34,
	15, 40, 41, 44, 0, 43, 36, 0, 0, 89,
	263, 0, 20, 239, 0, 195, 0, -2, 0, 0,
	0, 266, 262, 0, 266, 0, 0, 0, 0, 0,
	31, 0, 45, 0, 205, 0, 0, 38, 227, 0,
	88, 0, 0, 97, 0, 0, 0, 266, 0, 260,
	23, 0, 26, 0, 28, 245, 0, 266, 57, 46,
	48, 52, 0, 99, 0, 104, 107, 0, 140, 141,
	142, 143, 0, 0, 0, 0, 0, 195, 0, -2,
	170, 0, 105, 106, 198, 199, 200, 201, 202, 203,
	204, 238, 0, 181, 182, 183, 184, 185, 186, 172,
	173, 174, 175, 176, 177, 178, 179, 188, 264, 0,
	242, 243, 206, 207, 0, 0, 0, 87, 97, 90,
	212, 0, 240, 241, 196, 0, 21, 247, 0, 0,
	266, 256, 250, 251, 252, 253, 254, 255, 27, 29,
	30, 97, 0, 0, 49, 0, 0, 0, 53, 0,
	55, 56, 0, 102, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 125, 126, 127, 128, 129, 130, 131,
	118, 0, 0, 0, 159, 160, 161, 162, 0, 0,
	0, 0, 0, 0, 136, 0, 189, 14, 231, 0,
	234, 231, 0, 229, 212, 220, 0, 98, 0, 0,
	261, 24, 0, 257, 208, 47, 58, 59, 61, 62,
	72, 70, 0, 100, 101, 103, 119, 0, 121, 123,
	54, 50, 0, 108, 109, 110, 132, 133, 134, 0,
	0, 0, 0, 112, 114, 0, 144, 145, 146, 147,
	148, 149, 150, 151, 152, 153, 154, 155, 158, 138,
	139, 156, 157, 163, 164, 0, 0, 265, 167, 0,
	0, 135, 0, 237, 193, 190, 0, 16, 0, 233,
	235, 0, 17, 228, 0, 220, 19, 0, 0, 197,
	266, 25, 210, 0, 0, 0, 0, 0, 77, 0,
	0, 80, 0, 0, 0, 91, 73, 0, 75, 76,
	0, 71, 0, 120, 122, 124, 0, 111, 113, 115,
	0, 0, 165, 0, 168, 171, 137, 0, 191, 0,
	0, 0, 0, 230, 18, 221, 213, 214, 217, 22,
	212, 0, 0, 60, 66, 0, 69, 78, 79, 81,
	0, 83, 0, 85, 86, 63, 0, 0, 0, 74,
	64, 65, 51, 0, 116, 166, 187, 194, 0, 0,
	236, 0, 0, 0, 216, 218, 219, 220, 211, 209,
	0, 0, 82, 84, 0, 0, 0, 117, 192, 0,
	222, 223, 215, 224, 67, 68, 0, 0, 0, 232,
	13, 0, 0, 0, 95, 0, 0, 225, 0, 92,
	0, 93, 94, 0, 96, 226,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 64, 3, 3, 3, 91, 83, 3,
	46, 48, 88, 86, 47, 87, 98, 89, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	72, 71, 73, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 93, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 82, 3, 94,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 65,
	66, 67, 68, 69, 70, 74, 75, 76, 77, 78,
	79, 80, 81, 84, 85, 90, 92, 95, 96, 97,
	99, 100, 101, 102, 103, 104, 105, 106, 107, 108,
	109, 110, 111, 112, 113, 114, 115, 116, 117, 118,
	119, 120, 121, 122, 123, 124, 125, 126, 127, 128,
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
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].expr), GroupBy: GroupBy(yyDollar[8].exprs), Having: NewWhere(HavingStr, yyDollar[9].expr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:211
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{Expr: yyDollar[4].expr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:215
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:221
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:225
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[6].updateExprs {
				cols = append(cols, updateList.Name.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:237
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].expr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:243
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].expr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:249
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:255
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableIdent}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:259
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:264
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: NewTableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:270
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[4].tableIdent}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:274
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:279
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: NewTableIdent(yyDollar[3].colIdent.Lowered()), NewName: NewTableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:285
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:291
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:299
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:304
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: NewTableIdent(yyDollar[4].colIdent.Lowered()), IfExists: exists}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:314
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[3].tableIdent}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:320
		{
			yyVAL.statement = &Other{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:324
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:328
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:333
		{
			setAllowComments(yylex, true)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:337
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:343
		{
			yyVAL.bytes2 = nil
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:347
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:353
		{
			yyVAL.str = UnionStr
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:357
		{
			yyVAL.str = UnionAllStr
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:361
		{
			yyVAL.str = UnionDistinctStr
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:366
		{
			yyVAL.str = ""
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:370
		{
			yyVAL.str = DistinctStr
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:375
		{
			yyVAL.str = ""
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:379
		{
			yyVAL.str = StraightJoinHint
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:385
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:389
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:395
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:399
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:403
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Name: yyDollar[1].tableIdent}}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:407
		{
			yyVAL.selectExpr = &StarExpr{TableName: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}}
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:412
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:416
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:420
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:427
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:432
		{
			yyVAL.tableExprs = TableExprs{&AliasedTableExpr{Expr: &TableName{Name: NewTableIdent("dual")}}}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:436
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:442
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:446
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:456
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:460
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:464
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:477
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 67:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:481
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].expr}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:485
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:489
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:494
		{
			yyVAL.empty = struct{}{}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:496
		{
			yyVAL.empty = struct{}{}
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:499
		{
			yyVAL.tableIdent = NewTableIdent("")
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:503
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:507
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:514
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:520
		{
			yyVAL.str = JoinStr
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:524
		{
			yyVAL.str = JoinStr
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:528
		{
			yyVAL.str = JoinStr
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:532
		{
			yyVAL.str = StraightJoinStr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:538
		{
			yyVAL.str = LeftJoinStr
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:542
		{
			yyVAL.str = LeftJoinStr
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:546
		{
			yyVAL.str = RightJoinStr
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:550
		{
			yyVAL.str = RightJoinStr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:556
		{
			yyVAL.str = NaturalJoinStr
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:560
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:570
		{
			yyVAL.tableName = yyDollar[2].tableName
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:574
		{
			yyVAL.tableName = yyDollar[1].tableName
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:580
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:584
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:589
		{
			yyVAL.indexHints = nil
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:593
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:597
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:601
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:607
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:611
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:616
		{
			yyVAL.expr = nil
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:620
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:626
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:630
		{
			yyVAL.expr = &AndExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:634
		{
			yyVAL.expr = &OrExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:638
		{
			yyVAL.expr = &NotExpr{Expr: yyDollar[2].expr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:642
		{
			yyVAL.expr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].expr}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:646
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:652
		{
			yyVAL.boolVal = BoolVal(true)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:656
		{
			yyVAL.boolVal = BoolVal(false)
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:662
		{
			yyVAL.expr = yyDollar[1].boolVal
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:666
		{
			yyVAL.expr = &ComparisonExpr{Left: yyDollar[1].expr, Operator: yyDollar[2].str, Right: yyDollar[3].boolVal}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:670
		{
			yyVAL.expr = &ComparisonExpr{Left: yyDollar[1].expr, Operator: yyDollar[2].str, Right: yyDollar[3].expr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:674
		{
			yyVAL.expr = &ComparisonExpr{Left: yyDollar[1].expr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:678
		{
			yyVAL.expr = &ComparisonExpr{Left: yyDollar[1].expr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:682
		{
			yyVAL.expr = &ComparisonExpr{Left: yyDollar[1].expr, Operator: LikeStr, Right: yyDollar[3].expr}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:686
		{
			yyVAL.expr = &ComparisonExpr{Left: yyDollar[1].expr, Operator: NotLikeStr, Right: yyDollar[4].expr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:690
		{
			yyVAL.expr = &ComparisonExpr{Left: yyDollar[1].expr, Operator: RegexpStr, Right: yyDollar[3].expr}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:694
		{
			yyVAL.expr = &ComparisonExpr{Left: yyDollar[1].expr, Operator: NotRegexpStr, Right: yyDollar[4].expr}
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:698
		{
			yyVAL.expr = &RangeCond{Left: yyDollar[1].expr, Operator: BetweenStr, From: yyDollar[3].expr, To: yyDollar[5].expr}
		}
	case 117:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:702
		{
			yyVAL.expr = &RangeCond{Left: yyDollar[1].expr, Operator: NotBetweenStr, From: yyDollar[4].expr, To: yyDollar[6].expr}
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:706
		{
			yyVAL.expr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:712
		{
			yyVAL.str = IsNullStr
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:716
		{
			yyVAL.str = IsNotNullStr
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:720
		{
			yyVAL.str = IsTrueStr
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:724
		{
			yyVAL.str = IsNotTrueStr
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:728
		{
			yyVAL.str = IsFalseStr
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:732
		{
			yyVAL.str = IsNotFalseStr
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:738
		{
			yyVAL.str = EqualStr
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:742
		{
			yyVAL.str = LessThanStr
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:746
		{
			yyVAL.str = GreaterThanStr
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:750
		{
			yyVAL.str = LessEqualStr
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:754
		{
			yyVAL.str = GreaterEqualStr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:758
		{
			yyVAL.str = NotEqualStr
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:762
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:768
		{
			yyVAL.colTuple = yyDollar[1].valTuple
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:772
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:776
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:782
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:788
		{
			yyVAL.exprs = Exprs{yyDollar[1].expr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:792
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:798
		{
			yyVAL.str = string(yyDollar[1].bytes)
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:802
		{
			yyVAL.str = string(yyDollar[1].bytes)
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:808
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:812
		{
			yyVAL.expr = yyDollar[1].colName
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:816
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:820
		{
			yyVAL.expr = yyDollar[1].subquery
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:824
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: BitAndStr, Right: yyDollar[3].expr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:828
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: BitOrStr, Right: yyDollar[3].expr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:832
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: BitXorStr, Right: yyDollar[3].expr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:836
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: PlusStr, Right: yyDollar[3].expr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:840
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: MinusStr, Right: yyDollar[3].expr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:844
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: MultStr, Right: yyDollar[3].expr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:848
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: DivStr, Right: yyDollar[3].expr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:852
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: IntDivStr, Right: yyDollar[3].expr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:856
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: ModStr, Right: yyDollar[3].expr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:860
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: ModStr, Right: yyDollar[3].expr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:864
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: ShiftLeftStr, Right: yyDollar[3].expr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:868
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].expr, Operator: ShiftRightStr, Right: yyDollar[3].expr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:872
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONExtractOp, Right: yyDollar[3].expr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:876
		{
			yyVAL.expr = &BinaryExpr{Left: yyDollar[1].colName, Operator: JSONUnquoteExtractOp, Right: yyDollar[3].expr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:880
		{
			yyVAL.expr = &CollateExpr{Expr: yyDollar[1].expr, Charset: yyDollar[3].str}
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:884
		{
			if num, ok := yyDollar[2].expr.(*SQLVal); ok && num.Type == IntVal {
				yyVAL.expr = num
			} else {
				yyVAL.expr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].expr}
			}
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:892
		{
			if num, ok := yyDollar[2].expr.(*SQLVal); ok && num.Type == IntVal {
				// Handle double negative
				if num.Val[0] == '-' {
					num.Val = num.Val[1:]
					yyVAL.expr = num
				} else {
					yyVAL.expr = NewIntVal(append([]byte("-"), num.Val...))
				}
			} else {
				yyVAL.expr = &UnaryExpr{Operator: UMinusStr, Expr: yyDollar[2].expr}
			}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:906
		{
			yyVAL.expr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].expr}
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:910
		{
			yyVAL.expr = &UnaryExpr{Operator: BangStr, Expr: yyDollar[2].expr}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:914
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.expr = &IntervalExpr{Expr: yyDollar[2].expr, Unit: yyDollar[3].colIdent}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:922
		{
			yyVAL.expr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 165:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:926
		{
			yyVAL.expr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 166:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:930
		{
			yyVAL.expr = &FuncExpr{Name: yyDollar[1].colIdent, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:934
		{
			yyVAL.expr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 168:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:938
		{
			yyVAL.expr = &FuncExpr{Name: yyDollar[1].colIdent, Exprs: yyDollar[3].selectExprs}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:942
		{
			yyVAL.expr = &FuncExpr{Name: yyDollar[1].colIdent}
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:946
		{
			yyVAL.expr = yyDollar[1].caseExpr
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:950
		{
			yyVAL.expr = &ValuesFuncExpr{Name: yyDollar[3].colIdent}
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:958
		{
			yyVAL.colIdent = NewColIdent("current_timestamp")
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:962
		{
			yyVAL.colIdent = NewColIdent("current_date")
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:966
		{
			yyVAL.colIdent = NewColIdent("current_time")
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:970
		{
			yyVAL.colIdent = NewColIdent("utc_timestamp")
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:974
		{
			yyVAL.colIdent = NewColIdent("utc_time")
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:978
		{
			yyVAL.colIdent = NewColIdent("utc_date")
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:982
		{
			yyVAL.colIdent = NewColIdent("localtime")
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:986
		{
			yyVAL.colIdent = NewColIdent("localtimestamp")
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:995
		{
			yyVAL.colIdent = NewColIdent("if")
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:999
		{
			yyVAL.colIdent = NewColIdent("database")
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1003
		{
			yyVAL.colIdent = NewColIdent("unix_timestamp")
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1007
		{
			yyVAL.colIdent = NewColIdent("mod")
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1011
		{
			yyVAL.colIdent = NewColIdent("replace")
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1015
		{
			yyVAL.colIdent = NewColIdent("left")
		}
	case 187:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1021
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].expr, Whens: yyDollar[3].whens, Else: yyDollar[4].expr}
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1026
		{
			yyVAL.expr = nil
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1030
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1036
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1040
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1046
		{
			yyVAL.when = &When{Cond: yyDollar[2].expr, Val: yyDollar[4].expr}
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1051
		{
			yyVAL.expr = nil
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1055
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1061
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1065
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 197:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1069
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1075
		{
			yyVAL.expr = NewStrVal(yyDollar[1].bytes)
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1079
		{
			yyVAL.expr = NewHexVal(yyDollar[1].bytes)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1083
		{
			yyVAL.expr = NewIntVal(yyDollar[1].bytes)
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1087
		{
			yyVAL.expr = NewFloatVal(yyDollar[1].bytes)
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1091
		{
			yyVAL.expr = NewHexNum(yyDollar[1].bytes)
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1095
		{
			yyVAL.expr = NewValArg(yyDollar[1].bytes)
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1099
		{
			yyVAL.expr = &NullVal{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1105
		{
			// TODO(sougou): Deprecate this construct.
			if yyDollar[1].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.expr = NewIntVal([]byte("1"))
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1114
		{
			yyVAL.expr = NewIntVal(yyDollar[1].bytes)
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1118
		{
			yyVAL.expr = NewValArg(yyDollar[1].bytes)
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1123
		{
			yyVAL.exprs = nil
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1127
		{
			yyVAL.exprs = yyDollar[3].exprs
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1132
		{
			yyVAL.expr = nil
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1136
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1141
		{
			yyVAL.orderBy = nil
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1145
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1151
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 215:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1155
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1161
		{
			yyVAL.order = &Order{Expr: yyDollar[1].expr, Direction: yyDollar[2].str}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1166
		{
			yyVAL.str = AscScr
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1170
		{
			yyVAL.str = AscScr
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1174
		{
			yyVAL.str = DescScr
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1179
		{
			yyVAL.limit = nil
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1183
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].expr}
		}
	case 222:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1187
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].expr, Rowcount: yyDollar[4].expr}
		}
	case 223:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1191
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].expr, Rowcount: yyDollar[2].expr}
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1196
		{
			yyVAL.str = ""
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1200
		{
			yyVAL.str = ForUpdateStr
		}
	case 226:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1204
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
	case 227:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1217
		{
			yyVAL.columns = nil
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1221
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1227
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1231
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 231:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1236
		{
			yyVAL.updateExprs = nil
		}
	case 232:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1240
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1246
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1250
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1256
		{
			yyVAL.values = Values{yyDollar[1].valTuple}
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1260
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].valTuple)
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1266
		{
			yyVAL.valTuple = ValTuple(yyDollar[2].exprs)
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1272
		{
			if len(yyDollar[1].valTuple) == 1 {
				yyVAL.expr = &ParenExpr{yyDollar[1].valTuple[0]}
			} else {
				yyVAL.expr = yyDollar[1].valTuple
			}
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1282
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1286
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1292
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].expr}
		}
	case 244:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1301
		{
			yyVAL.byt = 0
		}
	case 245:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1303
		{
			yyVAL.byt = 1
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1306
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1308
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1311
		{
			yyVAL.str = ""
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1313
		{
			yyVAL.str = IgnoreStr
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1317
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1319
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1321
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1323
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1325
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1327
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1330
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1332
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1335
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1337
		{
			yyVAL.empty = struct{}{}
		}
	case 260:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1340
		{
			yyVAL.empty = struct{}{}
		}
	case 261:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1342
		{
			yyVAL.empty = struct{}{}
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1346
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1352
		{
			yyVAL.tableIdent = NewTableIdent(string(yyDollar[1].bytes))
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1358
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1367
		{
			decNesting(yylex)
		}
	case 266:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./go/vt/sqlparser/sql.y:1372
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
