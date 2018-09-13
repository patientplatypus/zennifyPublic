<template>
  <div class='DatePicker' :style="{textAlign: 'center'}">
    <table align='center' :style="{width: '100%', background:'var(--primary2)'}">
        <tr class="yearBox" align='center'>
            <td>
                <div class='arrow-down cursorPointer' :style="{transform: 'rotate(90deg)'}" @click="changeYear(-1)"/>
            </td>
            <td :style="{textAlign: 'center'}">
                {{dispYear}}
            </td>
            <td>
                <div class='arrow-down cursorPointer' :style="{transform: 'rotate(-90deg)'}" @click="changeYear(1)"/>
            </td>
        </tr>
        <tr class="monthBox" align='center'>
            <td>
                <div class='arrow-down cursorPointer' :style="{transform: 'rotate(90deg)'}" @click="changeMonth(-1)"/>
            </td>
            <td align='center'>
                {{dispMonth}}
            </td>
            <td>
               <div class='arrow-down cursorPointer' :style="{transform: 'rotate(-90deg)'}" @click="changeMonth(1)"/>
            </td>
        </tr>
    </table>
    <table align='center' cellspacing=0 class='tableBox'
    :style="{width: '90%', marginLeft: '5%', marginRight: '5%'}"
    >
        <tr>
            <td class="dayBox sunBox">
                Su
            </td>
            <td class="dayBox">
                M
            </td>
            <td class="dayBox">
                T
            </td>
            <td class="dayBox">
                W
            </td>
            <td class="dayBox">
                Th
            </td>
            <td class="dayBox">
                F
            </td>
            <td class="dayBox satBox">
                S
            </td>
        </tr>

<!-- :style="{color: handleDate(rowItem, colItem)===0?'transparent':'black'}" -->
            <tr v-for="rowItem in rows">
                <td v-for="colItem in cols" class="dateBox" 
                @click='handleDateClicked(rowItem, colItem)'
                :style="{color: handleDate(rowItem, colItem)===0?'transparent':'black'}">
                    <div class='todayDate'   
                    :class="{'selectDate':handleDate(rowItem, colItem)!=0&&checkDateDif(rowItem, colItem)===true, 'selectDateGrey': checkDateDif(rowItem, colItem)===false,'clickedDate':dateClicked.toString()===[rowItem, colItem].toString()&&monthYearClicked.toString()===[dispMonth, dispYear].toString(),'colorRecurring':handleRecurring(rowItem, colItem)}"    
                    v-if="(handleDate(rowItem, colItem)===currentDate)&&monthNames.indexOf(dispMonth) == currentMonth && dispYear == currentYear" >
                        {{handleDate(rowItem, colItem)}}
                    </div>
                    <div
                    :class="{'selectDate':handleDate(rowItem, colItem)!=0&&checkDateDif(rowItem, colItem)===true, 'selectDateGrey': checkDateDif(rowItem, colItem)===false, 'clickedDate':dateClicked.toString()===[rowItem, colItem].toString()&&monthYearClicked.toString()===[dispMonth, dispYear].toString(), 'colorRecurring':handleRecurring(rowItem, colItem)}"   
                    v-else>
                        {{handleDate(rowItem, colItem)}}
                    </div>
                    <!-- <div>
                        {{handleDate(rowItem, colItem)}}
                    </div> -->
                </td>
            </tr>


    </table>
  </div>
</template>

<script>



