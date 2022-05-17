CREATE TABLE users (
    username TEXT PRIMARY KEY,
    user_email TEXT,
    password TEXT
);

CREATE TABLE posts (
    post_id BIGSERIAL PRIMARY KEY,
    creator TEXT REFERENCES users(username),
    title TEXT NOT NULL,
    description TEXT,
    image_url TEXT
);

CREATE TABLE comments (
    comment_id BIGSERIAL PRIMARY KEY ,
    text TEXT NOT NULL,
    post BIGINT REFERENCES posts(post_id),
    creator TEXT REFERENCES users(username)
);

CREATE TABLE reactions (
    reactions_id BIGSERIAL PRIMARY KEY ,
    emoticon TEXT,
    creator TEXT REFERENCES users(username)
);

CREATE TABLE post_reactions (
    post_id BIGINT REFERENCES posts(post_id),
    reaction_id BIGINT REFERENCES reactions(reactions_id)
);

CREATE TABLE comment_reactions (
    comment_id BIGINT REFERENCES comments(comment_id),
    reaction_id BIGINT REFERENCES reactions(reactions_id)
);
