/*
Navicat MySQL Data Transfer

Source Server         : 10.2.122.22
Source Server Version : 50505
Source Host           : 10.2.122.22:3306
Source Database       : quant

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-11-03 11:27:41
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for trade
-- ----------------------------
DROP TABLE IF EXISTS `trade`;
CREATE TABLE `trade` (
	`deal_date`	int(8) NOT NULL DEFAULT 0 COMMENT '成交日期',
	`deal_time`	int(6)	NOT NULL DEFAULT 0 COMMENT '成交时间',
	`deal_no`	varchar(64)	NOT NULL DEFAULT '' COMMENT '成交编号',
	`batch_no`	int(8)	NOT NULL DEFAULT 0 COMMENT '委托批号',
	`entrust_no`	int(8)	NOT NULL DEFAULT 0 COMMENT '委托编号',
	`market_no`	varchar(3)	NOT NULL DEFAULT  '' COMMENT '交易市场',
	`stock_code`	varchar(16) NOT NULL DEFAULT ''	COMMENT '证券代码',
	`account_code`	varchar(32)	NOT NULL DEFAULT '' COMMENT '账户编号',
	`combi_no`	varchar(16)	NOT NULL DEFAULT '' COMMENT '组合编号',
	`stockholder_id`	varchar(20)	NOT NULL DEFAULT '' COMMENT '股东代码',
	`report_seat`	varchar(6)	NOT NULL DEFAULT '' COMMENT '申报席位',
	`entrust_direction`	varchar(4)	NOT NULL DEFAULT '' COMMENT '委托方向',
	`futures_direction`	varchar(1)	NOT NULL DEFAULT '' COMMENT '开平方向',
	`entrust_amount`	int(12)	NOT NULL DEFAULT 0 COMMENT '委托数量',
	`entrust_state`	varchar(1)	NOT NULL DEFAULT '' COMMENT '委托状态',
	`entrust_status`	varchar(1)	NOT NULL DEFAULT '' COMMENT '委托状态',
	`deal_amount`	int(16)	NOT NULL DEFAULT 0 COMMENT '本次成交数量',
	`deal_price`	double	NOT NULL DEFAULT 0 COMMENT '本次成交价格',
	`deal_balance`	double	NOT NULL DEFAULT 0 COMMENT '本次成交金额',
	`deal_fee`	double	NOT NULL DEFAULT 0 COMMENT '本次费用',
	`total_deal_amount`	int(16)	NOT NULL DEFAULT 0 COMMENT '累计成交数量',
	`total_deal_balance`	double	NOT NULL DEFAULT 0 COMMENT '累计成交金额',
	`cancel_amount`	int(12)	NOT NULL DEFAULT 0 COMMENT '撤销数量',
	`report_direction`	varchar(2)	NOT NULL DEFAULT '' COMMENT '申报方向',
	`extsystem_id`	int(8)	NOT NULL DEFAULT 0 COMMENT '第三方系统自定义号',
	`third_reff`	varchar(256)	NOT NULL DEFAULT '' COMMENT '第三方系统自定义说明',
	`time_stamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`deal_no`,`deal_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
