fetch("/api", { headers: { 'accept': 'text/plain' } })
    .then(response => response.body)
    .then(data => {
        document.getElementById("sub").innerHTML = data;
    }).catch(err => {
    msg = "There was a problem getting your IP address"
    document.getElementById("sub").innerHTML = msg
    console.error(err)
})