-- by pricebook id
select id,
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
where price_book_id = 3;

-- by company id
select id,
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
  and deleted_at IS NULL;

-- by company id and pricebook id
select id,
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
and deleted_at IS NULL;

-- create pricebook line item table DDL
create table rp_price_book_line_item
(
    id                int unsigned auto_increment
        primary key,
    company_id        int                                  not null,
    price_book_id     int unsigned                         null,
    old_is_labour     tinyint(1)      default 0            not null,
    product_code      varchar(250)                         null,
    product_name      varchar(500)                         null,
    unit_type         varchar(10)                          null,
    cost_price        decimal(19, 10) default 0.0000000000 not null,
    retail_price      decimal(19, 2)  default 0.00         not null,
    tax_rate_override decimal(19, 10)                      null,
    trade_price       decimal(19, 10)                      null,
    search_values     varchar(500)                         null,
    updated_at        datetime                             null,
    deleted_at        datetime                             null,
    deleted_by        int                                  null,
    is_taxable        tinyint         default 1            null,
    sales_account_id  int unsigned                         null,
    supplier_sku      varchar(30)                          null,
    constraint rp_price_book_line_item_ibfk_1
        foreign key (price_book_id) references rp_price_book (id)
            on update cascade on delete cascade,
    constraint rp_price_book_line_item_ibfk_2
        foreign key (company_id) references gl_company (id)
            on update cascade on delete cascade
);

create index company_id
    on rp_price_book_line_item (company_id);

create index idx_company_id_price_book_id_deleted_at
    on rp_price_book_line_item (company_id, price_book_id, deleted_at);

create index price_book_id
    on rp_price_book_line_item (price_book_id);

create index product_code
    on rp_price_book_line_item (product_code);

create index product_name
    on rp_price_book_line_item (product_name(255));

create index search_values
    on rp_price_book_line_item (search_values(255));

create index unit_type
    on rp_price_book_line_item (unit_type);
