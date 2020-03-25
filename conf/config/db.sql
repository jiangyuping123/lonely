--CREATE DATABASE IF NOT EXISTS cloudinfo;
--use cloudinfo;

use test

DROP TABLE IF EXISTS t_tenant_info;

--租户信息表:
CREATE TABLE IF NOT EXISTS t_tenant_info (
    tenant_id varchar(128) NOT NULL COMMENT '租户ID',
    flowlabel_id int(32) NOT NULL COMMENT '流量标签' ,
    tenant_name varchar(128) NOT NULL COMMENT '租户名称',
    PRIMARY KEY(tenant_id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--服务器主机信息:
CREATE TABLE IF NOT EXISTS t_hypervisors_info (
    id varchar(128) NOT NULL COMMENT '主机id',
    hypervisor_hostname varchar(128) NOT NULL COMMENT '主机名称',
    host_ip varchar(64) NOT NULL COMMENT '主机IP',
    status varchar(16)  NOT NULL COMMENT '主机状态',
    free_disk_gb DECIMAL(20,2) COMMENT '磁盘剩余空间',
    free_ram_mb DECIMAL(20,2) COMMENT '内存剩余',
    local_gb_used DECIMAL(20,2) COMMENT '磁盘使用',
    memory_mb_used DECIMAL(20,2) COMMENT '内存使用',
    PRIMARY KEY(id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--虚拟机信息:
CREATE TABLE IF NOT EXISTS t_vm_info (
    id varchar(128) NOT NULL COMMENT '虚拟机ID',
    name varchar(128) NOT NULL COMMENT '虚拟机名称',
    OS_EXT_SRV_ATTR_host varchar(128) NOT NULL COMMENT '所在节点',
    tenant_id varchar(128) NOT NULL COMMENT '所属租户ID',
    status varchar(16) NOT NULL COMMENT '状态',
    create_time varchar(64) NOT NULL COMMENT '创建时间',
    PRIMARY KEY(id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--虚拟机接口信息:
CREATE TABLE IF NOT EXISTS t_vminterface_info (
    id varchar(128) NOT NULL COMMENT '虚拟机ID',
    ip_address varchar(128) NOT NULL COMMENT '虚拟机IP地址',
    mac_addr varchar(128) NOT NULL COMMENT 'mac地址',
    net_id varchar(128) NOT NULL COMMENT '网络ID',
    port_id varchar(128) NOT NULL COMMENT '网络端口ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--网络信息:
CREATE TABLE IF NOT EXISTS t_network_info(
    id varchar(128) NOT NULL COMMENT '网络ID',
    name varchar(128) NOT NULL COMMENT '网络名称',
    router_external boolean NOT NULL COMMENT '是否是外部路由',
    tenant_id varchar(128) NOT NULL COMMENT '租户ID',
    subnets MEDIUMTEXT NOT NULL COMMENT '子网ID, 用逗号分隔',
    PRIMARY KEY(id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--子网信息:
CREATE TABLE IF NOT EXISTS t_subnet_info(
    id varchar(128) NOT NULL COMMENT '子网ID',    
    name varchar(128) NOT NULL COMMENT '子网名称',
    tenant_id varchar(128) NOT NULL COMMENT '租户ID',
    network_id varchar(128) NOT NULL COMMENT '子网属于的网络ID',
    gateway_ip varchar(128) NOT NULL COMMENT '网关IP',
    cidr varchar(128) NOT NULL COMMENT 'cidr',
    PRIMARY KEY(id) 
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--路由信息:
CREATE TABLE IF NOT EXISTS t_routes_info(
    id varchar(128) NOT NULL COMMENT '路由ID',
    tenant_id varchar(128) NOT NULL COMMENT '租户ID', 
    name varchar(128) NOT NULL COMMENT '路由名称',
    status varchar(16) NOT NULL COMMENT '状态',
    external_gateway_info MEDIUMTEXT NOT NULL COMMENT '外部网关信息',
    routes MEDIUMTEXT NOT NULL COMMENT '路由信息',
    PRIMARY KEY(id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--虚拟端口信息:
CREATE TABLE IF NOT EXISTS t_ports_info(
    id varchar(128) NOT NULL COMMENT '虚拟端口ID',
    name varchar(128) NOT NULL COMMENT '虚拟端口名称',
    device_owner varchar(128) NOT NULL COMMENT '设备类型',
    binding_host_id varchar(128) NOT NULL COMMENT '端口绑定的虚拟机ID',
    device_id varchar(128) NOT NULL COMMENT '端口使用设备的ID', 
    network_id varchar(128) NOT NULL COMMENT '网络ID',
    tenant_id varchar(128) NOT NULL COMMENT '租户ID',
    mac_address varchar(128) NOT NULl COMMENT 'mac地址信息',
    fixed_ips MEDIUMTEXT NOT NULl COMMENT 'IP地址和子网ID[JSON]',
    PRIMARY KEY(id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--浮动IP信息:
CREATE TABLE IF NOT EXISTS t_floatings_ip(
    id varchar(128) NOT NULL COMMENT '浮动IP的id',
    floating_ip_address varchar(128) NOT NULL COMMENT '浮动IP地址',
    dns_domain varchar(128) NOT NULL COMMENT 'DNS 域名',
    tenant_id varchar(128) NOT NULL COMMENT '租户ID',
    dns_name varchar(128) NOT NULL COMMENT 'dns名称',
    port_id varchar(128) NOT NULL COMMENT '端口ID',
    floating_network_id varchar(128) NOT NULL COMMENT '浮动网络ID',
    fixed_ip_address varchar(128) NOT NULL COMMENT '浮动IP地址',
    PRIMARY KEY(id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--loadbalancers信息:
CREATE TABLE IF NOT EXISTS t_loadbalancers_info(
    id varchar(128) NOT NULL COMMENT '均衡器id',
    tenant_id varchar(128) NOT NULL COMMENT '租户id',
    vip_address varchar(128) NOT NULl COMMENT '虚拟ip地址',
    vip_network_id varchar(128) NOT NULL COMMENT '虚拟网络id',
    vip_port_id varchar(128) NOT NULL COMMENT '虚拟端口id',
    PRIMARY KEY(id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--segments 信息
CREATE TABLE IF NOT EXISTS t_segments_info(
    id varchar(128) NOT NULL COMMENT 'id',
    network_id varchar(128) NOT NULL COMMENT '网络id',
    network_type varchar(128) NOT NULL COMMENT '网络类型',
    segmentation_id varchar(128) NOT NULL COMMENT '段id',
    PRIMARY KEY(id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
