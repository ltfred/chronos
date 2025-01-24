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

var WorkDays = map[string]Holiday{}
