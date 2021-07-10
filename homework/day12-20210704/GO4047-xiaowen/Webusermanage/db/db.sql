SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mycmdb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mycmdb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mycmdb` DEFAULT CHARACTER SET utf8mb4 ;
USE `mycmdb` ;

-- -----------------------------------------------------
-- Table `mycmdb`.`role`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mycmdb`.`role` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL DEFAULT '\"\"',
  PRIMARY KEY (`id`))
ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4;


-- -----------------------------------------------------
-- Table `mycmdb`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mycmdb`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL DEFAULT '\"\"',
  `password` VARCHAR(512) NOT NULL DEFAULT '\"\"',
  `phone` VARCHAR(45) NOT NULL DEFAULT '\"\"',
  `email` VARCHAR(45) NOT NULL DEFAULT '\"\"',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `role_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_name` (`name` ASC),
  INDEX `fk_user_role1_idx` (`role_id` ASC),
  CONSTRAINT `fk_user_role1`
    FOREIGN KEY (`role_id`)
    REFERENCES `mycmdb`.`role` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = '			\n\n\n\n\n' DEFAULT CHARACTER SET utf8mb4;


-- -----------------------------------------------------
-- Table `mycmdb`.`department`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mycmdb`.`department` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL DEFAULT '\"\"',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4;


-- -----------------------------------------------------
-- Table `mycmdb`.`asset`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mycmdb`.`asset` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ip` VARCHAR(45) NOT NULL DEFAULT '\"\"' COMMENT '资产ip',
  `addr` VARCHAR(45) NOT NULL DEFAULT '\"\"' COMMENT '资产位置',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4;


-- -----------------------------------------------------
-- Table `mycmdb`.`rel_role_per`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mycmdb`.`rel_role_per` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `role_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_rel_role_per_role_idx` (`role_id` ASC),
  CONSTRAINT `fk_rel_role_per_role`
    FOREIGN KEY (`role_id`)
    REFERENCES `mycmdb`.`role` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4;


-- -----------------------------------------------------
-- Table `mycmdb`.`permissions`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mycmdb`.`permissions` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `rel_role_per_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_permissions_rel_role_per1_idx` (`rel_role_per_id` ASC),
  CONSTRAINT `fk_permissions_rel_role_per1`
    FOREIGN KEY (`rel_role_per_id`)
    REFERENCES `mycmdb`.`rel_role_per` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

insert into user(name,password,created_at,updated_at,role_id) values("admin", "$2a$10$G/5TUKA/UyRIqTbFzCmOYOMayc1XE2hrFvIZofZjk0rXKrKB.9CGe",now(), now(), 1);