CREATE TABLE emails
(
    server_id         TEXT    NOT NULL COLLATE NOCASE,
    uid               INTEGER NOT NULL,
    expunged          BOOLEAN NOT NULL,
    mailbox_server_id TEXT    NOT NULL,
    user_id           TEXT    NOT NULL COLLATE NOCASE,
    UNIQUE (user_id, server_id)
);
