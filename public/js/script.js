'use strict';
var HOST = 'http://localhost:8880';
var GET_PERSONS_URL = HOST + '/api/persons';
var GET_DATE_URL = HOST + '/api/date';
var NEXT_DATE_URL = HOST + '/api/nextdate';
function onError(response) {
    console.log(response);
}

/*
function onSuccessPersons(response) {
    var personList = new Vue({
       el: '#personList',
       data: {
           persons: response.result.items
       }
    });
}
*/

var currentDate = new Vue({
    el: '#currentDate',
    data: {
        currentDate: ""
    },
    methods: {
        nextDate: function() {
            window.load(NEXT_DATE_URL, getDateFromServer, onError);
        }
    }
});

function onSuccessDate(curDate) {
    currentDate.currentDate = curDate.result;
}
//window.load(GET_PERSONS_URL, onSuccessPersons, onError);

function getDateFromServer() {
    window.load(GET_DATE_URL, onSuccessDate, onError);
}
getDateFromServer();