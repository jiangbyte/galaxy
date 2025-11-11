-- ============================================================ 用户与权限 ============================================================ --
-- ----------------------------
-- Table structure for auth_account (核心账户表)
-- ----------------------------
DROP TABLE IF EXISTS "auth_account";
CREATE TABLE "auth_account"
(
    "id"                   VARCHAR(32) PRIMARY KEY,
    "username"             VARCHAR(64)  NOT NULL,
    "password"             VARCHAR(100) NOT NULL,
    "email"                VARCHAR(128),
    "telephone"            VARCHAR(20),
    "group_id"             VARCHAR(32),               -- 用户组关系到数据权限

    -- 账户安全状态 --
    "status"               SMALLINT    DEFAULT 1,     -- 0:禁用 1:正常 2:锁定 3:冻结
    "is_verified"          BOOLEAN     DEFAULT FALSE, -- 是否已验证
    "is_active"            BOOLEAN     DEFAULT TRUE,  -- 是否激活

    -- 安全相关 --
    "password_strength"    SMALLINT    DEFAULT 0,     -- 密码强度 0-3
    "last_password_change" timestamptz,               -- 最后修改密码时间
    "security_question"    VARCHAR(255),              -- 安全问题
    "security_answer"      VARCHAR(255),              -- 安全答案

    -- 登录与活跃信息 --
    "last_login_time"      timestamptz,               -- 最后登录时间
    "last_login_ip"        VARCHAR(64),               -- 最后登录IP
    "last_active_time"     timestamptz,               -- 最后活跃时间
    "login_count"          INTEGER     DEFAULT 0,     -- 登录次数
    "failed_login_count"   SMALLINT    DEFAULT 0,     -- 连续失败次数
    "lock_until"           timestamptz,               -- 锁定截止时间

    --- 固定字段 ---
    "deleted"              BOOLEAN     DEFAULT FALSE,
    "create_time"          timestamptz DEFAULT NOW(),
    "create_user"          VARCHAR(32),
    "update_time"          timestamptz DEFAULT NOW(),
    "update_user"          VARCHAR(32)
);

CREATE UNIQUE INDEX "idx_account_username" ON "auth_account" ("username");
CREATE UNIQUE INDEX "idx_account_telephone" ON "auth_account" ("telephone");
CREATE UNIQUE INDEX "idx_account_email" ON "auth_account" ("email");
CREATE INDEX "idx_account_status" ON "auth_account" ("status");
CREATE INDEX "idx_account_last_login" ON "auth_account" ("last_login_time");
COMMENT ON TABLE "auth_account" IS '账户认证信息表';

-- ----------------------------
-- Table structure for user_info (用户基本信息表)
-- ----------------------------
DROP TABLE IF EXISTS "user_info";
CREATE TABLE "user_info"
(
    "id"           VARCHAR(32) PRIMARY KEY,
    "account_id"   VARCHAR(32)  NOT NULL,

    -- 基础身份信息 --
    "nickname"     VARCHAR(128) NOT NULL,
    "real_name"    VARCHAR(64),           -- 真实姓名
    "avatar"       VARCHAR(255),          -- 头像
    "gender"       SMALLINT    DEFAULT 0, -- 0:未知, 1:男, 2:女
    "birthday"     DATE,                  -- 生日

    -- 展示信息 --
    "display_name" VARCHAR(128),          -- 展示名称(可不同于昵称)
    "title"        VARCHAR(100),          -- 头衔 "算法大神"
    "signature"    VARCHAR(500),          -- 个性签名
    "quote"        VARCHAR(255),          -- 个人语录

    -- 基础统计(高频查询) --
    "level"        INTEGER     DEFAULT 1, -- 用户等级
    "exp"          BIGINT      DEFAULT 0, -- 经验值

    --- 固定字段 ---
    "deleted"      BOOLEAN     DEFAULT FALSE,
    "create_time"  timestamptz DEFAULT NOW(),
    "create_user"  VARCHAR(32),
    "update_time"  timestamptz DEFAULT NOW(),
    "update_user"  VARCHAR(32)
);

CREATE UNIQUE INDEX "idx_user_info_account" ON "user_info" ("account_id");
CREATE INDEX "idx_user_nickname" ON "user_info" ("nickname");
CREATE INDEX "idx_user_level" ON "user_info" ("level");
COMMENT ON TABLE "user_info" IS '用户基本信息表';

-- ----------------------------
-- Table structure for user_profile (用户档案详情表)
-- ----------------------------
DROP TABLE IF EXISTS "user_profile";
CREATE TABLE "user_profile"
(
    "id"             VARCHAR(32) PRIMARY KEY,
    "user_id"        VARCHAR(32) NOT NULL,

    -- 教育职业信息 --
    "school"         VARCHAR(100),              -- 学校
    "major"          VARCHAR(100),              -- 专业
    "student_id"     VARCHAR(50),               -- 学号
    "company"        VARCHAR(100),              -- 公司
    "job_title"      VARCHAR(100),              -- 职位
    "industry"       VARCHAR(100),              -- 行业

    -- 地理位置 --
    "country"        VARCHAR(50),               -- 国家
    "province"       VARCHAR(50),               -- 省份
    "city"           VARCHAR(50),               -- 城市
    "location"       VARCHAR(100),              -- 详细地址

    -- 个人背景 --
    "background"     VARCHAR(255),              -- 个人背景图
    "bio"            TEXT,                      -- 个人简介
    "interests"      JSONB,                     -- 兴趣标签 ["算法", "AI", "Web开发"]

    -- 社交链接 --
    "website"        VARCHAR(255),              -- 个人网站
    "github"         VARCHAR(100),              -- GitHub
    "blog"           VARCHAR(255),              -- 博客
    "weibo"          VARCHAR(100),              -- 微博

    -- 隐私设置 --
    "privacy_level"  SMALLINT    DEFAULT 1,     -- 隐私级别 1-3
    "show_real_name" BOOLEAN     DEFAULT FALSE, -- 是否显示真实姓名
    "show_birthday"  BOOLEAN     DEFAULT FALSE, -- 是否显示生日
    "show_location"  BOOLEAN     DEFAULT TRUE,  -- 是否显示位置

    --- 固定字段 ---
    "deleted"        BOOLEAN     DEFAULT FALSE,
    "create_time"    timestamptz DEFAULT NOW(),
    "create_user"    VARCHAR(32),
    "update_time"    timestamptz DEFAULT NOW(),
    "update_user"    VARCHAR(32)
);

