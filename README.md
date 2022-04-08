# DAT Debezium demo

Really simple setup to play around with debezium. ⚠️ **This sample comes without warranty, if you decide to connect to a production database and it burns to the ground i am not responsible for your decision. You have been warned** ⚠️.

## Postgres

- `/start.sh TRUE postgres postgres` - to start debezium with postgres and redis

### Update aplication.properties

Update the configuration to use postgres as the data source.

```txt
debezium.source.connector.class=io.debezium.connector.postgresql.PostgresConnector
debezium.source.offset.storage.file.filename=data/offsets.dat
debezium.source.offset.flush.interval.ms=0
debezium.source.database.hostname=postgres
debezium.source.database.port=5432
debezium.source.database.user=postgres
debezium.source.database.password=postgres
debezium.source.database.dbname=postgres
debezium.source.database.server.name=postgres
debezium.source.plugin.name=pgoutput
```

## Mysql

- `/start.sh TRUE` to start debezium services (kafka, zookeeper, redis, mysql) and debezium
- `/start.sh` to start just debezium (sometimes need to run this if mysql is not completed startup). See known issues
- `setup.sql` you can change this file to customize the databse scheme to anything you like and seed it with your own desired data.

## Redis

- `docker exec -it redis /bin/bash` to connect to redis container
- `redis-cli` to connect to redis from command line
- `SCAN 0 TYPE stream` show list of streams available to read from, useful to know if debezium has successfully read changes from your database and propogated
them to redis.

## Demo Node Client

- `npm install`
- `node index.js` to run the sample code that will read from the stream for data to get a single result. Runs continuosly untill `CTRL + C`
is pressed.

## insert.json

Sample insert event generated by debezium. Usefull to build your client applications.

- event is specific to mysql only though. might be same for postgres 🤷🏼‍♂️

## Known issues

When running my scripts if your machine is slow, mysql can take a bit longer to startup and run all the setup sql.

- if you encounter `io.debezium.DebeziumException: Unexpected error while connecting to MySQL and looking at BINLOG_FORMAT mode:` just run `docker start debezium` afterwards again and it should start up.
- Permission denied exception on `./data` folder. Unsure why but `chown` the folder/files and set permissions to `777` to fix.
