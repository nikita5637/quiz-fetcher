-- +goose Up

ALTER TABLE `game_type_match` ADD `name` VARCHAR(64) NULL AFTER `description`, ADD UNIQUE (`name`);
ALTER TABLE `game_type_match` CHANGE `description` `description` VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL;

-- +goose Down

ALTER TABLE `game_type_match` DROP `name`;
ALTER TABLE `game_type_match` CHANGE `description` `description` VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL;
