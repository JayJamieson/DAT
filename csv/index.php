<?php
require __DIR__ . '/vendor/autoload.php';

use League\Csv\Reader;
use League\Csv\Writer;

function convert($size)
 {
    $unit=array('b','kb','mb','gb','tb','pb');
    return @round($size/pow(1024,($i=floor(log($size,1024)))),2).' '.$unit[$i];
 }

const DB_DRIVER = "mysql";
const DB_HOST = "127.0.0.1";
const DB_NAME = "fergus_tmp_test";
const DB_LOGIN = "root";
const DB_PASS = "olddognewtricks";

$connectionString = DB_DRIVER.':host='.DB_HOST.';dbname='.DB_NAME;

//Connect to database.
$dbh = new PDO($connectionString, DB_LOGIN, DB_PASS);

// We fetch the info from a DB using a PDO object
$sth = $dbh->prepare(
    "select id,product_code, product_code, product_name, unit_type, cost_price, retail_price, tax_rate_override, trade_price, search_values, supplier_sku from rp_price_book_line_item where price_book_id = 3"
);

echo "pre load from db ".convert(memory_get_usage(true))."\n";
echo "peak pre process from db ".convert(memory_get_peak_usage(true))."\n";

// Because we don't want to duplicate the data for each row
// PDO::FETCH_NUM could also have been used
$sth->setFetchMode(PDO::FETCH_ASSOC);
$sth->execute();

echo "\n";

$tmp = new SplTempFileObject();
// We create the CSV into memory
$csv = Writer::createFromFileObject($tmp);

// The PDOStatement Object implements the Traversable Interface
// that's why Writer::insertAll can directly insert
// the data into the CSV
$csv->insertAll($sth);

echo "post load from db ".convert(memory_get_usage(true))."\n";
echo "post peak pre process from db ".convert(memory_get_peak_usage(true))."\n";

echo "\n";

echo "pre load file ".convert(memory_get_usage(true))."\n";
echo "peak pre process file ".convert(memory_get_peak_usage(true))."\n";

$csv = Reader::createFromPath('./data/jarussell_ZTF01NPL_010323.csv', 'r');

echo "post load file ".convert(memory_get_usage(true))."\n";
echo "peak post process file ".convert(memory_get_peak_usage(true))."\n";

$count = [];
$data = [];

$records = $csv->getRecords();
foreach ($records as $record) {
    $price = $record[7];
    $data[] = $record;

    if (!key_exists($price, $count)) {
        $count[$price] = 1;
    } else {
        $count[$price] += 1;
    }
}

echo "post process file ".convert(memory_get_usage(true))."\n";
echo "peak post process file ".convert(memory_get_peak_usage(true))."\n";
