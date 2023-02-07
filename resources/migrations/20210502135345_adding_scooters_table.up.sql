CREATE TABLE scooters (
     id MEDIUMINT NOT NULL AUTO_INCREMENT,
     uuid CHAR(36) NOT NULL,
     latitude FLOAT NOT NULL,
     longitude FLOAT NOT NULL,
     status ENUM('available', 'occupied') NOT NULL,
     created_at DATETIME NOT NULL,
     updated_at DATETIME NOT NULL,
     UNIQUE KEY `UNIQ_SCOOTER_UUID` (`uuid`),
     INDEX IDX_UUID (uuid),
     PRIMARY KEY (id)
) DEFAULT CHARACTER SET UTF8 COLLATE `UTF8_unicode_ci` ENGINE = InnoDB;


INSERT INTO scooters (uuid, latitude, longitude, status, created_at, updated_at)
VALUES
       ('cd651482-f10e-47d1-9f31-a77fd1fa343d', 31.4225, -122.084, 'available', NOW(), NOW()),
       ('d544915a-346b-4ea8-b2ee-dfc631f00ba5', 31.4225, -126.084, 'available', NOW(), NOW()),
       ('29d2513a-546f-4a41-8d8e-291c921e6ebe', 21.4225, -12.0840, 'available', NOW(), NOW())