CREATE UNIQUE INDEX "idx_user_profile_user" ON "user_profile" ("user_id");
COMMENT ON TABLE "user_profile" IS '用户档案详情表';

-- ----------------------------
-- Table structure for user_preference (用户偏好设置表)
-- ----------------------------
DROP TABLE IF EXISTS "user_preference";
CREATE TABLE "user_preference"
(
    "id"                   VARCHAR(32) PRIMARY KEY,
    "user_id"              VARCHAR(32) NOT NULL,

    -- 界面设置 --
    "theme"                VARCHAR(50) DEFAULT 'light',  -- 主题
    "language"             VARCHAR(10) DEFAULT 'zh-CN',  -- 语言
    "font_size"            SMALLINT    DEFAULT 14,       -- 字体大小
    "code_theme"           VARCHAR(50) DEFAULT 'github', -- 代码主题

    -- 通知设置 --
    "email_notifications"  BOOLEAN     DEFAULT TRUE,     -- 邮件通知
    "push_notifications"   BOOLEAN     DEFAULT TRUE,     -- 推送通知

    -- 隐私与展示 --
    "allow_direct_message" BOOLEAN     DEFAULT TRUE,     -- 允许私信

    --- 固定字段 ---
    "deleted"              BOOLEAN     DEFAULT FALSE,
    "create_time"          timestamptz DEFAULT NOW(),
    "create_user"          VARCHAR(32),
    "update_time"          timestamptz DEFAULT NOW(),
    "update_user"          VARCHAR(32)
);

CREATE UNIQUE INDEX "idx_user_preference_user" ON "user_preference" ("user_id");
COMMENT ON TABLE "user_preference" IS '用户偏好设置表';

-- ----------------------------
-- Table structure for user_stats (用户统计信息表)
-- ----------------------------
DROP TABLE IF EXISTS "user_stats";
CREATE TABLE "user_stats"
(
    "id"          VARCHAR(32) PRIMARY KEY,
    "user_id"     VARCHAR(32) NOT NULL,

    -- 等级与经验 --
    "level"       INTEGER     DEFAULT 1, -- 用户等级 (1-100)
    "exp"         BIGINT      DEFAULT 0, -- 经验值
    "title"       VARCHAR(100),          -- 用户头衔

    --- 固定字段 ---
    "deleted"     BOOLEAN     DEFAULT FALSE,
    "create_time" timestamptz DEFAULT NOW(),
    "create_user" VARCHAR(32),
    "update_time" timestamptz DEFAULT NOW(),
    "update_user" VARCHAR(32)
);

CREATE UNIQUE INDEX "idx_user_stats_user" ON "user_stats" ("user_id");
CREATE INDEX "idx_user_level" ON "user_stats" ("level");
COMMENT ON TABLE "user_stats" IS '用户统计信息表';

-- ----------------------------
-- Table structure for user_vip (VIP信息表)
-- ----------------------------
DROP TABLE IF EXISTS "user_vip";
CREATE TABLE "user_vip"
(
    "id"              VARCHAR(32) PRIMARY KEY,
    "user_id"         VARCHAR(32) NOT NULL,
    "vip_level"       SMALLINT    DEFAULT 0,     -- VIP等级 (0:普通用户, 1-10:VIP等级)
    "vip_expire_time" timestamptz,               -- VIP到期时间
    "vip_start_time"  timestamptz,               -- VIP开始时间
    "is_auto_renew"   BOOLEAN     DEFAULT FALSE, -- 是否自动续费

    --- 固定字段 ---
    "deleted"         BOOLEAN     DEFAULT FALSE,
    "create_time"     timestamptz DEFAULT NOW(),
    "create_user"     VARCHAR(32),
    "update_time"     timestamptz DEFAULT NOW(),
    "update_user"     VARCHAR(32)
);

CREATE UNIQUE INDEX "idx_user_vip_user" ON "user_vip" ("user_id");
CREATE INDEX "idx_vip_expire" ON "user_vip" ("vip_expire_time");
COMMENT ON TABLE "user_vip" IS '用户VIP信息表';

-- ----------------------------
-- Table structure for user_badge (用户徽章表)
-- ----------------------------
DROP TABLE IF EXISTS "user_badge";
CREATE TABLE "user_badge"
(
    "id"           VARCHAR(32) PRIMARY KEY,
    "user_id"      VARCHAR(32) NOT NULL,
    "badge_id"     VARCHAR(32) NOT NULL,      -- 徽章类型ID
    "badge_name"   VARCHAR(100),              -- 徽章名称
    "badge_icon"   VARCHAR(255),              -- 徽章图标
    "acquire_time" timestamptz DEFAULT NOW(), -- 获得时间
    "is_equipped"  BOOLEAN     DEFAULT FALSE, -- 是否佩戴

    --- 固定字段 ---
    "deleted"      BOOLEAN     DEFAULT FALSE,
    "create_time"  timestamptz DEFAULT NOW(),
    "create_user"  VARCHAR(32),
    "update_time"  timestamptz DEFAULT NOW(),
    "update_user"  VARCHAR(32)
);

CREATE INDEX "idx_user_badge_user" ON "user_badge" ("user_id");
CREATE INDEX "idx_badge_equipped" ON "user_badge" ("is_equipped");
COMMENT ON TABLE "user_badge" IS '用户徽章表';

