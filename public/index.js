const apiUrl = 'http://localhost:8880';

new Vue({
    el: '#app',
    data() {
        return {
            currentDate: {
                day: null,
                ten: null,
                month: null,
                year: null
            },
            events: [],
            lakes: [],
            fishers: [],
            time: 0,
            maxTime: 10,
            timer: null,
            timerChecked: false
        };
    },
    mounted() {
        this.fetchAll();
    },
    updated() {
        this.scrollTextArea();
    },
    methods: {
        fetchDate() {
            fetch(`${apiUrl}/api/date`)
                .then(stream => stream.json())
                .then(data => {
                    const date = data.result;
                    this.currentDate = {
                        day: date.day,
                        ten: date.ten,
                        month: date.month,
                        year: date.year,
                    };
                })
                .catch(error => console.error(error))
            },
            fetchLakes() {
                fetch(`${apiUrl}/api/lakes`)
                .then(stream => stream.json())
                .then(data => {
                    const lakes = [];
                    for (let key in data.result.items) {
                        const lake = data.result.items[key];
                        let size;
                        switch (lake.Size) {
                            case 1:
                                size = 'Крохотное';
                                break;
                            case 2:
                                size = 'Маленькое';
                                break;
                            case 3:
                                size = 'Среднее';
                                break
                            case 4:
                                size = 'Большое';
                                break
                            case 5:
                                size = 'Огромное';
                                break
                        }
                        lakes.push(
                            {
                                fishNum: lake.Capacity,
                                size: size
                            }
                            );
                        }
                        this.lakes = lakes;
                    })
                    .catch(error => console.error(error))
                },
                fetchFishers() {
                    fetch(`${apiUrl}/api/persons`)
                    .then(stream => stream.json())
                    .then(data => {
                    const fishers = [];
                    for (let key in data.result.items) {
                        const fisher = data.result.items[key];
                        //console.log(fisher)
                        fishers.push(
                            {
                                name: fisher.Name,
                                level: fisher.Skill
                            }
                        );
                    }
                    this.fishers = fishers;
                    //console.log(fishers)
                })
                .catch(error => console.error(error))
        },
        fetchEvents(){
            fetch(`${apiUrl}/api/events`)
                .then(stream => stream.json())
                .then(data => {
                    this.events.push(data.result.items);
                })
                .catch(error => console.error(error))
        },
        fetchAll() {
            this.fetchDate();
            this.fetchEvents();
            this.fetchLakes();
            this.fetchFishers();
        },
        fetchNewDay() {
            fetch(`${apiUrl}/api/nextdate`)
                .then(response => {
                    if(response.ok) {
                        this.resetTime();
                        this.fetchAll();
                    } else {
                        this.time = 0;
                    }
                })
                .catch(error => console.error(error))
        },
        resetTime() {
            this.time = 0;
        },
        setFetchTimeout() {
            this.timer = setTimeout(() => {
                this.time += 1;
                if (this.time >= this.maxTime) {
                    this.fetchNewDay();
                }
                this.timer = setTimeout(this.setFetchTimeout);
            }, 1000);
        },
        editTimer() {
            clearTimeout(this.timer);
            this.timer = null;
            this.time = 0;
        },
        handleCheckbox() {
            if (!this.timerChecked) {
                this.editTimer();
            }
        },
        scrollTextArea() {
            const textarea = document.getElementsByClassName('events')[0];
            textarea.scrollTop = textarea.scrollHeight;
        }
    },
    computed: {
        listWithNewline() {
            return this.events.toString().split(',').join('\n');
        }
    }
});