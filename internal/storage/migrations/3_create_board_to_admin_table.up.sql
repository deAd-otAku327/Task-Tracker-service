CREATE TABLE IF NOT EXISTS board_to_admin
(
    board_id INTEGER,
    admin_id INTEGER
);

ALTER TABLE board_to_admin ADD CONSTRAINT board_to_admin_pk PRIMARY KEY (board_id, admin_id);

ALTER TABLE board_to_admin ADD CONSTRAINT board_to_admin_board_id_fk
FOREIGN KEY (board_id) REFERENCES dashboards(id) ON DELETE CASCADE;

ALTER TABLE board_to_admin ADD CONSTRAINT board_to_admin_admin_id_fk
FOREIGN KEY (admin_id) REFERENCES users(id) ON DELETE CASCADE;