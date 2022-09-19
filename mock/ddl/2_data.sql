USE example;

START TRANSACTION;

INSERT INTO users (id, name) VALUES
                                (1, "Hoge"),
                                (2, "Fuga");


INSERT INTO reviews (id, text, user_id) VALUES
                                               (1 ,"test message id 1", 1),
                                               (2 ,"test message id 2", 1),
                                               (3 ,"test message id 3", 2),
                                               (4 ,"test message id 4", 2);

COMMIT;