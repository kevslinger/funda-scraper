CREATE TABLE funda_houses (
    time_seen       DATETIME,
    link            VARCHAR(256), -- TODO:
    house_id        INT,
    house_address   VARCHAR(256),
    price           INT,
    house_description VARCHAR(), -- TODO: Very large??
    listed_since    DATETIME,
    zip_code        VARCHAR(6),
    built_year      INT,
    total_size      INT,
    living_size     INT,
    house_type      VARCHAR(256),
    building_type   VARCHAR(256),
    num_rooms       INT,
    num_bedrooms    INT,
    layout          VARCHAR(256),
    energy_label    VARCHAR(8),
    insulation      VARCHAR(256),
    heating         VARCHAR(256),
    ownership_type  VARCHAR(256),
    exteriors       VARCHAR(256),
    parking         VARCHAR(256),
    neighbourhood_name  VARCHAR(256),
    date_listed     DATETIME,
    date_sold       DATETIME,
    term            INT,
    price_sold      INT,
    last_ask_price  INT,
    last_ask_price_m2   INT
);