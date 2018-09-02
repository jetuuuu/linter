var a = 10;
var b = 20;
var c = 30;

if (a < b && c === 30) {
    Sum(a, b , c);
}

c = undefined;

if (!c) {
    Sum(a, b);
}