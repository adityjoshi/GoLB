const http = require('http');

const servers = [
    { port: 8001, name: 'Server 1' },
    { port: 8002, name: 'Server 2' },
    { port: 8003, name: 'Server 3' },
    { port: 8004, name: 'Server 4' },
    { port: 8005, name: 'Server 5' }
];

// Start the 5 backend servers
servers.forEach(server => {
    http.createServer((req, res) => {
        res.writeHead(200, { 'Content-Type': 'text/plain' });
        res.end(`${server.name} is handling the request!`);
    }).listen(server.port, () => {
        console.log(`${server.name} is running on port ${server.port}`);
    });
});

