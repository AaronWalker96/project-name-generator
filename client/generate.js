var HttpClient = function() {
    this.get = function(aUrl, aCallback) {
        var anHttpRequest = new XMLHttpRequest();
        anHttpRequest.onreadystatechange = function() { 
            if (anHttpRequest.readyState == 4 && anHttpRequest.status == 200)
                aCallback(anHttpRequest.responseText);
        }

        anHttpRequest.open( "GET", aUrl, true );            
        anHttpRequest.send( null );
    }
}


function generate() {

    var client = new HttpClient();
    client.get('http://localhost:8080/api/generate', function(response) {
        document.querySelector('#generated-name').innerText = response.substring(1, response.length-1);
    });
    
}
