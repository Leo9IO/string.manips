package entities

var HtmlEntities = map[string]rune{
	"AElig":    198,
	"Aacute":   193,
	"Acirc":    194,
	"Agrave":   192,
	"Alpha":    913,
	"Aring":    197,
	"Atilde":   195,
	"Auml":     196,
	"Beta":     914,
	"Ccedil":   199,
	"Chi":      935,
	"Dagger":   8225,
	"Delta":    916,
	"ETH":      208,
	"Eacute":   201,
	"Ecirc":    202,
	"Egrave":   200,
	"Epsilon":  917,
	"Eta":      919,
	"Euml":     203,
	"Gamma":    915,
	"Iacute":   205,
	"Icirc":    206,
	"Igrave":   204,
	"Iota":     921,
	"Iuml":     207,
	"Kappa":    922,
	"Lambda":   923,
	"Mu":       924,
	"Ntilde":   209,
	"Nu":       925,
	"OElig":    338,
	"Oacute":   211,
	"Ocirc":    212,
	"Ograve":   210,
	"Omega":    937,
	"Omicron":  927,
	"Oslash":   216,
	"Otilde":   213,
	"Ouml":     214,
	"Phi":      934,
	"Pi":       928,
	"Prime":    8243,
	"Psi":      936,
	"Rho":      929,
	"Scaron":   352,
	"Sigma":    931,
	"THORN":    222,
	"Tau":      932,
	"Theta":    920,
	"Uacute":   218,
	"Ucirc":    219,
	"Ugrave":   217,
	"Upsilon":  933,
	"Uuml":     220,
	"Xi":       926,
	"Yacute":   221,
	"Yuml":     376,
	"Zeta":     918,
	"aacute":   225,
	"acirc":    226,
	"acute":    180,
	"aelig":    230,
	"agrave":   224,
	"alefsym":  8501,
	"alpha":    945,
	"and":      8743,
	"ang":      8736,
	"apos":     39,
	"aring":    229,
	"asymp":    8776,
	"atilde":   227,
	"auml":     228,
	"bdquo":    8222,
	"beta":     946,
	"brvbar":   166,
	"bull":     8226,
	"cap":      8745,
	"ccedil":   231,
	"cedil":    184,
	"cent":     162,
	"chi":      967,
	"circ":     710,
	"clubs":    9827,
	"cong":     8773,
	"copy":     169,
	"crarr":    8629,
	"cup":      8746,
	"curren":   164,
	"dArr":     8659,
	"dagger":   8224,
	"darr":     8595,
	"deg":      176,
	"delta":    948,
	"diams":    9830,
	"divide":   247,
	"eacute":   233,
	"ecirc":    234,
	"egrave":   232,
	"empty":    8709,
	"emsp":     8195,
	"ensp":     8194,
	"epsilon":  949,
	"equiv":    8801,
	"eta":      951,
	"eth":      240,
	"euml":     235,
	"euro":     8364,
	"exist":    8707,
	"fnof":     402,
	"forall":   8704,
	"frac12":   189,
	"frac14":   188,
	"frac34":   190,
	"frasl":    8260,
	"gamma":    947,
	"ge":       8805,
	"gt":       62,
	"hArr":     8660,
	"harr":     8596,
	"hearts":   9829,
	"hellip":   8230,
	"iacute":   237,
	"icirc":    238,
	"iexcl":    161,
	"igrave":   236,
	"image":    8465,
	"infin":    8734,
	"int":      8747,
	"iota":     953,
	"iquest":   191,
	"isin":     8712,
	"iuml":     239,
	"kappa":    954,
	"lArr":     8656,
	"lambda":   955,
	"lang":     9001,
	"laquo":    171,
	"larr":     8592,
	"lceil":    8968,
	"ldquo":    8220,
	"le":       8804,
	"lfloor":   8970,
	"lowast":   8727,
	"loz":      9674,
	"lrm":      8206,
	"lsaquo":   8249,
	"lsquo":    8216,
	"macr":     175,
	"mdash":    8212,
	"micro":    181,
	"middot":   183,
	"minus":    8722,
	"mu":       956,
	"nabla":    8711,
	"nbsp":     160,
	"ndash":    8211,
	"ne":       8800,
	"ni":       8715,
	"not":      172,
	"notin":    8713,
	"nsub":     8836,
	"ntilde":   241,
	"nu":       957,
	"oacute":   243,
	"ocirc":    244,
	"oelig":    339,
	"ograve":   242,
	"oline":    8254,
	"omega":    969,
	"omicron":  959,
	"oplus":    8853,
	"or":       8744,
	"ordf":     170,
	"ordm":     186,
	"oslash":   248,
	"otilde":   245,
	"otimes":   8855,
	"ouml":     246,
	"para":     182,
	"part":     8706,
	"permil":   8240,
	"perp":     8869,
	"phi":      966,
	"pi":       960,
	"piv":      982,
	"plusmn":   177,
	"pound":    163,
	"prime":    8242,
	"prod":     8719,
	"prop":     8733,
	"psi":      968,
	"quot":     34,
	"rArr":     8658,
	"radic":    8730,
	"rang":     9002,
	"raquo":    187,
	"rarr":     8594,
	"rceil":    8969,
	"rdquo":    8221,
	"real":     8476,
	"reg":      174,
	"rfloor":   8971,
	"rho":      961,
	"rlm":      8207,
	"rsaquo":   8250,
	"rsquo":    8217,
	"sbquo":    8218,
	"scaron":   353,
	"sdot":     8901,
	"sect":     167,
	"shy":      173,
	"sigma":    963,
	"sigmaf":   962,
	"sim":      8764,
	"spades":   9824,
	"sub":      8834,
	"sube":     8838,
	"sum":      8721,
	"sup":      8835,
	"sup1":     185,
	"sup2":     178,
	"sup3":     179,
	"supe":     8839,
	"szlig":    223,
	"tau":      964,
	"there4":   8756,
	"theta":    952,
	"thetasym": 977,
	"thinsp":   8201,
	"thorn":    254,
	"tilde":    732,
	"times":    215,
	"trade":    8482,
	"uArr":     8657,
	"uacute":   250,
	"uarr":     8593,
	"ucirc":    251,
	"ugrave":   249,
	"uml":      168,
	"upsih":    978,
	"upsilon":  965,
	"uuml":     252,
	"weierp":   8472,
	"xi":       958,
	"yacute":   253,
	"yen":      165,
	"yuml":     255,
	"zeta":     950,
	"zwj":      8205,
	"zwnj":     8204,
}

