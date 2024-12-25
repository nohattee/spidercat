CREATE TABLE IF NOT EXISTS scraped_item (
    id TEXT PRIMARY KEY,
    external_id TEXT NOT NULL,
    title TEXT,
    description TEXT,
    genres TEXT,
    authors TEXT,
    tags TEXT,
    chapters TEXT,
    image_urls TEXT,
    thumbnail_url TEXT,
    source_id TEXT NOT NULL,
    source_item_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),

    CONSTRAINT item__source_id_external_id__unique UNIQUE (source_id, external_id)
);