export default {
  name: 'DatePicker',
  props: ['pickerType', 'notBefore', 'recurring'],
  data: function(){
    return {
        dispDate: new Date(),
        dispMonth: null,
        dispYear: null,
        row1Start: null,
        recurringNum: null,
        currentMonth: null,
        currentYear: null,
        currentDate: null,
        endDateNum: null,
        handleDateState: [
            [0, 0, 0, 0, 0, 0, 0], 
            [0, 0, 0, 0, 0, 0, 0], 
            [0, 0, 0, 0, 0, 0, 0], 
            [0, 0, 0, 0, 0, 0, 0], 
            [0, 0, 0, 0, 0, 0, 0], 
            [0, 0, 0, 0, 0, 0, 0], 
            [0, 0, 0, 0, 0, 0, 0], 
        ],
        dateClicked: ['', ''],
        monthYearClicked: [0, 0],
        todayIndex: [0, 0],
        monthNames: ["January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"
        ],
        cols: [
            1, 2, 3, 4, 5, 6, 7
        ],
        rows: [
            1, 2, 3, 4, 5, 6
        ], 
        emittedK: null, 
        emitter: {
            recurrenceToEmit: [],
            dateReturn: {
                type: null,
                value: null
            }
        }
    }
  }, 
  beforeDestroy: function(){
  },
  updated: function () {
    this.$nextTick(function () {
        if(this.emitter.dateReturn.type!=null && this.emitter.dateReturn.value!=null){
            this.$emit('dateReturn', {type: this.emitter.dateReturn.type, value: this.emitter.dateReturn.value});
        }
        if(this.recurring!=null){
            this.$emit('pickerFinishedFunc', this.pickerType);
        }
        // if(this.emitter.recurrenceToEmit.length>0){
        //     this.$emit('emitRecurrence', this.emitter.recurrenceToEmit);
        // }
    })
  },
  mounted: function(){
      console.log("date.getMonth: ", this.dispDate.getMonth());
      console.log(this.monthNames[this.dispDate.getMonth()]);
      var monthNum = this.dispDate.getMonth();
      this.dispMonth =  this.monthNames[monthNum];
      this.dispYear = this.dispDate.getFullYear();
      this.currentYear = this.dispYear;
      this.currentMonth =  this.dispDate.getMonth();
      this.currentDate = this.dispDate.getDate();
  },
  watch: {
    $props: {
        handler() {
            this.parseData();
        },
        deep: true,
        immediate: true,    
    }
  },
  methods:{
      lastDateChecker: function(dateNumber){
          var lastD = new Date(this.dispDate.getFullYear(), this.dispDate.getMonth() + 1, 0);
          var lastDate = lastD.getDate();
          if (dateNumber>lastDate){
              return 0;
          }else{
              return dateNumber;
          }
      },
      handleDate: function(rowItem, colItem){
        var firstDate = new Date(this.dispDate.getFullYear(), this.dispDate.getMonth(), 1)
        var firstDay = firstDate.getDay();
        var lastDate = new Date(this.dispDate.getFullYear(), this.dispDate.getMonth(), 0).getDate();
        if (rowItem===1){
            if (this.row1Start != null){
                if(colItem === 7){
                    this.endDateNum = colItem - this.row1Start + 1;
                }
                if(colItem >= this.row1Start){
                    this.handleDateState[rowItem-1][colItem-1] = colItem - this.row1Start + 1;
                    return colItem - this.row1Start + 1;
                }
            }else if (this.row1Start === null){
                console.log('3433434333455445343 value of colItem: ', colItem);
                if((colItem)===firstDay+1){
                    console.log('3433434333455445343 value of firstDay: ', firstDay);
                    console.log('3433434333455445343 (colItem): ', (colItem));
                    console.log('3433434333455445343 (rowItem): ', (rowItem));
                    this.row1Start = colItem;
                    console.log('beta');
                    this.handleDateState[rowItem-1][colItem-1] = 1;
                    return 1;
                }else{
                    return 0;
                }
            }
        }else if(this.endDateNum != null){
            if ( this.endDateNum+7*(rowItem-2)+colItem <= lastDate){
                this.handleDateState[rowItem-1][colItem-1] = this.endDateNum+7*(rowItem-2)+colItem;
                return this.endDateNum+7*(rowItem-2)+colItem;
            }else{
                this.handleDateState[rowItem-1][colItem-1] = 0;
                return 0;
            }
        }
      }, 
      changeMonth(addVal){
          var newMonthIndex = (this.monthNames.indexOf(this.dispMonth) + addVal)%12;
          if (newMonthIndex === -1){
              newMonthIndex = 11;
          }
          var newMonthName = this.monthNames[newMonthIndex];
          this.dispMonth = newMonthName;
          this.updateDispDate();
      }, 
      changeYear(addVal){
          var newYear = Number(this.dispYear) + addVal;
          this.dispYear = newYear;
          this.updateDispDate();
      },
      updateDispDate(){
        var d = new Date(this.dispYear,this.monthNames.indexOf(this.dispMonth),1);
        this.dispDate = d;
        this.row1Start = null;
        this.$forceUpdate();
      },
      todaysDate(dateVal, colItem, rowItem){
        if(dateVal === this.currentDate && this.monthNames.indexOf(this.dispMonth) == this.currentMonth && this.dispYear == this.currentYear){
            this.todayIndex = [rowItem, colItem];
        }
        return;
      },
      handleDateClicked(rowItem, colItem){   
    
        console.log('83984937573945734957 inside dateClicked!')
        var dateSelected = this.handleDateState[rowItem-1][colItem-1];
        if (dateSelected>0){
            console.log("83984937573945734957 inside dateSelected>0")
            var dateString = (this.monthNames.indexOf(this.dispMonth)+1).toString()+"/"+dateSelected+"/"+this.dispYear+" 00:00:01";
            var k = new Date(dateString);
            if (this.notBefore != undefined || this.notBefore != null){
                console.log('83984937573945734957 inside first if');
                console.log('83984937573945734957 inside first if and Date.parse(this.notBefore): ', Date.parse(this.notBefore));
                console.log('83984937573945734957 inside first if and this.notBefore: ', this.notBefore);
                console.log('83984937573945734957 inside first if and Date.parse(k): ', Date.parse(k));
                if (Date.parse(this.notBefore) <= Date.parse(k)){
                    console.log('83984937573945734957 inside second if');
                    this.dateClicked = [rowItem, colItem];
                    this.monthYearClicked = [this.dispMonth, this.dispYear];
                    this.emittedK = k;
                    this.emitter.dateReturn.type = this.pickerType;
                    this.emitter.dateReturn.value = k;
                    // this.$emit('dateReturn', {type: this.pickerType, value: k});
                }
            }else{
                    console.log('83984937573945734957 inside else')
                    this.dateClicked = [rowItem, colItem];
                    this.monthYearClicked = [this.dispMonth, this.dispYear];
                    this.emittedK = k;
                    this.emitter.dateReturn.type = this.pickerType;
                    this.emitter.dateReturn.value = k;
                    // this.$emit('dateReturn', {type: this.pickerType, value: k});
            }
        } 
      },
      parseData: function(){
        if(Date.parse(this.emittedK) < Date.parse(this.notBefore)){
            this.emittedK = null;
            this.dateClicked = [0, 0];
            this.monthYearClicked = ['', ''];
            // this.$emit('dateReturn', {type: this.pickerType, value: ""});
            this.emitter.dateReturn.type = this.pickerType;
            this.emitter.dateReturn.value = "";
            this.$forceUpdate();
        }
      }, 
      checkDateDif: function(rowItem, colItem){
        var dateSelected = this.handleDateState[rowItem-1][colItem-1];
        var dateString = (this.monthNames.indexOf(this.dispMonth)+1).toString()+"/"+dateSelected+"/"+this.dispYear+" 00:00:01";
        var k = new Date(dateString);
        if (this.notBefore != undefined || this.notBefore != null){
            if(Date.parse(k) < Date.parse(this.notBefore)){
                return false;
            }else{
                return true;
            }
        }else{
            return true;
        }
      }, 
      handleRecurring: function(rowItem, colItem){
        // var lastD = new Date(this.dispDate.getFullYear(), this.dispDate.getMonth() + 1, 0);
        // var lastDate = lastD.getDate();
        if(this.notBefore != null && this.notBefore != undefined)
        {
            var boxDateVal = this.handleDateState[rowItem-1][colItem-1];
            var dateString = (this.monthNames.indexOf(this.dispMonth)+1).toString()+"/"+boxDateVal+"/"+this.dispYear+" 00:00:01";
            var k = new Date(dateString);
            // if(Date.parse(k)===Date.parse(this.emittedK) && boxDateVal === lastDate){
            //     this.$emit('recurrenceReturn', this.recurrenceToEmit);
            // }

            if (
                (this.recurring==='weekly') &&
                (boxDateVal != 0)
                ){
                    var beforeDay = this.notBefore.getDay() + 1;
                    if((colItem === beforeDay) && Date.parse(k) >= Date.parse(this.notBefore)){
                        if(this.emittedK!=null && Date.parse(k)>Date.parse(this.emittedK)){
                            return false;
                        }else{
                            // if(this.emitter.recurrenceToEmit.indexOf(k.toString())===-1){
                            //     this.emitter.recurrenceToEmit.push(k.toString())
                            // }
                            return true;
                        }
                    }
            }else if (
                (this.recurring==='monthly') &&
                (boxDateVal != 0)
                ){
                    var beforeDate = this.notBefore.getDate();
                    if((boxDateVal === beforeDate) && Date.parse(k) >= Date.parse(this.notBefore)){
                        if(this.emittedK!=null && Date.parse(k)>Date.parse(this.emittedK)){
                            return false;
                        }else{
                            // if(this.emitter.recurrenceToEmit.indexOf(k.toString())===-1){
                            //     this.emitter.recurrenceToEmit.push(k.toString())
                            // }
                            return true;
                        }
                    }
            }else if(
                (this.recurring==='yearly')&&
                (boxDateVal!=0)
                ){
                    var beforeDate = this.notBefore.getDate();
                    var beforeYear = this.notBefore.getFullYear();
                    var beforeMonth = this.notBefore.getMonth();

                    if(beforeYear <= this.dispYear &&
                        this.monthNames[beforeMonth] === this.dispMonth &&
                        beforeDate  === boxDateVal){
                            if(this.emittedK!=null && Date.parse(k)>Date.parse(this.emittedK)){
                                return false;
                            }else{
                                // if(this.emitter.recurrenceToEmit.indexOf(k.toString())===-1){
                                //     this.emitter.recurrenceToEmit.push(k.toString())
                                // }
                                return true;
                            }
                    }
            }
        }
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>


.tableBox{
    table-layout: fixed;
    height: 200px;
    width: 200px;
    text-align: center;
    background: var(--primary2);
    border-color: var(--secondary1);
    border-width: 1px;
    border-style: solid;
    /* empty-cells: hide; */
}

.dateBox{
    border-color: var(--secondary1);
    border-width: 1px;
    border-style: solid;
}

.dayBox{
    border-color: var(--secondary1);
    border-top-width: 1px;
    border-bottom-width: 1px;
    border-right-width: 0px;
    border-left-width: 0px;
    border-style: solid;
}

.satBox{
    border-color: var(--secondary1);
    border-right-width: 1px;
    border-style: solid;
}

.sunBox{
    border-color: var(--secondary1);
    border-left-width: 1px;
    border-style: solid;
}

.todayDate{
    background: var(--secondary1);
    color: var(--primary2);
}

.selectDate:hover{
    background: var(--secondary2);
    color: var(--primary1);
    cursor: pointer;
}

.selectDateGrey:hover{
    background: rgba(0,0,0,0.3);
    color: var(--primary1);
    cursor: pointer;
}

.cursorPointer:hover{
    cursor: pointer;
}

.clickedDate{
    background: var(--secondary2);
    color: var(--primary2);
}

.colorRecurring{
    font-weight: 900;
    color: var(--attention);
}

</style>
