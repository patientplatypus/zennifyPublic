const newInterestTimeSeries = (payload) => {

    function decimalize(x) {
        return Number(Number.parseFloat(x).toFixed(2));
    }

    function daysDiff(startDate, endDate) {
        const timeDiff  = (endDate - startDate);
        const days      = timeDiff / (1000 * 60 * 60 * 24);
        return days;
    }

    var Labels = [];
    var Data = [];

    var compounded = payload.compounded; //1, 3, 6, 12
    var amount     = payload.amount;
    var interest   = payload.interest/100;
    const startDate  = payload.startDate;
    var endDate    = payload.endDate;

    var iterateDate =  new Date(payload.startDate.getTime());
    var interestOwed = 0;
    var daysEllapsed = 0;
    var totalDaysDif = daysDiff(startDate, endDate);

    var daysIterate = 0;

    if (totalDaysDif < 500){
        daysIterate = 1;
    }else if (totalDaysDif < 1000){
        daysIterate = 7;
    }else if (totalDaysDif >= 1000){
        daysIterate = 30;
    }

    //0th case
    Labels.push(iterateDate.toString());

    //Nth case
    var daysCount = 0
    do{
        iterateDate.setDate(iterateDate.getDate()+daysIterate);
        Labels.push(iterateDate.toString());

        var daysEllapsed = daysDiff(startDate, new Date(iterateDate.toString()));
        console.log("value of daysEllapsed: ", daysEllapsed);
        var interestBase = (1+interest/(compounded));
        console.log("value of interestBase: ", interestBase);
        var interestPow = (compounded)*(daysEllapsed/365);
        console.log('value of interestPow: ', interestPow);
        var interestComposed = amount*Math.pow(interestBase, interestPow);
        console.log("value of interestComposed: ", interestComposed);
        var interestDecimalized = decimalize(interestComposed);
        console.log('value of interestDecimalized: ', interestDecimalized);
        Data.push(interestDecimalized)

        daysCount = daysCount + daysIterate
        console.log('inside doWhile loop and value of daysCount: ', daysCount);
    }while(daysCount<totalDaysDif);

    console.log('after do while loop and value of Labels: ', Labels);    
    console.log('after do while loop and value of Data: ', Data);

    return {labels: Labels, data: Data};
}

export default newInterestTimeSeries;