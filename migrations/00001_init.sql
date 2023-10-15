-- +goose Up

DROP TABLE IF EXISTS `sync_log`;
CREATE TABLE IF NOT EXISTS `sync_log` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `last_sync_at` datetime NOT NULL,
  `status` int(11) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `sync_log`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `sync_log`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

-- +goose Down

DROP TABLE IF EXISTS `sync_log`;
