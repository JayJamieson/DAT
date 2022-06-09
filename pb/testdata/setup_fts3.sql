create table pricebook
(
    C1  INTEGER, -- C0,
    C2  INTEGER, -- C1,
    C3  TEXT, -- C2,
    C4  TEXT, -- C3,
    C5  INTEGER, -- C4 ,
    C6  TEXT, -- C5,
    C7  TEXT, -- C6,
    C8  REAL, -- C7,
    C9  TEXT, -- C8,
    C10 TEXT, -- C9,
    C11 TEXT, -- C10,
    C12 TEXT, -- C11,
    C13 TEXT, -- C12,
    C14 REAL -- C13
);

.separator ','
.import ./testdata/jarussell.csv pricebook

create table jarussell (
    id  integer
        constraint pricebook_pk
            primary key autoincrement,
	`product_code` TEXT DEFAULT NULL,
	`product_name` TEXT DEFAULT NULL,
	`unit_type` TEXT  DEFAULT NULL,
	`cost_price` TEXT DEFAULT '0.0000000000' NOT NULL,
	`retail_price` TEXT DEFAULT '0.00' NOT NULL,
	`trade_price` TEXT DEFAULT '0.00' NOT NULL,
	`search_values` TEXT DEFAULT '' NOT NULL,
	`supplier_sku` TEXT NULL
);

CREATE VIRTUAL TABLE jarussell_fts USING fts3(
	`product_code`,
	`product_name` ,
	`unit_type`,
	`cost_price`,
	`retail_price`,
	`trade_price`,
	`search_values`,
	`supplier_sku`,
);


CREATE TRIGGER user_ai
    AFTER INSERT
    ON jarussell
BEGIN
    INSERT INTO jarussell_fts (rowid,
                          `product_code`,
                          `product_name`,
                          `unit_type`,
                          `cost_price`,
                          `retail_price`,
                          `trade_price`,
                          `search_values`,
                          `supplier_sku`)
    VALUES (new.id,
            new.`product_code`,
            new.`product_name`,
            new.`unit_type`,
            new.`cost_price`,
            new.`retail_price`,
            new.`trade_price`,
            new.`search_values`,
            new.`supplier_sku`);
END;

CREATE TRIGGER user_ad
    AFTER DELETE
    ON jarussell
BEGIN
    INSERT INTO jarussell_fts (jarussell_fts,
                          rowid,
                          `product_code`,
                          `product_name`,
                          `unit_type`,
                          `cost_price`,
                          `retail_price`,
                          `trade_price`,
                          `search_values`,
                          `supplier_sku`)
    VALUES ('delete',
            old.id,
            old.`product_code`,
            old.`product_name`,
            old.`unit_type`,
            old.`cost_price`,
            old.`retail_price`,
            old.`trade_price`,
            old.`search_values`,
            old.`supplier_sku`);
END;

CREATE TRIGGER user_au
    AFTER UPDATE
    ON jarussell
BEGIN
    INSERT INTO jarussell_fts (jarussell_fts,
                          rowid,
                          `product_code`,
                          `product_name`,
                          `unit_type`,
                          `cost_price`,
                          `retail_price`,
                          `trade_price`,
                          `search_values`,
                          `supplier_sku`)
    VALUES ('delete',
            old.id,
            old.`product_code`,
            old.`product_name`,
            old.`unit_type`,
            old.`cost_price`,
            old.`retail_price`,
            old.`trade_price`,
            old.`search_values`,
            old.`supplier_sku`);

    INSERT INTO jarussell_fts (rowid,
                          `product_code`,
                          `product_name`,
                          `unit_type`,
                          `cost_price`,
                          `retail_price`,
                          `trade_price`,
                          `search_values`,
                          `supplier_sku`)
    VALUES (new.id,
            new.`product_code`,
            new.`product_name`,
            new.`unit_type`,
            new.`cost_price`,
            new.`retail_price`,
            new.`trade_price`,
            new.`search_values`,
            new.`supplier_sku`);
END;

insert into jarussell (product_code, product_name, unit_type, cost_price, retail_price, trade_price, search_values,
                       supplier_sku) select C2,C3,C4,C14,C8,C8,C11,C1 from pricebook;
