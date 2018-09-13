

const newRecurringExpense = (payload) => {
    console.log('inisde newRecurringExpense and value of payload: ', payload);
    
    var dateTypes = payload[0];
    var dateTimes = payload[1];
    var paymentAmount = payload[2];
    var dataLabels = [];
    var dataPayment = [];
    var dataText = '';

    var monthNames = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];

    if(dateTypes === 'weekly'){
        dataText = 'Recurring Payment (Weekly)'
    }else if (dateTypes === 'monthly'){
        dataText = 'Recurring Payment (Monthly)'
    }else if (dateTypes === 'yearly'){
        dataText = 'Recurring Payment (Yearly)'
    }

    dateTimes.forEach(dateT => {
        var dateTime = new Date(dateT);
        var dateMonth = monthNames[dateTime.getMonth()];
        var dateDay = dateTime.getDate();
        var dateYear = dateTime.getFullYear();
        var dateString = dateMonth + ", " + dateDay + ", " + dateYear;
        dataPayment.push(paymentAmount);
        dataLabels.push(dateString)
    })

    return(
        {
            type: 'bar', 
            data: {
                labels: dataLabels, 
                datasets:[{
                    label: "Dollars", 
                    data: dataPayment
                }]
            }, 
            options: {
                legend: {
                    display: true,
                    position: 'bottom'
                },
                title: {
                    display: true,
                    text: dataText
                }
            }
        }
    )
}


export default newRecurringExpense;