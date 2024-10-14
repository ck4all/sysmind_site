CREATE TABLE master_prompts (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  uuid varchar(36) NOT NULL UNIQUE ,
  nama_prompt varchar(255) NOT NULL , 
  prompt text NOT NULL ,
  img_path varchar(255) NULL ,
  kategori_uuid varchar(36) NULL ,
  aspek_rasio_uuid varchar(36) NULL ,
  dimensi_uuid varchar(36) NULL ,
  aigen_uuid varchar(36) NULL ,
  user_uuid varchar(36) NULL ,
  deleted_at datetime(3) NULL ,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_master_prompts_created_at (created_at),
  KEY idx_master_prompts_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
