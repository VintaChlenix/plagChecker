CREATE TABLE metadata
(
    name      text,
    lab_id    text,
    variant   text,
    norm_code text,
    sum       text,
    tokens    text,
    url       text,
    PRIMARY KEY (name, lab_id, variant)
);
CREATE TABLE sendings
(
    id         serial PRIMARY KEY,
    name       text,
    lab_id     text,
    variant    text,
    results    float[],
    url        text,
    source_url text,
);

