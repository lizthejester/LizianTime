package ltime

/*
LEAP YEAR
	Menotheen: Jan 1 - 31 + Feb 1-5
	Lengten: Feb 6-29 + Mar 1-13
	Regen: Mar 14-31 + Apr 1-18
	Leorar: Apr 19-30 + May 1-25
	Mysund: May 26-31 + Jun 1-30
	Heisswerm: Jul 1-31 + Aug 1-6
	Largaheiss: Aug 7-31 + Sept 1-11
	Pommois: Sept 12-30 + Oct 1-18
	Spinnan: Oct 19-31 + Nov 1-23
	Kalt: Nov 24-30 + Dec 1-31 */
/*
NON LEAP YEAR
	Menotheen: Jan 1 - 31 + Feb 1-5
	Lengten: Feb 6-28 + Mar 1-14
	Regen: Mar 15-31 + Apr 1-19
	Leorar: Apr 20-30 + May 1-26
	Mysund: May 27-31 + Jun 1-30 + Jul 1
	Heisswerm: Jul 2-31 + Aug 1-7
	Largaheiss: Aug 8-31 + Sept 1-12
	Pommois: Sept 13-30 + Oct 1-19
	Spinnan: Oct 20-31 + Nov 1-24
	Kalt: Nov 25-30 + Dec 1-31 */

func GetDayOfWeek(lizDay int, lizMonth string) string {
	var dayOfYear int
	switch lizMonth {
	case "Menotheen":
		dayOfYear = lizDay
	case "Lengten":
		dayOfYear = lizDay + 36
	case "Regen":
		dayOfYear = lizDay + 73
	case "Leorar":
		dayOfYear = lizDay + 109
	case "Mysund":
		dayOfYear = lizDay + 146
	case "Heisswerm":
		dayOfYear = lizDay + 182
	case "Largaheiss":
		dayOfYear = lizDay + 219
	case "Pommois":
		dayOfYear = lizDay + 255
	case "Spinnan":
		dayOfYear = lizDay + 292
	case "Kalt":
		dayOfYear = lizDay + 328
	}
	daysOfTheWeek := [10]string{"Deca", "Monit", "Tweg", "Treag", "Tessafleur", "Sectaeg", "Deghex", "Siebaeg", "Octa", "Noywaeg"}

	return daysOfTheWeek[dayOfYear%10]
}

func GetDayMonth(year int, month string, day int) (string, int) {
	var lizMonth string
	var lizDay int
	leapYear := false
	if year%4 == 0 {
		leapYear = true
	}

	if leapYear {
		switch month {
		case "january", "jan", "01":
			lizMonth = "Menotheen"
			lizDay = day
		case "february", "feb", "02":
			if day <= 6 {
				lizMonth = "Menotheen"
				lizDay = day + 30
			} else {
				lizMonth = "Lengten"
				lizDay = day - 5
			}
		case "march", "mar", "03":
			if day <= 13 {
				lizMonth = "Lengten"
				lizDay = day + 24
			} else {
				lizMonth = "Regen"
				lizDay = day - 13
			}
		case "april", "apr", "04":
			if day <= 18 {
				lizMonth = "Regen"
				lizDay = day + 18
			} else {
				lizMonth = "Leorar"
				lizDay = day - 18
			}
		case "may", "05":
			if day <= 25 {
				lizMonth = "Leorar"
				lizDay = day + 12
			} else {
				lizMonth = "Mysund"
				lizDay = day - 25
			}
		case "june", "jun", "06":
			lizMonth = "Mysund"
			lizDay = day + 6
		case "july", "jul", "07":
			lizMonth = "Heisswerm"
			lizDay = day
		case "august", "aug", "08":
			if day <= 6 {
				lizMonth = "Heisswerm"
				lizDay = day + 31
			} else {
				lizMonth = "Largaheiss"
				lizDay = day - 6
			}
		case "september", "sept", "sep", "09":
			if day <= 11 {
				lizMonth = "Largaheiss"
				lizDay = day + 25
			} else {
				lizMonth = "Pommois"
				lizDay = day - 11
			}
		case "october", "oct", "10":
			if day <= 18 {
				lizMonth = "Pommois"
				lizDay = day + 19
			} else {
				lizMonth = "Spinnan"
				lizDay = day - 18
			}
		case "november", "nov", "11":
			if day <= 23 {
				lizMonth = "Spinnan"
				lizDay = day + 13
			} else {
				lizMonth = "Kalt"
				lizDay = day - 23
			}
		case "december", "dec", "12":
			lizMonth = "Kalt"
			lizDay = day + 7
		}
	} else {
		switch month {
		case "january", "jan", "01":
			lizMonth = "Menotheen"
			lizDay = day
		case "february", "feb", "02":
			if day <= 5 {
				lizMonth = "Menotheen"
				lizDay = day + 31
			} else {
				lizMonth = "Lengten"
				lizDay = day - 5
			}
		case "march", "mar", "03":
			if day <= 14 {
				lizMonth = "Lengten"
				lizDay = day + 23
			} else {
				lizMonth = "Regen"
				lizDay = day - 14
			}
		case "april", "apr", "04":
			if day <= 19 {
				lizMonth = "Regen"
				lizDay = day + 17
			} else {
				lizMonth = "Leorar"
				lizDay = day - 19
			}
		case "05", "may":
			if day <= 26 {
				lizMonth = "Leorar"
				lizDay = day + 11
			} else {
				lizMonth = "Mysund"
				lizDay = day - 26
			}
		case "june", "jun", "06":
			lizMonth = "Mysund"
			lizDay = day + 5
		case "july", "jul", "07":
			if day == 1 {
				lizMonth = "Mysund"
				lizDay = 36
			} else {
				lizMonth = "Heisswerm"
				lizDay = day - 1
			}
		case "august", "aug", "08":
			if day <= 7 {
				lizMonth = "Heisswerm"
				lizDay = day + 30
			} else {
				lizMonth = "Largaheiss"
				lizDay = day - 7
			}
		case "september", "sept", "sep", "09":
			if day <= 12 {
				lizMonth = "Largaheiss"
				lizDay = day + 24
			} else {
				lizMonth = "Pommois"
				lizDay = day - 12
			}
		case "october", "oct", "10":
			if day <= 19 {
				lizMonth = "Pommois"
				lizDay = day + 18
			} else {
				lizMonth = "Spinnan"
				lizDay = day - 19
			}
		case "november", "nov", "11":
			if day <= 24 {
				lizMonth = "Spinnan"
				lizDay = day + 12
			} else {
				lizMonth = "Kalt"
				lizDay = day - 24
			}
		case "december", "dec", "12":
			lizMonth = "Kalt"
			lizDay = day + 6
		}
	}
	return lizMonth, lizDay
}