-- ----------------------------
-- Table structure for badge_config (徽章配置表)
-- ----------------------------
DROP TABLE IF EXISTS "badge_config";
CREATE TABLE "badge_config"
(
    "id"              VARCHAR(32) PRIMARY KEY,
    "name"            VARCHAR(100) NOT NULL, -- 徽章名称
    "code"            VARCHAR(50)  NOT NULL, -- 徽章代码
    "icon"            VARCHAR(255),          -- 徽章图标
    "description"     VARCHAR(500),          -- 徽章描述
    "color"           VARCHAR(20),           -- 徽章颜色
    "rarity"          SMALLINT    DEFAULT 1, -- 稀有度 1-5
    "condition_type"  VARCHAR(50),           -- 获得条件类型
    "condition_value" VARCHAR(255),          -- 获得条件值
    --- 固定字段 ---
    "deleted"         BOOLEAN     DEFAULT FALSE,
    "create_time"     timestamptz DEFAULT NOW(),
    "create_user"     VARCHAR(32),
    "update_time"     timestamptz DEFAULT NOW(),
    "update_user"     VARCHAR(32)
);

CREATE UNIQUE INDEX "idx_badge_code" ON "badge_config" ("code");
COMMENT ON TABLE "badge_config" IS '徽章配置表';

-- ----------------------------
-- VIP等级权益表
-- ----------------------------
DROP TABLE IF EXISTS "user_vip_privilege";
CREATE TABLE "user_vip_privilege"
(
    "id"                 VARCHAR(32) PRIMARY KEY,
    "vip_level"          SMALLINT,
    "name"               VARCHAR(50),
    "color"              VARCHAR(20),
    "ai_chat_privilege"  BOOLEAN, -- AI对话特权
    --- 固定字段 ---
    "deleted"            BOOLEAN     DEFAULT FALSE,
    "create_time"        timestamptz DEFAULT NOW(),
    "create_user"        VARCHAR(32),
    "update_time"        timestamptz DEFAULT NOW(),
    "update_user"        VARCHAR(32)
);

-- ----------------------------
-- 用户等级成长配置表
-- ----------------------------
DROP TABLE IF EXISTS "user_level_config";
CREATE TABLE "user_level_config"
(
    "id"             VARCHAR(32) PRIMARY KEY,
    "level"          INTEGER NOT NULL, -- 等级值
    "level_name"     VARCHAR(100),     -- 等级名称
    "exp_required"   BIGINT,           -- 所需经验值
    "badge_unlocked" VARCHAR(100),     -- 解锁的徽章
    "title_unlocked" VARCHAR(100),     -- 解锁的头衔
    --- 固定字段 ---
    "deleted"        BOOLEAN     DEFAULT FALSE,
    "create_time"    timestamptz DEFAULT NOW(),
    "create_user"    VARCHAR(32),
    "update_time"    timestamptz DEFAULT NOW(),
    "update_user"    VARCHAR(32)
);

CREATE UNIQUE INDEX "idx_level_config_level" ON "user_level_config" ("level");

-- ----------------------------
-- Table structure for auth_account_role
-- ----------------------------
DROP TABLE IF EXISTS "auth_account_role";
CREATE TABLE "auth_account_role"
(
    "id"         VARCHAR(32) PRIMARY KEY,
    "account_id" VARCHAR(32) NOT NULL,
    "role_id"    VARCHAR(32) NOT NULL
);
COMMENT ON TABLE "auth_account_role" IS '账户-角色 关联表(1-N)';


-- ----------------------------
-- Table structure for auth_group
-- ----------------------------
DROP TABLE IF EXISTS "auth_group";
CREATE TABLE "auth_group"
(
    "id"             VARCHAR(32) PRIMARY KEY,
    "parent_id"      VARCHAR(32),
    "name"           VARCHAR(100),
    "code"           VARCHAR(50),
    "description"    VARCHAR(255),
    "sort"           SMALLINT    DEFAULT 99,
    "admin_id"       VARCHAR(32),
    "max_user_count" INTEGER, -- 最大用户数限制
    "is_system"      BOOLEAN     DEFAULT FALSE,
    --- 固定字段 ---
    "deleted"        BOOLEAN     DEFAULT FALSE,
    "create_time"    timestamptz DEFAULT NOW(),
    "create_user"    VARCHAR(32),
    "update_time"    timestamptz DEFAULT NOW(),
    "update_user"    VARCHAR(32)
);
CREATE UNIQUE INDEX "idx_group_code" ON "auth_group" ("code");
COMMENT ON TABLE "auth_group" IS '用户组表';

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS "auth_role";
CREATE TABLE "auth_role"
(
    "id"               VARCHAR(32) PRIMARY KEY,
    "name"             VARCHAR(255),
    "code"             VARCHAR(50),
    "data_scope"       VARCHAR(50), -- 角色级别的数据权限
    "description"      VARCHAR(255),
    "assign_group_ids" jsonb,       -- 分配管理的用户组范围
    --- 固定字段 ---
    "deleted"          BOOLEAN     DEFAULT FALSE,
    "create_time"      timestamptz DEFAULT NOW(),
    "create_user"      VARCHAR(32),
    "update_time"      timestamptz DEFAULT NOW(),
    "update_user"      VARCHAR(32)
);
CREATE INDEX "idx_role_code" ON "auth_role" ("code");
CREATE INDEX "idx_name" ON "auth_role" ("name");
CREATE INDEX "idx_data_scope" ON "auth_role" ("data_scope");
COMMENT ON TABLE "auth_role" IS '角色表';

-- ----------------------------
-- Table structure for auth_role_menu
-- ----------------------------
DROP TABLE IF EXISTS "auth_role_menu";
CREATE TABLE "auth_role_menu"
(
    "id"      VARCHAR(32) PRIMARY KEY,
    "role_id" VARCHAR(32) NOT NULL,
    "menu_id" VARCHAR(32) NOT NULL
);
COMMENT ON TABLE "auth_role_menu" IS '角色-菜单 关联表(1-N)';


