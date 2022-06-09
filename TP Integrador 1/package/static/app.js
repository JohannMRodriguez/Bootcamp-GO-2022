document.getElementById("myBtn").addEventListener("click", function() {
    var userInfo = document.getElementById("userInput").value
    localStorage.setItem("userInfo", userInfo);
    window.location = 'user-area.html'
});


document.getElementById("backBtn").addEventListener("click", function() {
    window.location = 'index.html'
});