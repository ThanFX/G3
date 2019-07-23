Vue.component("lake", {
    props: ["size", "fish-num"],
    template: `
        <div class="lake">
            <img class="lake-img" src="lake/lake.png">
            <div class="lake-info">
                <div>{{ size }}</div>
                <div>{{ fishNum }}</div>
            </div>
        </div>
    `
});