-- ============================================================ 系统基础 ============================================================ --
-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS "sys_config";
CREATE TABLE "sys_config"
(
    "id"             VARCHAR(32) PRIMARY KEY,
    "name"           VARCHAR(255),
    "code"           VARCHAR(255),
    "value"          VARCHAR(255),
    "component_type" VARCHAR(255),
    "description"    VARCHAR(255),
    "config_type"    VARCHAR(255),
    "sort"           INT         DEFAULT 0,
    --- 固定字段 ---
    "deleted"        BOOLEAN     DEFAULT FALSE,
    "create_time"    timestamptz DEFAULT NOW(),
    "create_user"    VARCHAR(32),
    "update_time"    timestamptz DEFAULT NOW(),
    "update_user"    VARCHAR(32)
);
CREATE UNIQUE INDEX "idx_code" ON "sys_config" ("code");
COMMENT ON TABLE "sys_config" IS '系统配置表';

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS "sys_dict";
CREATE TABLE "sys_dict"
(
    "id"          VARCHAR(32) PRIMARY KEY,
    "dict_type"   VARCHAR(64),
    "type_label"  VARCHAR(64),
    "dict_value"  VARCHAR(255),
    "dict_label"  VARCHAR(255),
    "sort_order"  INT         DEFAULT 0,
    --- 固定字段 ---
    "deleted"     BOOLEAN     DEFAULT FALSE,
    "create_time" timestamptz DEFAULT NOW(),
    "create_user" VARCHAR(32),
    "update_time" timestamptz DEFAULT NOW(),
    "update_user" VARCHAR(32)
);
CREATE UNIQUE INDEX "uk_type_code" ON "sys_dict" ("dict_type", "dict_value");
COMMENT ON TABLE "sys_dict" IS '系统字典表';

-- ----------------------------
-- Table structure for sys_log
-- ----------------------------
DROP TABLE IF EXISTS "sys_log";
CREATE TABLE "sys_log"
(
    "id"             VARCHAR(32) PRIMARY KEY,
    "user_id"        VARCHAR(32),
    "operation"      VARCHAR(255),
    "method"         VARCHAR(255),
    "params"         TEXT,
    "ip"             VARCHAR(255),
    "operation_time" timestamptz,
    "category"       VARCHAR(255),
    "module"         VARCHAR(255),
    "description"    VARCHAR(255),
    "status"         VARCHAR(255),
    "message"        TEXT,
    --- 固定字段 ---
    "deleted"        BOOLEAN     DEFAULT FALSE,
    "create_time"    timestamptz DEFAULT NOW(),
    "create_user"    VARCHAR(32),
    "update_time"    timestamptz DEFAULT NOW(),
    "update_user"    VARCHAR(32)
);
COMMENT ON TABLE "sys_log" IS '系统活动/日志记录表';

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS "sys_menu";
CREATE TABLE "sys_menu"
(
    "id"             VARCHAR(32) PRIMARY KEY,
    "pid"            VARCHAR(32) DEFAULT '0',
    "name"           VARCHAR(100),
    "path"           VARCHAR(200),
    "component_path" VARCHAR(500),
    "title"          VARCHAR(100),
    "icon"           VARCHAR(100),
    "keep_alive"     BOOLEAN     DEFAULT FALSE,
    "visible"        BOOLEAN     DEFAULT TRUE,
    "sort"           INT         DEFAULT 0,
    "pined"          BOOLEAN     DEFAULT FALSE,
    "menu_type"      INT         DEFAULT 0,
    "parameters"     VARCHAR(500),
    "extra_params"   jsonb,
    --- 固定字段 ---
    "deleted"        BOOLEAN     DEFAULT FALSE,
    "create_time"    timestamptz DEFAULT NOW(),
    "create_user"    VARCHAR(32),
    "update_time"    timestamptz DEFAULT NOW(),
    "update_user"    VARCHAR(32)
);
CREATE INDEX "idx_pid" ON "sys_menu" ("pid");
CREATE INDEX "idx_sort" ON "sys_menu" ("sort");
CREATE INDEX "idx_menu_type" ON "sys_menu" ("menu_type");
COMMENT ON TABLE "sys_menu" IS '菜单表';

-- ============================================================ 内容管理 ============================================================ --
-- ----------------------------
-- Table structure for cms_article
-- ----------------------------
DROP TABLE IF EXISTS "cms_article";
CREATE TABLE "cms_article"
(
    "id"          VARCHAR(32) PRIMARY KEY,
    "title"       VARCHAR(255),
    "subtitle"    VARCHAR(255),
    "cover"       VARCHAR(255),
    "author"      VARCHAR(255),
    "summary"     VARCHAR(255),
    "sort"        SMALLINT    DEFAULT 99,
    "to_url"      VARCHAR(255),
    "parent_id"   VARCHAR(32) DEFAULT '0',
    "type"        VARCHAR(32) DEFAULT '0',
    "category"    VARCHAR(32) DEFAULT '0',
    "content"     TEXT,
    --- 固定字段 ---
    "deleted"     BOOLEAN     DEFAULT FALSE,
    "create_time" timestamptz DEFAULT NOW(),
    "create_user" VARCHAR(32),
    "update_time" timestamptz DEFAULT NOW(),
    "update_user" VARCHAR(32)
);
COMMENT ON TABLE "cms_article" IS '系统文章表';