var HtmlEntitiesReverse = map[rune][]string{
	34:   []string{"quot"},
	39:   []string{"apos"},
	62:   []string{"gt"},
	160:  []string{"nbsp"},
	161:  []string{"iexcl"},
	162:  []string{"cent"},
	163:  []string{"pound"},
	164:  []string{"curren"},
	165:  []string{"yen"},
	166:  []string{"brvbar"},
	167:  []string{"sect"},
	168:  []string{"uml"},
	169:  []string{"copy"},
	170:  []string{"ordf"},
	171:  []string{"laquo"},
	172:  []string{"not"},
	173:  []string{"shy"},
	174:  []string{"reg"},
	175:  []string{"macr"},
	176:  []string{"deg"},
	177:  []string{"plusmn"},
	178:  []string{"sup2"},
	179:  []string{"sup3"},
	180:  []string{"acute"},
	181:  []string{"micro"},
	182:  []string{"para"},
	183:  []string{"middot"},
	184:  []string{"cedil"},
	185:  []string{"sup1"},
	186:  []string{"ordm"},
	187:  []string{"raquo"},
	188:  []string{"frac14"},
	189:  []string{"frac12"},
	190:  []string{"frac34"},
	191:  []string{"iquest"},
	192:  []string{"Agrave"},
	193:  []string{"Aacute"},
	194:  []string{"Acirc"},
	195:  []string{"Atilde"},
	196:  []string{"Auml"},
	197:  []string{"Aring"},
	198:  []string{"AElig"},
	199:  []string{"Ccedil"},
	200:  []string{"Egrave"},
	201:  []string{"Eacute"},
	202:  []string{"Ecirc"},
	203:  []string{"Euml"},
	204:  []string{"Igrave"},
	205:  []string{"Iacute"},
	206:  []string{"Icirc"},
	207:  []string{"Iuml"},
	208:  []string{"ETH"},
	209:  []string{"Ntilde"},
	210:  []string{"Ograve"},
	211:  []string{"Oacute"},
	212:  []string{"Ocirc"},
	213:  []string{"Otilde"},
	214:  []string{"Ouml"},
	215:  []string{"times"},
	216:  []string{"Oslash"},
	217:  []string{"Ugrave"},
	218:  []string{"Uacute"},
	219:  []string{"Ucirc"},
	220:  []string{"Uuml"},
	221:  []string{"Yacute"},
	222:  []string{"THORN"},
	223:  []string{"szlig"},
	224:  []string{"agrave"},
	225:  []string{"aacute"},
	226:  []string{"acirc"},
	227:  []string{"atilde"},
	228:  []string{"auml"},
	229:  []string{"aring"},
	230:  []string{"aelig"},
	231:  []string{"ccedil"},
	232:  []string{"egrave"},
	233:  []string{"eacute"},
	234:  []string{"ecirc"},
	235:  []string{"euml"},
	236:  []string{"igrave"},
	237:  []string{"iacute"},
	238:  []string{"icirc"},
	239:  []string{"iuml"},
	240:  []string{"eth"},
	241:  []string{"ntilde"},
	242:  []string{"ograve"},
	243:  []string{"oacute"},
	244:  []string{"ocirc"},
	245:  []string{"otilde"},
	246:  []string{"ouml"},
	247:  []string{"divide"},
	248:  []string{"oslash"},
	249:  []string{"ugrave"},
	250:  []string{"uacute"},
	251:  []string{"ucirc"},
	252:  []string{"uuml"},
	253:  []string{"yacute"},
	254:  []string{"thorn"},
	255:  []string{"yuml"},
	338:  []string{"OElig"},
	339:  []string{"oelig"},
	352:  []string{"Scaron"},
	353:  []string{"scaron"},
	376:  []string{"Yuml"},
	402:  []string{"fnof"},
	710:  []string{"circ"},
	732:  []string{"tilde"},
	913:  []string{"Alpha"},
	914:  []string{"Beta"},
	915:  []string{"Gamma"},
	916:  []string{"Delta"},
	917:  []string{"Epsilon"},
	918:  []string{"Zeta"},
	919:  []string{"Eta"},
	920:  []string{"Theta"},
	921:  []string{"Iota"},
	922:  []string{"Kappa"},
	923:  []string{"Lambda"},
	924:  []string{"Mu"},
	925:  []string{"Nu"},
	926:  []string{"Xi"},
	927:  []string{"Omicron"},
	928:  []string{"Pi"},
	929:  []string{"Rho"},
	931:  []string{"Sigma"},
	932:  []string{"Tau"},
	933:  []string{"Upsilon"},
	934:  []string{"Phi"},
	935:  []string{"Chi"},
	936:  []string{"Psi"},
	937:  []string{"Omega"},
	945:  []string{"alpha"},
	946:  []string{"beta"},
	947:  []string{"gamma"},
	948:  []string{"delta"},
	949:  []string{"epsilon"},
	950:  []string{"zeta"},
	951:  []string{"eta"},
	952:  []string{"theta"},
	953:  []string{"iota"},
	954:  []string{"kappa"},
	955:  []string{"lambda"},
	956:  []string{"mu"},
	957:  []string{"nu"},
	958:  []string{"xi"},
	959:  []string{"omicron"},
	960:  []string{"pi"},
	961:  []string{"rho"},
	962:  []string{"sigmaf"},
	963:  []string{"sigma"},
	964:  []string{"tau"},
	965:  []string{"upsilon"},
	966:  []string{"phi"},
	967:  []string{"chi"},
	968:  []string{"psi"},
	969:  []string{"omega"},
	977:  []string{"thetasym"},
	978:  []string{"upsih"},
	982:  []string{"piv"},
	8194: []string{"ensp"},
	8195: []string{"emsp"},
	8201: []string{"thinsp"},
	8204: []string{"zwnj"},
	8205: []string{"zwj"},
	8206: []string{"lrm"},
	8207: []string{"rlm"},
	8211: []string{"ndash"},
	8212: []string{"mdash"},
	8216: []string{"lsquo"},
	8217: []string{"rsquo"},
	8218: []string{"sbquo"},
	8220: []string{"ldquo"},
	8221: []string{"rdquo"},
	8222: []string{"bdquo"},
	8224: []string{"dagger"},
	8225: []string{"Dagger"},
	8226: []string{"bull"},
	8230: []string{"hellip"},
	8240: []string{"permil"},
	8242: []string{"prime"},
	8243: []string{"Prime"},
	8249: []string{"lsaquo"},
	8250: []string{"rsaquo"},
	8254: []string{"oline"},
	8260: []string{"frasl"},
	8364: []string{"euro"},
	8465: []string{"image"},
	8472: []string{"weierp"},
	8476: []string{"real"},
	8482: []string{"trade"},
	8501: []string{"alefsym"},
	8592: []string{"larr"},
	8593: []string{"uarr"},
	8594: []string{"rarr"},
	8595: []string{"darr"},
	8596: []string{"harr"},
	8629: []string{"crarr"},
	8656: []string{"lArr"},
	8657: []string{"uArr"},
	8658: []string{"rArr"},
	8659: []string{"dArr"},
	8660: []string{"hArr"},
	8704: []string{"forall"},
	8706: []string{"part"},
	8707: []string{"exist"},
	8709: []string{"empty"},
	8711: []string{"nabla"},
	8712: []string{"isin"},
	8713: []string{"notin"},
	8715: []string{"ni"},
	8719: []string{"prod"},
	8721: []string{"sum"},
	8722: []string{"minus"},
	8727: []string{"lowast"},
	8730: []string{"radic"},
	8733: []string{"prop"},
	8734: []string{"infin"},
	8736: []string{"ang"},
	8743: []string{"and"},
	8744: []string{"or"},
	8745: []string{"cap"},
	8746: []string{"cup"},
	8747: []string{"int"},
	8756: []string{"there4"},
	8764: []string{"sim"},
	8773: []string{"cong"},
	8776: []string{"asymp"},
	8800: []string{"ne"},
	8801: []string{"equiv"},
	8804: []string{"le"},
	8805: []string{"ge"},
	8834: []string{"sub"},
	8835: []string{"sup"},
	8836: []string{"nsub"},
	8838: []string{"sube"},
	8839: []string{"supe"},
	8853: []string{"oplus"},
	8855: []string{"otimes"},
	8869: []string{"perp"},
	8901: []string{"sdot"},
	8968: []string{"lceil"},
	8969: []string{"rceil"},
	8970: []string{"lfloor"},
	8971: []string{"rfloor"},
	9001: []string{"lang"},
	9002: []string{"rang"},
	9674: []string{"loz"},
	9824: []string{"spades"},
	9827: []string{"clubs"},
	9829: []string{"hearts"},
	9830: []string{"diams"},
}

// EOF
