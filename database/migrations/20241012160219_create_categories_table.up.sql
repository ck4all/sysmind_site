CREATE TABLE categories (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  uuid varchar(36) NOT NULL  UNIQUE ,
  name varchar(255) NOT NULL ,
  slug_name varchar(255) NULL ,
  urutan int8 NULL ,
  user_uuid varchar(36) NULL ,
  deleted_at datetime(3) NULL ,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_categories_created_at (created_at),
  KEY idx_categories_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
