console.log('*** inside the store/actions file ***');

import request from 'request';
import { setSingleVarMutation } from './mutations';
import { P } from 'glamorous';
var requestCookies = request.defaults({jar: true})

const urlSTRING = 'http://localhost:8000';

const urlINDEX = {
    loginUser: urlSTRING+'/loginUser',
    registerUser: urlSTRING+'/registerUser', 
    testJWT: urlSTRING+'/testJWT',
    cookieLogin: urlSTRING + '/cookieLogin', 
    moneySubmit_profit: urlSTRING + '/moneySubmit/profit', 
    moneySubmit_expense_oneTime: urlSTRING + '/moneySubmit/expense/oneTime',
    moneySubmit_expense_recurring: urlSTRING + '/moneySubmit/expense/recurring', 
    newChannel: urlSTRING + '/talk/newChannel',
    getChannels: urlSTRING + '/talk/getChannels',
    newSubChannel: urlSTRING + '/talk/newSubChannel',
    getSubChannels: urlSTRING + '/talk/getSubChannels', 
    newHashMsg: urlSTRING + '/hash/newHashMsg',
    getHashMsg: urlSTRING + '/hash/getHashMsg',
    getUsers: urlSTRING + '/users/getUsers', 
    sendUserMsg: urlSTRING +  '/users/sendMsg', 
    getMail: urlSTRING + '/users/getMail'   
}

window.getCookie = function(name) {
    var match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
    if (match) return match[2];
}

export const setSingleVar = ({commit}, {variableName, variableValue}) => {
    commit('setSingleVarMutation', {variableName, variableValue})
}

export const testAction = ({ commit }, {newString}) => {
    console.log('inside testAction and newString: ', newString);
    commit('testMutation', {newString})
}

export const Request = ({ commit }, {urlKEY, requestType, payload}) => {

    console.log('inside Request action');
    console.log('and value of payload: ', payload);

    var emailCookie = window.getCookie('zennifyEmail')

    if (payload.email!=null){
        console.log('payload.email is not null');
        commit('setSingleVarMutation', {variableName: 'email', variableValue: payload.email});
    }else if(typeof emailCookie !== 'undefined'){
        console.log('email cookie found!');
        commit('setSingleVarMutation', {variableName: 'email', variableValue: emailCookie});
    }

    function callback(error, response, body) {
        try{
            console.log('error: ', error)
            console.log('response from server ', response.headers);
            console.log('body from server :', body);

            var variableArray = []

            if (body.JWT!="" && body.JWT != undefined){
                console.log("inside if statement on JWT return")
                variableArray.push({variableName: 'localJWT', variableValue: body.JWT});
            }

            if ((body.Logintype==="Register"||body.Logintype==="Login") && body.Status==="Success"){
                variableArray.push({variableName: 'loggedIn', variableValue: true})
            }   

            if (body.Message === 'ChannelData'){
                variableArray.push({variableName: 'channelData', variableValue: body.ChannelData})
            }

            if (body.Message === 'SubChannelData'){
                variableArray.push({variableName: 'subChannelData', variableValue: body.SubChannelData})
            }

            if (body.Message === 'profileMsgArray'){
                variableArray.push({variableName: 'profileMsgArray', variableValue: body.Hashes})
            }

            if (body.Message === 'profileCanvas'){
                variableArray.push({variableName: 'profileCanvas', variableValue: body.Hashes})
            }

            if (body.Name === 'UsersList'){
                variableArray.push({variableName: 'usersList', variableValue: body.Value})
            }

            if (body.ReceivedMail!=null&&body.ReceivedMail!=undefined){
                variableArray.push({variableName: 'receivedMail', variableValue: body.ReceivedMail})
            }

            if (body.SentMail!=null&&body.SentMail!=undefined){
                variableArray.push({variableName: 'sentMail', variableValue: body.SentMail})
            }

            commit('setMultiVarMutation', {variableArray})
        }
        catch(err){
            console.log('there was some error in requestC callback: ', err)
        }
    }

    if (requestType==='get'){
        var options = {
            url: urlINDEX[urlKEY],
            json: true,
            jar: cookiejar,
            withCredentials: true,
            headers: {
                'Origin': 'http://localhost:8080'
            }
        }; 
        requestCookies.get(options, callback);
    }

    if (requestType==='post'){
        var cookiejar = request.jar();
        var options = {
            url: urlINDEX[urlKEY],
            body: payload, 
            json: true,
            jar: cookiejar,
            withCredentials: true,
            headers: {
                'Origin': 'http://localhost:8080'
            }
        };
        requestCookies.post(options, callback);
    }
}
