package constants

type Holiday string

const (
	NewYear            Holiday = "元旦"
	SpringFestival     Holiday = "春节"
	TombSweepingDay    Holiday = "清明节"
	LabourDay          Holiday = "劳动节"
	DragonBoatFestival Holiday = "端午节"
	NationalDay        Holiday = "国庆节"
	MidAutumnFestival  Holiday = "中秋节"
)

var Holidays = map[int]map[string]Holiday{
	2004: {
		"01-01": NewYear,
		"01-22": SpringFestival,
		"01-23": SpringFestival,
		"01-24": SpringFestival,
		"01-25": SpringFestival,
		"01-26": SpringFestival,
		"01-27": SpringFestival,
		"01-28": SpringFestival,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"05-05": LabourDay,
		"05-06": LabourDay,
		"05-07": LabourDay,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2005: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"02-09": SpringFestival,
		"02-10": SpringFestival,
		"02-11": SpringFestival,
		"02-12": SpringFestival,
		"02-13": SpringFestival,
		"02-14": SpringFestival,
		"02-15": SpringFestival,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"05-05": LabourDay,
		"05-06": LabourDay,
		"05-07": LabourDay,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2006: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"01-29": SpringFestival,
		"01-30": SpringFestival,
		"01-31": SpringFestival,
		"02-01": SpringFestival,
		"02-02": SpringFestival,
		"02-03": SpringFestival,
		"02-04": SpringFestival,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"05-05": LabourDay,
		"05-06": LabourDay,
		"05-07": LabourDay,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2007: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"02-18": SpringFestival,
		"02-19": SpringFestival,
		"02-20": SpringFestival,
		"02-21": SpringFestival,
		"02-22": SpringFestival,
		"02-23": SpringFestival,
		"02-24": SpringFestival,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2008: {
		"01-01": NewYear,
		"02-06": SpringFestival,
		"02-07": SpringFestival,
		"02-08": SpringFestival,
		"02-09": SpringFestival,
		"02-10": SpringFestival,
		"02-11": SpringFestival,
		"02-12": SpringFestival,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"06-07": DragonBoatFestival,
		"06-08": DragonBoatFestival,
		"06-09": DragonBoatFestival,
		"09-13": MidAutumnFestival,
		"09-14": MidAutumnFestival,
		"09-15": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2009: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"01-25": SpringFestival,
		"01-26": SpringFestival,
		"01-27": SpringFestival,
		"01-28": SpringFestival,
		"01-29": SpringFestival,
		"01-30": SpringFestival,
		"01-31": SpringFestival,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-28": DragonBoatFestival,
		"05-29": DragonBoatFestival,
		"05-30": DragonBoatFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": MidAutumnFestival,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2010: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"02-13": SpringFestival,
		"02-14": SpringFestival,
		"02-15": SpringFestival,
		"02-16": SpringFestival,
		"02-17": SpringFestival,
		"02-18": SpringFestival,
		"04-03": TombSweepingDay,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"06-14": DragonBoatFestival,
		"06-15": DragonBoatFestival,
		"06-16": DragonBoatFestival,
		"09-22": MidAutumnFestival,
		"09-23": MidAutumnFestival,
		"09-24": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2011: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"02-02": SpringFestival,
		"02-03": SpringFestival,
		"02-04": SpringFestival,
		"02-05": SpringFestival,
		"02-06": SpringFestival,
		"02-07": SpringFestival,
		"04-03": TombSweepingDay,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"04-30": LabourDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"06-04": DragonBoatFestival,
		"06-06": DragonBoatFestival,
		"09-10": MidAutumnFestival,
		"09-11": MidAutumnFestival,
		"09-12": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2012: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"01-22": SpringFestival,
		"01-23": SpringFestival,
		"01-24": SpringFestival,
		"01-25": SpringFestival,
		"01-26": SpringFestival,
		"01-27": SpringFestival,
		"01-28": SpringFestival,
		"04-02": TombSweepingDay,
		"04-03": TombSweepingDay,
		"04-04": TombSweepingDay,
		"04-29": LabourDay,
		"04-30": LabourDay,
		"05-01": LabourDay,
		"06-22": DragonBoatFestival,
		"06-24": DragonBoatFestival,
		"09-30": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2013: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"02-09": SpringFestival,
		"02-10": SpringFestival,
		"02-11": SpringFestival,
		"02-12": SpringFestival,
		"02-13": SpringFestival,
		"02-14": SpringFestival,
		"02-15": SpringFestival,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"04-29": LabourDay,
		"04-30": LabourDay,
		"05-01": LabourDay,
		"06-10": DragonBoatFestival,
		"06-11": DragonBoatFestival,
		"06-12": DragonBoatFestival,
		"09-19": MidAutumnFestival,
		"09-20": MidAutumnFestival,
		"09-21": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2014: {
		"01-01": NewYear,
		"01-31": SpringFestival,
		"02-01": SpringFestival,
		"02-02": SpringFestival,
		"02-03": SpringFestival,
		"02-04": SpringFestival,
		"02-05": SpringFestival,
		"02-06": SpringFestival,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"04-07": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"06-02": DragonBoatFestival,
		"09-08": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2015: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"02-18": SpringFestival,
		"02-19": SpringFestival,
		"02-20": SpringFestival,
		"02-21": SpringFestival,
		"02-22": SpringFestival,
		"02-23": SpringFestival,
		"02-24": SpringFestival,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"05-01": LabourDay,
		"06-20": DragonBoatFestival,
		"06-22": DragonBoatFestival,
		"09-27": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2016: {
		"01-01": NewYear,
		"02-07": SpringFestival,
		"02-08": SpringFestival,
		"02-09": SpringFestival,
		"02-10": SpringFestival,
		"02-11": SpringFestival,
		"02-12": SpringFestival,
		"02-13": SpringFestival,
		"04-04": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"06-09": DragonBoatFestival,
		"06-10": DragonBoatFestival,
		"06-11": DragonBoatFestival,
		"09-15": MidAutumnFestival,
		"09-16": MidAutumnFestival,
		"09-17": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2017: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-27": SpringFestival,
		"01-28": SpringFestival,
		"01-29": SpringFestival,
		"01-30": SpringFestival,
		"01-31": SpringFestival,
		"02-01": SpringFestival,
		"02-02": SpringFestival,
		"04-02": TombSweepingDay,
		"04-03": TombSweepingDay,
		"04-04": TombSweepingDay,
		"05-01": LabourDay,
		"05-28": DragonBoatFestival,
		"05-29": DragonBoatFestival,
		"05-30": DragonBoatFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": MidAutumnFestival,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2018: {
		"01-01": NewYear,
		"02-15": SpringFestival,
		"02-16": SpringFestival,
		"02-17": SpringFestival,
		"02-18": SpringFestival,
		"02-19": SpringFestival,
		"02-20": SpringFestival,
		"02-21": SpringFestival,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"04-07": TombSweepingDay,
		"04-29": LabourDay,
		"04-30": LabourDay,
		"05-01": LabourDay,
		"06-18": DragonBoatFestival,
		"09-24": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
		"12-30": NewYear,
		"12-31": NewYear,
	},
	2019: {
		"01-01": NewYear,
		"02-04": SpringFestival,
		"02-05": SpringFestival,
		"02-06": SpringFestival,
		"02-07": SpringFestival,
		"02-08": SpringFestival,
		"02-09": SpringFestival,
		"02-10": SpringFestival,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"04-07": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"06-07": DragonBoatFestival,
		"06-08": DragonBoatFestival,
		"06-09": DragonBoatFestival,
		"09-13": MidAutumnFestival,
		"09-14": MidAutumnFestival,
		"09-15": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2020: {
		"01-01": NewYear,
		"01-24": SpringFestival,
		"01-25": SpringFestival,
		"01-26": SpringFestival,
		"01-27": SpringFestival,
		"01-28": SpringFestival,
		"01-29": SpringFestival,
		"01-30": SpringFestival,
		"01-31": SpringFestival,
		"02-01": SpringFestival,
		"02-02": SpringFestival,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"05-05": LabourDay,
		"06-25": DragonBoatFestival,
		"06-26": DragonBoatFestival,
		"06-27": DragonBoatFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
		"10-08": NationalDay,
	},
	2021: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"02-11": SpringFestival,
		"02-12": SpringFestival,
		"02-13": SpringFestival,
		"02-14": SpringFestival,
		"02-15": SpringFestival,
		"02-16": SpringFestival,
		"02-17": SpringFestival,
		"04-03": TombSweepingDay,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"05-05": LabourDay,
		"06-12": DragonBoatFestival,
		"06-13": DragonBoatFestival,
		"06-14": DragonBoatFestival,
		"09-19": MidAutumnFestival,
		"09-20": MidAutumnFestival,
		"09-21": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2022: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-03": NewYear,
		"01-31": SpringFestival,
		"02-01": SpringFestival,
		"02-02": SpringFestival,
		"02-03": SpringFestival,
		"02-04": SpringFestival,
		"02-05": SpringFestival,
		"02-06": SpringFestival,
		"04-03": TombSweepingDay,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"04-30": LabourDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"06-03": DragonBoatFestival,
		"06-04": DragonBoatFestival,
		"06-05": DragonBoatFestival,
		"09-10": MidAutumnFestival,
		"09-11": MidAutumnFestival,
		"09-12": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
		"12-31": NewYear,
	},
	2023: {
		"01-01": NewYear,
		"01-02": NewYear,
		"01-21": SpringFestival,
		"01-22": SpringFestival,
		"01-23": SpringFestival,
		"01-24": SpringFestival,
		"01-25": SpringFestival,
		"01-26": SpringFestival,
		"01-27": SpringFestival,
		"04-05": TombSweepingDay,
		"04-29": LabourDay,
		"04-30": LabourDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"06-22": DragonBoatFestival,
		"06-23": DragonBoatFestival,
		"06-24": DragonBoatFestival,
		"09-29": MidAutumnFestival,
		"09-30": NationalDay,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"12-30": NewYear,
		"12-31": NewYear,
	},
	2024: {
		"01-01": NewYear,
		"02-10": SpringFestival,
		"02-11": SpringFestival,
		"02-12": SpringFestival,
		"02-13": SpringFestival,
		"02-14": SpringFestival,
		"02-15": SpringFestival,
		"02-16": SpringFestival,
		"02-17": SpringFestival,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"05-05": LabourDay,
		"06-10": DragonBoatFestival,
		"09-15": MidAutumnFestival,
		"09-16": MidAutumnFestival,
		"09-17": MidAutumnFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": NationalDay,
		"10-07": NationalDay,
	},
	2025: {
		"01-01": NewYear,
		"01-28": SpringFestival,
		"01-29": SpringFestival,
		"01-30": SpringFestival,
		"01-31": SpringFestival,
		"02-01": SpringFestival,
		"02-02": SpringFestival,
		"02-03": SpringFestival,
		"02-04": SpringFestival,
		"04-04": TombSweepingDay,
		"04-05": TombSweepingDay,
		"04-06": TombSweepingDay,
		"05-01": LabourDay,
		"05-02": LabourDay,
		"05-03": LabourDay,
		"05-04": LabourDay,
		"05-05": LabourDay,
		"05-31": DragonBoatFestival,
		"06-01": DragonBoatFestival,
		"06-02": DragonBoatFestival,
		"10-01": NationalDay,
		"10-02": NationalDay,
		"10-03": NationalDay,
		"10-04": NationalDay,
		"10-05": NationalDay,
		"10-06": MidAutumnFestival,
		"10-07": NationalDay,
		"10-08": NationalDay,
	},
}

var WorkDays = map[string]Holiday{
	// 2004
	"2004-01-17": SpringFestival,
	"2004-01-18": SpringFestival,
	"2004-05-08": LabourDay,
	"2004-05-09": LabourDay,
	"2004-10-09": NationalDay,
	"2004-10-10": NationalDay,
	// 2005
	"2005-02-05": SpringFestival,
	"2005-02-06": SpringFestival,
	"2005-04-30": LabourDay,
	"2005-05-08": LabourDay,
	"2005-10-08": NationalDay,
	"2005-10-09": NationalDay,
	// 2006
	"2006-01-28": SpringFestival,
	"2006-02-05": SpringFestival,
	"2006-04-29": LabourDay,
	"2006-04-30": LabourDay,
	"2006-09-30": NationalDay,
	"2006-10-08": NationalDay,
	"2006-12-30": NewYear,
	"2006-12-31": NewYear,
	// 2007
	"2007-02-17": SpringFestival,
	"2007-02-25": SpringFestival,
	"2007-04-28": LabourDay,
	"2007-05-29": LabourDay,
	"2007-09-29": NationalDay,
	"2007-09-30": NationalDay,
	"2007-12-29": NewYear,
	// 2008
	"2008-02-02": SpringFestival,
	"2008-02-03": SpringFestival,
	"2008-05-04": LabourDay,
	"2008-09-27": NationalDay,
	"2008-09-28": NationalDay,
	// 2009
	"2009-01-04": NewYear,
	"2009-01-24": SpringFestival,
	"2009-02-01": SpringFestival,
	"2009-05-31": DragonBoatFestival,
	"2009-09-27": NationalDay,
	"2009-10-10": NationalDay,
	// 2010
	"2010-02-20": SpringFestival,
	"2010-02-21": SpringFestival,
	"2010-06-12": DragonBoatFestival,
	"2010-06-13": DragonBoatFestival,
	"2010-09-19": MidAutumnFestival,
	"2010-09-25": MidAutumnFestival,
	"2010-09-26": MidAutumnFestival,
	"2010-10-09": NationalDay,
	// 2011
	"2011-01-31": SpringFestival,
	"2011-02-12": SpringFestival,
	"2011-04-02": TombSweepingDay,
	"2011-10-08": TombSweepingDay,
	"2011-10-09": TombSweepingDay,
	"2011-12-31": NewYear,
	// 2012
	"2012-01-21": SpringFestival,
	"2012-01-29": SpringFestival,
	"2012-03-31": TombSweepingDay,
	"2012-04-01": TombSweepingDay,
	"2012-04-28": LabourDay,
	"2012-09-29": NationalDay,
	// 2013
	"2013-01-05": NewYear,
	"2013-01-06": NewYear,
	"2013-02-16": SpringFestival,
	"2013-02-17": SpringFestival,
	"2013-04-07": TombSweepingDay,
	"2013-04-27": LabourDay,
	"2013-04-28": LabourDay,
	"2013-06-08": DragonBoatFestival,
	"2013-06-09": DragonBoatFestival,
	"2013-09-22": MidAutumnFestival,
	"2013-09-29": NationalDay,
	"2013-10-12": NationalDay,
	// 2014
	"2014-01-26": SpringFestival,
	"2014-02-08": SpringFestival,
	"2014-05-04": LabourDay,
	"2014-09-28": NationalDay,
	"2014-10-11": NationalDay,
	// 2015
	"2015-01-04": NewYear,
	"2015-02-15": SpringFestival,
	"2015-02-28": SpringFestival,
	"2015-10-10": NationalDay,
	// 2016
	"2016-02-06": SpringFestival,
	"2016-02-14": SpringFestival,
	"2016-06-12": DragonBoatFestival,
	"2016-09-18": MidAutumnFestival,
	"2016-10-08": NationalDay,
	"2016-10-09": NationalDay,
	// 2017
	"2017-01-22": SpringFestival,
	"2017-02-04": SpringFestival,
	"2017-04-01": TombSweepingDay,
	"2017-05-27": DragonBoatFestival,
	"2017-09-30": NationalDay,
	// 2018
	"2018-02-11": SpringFestival,
	"2018-02-24": SpringFestival,
	"2018-04-08": TombSweepingDay,
	"2018-04-28": LabourDay,
	"2018-09-29": NationalDay,
	"2018-09-30": NationalDay,
	"2018-12-29": NewYear,
	// 2019
	"2019-02-02": SpringFestival,
	"2019-02-03": SpringFestival,
	"2019-04-28": LabourDay,
	"2019-05-05": LabourDay,
	"2019-09-29": NationalDay,
	"2019-10-12": NationalDay,
	// 2020
	"2020-01-19": SpringFestival,
	"2020-04-26": LabourDay,
	"2020-05-09": LabourDay,
	"2020-06-28": DragonBoatFestival,
	"2020-09-27": NationalDay,
	"2020-10-10": NationalDay,
	// 2021
	"2021-02-07": SpringFestival,
	"2021-02-20": SpringFestival,
	"2021-04-25": LabourDay,
	"2021-05-08": LabourDay,
	"2021-09-18": MidAutumnFestival,
	"2021-09-26": NationalDay,
	"2021-10-09": NationalDay,
	// 2022
	"2022-01-29": SpringFestival,
	"2022-01-30": SpringFestival,
	"2022-04-02": TombSweepingDay,
	"2022-04-24": LabourDay,
	"2022-05-07": LabourDay,
	"2022-10-08": NationalDay,
	"2022-10-09": NationalDay,
	// 2023
	"2023-01-28": SpringFestival,
	"2023-01-29": SpringFestival,
	"2023-04-23": LabourDay,
	"2023-05-06": LabourDay,
	"2023-06-25": DragonBoatFestival,
	"2023-10-07": NationalDay,
	"2023-10-08": NationalDay,
	// 2024
	"2024-02-04": SpringFestival,
	"2024-02-18": SpringFestival,
	"2024-04-07": TombSweepingDay,
	"2024-04-28": LabourDay,
	"2024-05-11": LabourDay,
	"2024-09-14": MidAutumnFestival,
	"2024-09-29": NationalDay,
	"2024-10-12": NationalDay,
	// 2025
	"2025-01-26": SpringFestival,
	"2025-02-08": SpringFestival,
	"2025-04-27": LabourDay,
	"2025-09-28": NationalDay,
	"2025-10-11": NationalDay,
}
