CREATE TABLE app_infos (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  uuid varchar(36) NOT NULL UNIQUE,
  app_name varchar(100) NOT NULL,
  app_ver varchar(10) NOT NULL ,
  app_desc varchar(255),
  app_logo varchar(250) DEFAULT 'logo.png',
  app_theme varchar(36) DEFAULT 'light',
  app_color varchar(36) DEFAULT 'indigo',
  app_company varchar(255),
  app_slogan varchar(500),
  app_address varchar(255),
  app_website varchar(255),
  app_phone varchar(18),
  app_email varchar(200),
  app_tw varchar(200),
  app_fb varchar(200),
  app_ig varchar(200),
  app_in varchar(200),
  deleted_at datetime(3) NULL ,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_app_infos_created_at (created_at),
  KEY idx_app_infos_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
