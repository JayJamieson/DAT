<?php

namespace App\Utils;

use Iterator;

class Arrays
{


  public static function groupBy(Iterator &$elements, callable $keyFunction): array
  {
    $map = [];

    foreach ($elements as $element) {
      $key = call_user_func($keyFunction, $element);
      $map[$key][] = $element;
    }
    return $map;
  }

  public static function groupByWithRef(Iterator &$elements, callable $keyFunction, callable $extractFunc): array
  {
    $map = [];

    foreach ($elements as $element) {
      $key = call_user_func($keyFunction, $element);
      $map[$key][] = call_user_func($extractFunc, $element);
    }
    return $map;
  }
}
