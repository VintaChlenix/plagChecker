CREATE TABLE metadata
(
    name      text,
    lab_id    text,
    variant   text,
    norm_code text,
    sum       text,
    tokens    text,
    PRIMARY KEY (name, lab_id, variant)
)
