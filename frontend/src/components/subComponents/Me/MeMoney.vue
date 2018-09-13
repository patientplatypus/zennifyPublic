<template>
  <div>

    <!-- BEGIN money menu navigation -->

    <div class='me money' :style="{position: 'fixed', top: 0, left: 0}">
        <div class='meMoneyMenu card1'>  
            <h2>
                money menu
            </h2>
            <div class='flexbox-container' :style="{flexDirection: 'row'}">
                <div :style="{flex: 5, textAlign: 'right', marginRight: '2vw'}">
                    View Data
                </div>
                <div :style="{flex: 1}">
                    <input readonly class='button1' 
                    @click="changeNav('viewData')"
                    :class="{toggled: moneyToggled.viewData}"
                    value='data'/>
                </div>
            </div>
            <div class='flexbox-container' :style="{flexDirection: 'row'}">
                <div :style="{flex: 5, textAlign: 'right', marginRight: '2vw'}">
                    Add Profit
                </div>
                <div :style="{flex: 1}">
                    <input readonly class='button1' 
                    @click="changeNav('addProfit')"
                    :class="{toggled: moneyToggled.addProfit}"
                    value='profit'/>
                </div>
            </div>
            <div class='flexbox-container' :style="{flexDirection: 'row'}">
                <div :style="{flex: 5, textAlign: 'right', marginRight: '2vw'}">
                    Add Expense
                </div>
                <div :style="{flex: 1}">
                    <input readonly class='button1' 
                    @click="changeNav('addExpense')"
                    :class="{toggled: moneyToggled.addExpense}"
                    value='expense'/>
                </div>
            </div>
            <div class='flexbox-container' :style="{flexDirection: 'row'}">
                <div :style="{flex: 5, textAlign: 'right', marginRight: '2vw'}">
                    Settings
                </div>
                <div :style="{flex: 1}">
                    <input readonly class='button1' 
                    @click="changeNav('settings')"
                    :class="{toggled: moneyToggled.settings}"
                    value='settings'/>
                </div>
            </div>
        </div>
    </div>

    <!-- END money menu navigation -->

    <!-- BEGIN expense user flow -->

    <div v-if="moneyToggled.addExpense===true">
        <div class='me money' :style="{position: 'fixed', top: 0, left: 0}">
            <div class ='meMoneyMainScreen card1' :style="{width: '98%', marginLeft: '1%', position:'relative'}">  
                <h2>
                    add an expense
                </h2>
                <table :style="{width: '100%', maxWidth: '100%', borderCollapse: 'collapse', borderSpacing: '0 50px', tableLayout: 'fixed'}">
                    <tr class="spaceUnder" :style="{width: '80%', paddingBottom: '10px'}">
                        <td width='40%' :style="{textAlign: 'center'}">
                            name
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <input class='input1' type='text' v-model="expenseName" :style="{width: '99%'}"> 
                        </td>
                    </tr>
                    <tr class="spaceUnder">
                        <td width='40%' :style="{textAlign: 'center'}">
                            owed to
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <input class='input1' type='text' v-model="expenseOwedTo" :style="{width: '99%'}"> 
                        </td>
                    </tr>
                    <tr class="spaceUnder">
                        <td width='40%' :style="{textAlign: 'center'}">
                            description
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <textarea class='textArea1' v-model="expenseDescription" rows="4" cols="50" type='text' :style="{width: '99%', height: '20vh', resize: 'none'}"/> 
                        </td>
                    </tr>
                    <tr class="spaceUnder" :style="{width: '80%', paddingBottom: '10px'}">
                        <td width='30%' :style="{textAlign: 'center'}">
                            type
                        </td>
                        <td width='70%' :style="{textAlign: 'center'}">
                            <Dropdown
                            :style="{width: '100%'}"
                            v-on:selectReturn='selectReturn'
                            :dropDataArray="expenseType"/>
                        </td>
                    </tr>
                    <tr v-if="selectValues.expenseType==='loan'" class="spaceUnder" :style="{width: '80%', paddingBottom: '10px'}">
                        <td width='30%' :style="{textAlign: 'center'}">
                            loan type
                        </td>
                        <td width='70%' :style="{textAlign: 'center'}">
                            <Dropdown
                            :style="{width: '100%'}"
                            v-on:selectReturn='selectReturn'
                            :dropDataArray="loanType"/>
                        </td>
                    </tr>
                    <!-- new loan flow -->
                    <tr v-if="selectValues.loanType==='new loan'&&selectValues.expenseType==='loan'" class="spaceUnder">
                        <td width='40%' :style="{textAlign: 'center'}">
                            loan amount
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <input class='input1' type='text' :style="{width: '99%'}"
                            v-model="loanAmount"> 
                        </td>
                    </tr>
                    <tr v-if="selectValues.loanType==='new loan'&&selectValues.expenseType==='loan'" class="spaceUnder">
                        <td width='40%' :style="{textAlign: 'center'}">
                            interest
                        </td>
                        <td width='60%' :style="{textAlign: 'center', outline: 'none'}">
                            <input class='slider' step="0.1" min="0" max="30" type='range' v-model="interestSlider"> 
                        </td>
                    </tr>
                    <tr v-if="selectValues.loanType==='new loan'&&selectValues.expenseType==='loan'" class="spaceUnder">
                        <td colspan='2'>
                            compounded every...
                        </td>
                    </tr>
                    <tr v-if="selectValues.loanType==='new loan'&&selectValues.expenseType==='loan'" class="spaceUnder">
                        <td colspan='2' :style="{textAlign: 'center'}">
                            <div class='flexbox-container' :style="{flexDirection: 'row'}">
                                <div :style="{flex: 1}">
                                    <label class="radioContainer">
                                        one month
                                        <input type="radio" value="oneMonth" v-model="interestToggled">
                                        <span class="checkmarkRadio"></span>
                                    </label>
                                </div>
                                <div :style="{flex: 1}">
                                    <label class="radioContainer">
                                        three months
                                        <input type="radio" value="threeMonths" v-model="interestToggled">
                                        <span class="checkmarkRadio"></span>
                                    </label>
                                </div>
                                <div :style="{flex: 1}">
                                    <label class="radioContainer">
                                        six months
                                        <input type="radio" value="sixMonths" v-model="interestToggled">
                                        <span class="checkmarkRadio"></span>
                                    </label>
                                </div>
                                <div :style="{flex: 1}">
                                    <label class="radioContainer">
                                        twelve months
                                        <input type="radio" value="oneYear" v-model="interestToggled">
                                        <span class="checkmarkRadio"></span>
                                    </label>
                                </div>
                            </div>
                        </td>
                    </tr>
                    <!-- new loan flow -->
                    <!-- loan payment flow -->
                    <tr v-if="selectValues.loanType==='recurring payment' || selectValues.loanType==='one time payment'||selectValues.expenseType==='recurring'||selectValues.expenseType==='one time'" class="spaceUnder" :style="{width: '80%'}">
                        <td width='40%' :style="{textAlign: 'center'}">
                            payment
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <input class='input1' type='text' :style="{width: '99%'}"
                            v-model="paymentAmount"> 
                        </td>
                    </tr>
                    <tr v-if="selectValues.loanType==='recurring payment'||selectValues.expenseType==='recurring'" class="spaceUnder" :style="{width: '80%'}">
                        <td colspan='2' :style="{textAlign: 'center'}">
                            how often recurring payment occurs
                        </td>
                    </tr>
                    <tr v-if="selectValues.loanType==='recurring payment'||selectValues.expenseType==='recurring'" class="spaceUnder">
                        <td colspan='2' :style="{textAlign: 'center'}">
                            <div class='flexbox-container' :style="{flexDirection: 'row'}">
                                <div :style="{flex: 1}">
                                    <label class="radioContainer">
                                        weekly
                                        <input type="radio" value="weekly" v-model="recurringCheck.payment">
                                        <span class="checkmarkRadio"></span>
                                    </label>
                                </div>
                                <div :style="{flex: 1}">
                                    <label class="radioContainer">
                                        monthly
                                        <input type="radio" value="monthly" v-model="recurringCheck.payment">
                                        <span class="checkmarkRadio"></span>
                                    </label>
                                </div>
                                <div :style="{flex: 1}">
                                    <label class="radioContainer">
                                        yearly
                                        <input type="radio" value="yearly" v-model="recurringCheck.payment">
                                        <span class="checkmarkRadio"></span>
                                    </label>
                                </div>
                            </div>
                        </td>
                    </tr>
                    <div v-if="selectValues.loanType==='recurring payment' || selectValues.loanType==='one time payment'" class="spaceUnder">
                         <div :style="{textAlign: 'right'}">
                            <input readonly class='button1' value="select loan"/>
                        </div>
                    </div>
                    <!-- recurring loan payment flow -->
                </table>
                <!-- new loan flow -->
                <div :style="{position: 'absolute', right: '1vw', bottom: '2vh'}">
                    <div :style="{textAlign: 'right'}">
                        <input readonly class='button1' value="submit" @click="submitHandler"/>
                    </div>
                </div>
            </div>
            <!-- new loan flow -->
            <div class='meMoneyTimeBox card1' :style="{width: '92.5%', marginLeft: '2.5%'}">
                <div v-if='selectValues.expenseType==="one time"||selectValues.loanType==="one time payment"' :style="{height: '100%', width: '100%', textAlign: 'center'}">
                    <DatePicker
                    :style="{width: '100%'}"
                    :pickerType="'start'"
                    v-on:dateReturn='dateReturn'/>
                    <div>
                        Select Date
                    </div>
                </div>
                <div v-else-if="selectValues.loanType==='new loan'&&selectValues.expenseType==='loan'" class='flexbox-container' :style="{flexDirection: 'row'}">
                    <div :style="{flex: 1, textAlign:'center'}">
                        <DatePicker
                        :style="{width: '100%'}"
                        :pickerType="'start'"
                        v-on:dateReturn='dateReturn'/>
                        <div>
                            Select Start Date
                        </div>
                    </div>
                    <div :style="{flex: 1, textAlign:'center'}">
                        <DatePicker
                        :style="{width: '100%'}"
                        :pickerType="'end'"
                        :notBefore="dateVals.start"
                        v-on:dateReturn='dateReturn'
                        />
                        <div>
                            Select End Date
                        </div>
                    </div>
                </div>
                <div v-else-if="selectValues.expenseType==='recurring'||selectValues.loanType==='recurring payment'"
                class='flexbox-container' :style="{flexDirection: 'row'}"
                >
                    <div :style="{flex: 1, textAlign:'center'}">
                        <DatePicker
                        :style="{width: '100%'}"
                        :pickerType="'start'"
                        v-on:dateReturn='dateReturn'/>
                        <div>
                            Select Start Date
                        </div>
                    </div>
                    <div :style="{flex: 1, textAlign:'center'}">
                        <DatePicker
                        :style="{width: '100%'}"
                        :pickerType="'end'"
                        :recurring="recurringCheck.payment"
                        :notBefore="dateVals.start"
                        v-on:dateReturn='dateReturn'
                        v-on:pickerFinishedFunc='pickerFinishedFunc'
                        v-on:emitRecurrence='emitRecurrence'
                        />
                        <div>
                            Select End Date
                        </div>
                    </div>
                </div>
                <div v-else :style="{textAlign: 'center', marginTop: '10vh'}">
                    <h3>
                        Choose your money options in order to populate the calendar!
                    </h3>
                </div>
            </div>

            <div class='meMoneyGraphBox paper1' :style="{width: '92.5%', marginLeft: '2.5%',
            fontSize: '1.7rem', lineHeight: '1.9rem', position: 'relative'}">
                <div v-if="selectValues.loanType==='new loan'&&selectValues.expenseType==='loan'">
                    <div :style="{position: 'absolute', top: '5vh', left: '5vh'}">
                        <p>
                            <span :style="{fontWeight: '900'}">${{loanAmount}}</span> at
                        </p>
                        <p>
                            <span :style="{fontWeight: '900'}">{{interestSliderReturn}}%</span> interest compounded
                        </p>
                        <p v-if="interestToggled==='oneYear'">
                            <span :style="{fontWeight: '900'}">every year</span> from
                        </p>
                        <p v-if="interestToggled==='oneMonth'">
                            <span :style="{fontWeight: '900'}">every month</span> from
                        </p>
                        <p v-if="interestToggled==='threeMonths'">
                            <span :style="{fontWeight: '900'}">every three months</span> from
                        </p>
                        <p v-if="interestToggled==='sixMonths'">
                            <span :style="{fontWeight: '900'}">every six months</span> from
                        </p>
                        <p v-if="dateVals.start!=null">
                            <span :style="{fontWeight: '900'}">{{dateVals.month.start}}/{{dateVals.date.start}}/{{dateVals.year.start}}</span> to  <span :style="{fontWeight: '900'}" v-if="dateVals.date.end!=null">{{dateVals.month.end}}/{{dateVals.date.end}}/{{dateVals.year.end}}</span>
                        </p>
                    </div>
                    <div 
                    :style="{width: '35%', overflow: 'visible', position: 'absolute', top: 0, right: 0}">
                        <canvas id="loanPie" width="100%" height="100%"></canvas>
                    </div>
                    <div 
                    :style="{width: '100%', position: 'absolute', left: 0, top: '25vh'}">
                        <canvas id="loanChart" width="100%" height="30vh"></canvas>
                    </div>
                </div>
                <div v-else-if="selectValues.expenseType==='one time'">
                    <div :style="{position: 'absolute', top: '5vh', left: '5vh'}">
                        <p>
                            <span :style="{fontWeight: '900'}">One time</span> expense of
                        </p>
                        <p v-if="paymentAmount!=null">
                            <span :style="{fontWeight: '900'}">${{paymentAmount}}</span> on
                        </p>
                        <p v-if="dateVals.start!=null && paymentAmount!=null">
                            <span :style="{fontWeight: '900'}">{{dateVals.month.start}}/{{dateVals.date.start}}/{{dateVals.year.start}}</span> 
                        </p>
                    </div>
                </div>
                <div v-else-if="selectValues.expenseType==='recurring'">
                    <div :style="{position: 'absolute', top: '5vh', left: '5vh'}">
                        <p>
                            <span :style="{fontWeight: '900'}">Recurring</span> expense of
                        </p>
                        <p v-if="paymentAmount!=null">
                            <span :style="{fontWeight: '900'}">${{paymentAmount}}</span> recurring
                        </p>
                        <p  v-if="paymentAmount!=null">
                            <span :style="{fontWeight: '900'}">{{recurringCheck.payment}}</span> from 
                        </p>
                        <p v-if="dateVals.start!=null">
                            <span :style="{fontWeight: '900'}">{{dateVals.month.start}}/{{dateVals.date.start}}/{{dateVals.year.start}}</span> to  <span :style="{fontWeight: '900'}" v-if="dateVals.date.end!=null">{{dateVals.month.end}}/{{dateVals.date.end}}/{{dateVals.year.end}}</span>
                        </p>
                    </div>
                    <div v-if="showRecurringExpenseChart===true" :style="{width: '100%', position: 'absolute', left: 0, top: '25vh'}">
                        <canvas id="recurringExpenseChart" width="100%" height="30vh"></canvas>
                    </div>
                </div>
                <div v-else :style="{textAlign: 'center', marginTop: '10vh', fontWeight: '900'}">
                    <h3>
                        Choose your money options in order to populate the statistics graphs!
                    </h3>
                </div>
            </div>

        </div>
    </div>

    <!-- END expense user flow -->

    <!-- BEGIN profit user flow -->

    <div v-if="moneyToggled.addProfit===true">
        <div class='me money' :style="{position: 'fixed', top: 0, left: 0}">
            <div class ='meMoneyMainScreen card1' :style="{width: '98%', marginLeft: '1%', position: 'relative'}">  
                <h2>
                    add a profit
                </h2>
                <table :style="{width: '100%', maxWidth: '100%', borderCollapse: 'collapse', borderSpacing: '0 50px', tableLayout: 'fixed'}">
                    <tr class="spaceUnder" :style="{width: '80%', paddingBottom: '10px'}">
                        <td width='40%' :style="{textAlign: 'center'}">
                            name
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <input class='input1' type='text' v-model="profitName" :style="{width: '99%'}"> 
                        </td>
                    </tr>
                    <tr class="spaceUnder">
                        <td width='40%' :style="{textAlign: 'center'}">
                            profit from
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <input class='input1' type='text' v-model="profitFrom" :style="{width: '99%'}"> 
                        </td>
                    </tr>
                    <tr class="spaceUnder">
                        <td width='40%' :style="{textAlign: 'center'}">
                            description
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <textarea class='textArea1' v-model="profitDescription" rows="4" cols="50" type='text' :style="{width: '99%', height: '20vh', resize: 'none'}"/> 
                        </td>
                    </tr>
                    <tr class="spaceUnder" :style="{width: '80%', paddingBottom: '10px'}">
                        <td width='30%' :style="{textAlign: 'center'}">
                            type
                        </td>
                        <td width='70%' :style="{textAlign: 'center'}">
                            <Dropdown
                            :style="{width: '100%'}"
                            v-on:selectReturn='selectReturn'
                            :dropDataArray="profitType"/>
                        </td>
                    </tr>
                </table>
                <table :style="{width: '100%', maxWidth: '100%', borderCollapse: 'collapse', borderSpacing: '0 50px', tableLayout: 'fixed'}">
                    <tr v-if="selectValues.profitType==='one time'">
                        <td width='40%' :style="{textAlign: 'center'}">
                            profit amount
                        </td>
                        <td width='60%' :style="{textAlign: 'center'}">
                            <input class='input1' type='text' v-model="profitAmount" :style="{width: '99%'}"> 
                        </td>  
                    </tr>
                </table>
                <div :style="{position: 'absolute', right: '1vw', bottom: '2vh'}">
                    <div :style="{textAlign: 'right'}">
                        <input readonly class='button1' value="submit" @click="submitHandler"/>
                    </div>
                </div>
            </div>
            <div class='meMoneyTimeBox card1' :style="{width: '92.5%', marginLeft: '2.5%'}">
                <div v-if="selectValues.profitType==='one time'" :style="{height: '100%', width: '100%', textAlign: 'center'}">
                    <DatePicker
                    :style="{width: '100%'}"
                    :pickerType="'start'"
                    v-on:dateReturn='dateReturn'/>
                    <div>
                        Select Date
                    </div>
                </div>
                <div v-else :style="{textAlign: 'center', marginTop: '10vh'}">
                    <h3>
                        Choose your money options in order to populate the calendar!
                    </h3>
                </div>
            </div>
            <div class='meMoneyGraphBox paper1' :style="{width: '92.5%', marginLeft: '2.5%',
            fontSize: '1.7rem', lineHeight: '1.9rem', position: 'relative'}">
                <div v-if="selectValues.profitType==='one time'">
                    <div :style="{position: 'absolute', top: '5vh', left: '5vh'}">
                         <p>
                            <span :style="{fontWeight: '900'}">One time</span> profit of
                        </p>
                        <p v-if="profitAmount!=null">
                            <span :style="{fontWeight: '900'}">${{profitAmount}}</span> on
                        </p>
                        <p v-if="dateVals.start!=null && profitAmount!=null">
                            <span :style="{fontWeight: '900'}">{{dateVals.month.start}}/{{dateVals.date.start}}/{{dateVals.year.start}}</span> 
                        </p>
                    </div>
                </div>
                <div v-else :style="{textAlign: 'center', marginTop: '10vh', fontWeight: '900'}">
                    <h3>
                        Choose your money options in order to populate the statistics graphs!
                    </h3>
                </div>
            </div>
        </div>
    </div>

    <!-- END profit user flow -->

  </div>
