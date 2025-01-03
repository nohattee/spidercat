CREATE TABLE IF NOT EXISTS item (
    id TEXT PRIMARY KEY,
    external_id TEXT,
    name TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),

    CONSTRAINT item__external_id__unique UNIQUE (external_id)
);

CREATE TABLE IF NOT EXISTS category (
    id TEXT PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),

    CONSTRAINT category__name__unique UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS tag (
    id TEXT PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),

    CONSTRAINT tag__name__unique UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS author (
    id TEXT PRIMARY KEY,
    name TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    
    CONSTRAINT author__name__unique UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS item_author (
    item_id TEXT,
    author_id TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),

    PRIMARY KEY (item_id, author_id)
);

CREATE TABLE IF NOT EXISTS item_category (
    item_id TEXT,
    category_id TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),

    PRIMARY KEY (item_id, category_id)
);

CREATE TABLE IF NOT EXISTS item_tag (
    item_id TEXT,
    tag_id TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT TIMEZONE('UTC'::TEXT, NOW()),

    PRIMARY KEY (item_id, tag_id)
);
