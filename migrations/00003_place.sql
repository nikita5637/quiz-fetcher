-- +goose Up

CREATE TABLE IF NOT EXISTS `place` (
	`id` int(11) NOT NULL,
	`external_id` int(11) NOT NULL,
	`name` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
	`address` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `place`
	ADD PRIMARY KEY (`id`);

ALTER TABLE `place`
	MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

-- +goose Down

DROP TABLE `place`;
