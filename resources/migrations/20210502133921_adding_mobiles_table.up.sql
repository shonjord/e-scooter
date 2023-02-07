CREATE TABLE mobiles (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    uuid CHAR(36) NOT NULL ,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    INDEX IDX_UUID (uuid),
    UNIQUE KEY `UNIQ_MOBILE_UUID` (`uuid`),
    PRIMARY KEY (id)
) DEFAULT CHARACTER SET UTF8 COLLATE `UTF8_unicode_ci` ENGINE = InnoDB;

INSERT INTO mobiles (uuid, created_at, updated_at)
VALUES
       ('20587b2c-3969-49b6-add1-27fe09006ef9', NOW(), NOW()),
       ('1c5f944a-b9fa-41e1-a83c-8e6c5dea3a82', NOW(), NOW()),
       ('6c6a61d0-7780-42e5-a02e-6d0127a87252', NOW(), NOW())
