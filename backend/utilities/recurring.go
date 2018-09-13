package util

import (
	// standard library
	"time"
	"fmt"
	// "log"
	// "bytes"
	// "net/smtp"
	// sp "github.com/SparkPost/gosparkpost"
)

func inTimeSpan(start, end, check time.Time) bool {
    return (check.After(start) && check.Before(end))
}

func Parse_Recurring(startDate string, endDate string, recurring string)[]string{
	fmt.Println("inside Parse_Recurring");
	fmt.Println("value of startDate: ", startDate);
	fmt.Println("value of endDate: ", endDate);

	var returnDateArray []string;
	
	startDateTime, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(startDateTime)

	endDateTime, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(endDateTime)

	loopDate, err := time.Parse(time.RFC3339, startDate);
	addDays := 0;
	addMonths := 0;
	addYears := 0;
	fmt.Println("value of recurring: ", recurring)
	if recurring=="weekly"{
		fmt.Println("inside recurring if 1")
		addDays = 7
	}else if recurring=="monthly"{
		fmt.Println("inside recurring if 2")
		addMonths = 1
	}else if recurring=="yearly"{
		fmt.Println("inside recurring if 3")
		addYears = 1
	}

	addTime := true
	// firstLoop := 0
	var startBool bool
	var endBool bool
	var composedBool bool
	fmt.Println("before loop and value of addDays: ", addDays)
	for addTime == true{

		// fmt.Println("inside iteration of addTime loop in Parse_Recurring")

		startBool = loopDate.After(startDateTime) || (loopDate.After(startDateTime) || !(loopDate.Before(startDateTime) || loopDate.After(startDateTime)))		
		// fmt.Println("value of startBool : ", startBool)
		
		endBool = (loopDate.Before(endDateTime) || !(loopDate.Before(startDateTime) || loopDate.After(startDateTime)))
		// fmt.Println("value of endBool : ", endBool)

		composedBool = startBool && endBool
		// fmt.Println("value of composedBool: ", composedBool)

		if composedBool == true{
			// fmt.Println("adding loopDate to returnDateArray and value of loopDate: ", loopDate.String())
			returnDateArray = append(returnDateArray, loopDate.String());
		}else{
			// fmt.Println("failed if statement with a loopDate value of ", loopDate.String())
			addTime = false;
		}

		fmt.Println("value of loopDate BEFORE adding time: ", loopDate.String())
		loopDate = loopDate.AddDate(addYears, addMonths, addDays)
		fmt.Println("value of loopDate AFTER adding time: ", loopDate.String())

	}

	fmt.Println("value of returnDateArray before returning: ")
	fmt.Println(returnDateArray)

	return returnDateArray;

}