-- ----------------------------
-- Table structure for cms_banner
-- ----------------------------
DROP TABLE IF EXISTS "cms_banner";
CREATE TABLE "cms_banner"
(
    "id"                  VARCHAR(32) PRIMARY KEY,
    "title"               VARCHAR(255),
    "banner"              VARCHAR(255),
    "button_text"         VARCHAR(255),
    "is_visible_button"   BOOLEAN     DEFAULT TRUE,
    "jump_module"         VARCHAR(255),
    "jump_type"           VARCHAR(255),
    "jump_target"         VARCHAR(255),
    "target_blank"        BOOLEAN     DEFAULT FALSE,
    "sort"                SMALLINT    DEFAULT 99,
    "subtitle"            VARCHAR(255),
    "is_visible_subtitle" BOOLEAN     DEFAULT TRUE,
    "is_visible"          BOOLEAN     DEFAULT TRUE,
    --- 固定字段 ---
    "deleted"             BOOLEAN     DEFAULT FALSE,
    "create_time"         timestamptz DEFAULT NOW(),
    "create_user"         VARCHAR(32),
    "update_time"         timestamptz DEFAULT NOW(),
    "update_user"         VARCHAR(32)
);
COMMENT ON TABLE "cms_banner" IS '横幅表';

-- ----------------------------
-- Table structure for cms_category
-- ----------------------------
DROP TABLE IF EXISTS "cms_category";
CREATE TABLE "cms_category"
(
    "id"          VARCHAR(32) PRIMARY KEY,
    "name"        VARCHAR(255),
    "is_visible"  BOOLEAN     DEFAULT TRUE,
    --- 固定字段 ---
    "deleted"     BOOLEAN     DEFAULT FALSE,
    "create_time" timestamptz DEFAULT NOW(),
    "create_user" VARCHAR(32),
    "update_time" timestamptz DEFAULT NOW(),
    "update_user" VARCHAR(32)
);
COMMENT ON TABLE "cms_category" IS '分类表';

-- ----------------------------
-- Table structure for cms_notice
-- ----------------------------
DROP TABLE IF EXISTS "cms_notice";
CREATE TABLE "cms_notice"
(
    "id"          VARCHAR(32) PRIMARY KEY,
    "title"       VARCHAR(64),
    "cover"       VARCHAR(255),
    "url"         VARCHAR(255),
    "sort"        SMALLINT    DEFAULT 99,
    "content"     TEXT,
    "is_visible"  BOOLEAN     DEFAULT TRUE,
    --- 固定字段 ---
    "deleted"     BOOLEAN     DEFAULT FALSE,
    "create_time" timestamptz DEFAULT NOW(),
    "create_user" VARCHAR(32),
    "update_time" timestamptz DEFAULT NOW(),
    "update_user" VARCHAR(32)
);
COMMENT ON TABLE "cms_notice" IS '公告表';

-- ----------------------------
-- Table structure for cms_tag
-- ----------------------------
DROP TABLE IF EXISTS "cms_tag";
CREATE TABLE "cms_tag"
(
    "id"          VARCHAR(32) PRIMARY KEY,
    "name"        VARCHAR(255),
    "is_visible"  BOOLEAN     DEFAULT TRUE,
    --- 固定字段 ---
    "deleted"     BOOLEAN     DEFAULT FALSE,
    "create_time" timestamptz DEFAULT NOW(),
    "create_user" VARCHAR(32),
    "update_time" timestamptz DEFAULT NOW(),
    "update_user" VARCHAR(32)
);
COMMENT ON TABLE "cms_tag" IS '标签表';

-- ============================================================ 竞赛与活动 ============================================================ --

-- ----------------------------
-- Table structure for contest_info
-- ----------------------------
DROP TABLE IF EXISTS "contest_info";
CREATE TABLE "contest_info"
(
    "id"                  VARCHAR(32) PRIMARY KEY,
    "title"               VARCHAR(255) NOT NULL,
    "description"         TEXT,
    "contest_type"        VARCHAR(50),
    "rule_type"           VARCHAR(50),
    "category"            VARCHAR(32),
    "cover"               VARCHAR(255),
    "max_team_members"    INT         DEFAULT 1,
    "is_team_contest"     BOOLEAN     DEFAULT FALSE,
    "is_visible"          BOOLEAN     DEFAULT TRUE,
    "is_public"           BOOLEAN     DEFAULT FALSE,
    "password"            VARCHAR(100),
    "register_start_time" timestamptz,
    "register_end_time"   timestamptz,
    "contest_start_time"  timestamptz  NOT NULL,
    "contest_end_time"    timestamptz  NOT NULL,
    "frozen_time"         INT         DEFAULT 0,
    "penalty_time"        INT         DEFAULT 20,
    "allowed_languages"   jsonb,
    "status"              VARCHAR(32) DEFAULT NULL,
    "sort"                INT         DEFAULT 0,
    "deleted"             BOOLEAN     DEFAULT FALSE,
    "create_time"         timestamptz DEFAULT NOW(),
    "create_user"         VARCHAR(32),
    "update_time"         timestamptz DEFAULT NOW(),
    "update_user"         VARCHAR(32)
);
COMMENT ON TABLE "contest_info" IS '竞赛表';

-- ----------------------------
-- Table structure for contest_auth
-- ----------------------------
DROP TABLE IF EXISTS "contest_auth";
CREATE TABLE "contest_auth"
(
    "id"          VARCHAR(32) PRIMARY KEY,
    "contest_id"  VARCHAR(32) NOT NULL,
    "user_id"     VARCHAR(32) NOT NULL,
    "is_auth"     BOOLEAN     DEFAULT FALSE,
    --- 固定字段 ---
    "deleted"     BOOLEAN     DEFAULT FALSE,
    "create_time" timestamptz DEFAULT NOW(),
    "create_user" VARCHAR(32),
    "update_time" timestamptz DEFAULT NOW(),
    "update_user" VARCHAR(32)
);
CREATE UNIQUE INDEX "uk_contest_user" ON "contest_auth" ("contest_id", "user_id");
CREATE INDEX "idx_contest_id" ON "contest_auth" ("contest_id");
CREATE INDEX "idx_user_id" ON "contest_auth" ("user_id");
COMMENT ON TABLE "contest_auth" IS '竞赛认证表';

