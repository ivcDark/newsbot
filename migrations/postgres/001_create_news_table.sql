CREATE TABLE IF NOT EXISTS news (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    subtitle TEXT,
    url TEXT UNIQUE NOT NULL,
    image_url TEXT,
    content TEXT,
    published TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    is_published BOOLEAN NOT NULL DEFAULT FALSE
);