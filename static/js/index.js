function navBarResponse() {
    console.log("func")
    let x = document.getElementById("index-nav-top");
    if (x.className === "nav-page-top") {
        x.className += " responsive";
    } else {
        x.className = "nav-page-top";
    }
};

$('.nav-page-top a').click(navController);

function navController() {
    console.log("cliuck");
    $(".active").toggleClass("active");
    $(this).toggleClass("active");
}


$(document).ready(function() {
    console.log("load");
});