-- ----------------------------
-- Table structure for contest_participant
-- ----------------------------
DROP TABLE IF EXISTS "contest_participant";
CREATE TABLE "contest_participant"
(
    "id"             VARCHAR(32) PRIMARY KEY,
    "contest_id"     VARCHAR(32) NOT NULL,
    "user_id"        VARCHAR(32) NOT NULL,
    "team_id"        VARCHAR(32),
    "team_name"      VARCHAR(255),
    "is_team_leader" BOOLEAN     DEFAULT FALSE,
    "register_time"  timestamptz,
    "status"         VARCHAR(32),
    --- 固定字段 ---
    "deleted"        BOOLEAN     DEFAULT FALSE,
    "create_time"    timestamptz DEFAULT NOW(),
    "create_user"    VARCHAR(32),
    "update_time"    timestamptz DEFAULT NOW(),
    "update_user"    VARCHAR(32)
);
CREATE UNIQUE INDEX "uk_contest_user" ON "contest_participant" ("contest_id", "user_id");
CREATE INDEX "idx_contest_id" ON "contest_participant" ("contest_id");
CREATE INDEX "idx_user_id" ON "contest_participant" ("user_id");
COMMENT ON TABLE "contest_participant" IS '竞赛参与表';

-- ----------------------------
-- Table structure for contest_problem
-- ----------------------------
DROP TABLE IF EXISTS "contest_problem";
CREATE TABLE "contest_problem"
(
    "id"           VARCHAR(32) PRIMARY KEY,
    "contest_id"   VARCHAR(32) NOT NULL,
    "display_id"   VARCHAR(32),
    "problem_code" VARCHAR(10) NOT NULL,
    "problem_id"   VARCHAR(32) NOT NULL,
    "score"        INT         DEFAULT 0,
    "sort"         INT         DEFAULT 0,
    --- 固定字段 ---
    "deleted"      BOOLEAN     DEFAULT FALSE,
    "create_time"  timestamptz DEFAULT NOW(),
    "create_user"  VARCHAR(32),
    "update_time"  timestamptz DEFAULT NOW(),
    "update_user"  VARCHAR(32)
);
CREATE INDEX "idx_contest_id" ON "contest_problem" ("contest_id");
CREATE INDEX "idx_problem_code" ON "contest_problem" ("problem_code");
COMMENT ON TABLE "contest_problem" IS '竞赛题目表';

-- ============================================================ 题目核心 ============================================================ --

-- ----------------------------
-- Table structure for problem_info
-- ----------------------------
DROP TABLE IF EXISTS "problem_info";
CREATE TABLE "problem_info"
(
    "id"            VARCHAR(32) PRIMARY KEY,
    "display_id"    VARCHAR(32),
    "category_id"   VARCHAR(32)    DEFAULT '0',
    "title"         VARCHAR(255),
    "source"        VARCHAR(255),
    "url"           VARCHAR(255),
    "time_limit"    INT            DEFAULT 0,
    "memory_limit"  INT            DEFAULT 0,
    "description"   TEXT,
    "languages"     jsonb,
    "difficulty"    INT            DEFAULT 1,
    "threshold"     DECIMAL(10, 2) DEFAULT 0.50,
    "use_template"  BOOLEAN        DEFAULT FALSE,
    "code_template" jsonb,
    "is_public"     BOOLEAN        DEFAULT FALSE,
    "is_visible"    BOOLEAN        DEFAULT TRUE,
    "use_ai"        BOOLEAN        DEFAULT FALSE,
    --- 固定字段 ---
    "deleted"       BOOLEAN        DEFAULT FALSE,
    "create_time"   timestamptz    DEFAULT NOW(),
    "create_user"   VARCHAR(32),
    "update_time"   timestamptz    DEFAULT NOW(),
    "update_user"   VARCHAR(32)
);
CREATE INDEX "idx_category_id" ON "problem_info" ("category_id");
CREATE INDEX "idx_is_public" ON "problem_info" ("is_public");
CREATE INDEX "idx_is_visible" ON "problem_info" ("is_visible");
CREATE INDEX "idx_difficulty" ON "problem_info" ("difficulty");
COMMENT ON TABLE "problem_info" IS '题目表';

-- ----------------------------
-- Table structure for problem_tag_rel
-- ----------------------------
DROP TABLE IF EXISTS "problem_tag_rel";
CREATE TABLE "problem_tag_rel"
(
    "id"         VARCHAR(32) PRIMARY KEY,
    "problem_id" VARCHAR(32) NOT NULL,
    "tag_id"     VARCHAR(32) NOT NULL
);
COMMENT ON TABLE "problem_tag_rel" IS '题目标签关联表';

-- ----------------------------
-- Table structure for problem_test_case
-- ----------------------------
DROP TABLE IF EXISTS "problem_test_case";
CREATE TABLE "problem_test_case"
(
    "id"               VARCHAR(32) PRIMARY KEY,
    "problem_id"       VARCHAR(32) NOT NULL,
    "case_sign"        VARCHAR(255),
    "input_data"       TEXT,
    "expected_output"  TEXT,
    "input_file_path"  VARCHAR(500),
    "input_file_size"  BIGINT         DEFAULT 0,
    "output_file_path" VARCHAR(500),
    "output_file_size" BIGINT         DEFAULT 0,
    "is_sample"        BOOLEAN        DEFAULT FALSE,
    "score"            DECIMAL(10, 2) DEFAULT 0.00,
    --- 固定字段 ---
    "deleted"          BOOLEAN        DEFAULT FALSE,
    "create_time"      timestamptz    DEFAULT NOW(),
    "create_user"      VARCHAR(32),
    "update_time"      timestamptz    DEFAULT NOW(),
    "update_user"      VARCHAR(32)
);
CREATE INDEX "idx_problem_id" ON "problem_test_case" ("problem_id");
COMMENT ON TABLE "problem_test_case" IS '题目测试用例表';

-- ============================================================ 提交与判题 ============================================================ --

