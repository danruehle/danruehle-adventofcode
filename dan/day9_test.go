package dan

import (
	"errors"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var day9Input = []int64{
	18,
	4,
	30,
	21,
	7,
	39,
	11,
	42,
	12,
	31,
	1,
	44,
	22,
	5,
	10,
	48,
	14,
	40,
	36,
	2,
	37,
	9,
	24,
	28,
	35,
	8,
	43,
	3,
	13,
	17,
	4,
	6,
	23,
	7,
	11,
	12,
	15,
	16,
	72,
	38,
	10,
	59,
	5,
	41,
	14,
	18,
	19,
	21,
	24,
	9,
	25,
	8,
	20,
	28,
	43,
	53,
	13,
	31,
	17,
	22,
	33,
	23,
	49,
	26,
	27,
	15,
	29,
	34,
	30,
	35,
	38,
	42,
	51,
	21,
	32,
	48,
	25,
	37,
	36,
	40,
	41,
	28,
	39,
	56,
	44,
	43,
	45,
	46,
	47,
	52,
	108,
	49,
	53,
	60,
	79,
	57,
	65,
	77,
	58,
	73,
	80,
	61,
	138,
	64,
	82,
	67,
	71,
	83,
	94,
	87,
	88,
	91,
	140,
	96,
	101,
	109,
	144,
	115,
	118,
	119,
	121,
	125,
	206,
	167,
	128,
	131,
	132,
	146,
	227,
	138,
	190,
	154,
	170,
	175,
	178,
	344,
	200,
	340,
	205,
	210,
	230,
	233,
	236,
	237,
	469,
	246,
	253,
	403,
	378,
	259,
	263,
	270,
	373,
	380,
	316,
	385,
	463,
	345,
	448,
	383,
	405,
	410,
	415,
	435,
	696,
	466,
	639,
	553,
	644,
	795,
	650,
	523,
	579,
	522,
	529,
	533,
	798,
	790,
	661,
	699,
	1261,
	1429,
	750,
	853,
	933,
	840,
	825,
	1306,
	1051,
	1162,
	988,
	1075,
	1045,
	1701,
	1374,
	1052,
	2690,
	1062,
	1055,
	1382,
	1232,
	1360,
	2381,
	1449,
	1524,
	1575,
	2159,
	3221,
	1665,
	1758,
	1885,
	1813,
	2666,
	4886,
	2437,
	2063,
	2405,
	2107,
	3992,
	8107,
	2630,
	2287,
	4488,
	3047,
	5428,
	2935,
	2809,
	3478,
	3512,
	3099,
	3333,
	3423,
	3550,
	4250,
	4737,
	3698,
	4724,
	4350,
	6432,
	4170,
	6146,
	4394,
	4916,
	4917,
	5096,
	5439,
	5856,
	7031,
	5744,
	5908,
	6649,
	7593,
	7269,
	6522,
	6756,
	6883,
	6973,
	7248,
	8894,
	8092,
	9833,
	8520,
	9914,
	8564,
	9086,
	13770,
	9310,
	13729,
	10013,
	10535,
	18577,
	19473,
	12775,
	11652,
	12881,
	15650,
	14004,
	19099,
	21252,
	32248,
	13856,
	14221,
	15340,
	16612,
	16656,
	17084,
	23869,
	21339,
	17650,
	19621,
	19323,
	29431,
	20548,
	21665,
	22187,
	42735,
	28077,
	25656,
	33740,
	26737,
	27860,
	29196,
	33696,
	29561,
	35979,
	30940,
	36277,
	31952,
	48024,
	38843,
	41286,
	36973,
	37271,
	45277,
	56298,
	57521,
	49742,
	60501,
	43852,
	53516,
	59352,
	54597,
	57608,
	55933,
	57056,
	57421,
	65838,
	80682,
	78557,
	72952,
	62892,
	68229,
	70795,
	124162,
	116960,
	74244,
	80825,
	93569,
	111653,
	145039,
	93594,
	97368,
	98449,
	99785,
	108113,
	134490,
	114664,
	123259,
	112989,
	193354,
	208233,
	141449,
	241234,
	131121,
	139024,
	183784,
	142473,
	151620,
	155069,
	247137,
	167813,
	242603,
	187163,
	238809,
	331492,
	190962,
	195817,
	207898,
	278328,
	222777,
	227653,
	359853,
	316613,
	244110,
	270145,
	379273,
	272570,
	417352,
	273594,
	290644,
	338853,
	418594,
	354976,
	363630,
	665731,
	524511,
	378125,
	398860,
	638181,
	386779,
	570235,
	602050,
	430675,
	603963,
	542715,
	752164,
	660373,
	514255,
	1235966,
	543739,
	672454,
	546164,
	817333,
	564238,
	629497,
	693829,
	718606,
	794305,
	741755,
	764904,
	776985,
	785639,
	1216314,
	817454,
	901034,
	1338044,
	1118218,
	944930,
	1056970,
	1088879,
	1204112,
	1425277,
	1057994,
	1089903,
	1175661,
	1833609,
	1993299,
	2033809,
	1952646,
	1323326,
	1458733,
	1460361,
	1506659,
	1518740,
	1969016,
	2517331,
	1686673,
	2629389,
	1718488,
	1845964,
	2001900,
	2146873,
	2232631,
	4551140,
	2549240,
	2483271,
	2518355,
	2147897,
	3841681,
	3739290,
	3555957,
	2782059,
	3147034,
	2967020,
	2783687,
	3654556,
	2979101,
	3520640,
	4380528,
	5035686,
	3405161,
	3564452,
	4981001,
	3720388,
	3847864,
	10016687,
	5266958,
	7162587,
	4631168,
	5265330,
	5114917,
	6338016,
	4929956,
	5565746,
	5749079,
	5750707,
	5761160,
	8094018,
	5762788,
	8151808,
	7125549,
	6384262,
	6925801,
	6969613,
	7568252,
	14693801,
	8701389,
	9113194,
	8351556,
	11649592,
	9561124,
	9746085,
	9896498,
	10044873,
	10195286,
	13816306,
	18262513,
	10495702,
	12133341,
	19642583,
	12134969,
	11523948,
	12147050,
	14114344,
	13310063,
	18063954,
	13952514,
	13895414,
	23670998,
	15919808,
	24309630,
	26028755,
	19307209,
	18847258,
	20241787,
	19457622,
	25638292,
	33571966,
	23658917,
	24391116,
	36326467,
	23805765,
	22019650,
	23657289,
	25476462,
	29229871,
	25419362,
	37609803,
	27205477,
	27262577,
	57230883,
	45661149,
	29815222,
	42518256,
	34767066,
	43116539,
	47463054,
	38154467,
	38304880,
	39699409,
	41477272,
	45678567,
	45676939,
	45825415,
	62695996,
	53620987,
	47439012,
	72921533,
	49076651,
	55234584,
	54649233,
	69780833,
	54468054,
	61972543,
	108855571,
	64582288,
	83815616,
	73071946,
	74466475,
	76244338,
	91355506,
	77853876,
	87155839,
	78004289,
	108521411,
	102088245,
	116316983,
	131478922,
	93264427,
	96515663,
	101907066,
	109411555,
	203995311,
	103544705,
	109117287,
	127721179,
	144247308,
	116440597,
	151076235,
	137654234,
	175171122,
	147538421,
	149316284,
	154098214,
	154248627,
	155858165,
	165009715,
	200060368,
	171268716,
	189780090,
	195171493,
	212661992,
	252860989,
	225852152,
	198422729,
	300392519,
	212956260,
	219985302,
	225557884,
	236838466,
	244161776,
	254094831,
	263979018,
	296854705,
	285192655,
	305174449,
	301787048,
	303414498,
	392696631,
	310106792,
	320867880,
	377671707,
	361048806,
	546720032,
	432647294,
	393594222,
	411084721,
	687778499,
	703701014,
	411378989,
	432941562,
	438514144,
	489536902,
	462396350,
	481000242,
	498256607,
	1136642576,
	549171673,
	582047360,
	662835854,
	605201546,
	622654928,
	664463304,
	507622668,
	630974672,
	867587912,
	810613269,
	1120911535,
	1112824214,
	804678943,
	804973211,
	822463710,
	946136812,
	931198169,
	844320551,
	871455706,
	900910494,
	943396592,
	1293810526,
	979256849,
	1047428280,
	1312301611,
	1285490782,
	1089670028,
	1694871574,
	1739043618,
	1138597340,
	1643720153,
	1312595879,
	1957144765,
	2499844785,
	1609652154,
	1649293762,
	1891748831,
	4194716359,
	1627436921,
	1790457363,
	1715776257,
	1745231045,
	3276730683,
	1772366200,
	3028372136,
	2332919062,
	2026685129,
	2359729891,
	2375160810,
	2228267368,
	2910963540,
	2402265907,
	3636979876,
	3237089075,
	3084962079,
	4589326562,
	2922248033,
	4920120967,
	4294014738,
	3258945916,
	3343213178,
	3372667966,
	3488142457,
	4132096091,
	5256639504,
	3461007302,
	3517597245,
	3799051329,
	4386415020,
	4359604191,
	8341601583,
	4254952497,
	5892758055,
	5465356443,
	11149397559,
	5487227986,
	5324513940,
	6007210112,
	6159337108,
	6181193949,
	6265461211,
	8926363745,
	6602159094,
	6978604547,
	8838024409,
	12985814659,
	10439611849,
	6949149759,
	7715959799,
	7260058631,
	7316648574,
	7877201436,
	8054003826,
	8614556688,
	9579466437,
	18993024771,
	9720308940,
	18316813285,
	10789870383,
	11668421935,
	10811741926,
	11331724052,
	12166547220,
	12340531057,
	13159798496,
	16181625531,
	16558070984,
	13551308853,
	13927754306,
	25942174530,
	23507220743,
	18049929014,
	14209208390,
	14576707205,
	16668560514,
	15931205262,
	16491758124,
	19722425761,
	25761091968,
	20369336820,
	20510179323,
	29717869480,
	21601612309,
	22121594435,
	22143465978,
	22978289146,
	36906043452,
	28724618204,
	30109379837,
	35904051292,
	28136962696,
	27479063159,
	43660688690,
	28504461511,
	28785915595,
	30140413652,
	38909494408,
	69421780658,
	85913538782,
	32599765776,
	49738575005,
	36861094944,
	40232605084,
	58644875163,
	50626055946,
	50647927489,
	43723206744,
	50907510030,
	49600657594,
	91140115114,
	50457352305,
	55616025855,
	62740179428,
	72165150201,
	64408512803,
	78386573189,
	57290377106,
	58926329247,
	61104227287,
	61385681371,
	94371134233,
	69460860720,
	108383450168,
	72832370860,
	87318447249,
	77093700028,
	87768604974,
	83955811828,
	120108788209,
	108197887136,
	93323864338,
	101364862335,
	114009170397,
	185291587164,
	141126752617,
	130565088007,
	120312010618,
	116216706353,
	122489908658,
	166155178163,
	118394604393,
	138197927315,
	120030556534,
	130846542091,
	200172518181,
	157229465694,
	142293231580,
	160600975834,
	149926070888,
	161049511856,
	250491118716,
	234321181015,
	230687795794,
	194688726673,
	201521751474,
	207333034735,
	254414633668,
	316081249051,
	242520465192,
	269044469406,
	234611310746,
	438020830529,
	273139773671,
	295427393009,
	249241146484,
	256592531708,
	250877098625,
	393170330205,
	280772612979,
	292219302468,
	299522697274,
	302894207414,
	451281258381,
	523459103074,
	368382546591,
	425376522467,
	396210478147,
	402021761408,
	552135353898,
	585125718457,
	458210133360,
	832907966877,
	608300551519,
	477131775938,
	483852457230,
	548763843758,
	500118245109,
	576200005988,
	505833678192,
	507469630333,
	531649711604,
	543096401093,
	572991915447,
	580295310253,
	591741999742,
	960124546333,
	728270729881,
	764593024738,
	874216224783,
	770404307999,
	798232239555,
	1104641627051,
	1115770181852,
	1086128988445,
	1593598618778,
	977250021047,
	960984233168,
	1298350484664,
	982965454130,
	983970702339,
	1005951923301,
	1007587875442,
	1013303308525,
	1037483389796,
	1039119341937,
	1320012729623,
	1116088316540,
	1629225389538,
	1454511535036,
	1858186927122,
	1492863754619,
	1562825264293,
	1534997332737,
	1568636547554,
	1731388541167,
	1759216472723,
	1960215475177,
	1938234254215,
	1943949687298,
	1966936156469,
	1944954935507,
	2529620780722,
	2467814843561,
	1988917377431,
	3308930107054,
	4064618113459,
	2020891183967,
	2050786698321,
	2076602731733,
	2155207658477,
	2436101046163,
	2570599851576,
	2947375289655,
	5006700897739,
	3027861087356,
	3294213805460,
	3097822597030,
	5517975141231,
	3697450726938,
	3675338228465,
	3883189189722,
	3882183941513,
	6644826016593,
	3888904622805,
	6241589095115,
	5686368104369,
	4144125035908,
	6111439274628,
	10533730639398,
	4071677882288,
	4176098842444,
	4097493915700,
	4127389430054,
	4512703777896,
	4591308704640,
	5383476335818,
	7091500325563,
	7019053171943,
	8704151624677,
	6969552033925,
	6392036402490,
	11494915610446,
	9383818831307,
	11560860738565,
	7557522169978,
	8303488272498,
	8169171797988,
	7960582505093,
	7986398538505,
	12104707541001,
	14527074203903,
	13896522609203,
	8199067312342,
	8224883345754,
	8247776724732,
	23910893035210,
	8640093207950,
	13988605205868,
	9104012482536,
	16664734129770,
	14061052359488,
	22292093478366,
	15726693967966,
	22285935705242,
	13361588436415,
	13949558572468,
	15782405515732,
	25768746612306,
	27287554706531,
	15518104675071,
	15946981043598,
	16159649817435,
	24622117157607,
	16185465850847,
	16423950658096,
	23165064842024,
	37860451607678,
	22589651780418,
	24799743025385,
	24912510854502,
	17744105690486,
	29787746327454,
	22465600918951,
	25050993526134,
	54838739853588,
	36415159491419,
	61025516449702,
	38107756455489,
	32345115668282,
	93370632117984,
	68788298426056,
	31677754492506,
	32609416508943,
	50089221358768,
	86516494346094,
	32106630861033,
	33929571541333,
	38651066769798,
	34168056348582,
	38889551577047,
	70344731032752,
	52253347246405,
	49850736551519,
}

func firstBadNumber(t *testing.T) (int, int64, error) {
	for i := 25; i < len(day9Input); i++ {
		sum := day9Input[i]
		var found bool
		for j := i - 25; j < i-1; j++ {
			for k := j + 1; k < i; k++ {
				if day9Input[j]+day9Input[k] == sum {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return i, sum, nil
		}
	}
	return 0, 0, errors.New("Did not find number")
}

func TestDay9a(t *testing.T) {
	i, s, err := firstBadNumber(t)
	require.NoError(t, err)
	t.Logf("Number %d is not the sum of any of the previous 25 numbers, its value is %d", i, s)
}

func sum(n []int64) int64 {
	var sum int64
	for _, i := range n {
		sum += i
	}
	return sum
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestDay9b(t *testing.T) {
	i, s, err := firstBadNumber(t)
	require.NoError(t, err)
	t.Logf("Number %d is not the sum of any of the previous 25 numbers, its value is %d", i, s)
	for start := 0; start < i-1; start++ {
		for end := start + 2; end < i; end++ {
			numbers := Int64Slice(day9Input[start:end])
			if sum(numbers) == s {
				t.Logf("Contiguous range that adds to %d starts with number %d and ends with number %d", s, start, end-1)
				t.Logf("Numbers: %v", numbers)
				sort.Sort(numbers)
				t.Logf("Sum of first and last: %d", numbers[0]+numbers[len(numbers)-1])
				return
			}
		}
	}
	t.Error("Did not find range of numbers")
}
