/*
Navicat MySQL Data Transfer

Source Server         : 10.2.122.22
Source Server Version : 50505
Source Host           : 10.2.122.22:3306
Source Database       : quant

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-11-03 11:44:02
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
	`tactic_id` varchar(120) NOT NULL DEFAULT '' COMMENT '策略ID',
	`tactic_type` varchar(120) NOT NULL DEFAULT '' COMMENT '策略类型',
	`account_code`	varchar(32) NOT NULL DEFAULT '' COMMENT '账户编号',
	`batch_no`	int(8) NOT NULL DEFAULT 0 COMMENT '委托批号',
	`business_date`	int(8) NOT NULL DEFAULT 0 COMMENT '委托日期',
	`business_time`	int(6) NOT NULL DEFAULT 0 COMMENT '委托时间',
	`cancel_entrust_no`	int(8) NOT NULL DEFAULT 0 COMMENT '委托编号',
	`combi_no`	varchar(16) NOT NULL DEFAULT '' COMMENT '组合编号',
	`confirm_no`	varchar(32) NOT NULL DEFAULT '' COMMENT '委托确认号',
	`entrust_amount`	int(12) NOT NULL DEFAULT 0 COMMENT '委托数量',
	`cancel_amount`	int(12) NOT NULL DEFAULT 0 COMMENT '撤销数量',
	`entrust_direction`	varchar(4) NOT NULL DEFAULT '' COMMENT '委托方向',
	`entrust_no`	int(8) NOT NULL DEFAULT 0 COMMENT '委托编号',
	`entrust_price`	double NOT NULL DEFAULT 0 COMMENT '委托价格',
	`entrust_status`	varchar(1) NOT NULL DEFAULT '' COMMENT '委托状态',
	`deal_amount`	int(16) NOT NULL DEFAULT 0 COMMENT '成交数量',
	`deal_balance`	double NOT NULL DEFAULT 0 COMMENT '成交金额',
	`deal_price`	double NOT NULL DEFAULT 0 COMMENT '成交均价',
	`futures_direction`	varchar(1) NOT NULL DEFAULT '' COMMENT '开平方向',
	`invest_type`	varchar(1) NOT NULL DEFAULT '' COMMENT '投资类型',
	`market_no`	varchar(3) NOT NULL DEFAULT '' COMMENT '交易市场',
	`price_type`	varchar(1) NOT NULL DEFAULT '' COMMENT '委托价格类型',
	`report_no`	varchar(64) NOT NULL DEFAULT '' COMMENT '申报编号',
	`report_seat`	varchar(6) NOT NULL DEFAULT '' COMMENT '申报席位',
	`revoke_cause`	varchar(256) NOT NULL DEFAULT '' COMMENT '废单原因',
	`stock_code`	varchar(16) NOT NULL DEFAULT '' COMMENT '证券代码',
	`stockholder_id`	varchar(20) NOT NULL DEFAULT '' COMMENT '股东代码',
	`insert_date` date DEFAULT NULL COMMENT '插入日期',
	`insert_time` time DEFAULT NULL COMMENT '插入时间',
	`stop_price` double DEFAULT NULL COMMENT '止损价',
	`time_condition` varchar(32) DEFAULT NULL COMMENT '时间条件，达到后撤单',
	`volume_condition` varchar(32) DEFAULT NULL COMMENT '成交量条件，达到后xx',
	`third_reff`	varchar(256) NOT NULL DEFAULT '' COMMENT '第三方系统自定义说明',
	`extsystem_id`	int(8) NOT NULL DEFAULT 0 COMMENT '第三方系统自定义号',
	`unix_time` bigint DEFAULT NULL,
	PRIMARY KEY (`tactic_id`,`entrust_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
