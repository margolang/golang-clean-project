-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "public".Application
(
    id          UUID DEFAULT "public".uuid_generate_v4() NOT NULL PRIMARY KEY,
    first_name  VARCHAR(255)                             NOT NULL,
    last_name   VARCHAR(255)                             NOT NULL,
    phone       VARCHAR(50)                              NOT NULL,
    age         INTEGER                                  NOT NULL,
    car_brand   VARCHAR(255)                             NOT NULL,
    car_model   VARCHAR(255)                             NOT NULL,
    car_power   VARCHAR(255)                             NOT NULL,
    is_electric BOOLEAN                                  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public".Application;
-- +goose StatementEnd
