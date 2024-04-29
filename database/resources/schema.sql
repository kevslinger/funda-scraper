CREATE TABLE funda_houses (
    time_seen           TIMESTAMP,
    link                TEXT PRIMARY KEY,
    house_id            INT,
    house_address       VARCHAR(256),
    price               INT,
    house_description   TEXT,
    listed_since        TIMESTAMP,
    zip_code            VARCHAR(256),
    built_year          INT,
    total_size          INT,
    living_size         INT,
    house_type          VARCHAR(256),
    building_type       VARCHAR(256),
    num_rooms           INT,
    num_bedrooms        INT,
    layout              VARCHAR(256),
    energy_label        VARCHAR(256),
    insulation          VARCHAR(256),
    heating             VARCHAR(256),
    ownership_type      VARCHAR(256),
    exteriors           VARCHAR(256),
    parking             VARCHAR(256),
    neighbourhood_name  VARCHAR(256),
    date_listed         TIMESTAMP,
    date_sold           TIMESTAMP,
    term                INT,
    price_sold          INT,
    last_ask_price      INT,
    last_ask_price_m2   INT
);
