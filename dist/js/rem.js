var html = document.querySelector("html");
//获取当前屏幕宽度的十分之一
console.log(document.documentElement.clientWidth);
var remSize;
var dWidth = document.documentElement.clientWidth;
    // dWidth/remSize = 1920/100;
    remSize = dWidth*100/1920;
// var remSize = document.documentElement.clientWidth / 10;

html.style.fontSize = remSize + "px";
// //当屏幕宽度改变的时候，要将html标签的font-size属性进行更改，改成当前屏幕的宽度的十分之一
window.addEventListener("resize", function () {
    var html = document.querySelector("html");
    // console.log(html);
    //获取当前屏幕宽度的十分之一
    console.log(document.documentElement.clientWidth);
    // var remSize = document.documentElement.clientWidth *(1920/ 100);
    var dWidth = document.documentElement.clientWidth;
    var remSize = dWidth*100/1920;
    // console.log(remSize);
    html.style.fontSize = remSize + "px";
})
// console.log(window);
