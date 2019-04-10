INSERT INTO users (id, email, first_name, last_name)
VALUES ('C83F5B51-1031-45F3-ADF9-2CDA751BF8D1','chance@fake.com', 'Chance', 'Dean');
INSERT INTO users (id, email, first_name, last_name)
VALUES ('449C8C03-211D-498D-9666-C9208BD1F393','emma@fake.com', 'Emma', 'Renee');


INSERT INTO games (id, title)
VALUES ('1FC5D33E-02E1-419F-AE0B-4B98AE79F1B9','First');
INSERT INTO games (id, title)
VALUES ('8066938D-80C8-42E7-BEFD-2312128830E7','Second');
INSERT INTO games (id, title)
VALUES ('1C630FB1-F581-4454-B599-38E4178DFBC2','Third');

INSERT INTO messages (message_text, created_at, user_id, game_id)
VALUES ('hey there', '2018-12-19 22:35:06 -6:00', 'C83F5B51-1031-45F3-ADF9-2CDA751BF8D1', '1FC5D33E-02E1-419F-AE0B-4B98AE79F1B9');
INSERT INTO messages (message_text, created_at, user_id, game_id)
VALUES ('well hello', '2018-12-19 22:35:15 -6:00', '449C8C03-211D-498D-9666-C9208BD1F393', '1FC5D33E-02E1-419F-AE0B-4B98AE79F1B9');
