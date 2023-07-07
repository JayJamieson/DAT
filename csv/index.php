<?php
require __DIR__ . '/vendor/autoload.php';

use App\Utils\Arrays;
use League\Csv\Reader;
use League\Csv\Writer;

function convert($size)
{
    $unit = array('b', 'kb', 'mb', 'gb', 'tb', 'pb');
    return @round($size / pow(1024, ($i = floor(log($size, 1024)))), 2) . ' ' . $unit[$i];
}

// const DB_DRIVER = "mysql";
// const DB_HOST = "127.0.0.1";
// const DB_NAME = "fergus_tmp_test";
// const DB_LOGIN = "root";
// const DB_PASS = "olddognewtricks";

const DB_DRIVER = "mysql";
const DB_HOST = "127.0.0.1";
const DB_NAME = "fg_production";
const DB_LOGIN = "fg_jay_ro";
const DB_PASS = "qvidMMz2q6oU@";
const PORT = "3336";

$connectionString = DB_DRIVER . ':host=' . DB_HOST . ';port=' . PORT . ';dbname=' . DB_NAME;

//Connect to database.
$dbh = new PDO($connectionString, DB_LOGIN, DB_PASS, array(
    PDO::ATTR_PERSISTENT => false
));

// We fetch the info from a DB using a PDO object

$sth = $dbh->prepare("select id,
    product_code,
    product_name,
    unit_type,
    cost_price,
    retail_price,
    tax_rate_override,
    trade_price,
    search_values,
    supplier_sku
from rp_price_book_line_item
where company_id = 43764
and price_book_id = 151268
and deleted_at IS NULL");

echo "pre load from db " . convert(memory_get_usage(true)) . "\n";
echo "peak pre process from db " . convert(memory_get_peak_usage(true)) . "\n";

// Because we don't want to duplicate the data for each row
// PDO::FETCH_NUM could also have been used
$sth->setFetchMode(PDO::FETCH_ASSOC);
$sth->execute();

echo "\n";

$tmp = new SplTempFileObject();
// We create the CSV into memory
$db_csv = Writer::createFromFileObject($tmp);

// The PDOStatement Object implements the Traversable Interface
// that's why Writer::insertAll can directly insert
// the data into the CSV
$db_csv->insertAll($sth);

echo "post load from db " . convert(memory_get_usage(true)) . "\n";
echo "post peak pre process from db " . convert(memory_get_peak_usage(true)) . "\n";

echo "\n";

echo "pre load file " . convert(memory_get_usage(true)) . "\n";
echo "peak pre process file " . convert(memory_get_peak_usage(true)) . "\n";

$new_csv = Reader::createFromPath('./data/jarussell_ZTF01NPL_010323.csv', 'r');
$old_csv = Reader::createFromString($db_csv->toString());

echo "post load file " . convert(memory_get_usage(true)) . "\n";
echo "peak post process file " . convert(memory_get_peak_usage(true)) . "\n";

$count = [];
// $data = [];

$records = $old_csv->getRecords();
// foreach ($records as $record) {
//     $data[] = $record;
// }

$groudedData = Arrays::groupByWithRef($records, function ($record) {
    return $record[1] . $record[2] . $record[3];
}, function ($record) {
    return $record[0];
});

echo "post process file " . convert(memory_get_usage(true)) . "\n";
echo "peak post process file " . convert(memory_get_peak_usage(true)) . "\n";

foreach ($groudedData as $key => $val) {
    $cnt = count($val);

    if ($cnt > 1) {
        echo "$key: $cnt" . "\n";
    }
}

// echo count($groudedData)."\n";

$dbh = null;
$sth = null;
