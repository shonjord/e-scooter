CREATE TABLE trips (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    uuid CHAR(36) NOT NULL,
    scooter_uuid CHAR(36) NOT NULL,
    mobile_uuid CHAR(36) NOT NULL,
    status ENUM('in_progress', 'finished') NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    UNIQUE INDEX IDX_UUID (uuid),
    PRIMARY KEY (id),
    CONSTRAINT FK_SCOOTER FOREIGN KEY (`scooter_uuid`) REFERENCES scooters (uuid),
    CONSTRAINT FK_MOBILES FOREIGN KEY (`mobile_uuid`) REFERENCES mobiles (uuid)
) DEFAULT CHARACTER SET UTF8 COLLATE `UTF8_unicode_ci` ENGINE = InnoDB;
