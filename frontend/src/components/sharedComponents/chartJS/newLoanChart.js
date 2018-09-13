

const newLoanChart = (payload) => {
    return(
        {
            type: 'bar', 
            data: {
                datasets: [{
                    label: 'Original Loan',
                    data: [payload.bar.data], 
                    backgroundColor: ["#3e95cd"]
                }, {
                    type: 'line',
                    label: 'Loan + Interest', 
                    data: payload.line.data, 
                    backgroundColor: ["rgb(142, 94, 162, 0.5)"]
                }], 
                labels: payload.line.labels
            }, 
            options: {
                fill: false,
                responsive: true,
                scales: {
                    xAxes: [{
                        type: 'time',
                        display: true,
                        scaleLabel: {
                            display: true,
                            labelString: "Date",
                        }
                    }],
                    yAxes: [{
                        ticks: {
                            beginAtZero: true,
                        },
                        display: true,
                        scaleLabel: {
                            display: true,
                            labelString: "Dollars",
                        }
                    }]
                }
            }
        }
    );
}

export default newLoanChart;



// var mixedChart = new Chart(ctx, {
//     type: 'bar',
//     data: {
//       datasets: [{
//             label: 'Bar Dataset',
//             data: [10, 20, 30, 40]
//           }, {
//             label: 'Line Dataset',
//             data: [50, 50, 50, 50],
  
//             // Changes this dataset to become a line
//             type: 'line'
//           }],
//       labels: ['January', 'February', 'March', 'April']
//     },
//     options: options
//   });
  