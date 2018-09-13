

const newLoanPie = (payload) => {

    function decimalize(x) {
        return Number(Number.parseFloat(x).toFixed(2));
    }
    var cleanPrinciple = decimalize(payload.principle);
    var cleanInterestAmount = decimalize(payload.interestAmount);
    var cleanDenominator = cleanPrinciple + cleanInterestAmount;
    var cleanIntFraction = cleanInterestAmount / cleanDenominator;
    var cleanPriFraction = cleanPrinciple / cleanDenominator;
    console.log('893489034894389 in newLoanPie and value of payload: ', payload);
    console.log("893489034894389 value of cleanPrinciple: ", cleanPrinciple);
    console.log("893489034894389 value of cleanInterestAmount: ", cleanInterestAmount);
    return(
        {
            type: 'pie', 
            data: {
                labels:["Principal", "Interest"], 
                datasets:[{
                    label: "Dollars", 
                    backgroundColor: ["#3e95cd", "#8e5ea2"], 
                    data: [cleanPrinciple, cleanInterestAmount]
                }]
            }, 
            options: {
                legend: {
                    display: true,
                    position: 'bottom'
                },
                title: {
                    display: true,
                    text: 'Loan Breakdown'
                }
            }
        }
    )
}


export default newLoanPie;