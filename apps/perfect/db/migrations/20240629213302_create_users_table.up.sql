CREATE TABLE users
(
    id          uuid   DEFAULT gen_random_uuid() PRIMARY KEY,
    name        text NOT NULL,
    age         smallint,
    create_time bigint DEFAULT extract(epoch from now())
)
