CREATE TABLE pick_points
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ      DEFAULT now(),
    city       TEXT NOT NULL
);

