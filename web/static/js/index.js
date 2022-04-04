$(document).ready(function() {
    $(".task").dblclick(function() {
        if (!this.querySelector('input')){
            this.innerHTML = `
            <form action="/update" method="post">
                <input hidden name="id" value="${this.id}">
                <input type="text" name="text" value="${this.innerText}">
                <input type="submit" hidden>
            </form>`
        };
    });
});