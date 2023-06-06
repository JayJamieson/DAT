<?php
require __DIR__ . '/vendor/autoload.php';

use League\Csv\Reader;
use League\Csv\Statement;


$csv = Reader::createFromPath('test.csv', 'r');
$csv->setHeaderOffset(0); //set the CSV header offset

//get 25 records starting from the 11th row
$stmt = Statement::create()
    ->offset(10)
    ->limit(25)
;

$records = $stmt->process($csv);
foreach ($records as $record) {
    //do something here
}
