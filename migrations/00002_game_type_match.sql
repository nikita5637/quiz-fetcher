-- +goose Up

DROP TABLE IF EXISTS `game_type_match`;
CREATE TABLE IF NOT EXISTS `game_type_match` (
  `id` int(11) NOT NULL,
  `description` varchar(255) NOT NULL,
  `game_type` tinyint(11) unsigned NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `game_type_match`
 ADD PRIMARY KEY (`id`), ADD UNIQUE KEY `description` (`description`);

ALTER TABLE `game_type_match`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

-- +goose Down

DROP TABLE IF EXISTS `game_type_match`;
