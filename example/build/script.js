let counter = 0
function incrementCounter(){
counter = counter + 1
document.getElementById('counter-display').innerText = counter
}
document.getElementById('increment-btn').addEventListener('click',incrementCounter);
