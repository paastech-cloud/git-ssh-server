// node http server that returns hello world

// import http module
var http = require('http');

// create http server
http.createServer(function (req, res) {
    // write response header
    res.writeHead(200, {'Content-Type': 'text/html'});
    // write response body
    res.end('<h1>Hello World<h1>');
    }
).listen(8080, '0.0.0.0');
