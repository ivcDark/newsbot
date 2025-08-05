CREATE TABLE IF NOT EXISTS news (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    subtitle TEXT,
    url TEXT NOT NULL UNIQUE,
    image_url TEXT,
    content TEXT,
    published DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);