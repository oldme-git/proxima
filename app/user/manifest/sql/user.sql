CREATE TABLE `user` (
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `passport`  varchar(45) NOT NULL COMMENT 'User Passport',
    `password`  varchar(45) NOT NULL COMMENT 'User Password',
    `Email`     varchar(45) NOT NULL COMMENT 'User Email Address',
    `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