-- ----------------------------
-- Table structure for judge_case
-- ----------------------------
DROP TABLE IF EXISTS "judge_case";
CREATE TABLE "judge_case"
(
    "id"               VARCHAR(32) PRIMARY KEY,
    "submit_id"        VARCHAR(32) NOT NULL,
    "case_sign"        VARCHAR(255),
    "input_data"       TEXT,
    "output_data"      TEXT,
    "expected_output"  TEXT,
    "input_file_path"  VARCHAR(500),
    "input_file_size"  BIGINT         DEFAULT 0,
    "output_file_path" VARCHAR(500),
    "output_file_size" BIGINT         DEFAULT 0,
    "max_time"         DECIMAL(10, 2) DEFAULT 0.00,
    "max_memory"       DECIMAL(10, 2) DEFAULT 0.00,
    "is_sample"        BOOLEAN        DEFAULT FALSE,
    "score"            DECIMAL(10, 2) DEFAULT 0.00,
    "status"           VARCHAR(32),
    "message"          TEXT,
    "exit_code"        BIGINT         DEFAULT 0,
    --- 固定字段 ---
    "deleted"          BOOLEAN        DEFAULT FALSE,
    "create_time"      timestamptz    DEFAULT NOW(),
    "create_user"      VARCHAR(32),
    "update_time"      timestamptz    DEFAULT NOW(),
    "update_user"      VARCHAR(32)
);
COMMENT ON TABLE "judge_case" IS '判题结果用例表';


-- ----------------------------
-- Table structure for judge_submit
-- ----------------------------
DROP TABLE IF EXISTS "judge_submit";
CREATE TABLE "judge_submit"
(
    "id"              VARCHAR(32) PRIMARY KEY,
    "user_id"         VARCHAR(32),
    "module_type"     VARCHAR(32),
    "module_id"       VARCHAR(32),
    "problem_id"      VARCHAR(32),
    "language"        VARCHAR(64),
    "code"            TEXT,
    "code_length"     INT         DEFAULT 0,
    "is_test_submit"  BOOLEAN     DEFAULT FALSE,
    "is_admin_submit" BOOLEAN     DEFAULT FALSE,
    "max_time"        INT         DEFAULT 0,
    "max_memory"      INT         DEFAULT 0,
    "message"         TEXT,
    "status"          VARCHAR(32),
    "is_finish"       BOOLEAN     DEFAULT FALSE,
    "task_id"         VARCHAR(32),
    --- 固定字段 ---
    "deleted"         BOOLEAN     DEFAULT FALSE,
    "create_time"     timestamptz DEFAULT NOW(),
    "create_user"     VARCHAR(32),
    "update_time"     timestamptz DEFAULT NOW(),
    "update_user"     VARCHAR(32)
);
CREATE INDEX "idx_user_id" ON "judge_submit" ("user_id");
CREATE INDEX "idx_problem_id" ON "judge_submit" ("problem_id");
CREATE INDEX "idx_language" ON "judge_submit" ("language");
COMMENT ON TABLE "judge_submit" IS '提交表';


-- ============================================================ 用户学习记录 ============================================================ --
-- ----------------------------
-- Table structure for user_code_library
-- ----------------------------
DROP TABLE IF EXISTS "user_code_library";
CREATE TABLE "user_code_library"
(
    "id"               VARCHAR(32) PRIMARY KEY,
    "user_id"          VARCHAR(32),
    "module_type"      VARCHAR(32),
    "module_id"        VARCHAR(32),
    "problem_id"       VARCHAR(32),
    "submit_id"        VARCHAR(32),
    "submit_time"      timestamptz,
    "language"         VARCHAR(64),
    "code"             TEXT,
    "code_token"       jsonb,
    "code_token_name"  jsonb,
    "code_token_texts" jsonb,
    "code_length"      INT         DEFAULT 0,
    "access_count"     INT         DEFAULT 0,
    --- 固定字段 ---
    "deleted"          BOOLEAN     DEFAULT FALSE,
    "create_time"      timestamptz DEFAULT NOW(),
    "create_user"      VARCHAR(32),
    "update_time"      timestamptz DEFAULT NOW(),
    "update_user"      VARCHAR(32)
);
CREATE INDEX "idx_set_id" ON "user_code_library" ("module_type", "module_id");
CREATE INDEX "idx_problem_id" ON "user_code_library" ("problem_id");
CREATE INDEX "idx_language" ON "user_code_library" ("language");
COMMENT ON TABLE "user_code_library" IS '用户提交代码库';

-- ----------------------------
-- Table structure for user_progress
-- ----------------------------
DROP TABLE IF EXISTS "user_progress";
CREATE TABLE "user_progress"
(
    "id"          VARCHAR(32) PRIMARY KEY,
    "user_id"     VARCHAR(32),
    "module_type" VARCHAR(32),
    "module_id"   VARCHAR(32),
    "problem_id"  VARCHAR(32),
    "status"      VARCHAR(32),
    "is_finish"   BOOLEAN     DEFAULT FALSE,
    "finish_time" timestamptz,
    --- 固定字段 ---
    "deleted"     BOOLEAN     DEFAULT FALSE,
    "create_time" timestamptz DEFAULT NOW(),
    "create_user" VARCHAR(32),
    "update_time" timestamptz DEFAULT NOW(),
    "update_user" VARCHAR(32)
);
COMMENT ON TABLE "user_progress" IS '题集进度表';


