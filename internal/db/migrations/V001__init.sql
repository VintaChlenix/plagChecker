CREATE TABLE metadata(
    name text,
    lab_id integer,
    variant integer,
    norm_code text,
    sum text,
    PRIMARY KEY(name, lab_id, variant)
)
