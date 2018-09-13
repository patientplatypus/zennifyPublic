
const exampleChartData = {
    type: 'pie',
    data: {
        labels: ["Awesome", "Sick", "Winning"],
        datasets: [{
            label: "Cool Factor",
            backgroundColor: ["#3e95cd", "#8e5ea2","#3cba9f"],
            data: [33,33,33]
        }]
    },
    options: {
        legend: {
            display: true,
            position: 'bottom'
        },
        title: {
            display: true,
            text: 'Amount of Being'
        }
    }
}

export default exampleChartData;