</template>

<script>
import Dropdown from '../../sharedComponents/composedElements/Dropdown';
import DatePicker from '../../sharedComponents/composedElements/DatePicker';
import Worker from 'worker-loader!./../../sharedComponents/equations/recurringTimeComposer.js';

import newInterestTimeSeries from '../../sharedComponents/equations/newInterestTimeSeries';

// import recurringTimeComposer from 
// '../../sharedComponents/equations/recurringTimeComposer';

import Chart from 'chart.js';
import newLoanChart from '../../sharedComponents/chartJS/newLoanChart';
import newRecurringExpenseChart from '../../sharedComponents/chartJS/newRecurringExpenseChart';
import newLoanPie from '../../sharedComponents/chartJS/newLoanPie';
import recurringTimeComposer from '../../sharedComponents/equations/recurringTimeComposer';

import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'MeMoney',
  components:{
    Dropdown,
    DatePicker
  },
  data(){
    return {
        expenseName: '',
        expenseDescription: '', 
        expenseOwedTo: '', 
        profitName: '',
        profitDescription: '', 
        profitFrom: '', 
        moneyToggled: {
            addExpense: false,
            addProfit: false, 
            viewData: true,
            settings: false,
        },
        interestToggled: 'oneMonth',
        expenseType: [
            'expenseType',
            'one time',
            'recurring',
            'loan'
        ], 
        loanType: [
            'loanType',
            'new loan',
            'recurring payment',
            'one time payment'
        ],
        profitType: [
            'profitType',
            'one time',
            'recurring',
            'investment'
        ],
        selectArrayNames: {
            select1: 'expenseType',
            select2: 'loanType'
        },
        selectValues:{
            profitType: 'please choose',
            interestType: 'please choose',
            expenseType: 'please choose',
            loanType: 'please choose',
        }, 
        dateVals: {
            // requestTime: null,
            start: null, 
            end: null,
            date: {
                start: null, 
                end: null
            },
            month: {
                start: null, 
                end: null
            }, 
            year: {
                start: null, 
                end: null
            }
        }, 
        recurrenceWorker: null,
        interestSlider: 0, 
        loanAmount: null, 
        paymentAmount: null, 
        profitAmount: null,
        recurringCheck:{
            payment: 'weekly'
        }, 
        recurrenceArray: null, 
        pickerFinished: {
            start: null, 
            end: null
        }, 
        newLoanChartData: null, 
        newLoanPieData: null,
        newRecurringExpenseChartData: null,
        showRecurringExpenseChart: false
    }
  }, 
  watch: {
    returnRecurrenceObj: function(newVal, oldVal){
        // console.log("328938928932892382937 inside returnRecurrenceObj and value of ")
        // console.log("328938928932892382937  arg.interestSlider!=null && arg.interestSlider!=0 && arg.loanAmount!=null && arg.loanAmount!=0 && arg.dateVals.start!=null && arg.dateVals.end!=null")
        // console.log('328938928932892382937: ',  newVal.interestSlider!=null && newVal.interestSlider!=0 &&
        //             newVal.paymentAmount!=null && newVal.paymentAmount!=0 &&
        //             newVal.dateVals.start!=null &&
        //             newVal.dateVals.end!=null)
        // console.log('328938928932892382937 newVal.interestSlider!=null', newVal.interestSlider!=null)
        // console.log('328938928932892382937 newVal.interestSlider!=0 ', newVal.interestSlider!=0 )
        // console.log('328938928932892382937 newVal.paymentAmount!=null', newVal.paymentAmount!=null)
        // console.log('328938928932892382937 newVal.paymentAmount!=0 ',  newVal.paymentAmount!=0 )
        // console.log('328938928932892382937 newVal.dateVals.start!=null', newVal.dateVals.start!=null)
        // console.log('328938928932892382937 newVal.dateVals.end!=null', newVal.dateVals.end!=null)
        
        this.$nextTick(()=>{
            this.$worker.run((arg)=>{
                console.log("3289389289328923829375 (arg.paymentAmount != null && arg.startDate != null && arg.endDate != null) ", (arg.paymentAmount != null && arg.startDate != null && arg.endDate != null));
                console.log("3289389289328923829375  (arg.loanAmount != null && arg.interestSlider!= 0 && arg.startDate != null && arg.endDate != null) ",  (arg.loanAmount != null && arg.interestSlider!= 0 && arg.startDate != null && arg.endDate != null));
                if (    
                    (arg.paymentAmount != null && arg.startDate != null && arg.endDate != null) ||
                    (arg.loanAmount != null && arg.interestSlider!= 0 && arg.startDate != null && arg.endDate != null)
                ){
                    if(arg.recurring==='weekly'){
                        console.log('38923982398328932982398 inside arg.recurring === weekly')
                        var returnTime = [];
                        var loopWeekly = true;
                        var startDate = arg.startDate;
                        var endDate = arg.endDate;
                        var startDay = startDate.getDay();
                        var loopDate = new Date(startDate.getTime());
                        do {
                            console.log('38923982398328932982398 value of loopDate; ', loopDate);
                            if(loopDate.getDay()===startDay){
                                console.log('38923982398328932982398 loopDate.getDay(): ', loopDate.getDay());
                                loopDate.setDate(loopDate.getDate()+5);
                                returnTime.push(loopDate.toString());
                            }else{
                                loopDate.setDate(loopDate.getDate()+1);
                            }
                            console.log('38923982398328932982398 before if');
                            console.log('38923982398328932982398 loopDate: ', loopDate);
                            console.log('38923982398328932982398 value of arg.endDate: ', arg.endDate);
                            if(Date.parse(loopDate)>Date.parse(arg.endDate)){
                                loopWeekly = false;
                                console.log('value of returnTime: ', returnTime);
                            }
                        }while(loopWeekly === true);
                        return ['weekly', returnTime, arg.paymentAmount];
                    }else if(arg.recurring==='monthly'){
                        console.log('38923982398328932982398 inside arg.recurring === monthly')
                        var returnTime = [];
                        var startDate = arg.startDate;
                        var loopDate = new Date(startDate.getTime());
                        var loopMonthly = true;
                        do{
                            if(loopDate.getDate()===startDate.getDate()){
                                returnTime.push(loopDate.toString());
                                loopDate.setDate(loopDate.getDate()+28);
                            }else{
                                loopDate.setDate(loopDate.getDate()+1);
                            }
                            if(Date.parse(loopDate)>Date.parse(arg.endDate)){
                                loopMonthly = false;
                                console.log('value of returnTime: ', returnTime);
                            }
                        }while(loopMonthly === true);
                        return ['monthly', returnTime, arg.paymentAmount];
                    }else if(arg.recurring==='yearly'){
                        console.log('38923982398328932982398 inside arg.recurring === yearly')
                        var returnTime = [];
                        var startDate = arg.startDate;
                        var loopDate = new Date(startDate.getTime());
                        var loopYearly = true;
                        do{
                            if(
                                loopDate.getDate() === startDate.getDate() &&
                                loopDate.getMonth() === startDate.getMonth()
                            ){
                                returnTime.push(loopDate.toString());
                                loopDate.setDate(loopDate.getDate() + 360)
                            }else{
                                loopDate.setDate(loopDate.getDate() + 1);
                            }
                            if(Date.parse(loopDate)>Date.parse(arg.endDate)){
                                loopYearly = false;
                                console.log('value of returnTime: ', returnTime);
                            }
                        }while(loopYearly===true);
                        return ['yearly', returnTime, arg.paymentAmount];
                    }
                }else{
                    return 'some fields empty';
                }
            }, [newVal])
            .then(result => {
                console.log('3289389289328923829375 value of result: ', result);
                if (result != 'some fields empty'){
                    this.showRecurringExpenseChart = true;
                    this.newRecurringExpenseChartData = newRecurringExpenseChart(result);
                    if(window.chart.recurringExpenseChart && window.chart.recurringExpenseChart !== null){
                        window.chart.recurringExpenseChart.destroy();
                    }
                    this.createChart("recurringExpenseChart", this.newRecurringExpenseChartData);
                }else{
                    this.showRecurringExpenseChart = false;
                }
            })
            .catch(e => {
                console.error(e)
            })
        })
    },
    newLoanReturnObj: function (newVal, oldVal) {
        var isAnyEmpty = false;
        Object.keys(newVal).some((key, index) => {
            if(newVal[key]==="" || newVal[key]===null || newVal[key]===0){
                console.log('7378373737373737 inside isAnyEmpty true')
                isAnyEmpty = true;
            }
            return isAnyEmpty === true
        })
        if (isAnyEmpty===false && newVal!=null){
            console.log("&&&&%%%%^^^^^%%%%% newLoanReturnObj is completed!");
            console.log("&&&&%%%%^^^^^%%%%% value of newLoanReturnObj: ", newVal);
            
            var payload = {
                amount: Number(newVal.amount),
                compounded: newVal.compounded,
                interest: Number(newVal.interest),
                startDate: newVal.startDate,
                endDate: newVal.endDate,
            }
            
            var newInterestTimeSeriesReturn = newInterestTimeSeries(payload);
            
            var chartPayload = {
                line: newInterestTimeSeriesReturn,
                bar: {
                    labels: newVal.startDate,
                    data: Number(newVal.amount)
                }
            }

            var timeSeriesReturnData = newInterestTimeSeriesReturn.data;
            // console.log("575757578348948459067 value of timeSeriesReturnData: ", timeSeriesReturnData);

            var timeSeriesReturnLastValue = timeSeriesReturnData[timeSeriesReturnData.length - 1];
            // console.log("575757578348948459067 timeSeriesReturnLastValue: ", timeSeriesReturnLastValue);

            // console.log('575757578348948459067 value of Number(newVal.amount): ', Number(newVal.amount));

            var interestDif = timeSeriesReturnLastValue - newVal.amount;
            // console.log("575757578348948459067 value of interestDif: ", interestDif);

            var piePayload = {
                principle: Number(newVal.amount), 
                interestAmount:  interestDif
            }

            console.log('value of chartPayload: ', chartPayload);
            
            this.newLoanChartData = newLoanChart(chartPayload);
            if(window.chart.loanChart && window.chart.loanChart !== null){
                window.chart.loanChart.destroy();
            }
            this.createChart("loanChart", this.newLoanChartData);

            this.newLoanPieData = newLoanPie(piePayload);
            if(window.chart.loanPie && window.chart.loanPie !== null){
                window.chart.loanPie.destroy();
            }
            // if(!isNaN(newLoanPieData.interestAmount)){
            this.createChart("loanPie", this.newLoanPieData);
            // }
        }
    }
  },
  mounted: function(){
      window.chart = [];
    //   this.recurrenceWorker = new Worker('../../sharedComponents/equations/recurringTimeComposer.js');
  },
  methods: {
    ...mapActions([
        'Request',
    ]),
    submitHandler: function(){
        
        console.log("392813974692 inside submitHandler");

        const findType = () => {
            return new Promise(resolve => {
                var payload = {};
                var urlKEY = 'moneySubmit';
                if (this.selectValues.profitType==='one time'){
                    payload.type = 'oneTime';
                    payload.profit = Number(this.profitAmount);
                    payload.startDate = this.dateVals.start;
                    payload.localJWT = this.localJWT;
                    payload.dateArray = ["NA"]
                    payload.email = this.email;
                    payload.name = this.profitName;
                    payload.description = this.profitDescription;
                    payload.profitFrom = this.profitFrom;
                    urlKEY = urlKEY + '_profit';
                    console.log('392813974692 value of payload in findType: ', payload);
                    resolve({payload: payload, urlKEY: urlKEY})
                }else if (this.selectValues.expenseType==='one time'){
                    payload.type = "oneTime";
                    payload.expense = Number(this.paymentAmount);
                    payload.startDate = this.dateVals.start;
                    payload.localJWT = this.localJWT;
                    payload.dateArray = ["NA"];
                    payload.email = this.email;
                    payload.name = this.expenseName;
                    payload.description = this.profitDescription;
                    payload.expenseOwedTo = this.expenseOwedTo;
                    urlKEY = urlKEY + '_expense_oneTime';
                    resolve({payload: payload, urlKEY: urlKEY});
                }else if(this.selectValues.expenseType==="recurring"){
                    payload.type = "recurring",
                    payload.expense = Number(this.paymentAmount);
                    payload.startDate = this.dateVals.start;
                    payload.endDate = this.dateVals.end;
                    payload.localJWT = this.localJWT;
                    payload.dateArray = ["NA"];
                    payload.email = this.email;
                    payload.name = this.profitName;
                    payload.description = this.expenseDescription;
                    payload.recurring = this.recurringCheck.payment;
                    urlKEY = urlKEY + '_expense_recurring';
                    resolve({payload: payload, urlKEY: urlKEY});
                }
            })
        }

        const payloadChecker = (payload) => {
            return new Promise(resolve => {
                var payloadOK = true;
                Object.keys(payload).some((key, index)=>{
                    console.log('392813974692 value of payload[key]: ', payload[key]);
                    if (payload[key] === null){
                        console.log("392813974692 ERROR: null payload in submit handler.");
                        console.log("392813974692 ERROR: value in payload that is null : ", payload[key]);
                        payloadOK = false;
                    }
                });
                resolve(payloadOK);
            })
        }

        const requestHook = (payloadOK, urlKEY, payload) => {
            if (payloadOK===true){
                console.log('392813974692 inside requestHook right before send and values: ');
                console.log('392813974692 urlKEY: ', urlKEY);
                console.log('392813974692 payload: ', payload);
                this.Request({urlKEY: urlKEY, requestType:'post', payload: payload})
            }
        }

        const asyncCall = async () => {
            var findTypeRetObj = await findType();
            var payloadOK = await payloadChecker(findTypeRetObj.payload);
            requestHook(payloadOK, findTypeRetObj.urlKEY, findTypeRetObj.payload);
        }
        asyncCall();

    },
    createChart(chartID, chartData){
      const ctx = document.getElementById(chartID).getContext('2d');
      window.chart[chartID] = new Chart(ctx, {
        type: chartData.type,
        data: chartData.data,
        options: chartData.options,
      });
    },
    changeNav(navType){
      for (var key in this.moneyToggled) {
        console.log('inside changeNav and value of navType: ', navType);
        console.log('value of key:', key)
        if (key === navType){
          this.moneyToggled[key] = true
        }else{
          this.moneyToggled[key] = false
        }
      }
      if(this.moneyToggled.addExpense === true){
          this.selectValues = {
            expenseType: 'please choose',
            loanType: 'please choose',
          }
      }
    },
    selectReturn: function(selectValReturn){
        console.log('328938928932892382937 inside selectReturn')
        this.selectValues[selectValReturn.name] = selectValReturn.value;
        this.interestSlider = 0;
        this.loanAmount = null;
        this.profitAmount = null;
        this.paymentAmount = null;
        this.dateVals.start = null;
        this.dateVals.end = null;
        this.$forceUpdate();
    }, 
    pickerFinishedFunc: function(pickerType){
        this.pickerFinished[pickerType] = Date.now();
    },
    dateReturn: function(dateValReturn){

        console.log('893893484939848973 inside dateReturn and value of dateValReturn: ', dateValReturn);

        if (dateValReturn.value!=""){

            var dateMonth = dateValReturn.value.getMonth();
            this.dateVals.month[dateValReturn.type] = dateMonth;

            var dateDate = dateValReturn.value.getDate();
            this.dateVals.date[dateValReturn.type] = dateDate;

            var dateYear = dateValReturn.value.getFullYear();
            this.dateVals.year[dateValReturn.type] = dateYear;

            setTimeout(()=>{
                this.dateVals[dateValReturn.type] = dateValReturn.value;
            }, 100)
           
        }else{
            this.dateVals.month[dateValReturn.type] = null;
            this.dateVals.date[dateValReturn.type] = null;
            this.dateVals.year[dateValReturn.type] = null;
            setTimeout(()=>{
                this.dateVals[dateValReturn.type] = null;
            }, 100)
        }
        // this.dateVals.requestTime = Date.now();
    }, 
    // recurrenceReturn: function(recurrenceArray){
    //     this.recurrenceArray = recurrenceArray;
    //     console.log('372019478202383 value of this.recurrenceArray: ', this.recurrenceArray);
    // }
    emitRecurrence: function(emittedDates){
        console.log('inside emitRecurrence and value of emmittedDates: ', emittedDates);
    }
  }, 
  computed:{
    ...mapGetters([
     'localJWT', 
     'email'
    ]),
    mainScreenTitle: function(){
        if(this.moneyToggled.addExpense === true){
            return("Add an Expense");
        }else if(this.moneyToggled.addProfit === true){
            return("Add a Profit");
        }else if(this.moneyToggled.viewData === true){
            return("View Data Screen")
        }else{
            return("Settings Screen");
        }
    }, 
    interestSliderReturn: function(){
        return this.interestSlider;
    },
    returnRecurrenceObj: function(){
        return {
            paymentAmount: this.paymentAmount,
            recurring: this.recurringCheck.payment,
            startDate: this.dateVals.start,
            endDate: this.dateVals.end
        }
    },
    newLoanReturnObj: function(){
        var compoundedNum = 0;
        if (this.interestToggled === 'oneMonth'){
            compoundedNum = 12;
        }else if (this.interestToggled === 'threeMonths'){
            compoundedNum = 3;
        }else if (this.interestToggled === 'sixMonths'){
            compoundedNum = 6;
        }else if (this.interestToggled === 'oneYear'){
            compoundedNum = 1;
        }

        return {
            compounded: compoundedNum,
            interest: this.interestSlider,
            amount: this.loanAmount, 
            startDate: this.dateVals.start,
            endDate: this.dateVals.end
        }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@import '../../../assets/styles/screens/me.css';
</style>
