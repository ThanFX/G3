var connection = new WebSocket('ws://localhost:8880/ws/events/');
connection.onopen = function () {
    console.log('Connection open!');
    connection.send(JSON.stringify({
        "event": "ping",
        "data": "ping"
    }));
};
connection.onclose = function () {
    console.log('Connection closed');
};
connection.onerror = function (error) {
    console.log('Error detected: ' + error);
};
connection.onmessage = function (e) {
    //var t = JSON.parse(e.data);
    console.log(e.data);
};