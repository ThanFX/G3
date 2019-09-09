Vue.component("fisher", {
    props: ["name", "level"],
    template: `
        <div class="fisher">
            <img class="fisher-img" src="fisher/fisher.png">
                <div class="fisher-info">
                    <div class="bold">{{ name }}</div>
                    <div>{{ level }}</div>
                </div>
        </div>
    `
});