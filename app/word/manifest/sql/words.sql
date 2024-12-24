CREATE TABLE `words` (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    uid INT UNSIGNED NOT NULL,
    word VARCHAR ( 255 ) NOT NULL,
    definition TEXT,
    example_sentence TEXT,
    chinese_translation VARCHAR ( 255 ),
    pronunciation VARCHAR ( 255 ),
    created_at DATETIME,
    updated_at DATETIME
);