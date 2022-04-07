let redis = require('redis');
let process = require('process');

var rl = require("readline").createInterface({
    input: process.stdin,
    output: process.stdout
  });

if (process.platform === "win32") {
    console.log('Detected windows platform');
    rl.on("SIGINT", function () {
        console.log("shutting down");
        process.emit("SIGINT");
    });
  }

process.on("SIGINT", async function () {
    //graceful shutdown
    console.log("shutting down");
    process.exit();
});

let redisClient = redis.createClient({
    url: 'redis://localhost:8080'
});

async function readData(cancel) {
    while (true) {
        let result = await redisClient.XLEN('mysql.inventory.products');

        let data = await redisClient.XREAD({
            key: 'mysql.inventory.products',
            id: '$'
        }, {
            COUNT: 1,
            BLOCK: 5000
        });

        if (data != null) {
            let key = Object.keys(data[0].messages[0].message)[0];
            let change = data[0].messages[0].message[key];
            let parsed = JSON.parse(change);

            console.log(parsed);
            console.log('size', result);
        }
    }
}


(async () => {
    redisClient.on('error', (err) => console.log('Redis Client Error', err));
    await redisClient.connect();

    readData();
    console.log("Press CTRL + C to exit");
})();