-- ----------------------------
-- Table structure for user_solved
-- ----------------------------
DROP TABLE IF EXISTS "user_solved";
CREATE TABLE "user_solved"
(
    "id"                VARCHAR(32) PRIMARY KEY,
    "module_type"       VARCHAR(32),
    "module_id"         VARCHAR(32),
    "user_id"           VARCHAR(32),
    "problem_id"        VARCHAR(32),
    "submit_id"         VARCHAR(32),
    "is_solved"         BOOLEAN     DEFAULT FALSE,
    "first_solved_time" timestamptz,
    "solved_time"       timestamptz,
    "first_submit_time" timestamptz,
    --- 固定字段 ---
    "deleted"           BOOLEAN     DEFAULT FALSE,
    "create_time"       timestamptz DEFAULT NOW(),
    "create_user"       VARCHAR(32),
    "update_time"       timestamptz DEFAULT NOW(),
    "update_user"       VARCHAR(32)
);
COMMENT ON TABLE "user_solved" IS '用户解决表';

-- ============================================================ 代码相似度 ============================================================ --
-- ----------------------------
-- Table structure for similarity_stat
-- ----------------------------
DROP TABLE IF EXISTS "similarity_stat";
CREATE TABLE "similarity_stat"
(
    "id"                      VARCHAR(32) PRIMARY KEY,
    "statistics_type"         INT            DEFAULT 0,
    "task_id"                 VARCHAR(32),
    "module_type"             VARCHAR(32),
    "module_id"               VARCHAR(32),
    "problem_id"              VARCHAR(32),
    "sample_count"            INT            DEFAULT 0,
    "similarity_group_count"  INT            DEFAULT 0,
    "avg_similarity"          DECIMAL(10, 2) DEFAULT 0.00,
    "max_similarity"          DECIMAL(10, 2) DEFAULT 0.00,
    "threshold"               DECIMAL(10, 2) DEFAULT 0.50,
    "similarity_distribution" jsonb,
    "degree_statistics"       jsonb,
    --- 固定字段 ---
    "deleted"                 BOOLEAN        DEFAULT FALSE,
    "create_time"             timestamptz    DEFAULT NOW(),
    "create_user"             VARCHAR(32),
    "update_time"             timestamptz    DEFAULT NOW(),
    "update_user"             VARCHAR(32)
);
CREATE INDEX "idx_task_id" ON "similarity_stat" ("task_id");
CREATE INDEX "idx_problem_id" ON "similarity_stat" ("problem_id");
COMMENT ON TABLE "similarity_stat" IS '相似统计表';

-- ----------------------------
-- Table structure for data_similarity
-- ----------------------------
DROP TABLE IF EXISTS "similarity_segment";
CREATE TABLE "similarity_segment"
(
    "id"                       VARCHAR(32) PRIMARY KEY,
    "task_id"                  VARCHAR(32),
    "task_type"                BOOLEAN        DEFAULT FALSE,
    "problem_id"               VARCHAR(32),
    "module_type"              VARCHAR(32),
    "module_id"                VARCHAR(32),
    "language"                 VARCHAR(64),
    "similarity"               DECIMAL(10, 2) DEFAULT 0.00,

    "submit_user"              VARCHAR(32),
    "submit_code"              TEXT,
    "submit_code_length"       INT            DEFAULT 0,
    "submit_id"                VARCHAR(32),
    "submit_time"              timestamptz,
    "submit_code_token"        jsonb,
    "submit_code_token_name"   jsonb,
    "submit_code_token_texts"  jsonb,

    "library_user"             VARCHAR(32),
    "library_code"             TEXT,
    "library_code_length"      INT            DEFAULT 0,
    "library_id"               VARCHAR(32),
    "library_time"             timestamptz,
    "library_code_token"       jsonb,
    "library_code_token_name"  jsonb,
    "library_code_token_texts" jsonb,

    --- 固定字段 ---
    "deleted"                  BOOLEAN        DEFAULT FALSE,
    "create_time"              timestamptz    DEFAULT NOW(),
    "create_user"              VARCHAR(32),
    "update_time"              timestamptz    DEFAULT NOW(),
    "update_user"              VARCHAR(32)
);
COMMENT ON TABLE "similarity_segment" IS '相似片段表';

-- ============================================================ AI对话 ============================================================ --

-- ----------------------------
-- Table structure for ai_chat_memory
-- ----------------------------
DROP TABLE IF EXISTS "ai_chat_memory";
CREATE TABLE "ai_chat_memory"
(
    "id"              BIGSERIAL PRIMARY KEY,
    "conversation_id" VARCHAR(256) NOT NULL,
    "content"         TEXT         NOT NULL,
    "type"            VARCHAR(100) NOT NULL,
    "timestamp"       timestamptz  NOT NULL,
    CONSTRAINT "chk_message_type" CHECK ("type" IN ('USER', 'ASSISTANT', 'SYSTEM', 'TOOL'))
);


-- ----------------------------
-- Table structure for ai_conversation
-- ----------------------------
DROP TABLE IF EXISTS "ai_conversation";
CREATE TABLE "ai_conversation"
(
    "id"                 VARCHAR(32) PRIMARY KEY,
    "conversation_id"    VARCHAR(64),
    "problem_id"         VARCHAR(32),
    "set_id"             VARCHAR(32),
    "is_set"             BOOLEAN     DEFAULT FALSE,
    "user_id"            VARCHAR(32),
    "message_type"       VARCHAR(255),
    "message_role"       VARCHAR(32),
    "message_content"    TEXT,
    "user_code"          TEXT,
    "language"           VARCHAR(255),
    "prompt_tokens"      INT         DEFAULT 0,
    "completion_tokens"  INT         DEFAULT 0,
    "total_tokens"       INT         DEFAULT 0,
    "response_time"      timestamptz,
    "streaming_duration" INT,
    "status"             VARCHAR(32),
    "error_message"      TEXT,
    "user_platform"      VARCHAR(255),
    "ip_address"         VARCHAR(255),
    --- 固定字段 ---
    "deleted"            BOOLEAN     DEFAULT FALSE,
    "create_time"        timestamptz DEFAULT NOW(),
    "create_user"        VARCHAR(32),
    "update_time"        timestamptz DEFAULT NOW(),
    "update_user"        VARCHAR(32)
);
COMMENT ON TABLE "ai_conversation" IS '大模型对话表';


