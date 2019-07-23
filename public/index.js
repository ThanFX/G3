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
                fetch(`${apiUrl}`)
                .then(stream => stream.json())
                .then(data => {
                    const lakes = [];
                    for (let key in data.bpi) {
                        const lake = data.bpi[key];
                        lakes.push(
                            {
                                fishNum: lake.rate,
                                size: lake.rate_float
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
                    console.log(data);
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
            fetch(`${apiUrl}`)
                .then(stream => stream.json())
                .then(data => {
                    //this.events = data.time;

                    this.events.push('Иоган купил лошадь', 'Мольберт женился', 'Ивана убило');
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
        }
    },
    computed: {
        listWithNewline() {
            return this.events.toString().split(',').join('\n');
        }
    }
});