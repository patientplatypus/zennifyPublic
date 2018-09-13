

// const noneEmptyFunc = (newVal) => {
//     var promise = new Promise(resolve=>{
//         console.log('inside noneEmptyFunc and value of newVal: ', newVal);
//         // if (newVal.recurring==='weekly'){
    
//         // }
//         resolve('nonempty return')
//     })
//     return promise;
// }

const recurringTimeComposer = (newVal) => {
    console.log('inside recurringTimeComposer and value of payload: ', newVal);
    var isAnyEmpty = false;
    Object.keys(newVal).some((key, index)=>{
        if(newVal[key]==="" || newVal[key]===null || newVal[key]===0){
            isAnyEmpty = true;
        }
        return isAnyEmpty === true            
    });
    if (isAnyEmpty===false){
        if(newVal.recurring==='weekly'){
            var returnTime = [];
            var loopWeekly = true;
            var startDate = newVal.startDate;
            var endDate = newVal.endDate;
            var startDay = startDate.getDay();
            var loopDate = new Date(startDate.getTime());
            do {
                console.log('value of loopDate; ', loopDate);
                if(loopDate.getDay()===startDay){
                    console.log('loopDate.getDay(): ', loopDate.getDay());
                    returnTime.push(loopDate.toString());
                }
                loopDate.setDate(loopDate.getDate()+1);
                if(Date.parse(loopDate)>Date.parse(newVal.endDate)){
                    loopWeekly = false;
                    console.log('value of returnTime: ', returnTime);
                }
            }while(loopWeekly === true);
            return returnTime;
        }else{
            return 'weekly is false';
        }
    }else{
        return 'some fields empty';
    }
}

export default recurringTimeComposer;






// onmessage = function(e) {
//     console.log('Message received from main script: ', e.data);
//     // var workerResult = 'Result: ' + e.data;
//     // console.log('Posting message back to main script');
//     // postMessage('hello back handsome');
//     close();
// }

// worker.onmessage = e => {
//     console.log(e.data)
//   }
  
//   worker.onerror = e => {
//     console.error(e.message)
//   }

// import moment from 'moment';


// const recurringTimeComposer = (payload) => {
//     console.log('437823872378237 inside recurringTimeComposer and payload: ', payload);
//     console.log('payload.recurring: ', payload.recurring);
//     console.log('payload.recurring==="weekly": ', payload.recurring==="weekly");
//     var returnTime = [];
    
//     if(payload.recurring==='weekly'){

//         var loopWeekly = true;
//         var startDate = payload.startDate;
//         var endDate = payload.endDate;
//         var startDay = startDate.getDay();

//         async function processWeekly(){
//             for(var loopDate = startDate; Date.parse(loopDate)<=Date.parse(payload.endDate); loopDate.setDate(loopDate.getDate()+1)){
//                 const promise = new Promise(resolve=>{

//                 })
//             } 
//         }

        // moment().add(1, 'days').calendar();  

        // for(var loopDate = startDate; Date.parse(loopDate)<=Date.parse(payload.endDate); loopDate.setDate(loopDate.getDate()+1)){
        //     if(loopDate.getDay()===startDay){
        //         returnTime.push(loopDate);
        //     }
        //     var nextLoop = loopDate.setDate(loopDate.getDate()+1);
        //     if(Date.parse(nextLoop)>Date.parse(payload.endDate)){
        //         console.log("437823872378237 value of returnTime: ", returnTime);
        //         return returnTime;
        //     }
        // }
        // do {
        //     console.log('437823872378237 inside do while and loopDate: ', loopDate);
        //     console.log('437823872378237 inside do while and loopDate.getDay(): ', loopDate.getDay());
        //     console.log('437823872378237 inside do while and startDay: ', startDay);
        //     console.log('437823872378237 inside do while and loopDate: ', loopDate);
        //     console.log('437823872378237 and value of returnTime: ', returnTime);
        //     if(loopDate.getDay()===startDay){
        //         returnTime.push(loopDate);
        //     }
        //     loopDate.setDate(loopDate.getDate()+1);
        //     console.log('437823872378237 after loopDate and 1 day and value of loopDate: ', loopDate);
        //     if(Date.parse(loopDate)>Date.parse(payload.endDate)){
        //         loopWeekly = false;
        //     }
        // }while(loopWeekly === true);

        // console.log('437823872378237 right before return and value of returnTime is : ', returnTime);
        // return returnTime;

//     }else if (payload.recurring==='monthly'){

//     }else if (payload.recurring==='yearly'){

//     }

//     return 0;
//     // return {labels: Labels, data: Data};
// }

// export default recurringTimeComposer;