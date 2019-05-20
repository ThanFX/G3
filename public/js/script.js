'use strict';
var HOST = 'http://localhost:8880';
var GET_PERSONS_URL = HOST + '/api/persons';
var GET_DATE_URL = HOST + '/api/date';
function onError(response) {
    console.log("Ошибка!!!/n" + response);
}
function onSuccessPersons(response) {
    var personList = new Vue({
       el: '#personList',
       data: {
           persons: response.result.items
       }
    });
}
function onSuccessDate(curDate) {
    var currentDate = new Vue({
        el: '#currentDate',
        data: {
            currentDate: curDate.result
        }
    });
    console.log(currentDate.data);
}
window.load(GET_PERSONS_URL, onSuccessPersons, onError);
window.load(GET_DATE_URL, onSuccessDate, onError);