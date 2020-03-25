

console.log("Hello, WOrld")


function showTime()
{
    var today = new Date();
    alert("The time is: " + today.toString());
    setTimeout("showTime()", 5000);
}

showTime()
