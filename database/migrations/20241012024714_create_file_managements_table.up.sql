CREATE TABLE file_managements (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  uuid varchar(36) NOT NULL UNIQUE,
  folder_name varchar(100),
  file_name varchar(100)NOT NULL ,
  ext varchar(8),
  size bigint,
  type varchar(30),
  used int default 0,
  user_uuid varchar(36),
  deleted_at datetime(3),
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_file_managements_created_at (created_at),
  KEY idx_file_managements_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
