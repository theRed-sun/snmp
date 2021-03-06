package mibs

import (
	"github.com/soniah/gosnmp"
)

var oid518 = []OIDAttr{
	OIDAttr{Name: "k518_software_version", OID: "1800.51.10.1", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_phone_style", OID: "1800.51.10.2", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_current_time", OID: "1800.51.10.3", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_current_net_status", OID: "1800.51.10.4", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_current_net_ip", OID: "1800.51.10.5", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_current_net_mac", OID: "1800.51.10.6", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net0_link_status", OID: "1800.51.10.7", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net0_adptive_status", OID: "1800.51.10.8", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net0_rate_status", OID: "1800.51.10.9", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net0_mode_status", OID: "1800.51.10.10", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net1_link_status", OID: "1800.51.10.11", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net1_adptive_status", OID: "1800.51.10.12", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net1_rate_status", OID: "1800.51.10.13", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net1_mode_status", OID: "1800.51.10.14", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_sip_current_account", OID: "1800.51.11.1", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_sip_current_server", OID: "1800.51.11.2", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_sip_current_port", OID: "1800.51.11.3", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_sip_current_status", OID: "1800.51.11.4", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "k518_net_mode", OID: "1800.51.12.1", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518NetMode},
	OIDAttr{Name: "k518_net_static_ip", OID: "1800.51.12.2", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518NetStaticIp},
	OIDAttr{Name: "k518_net_static_gateway", OID: "1800.51.12.3", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518NetStaticGateway},
	OIDAttr{Name: "k518_net_static_mask", OID: "1800.51.12.4", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518NetStaticMask},
	OIDAttr{Name: "k518_pppoe_account", OID: "1800.51.12.5", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518PppoeAccount},
	OIDAttr{Name: "k518_pppoe_password", OID: "1800.51.12.6", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518PppoePassword},
	OIDAttr{Name: "k518_dns_mode", OID: "1800.51.12.7", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518DnsMode},
	OIDAttr{Name: "k518_master_dns", OID: "1800.51.12.8", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518MasterDns},
	OIDAttr{Name: "k518_slave_dns", OID: "1800.51.12.9", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SlaveDns},
	OIDAttr{Name: "k518_net0_adptive", OID: "1800.51.12.10", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Net0Adptive},
	OIDAttr{Name: "k518_net0_rate", OID: "1800.51.12.11", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Net0Rate},
	OIDAttr{Name: "k518_net0_mode", OID: "1800.51.12.12", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Net0Mode},
	OIDAttr{Name: "k518_net1_adptive", OID: "1800.51.12.13", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Net1Adptive},
	OIDAttr{Name: "k518_net1_rate", OID: "1800.51.12.14", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Net1Rate},
	OIDAttr{Name: "k518_net1_mode", OID: "1800.51.12.15", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Net1Mode},
	OIDAttr{Name: "k518_admin_name", OID: "1800.51.12.16", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518AdminName},
	OIDAttr{Name: "k518_admin_password", OID: "1800.51.12.17", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518AdminPassword},
	OIDAttr{Name: "k518_language_mode", OID: "1800.51.12.18", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518LanguageMode},
	OIDAttr{Name: "k518_http_mode", OID: "1800.51.12.19", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518HttpMode},
	OIDAttr{Name: "k518_http_port", OID: "1800.51.12.20", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518HttpPort},
	OIDAttr{Name: "k518_telnet_port", OID: "1800.51.12.21", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518TelnetPort},
	OIDAttr{Name: "k518_sip_local_port", OID: "1800.51.12.22", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipLocalPort},
	OIDAttr{Name: "k518_log_ouput_mode", OID: "1800.51.12.23", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518LogOuputMode},
	OIDAttr{Name: "k518_ntp_mode", OID: "1800.51.12.24", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518NtpMode},
	OIDAttr{Name: "k518_primary_ntp_address", OID: "1800.51.12.25", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518PrimaryNtpAddress},
	OIDAttr{Name: "k518_primary_ntp_port", OID: "1800.51.12.26", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518PrimaryNtpPort},
	OIDAttr{Name: "k518_second_ntp_address", OID: "1800.51.12.27", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SecondNtpAddress},
	OIDAttr{Name: "k518_second_ntp_port", OID: "1800.51.12.28", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SecondNtpPort},
	OIDAttr{Name: "k518_time_zone", OID: "1800.51.12.29", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518TimeZone},
	OIDAttr{Name: "k518_snmp_enable", OID: "1800.51.12.30", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SnmpEnable},
	OIDAttr{Name: "k518_snmp_server_ip", OID: "1800.51.12.31", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SnmpServerIp},
	OIDAttr{Name: "k518_snmp_server_port", OID: "1800.51.12.32", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SnmpServerPort},
	OIDAttr{Name: "k518_snmp_device_zone", OID: "1800.51.12.33", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SnmpDeviceZone},
	OIDAttr{Name: "k518_snmp_device_name", OID: "1800.51.12.34", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SnmpDeviceName},
	OIDAttr{Name: "k518_tr069_work_mode", OID: "1800.51.13.1", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Tr069WorkMode},
	OIDAttr{Name: "k518_tr069_acs_server", OID: "1800.51.13.2", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069AcsServer},
	OIDAttr{Name: "k518_tr069_acs_authmode", OID: "1800.51.13.3", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Tr069AcsAuthmode},
	OIDAttr{Name: "k518_tr069_acs_account", OID: "1800.51.13.4", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069AcsAccount},
	OIDAttr{Name: "k518_tr069_acs_password", OID: "1800.51.13.5", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069AcsPassword},
	OIDAttr{Name: "k518_tr069_connect_period", OID: "1800.51.13.6", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Tr069ConnectPeriod},
	OIDAttr{Name: "k518_tr069_cpe_authmode", OID: "1800.51.13.7", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Tr069CpeAuthmode},
	OIDAttr{Name: "k518_tr069_cpe_req_name", OID: "1800.51.13.8", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069CpeReqName},
	OIDAttr{Name: "k518_tr069_cpe_req_password", OID: "1800.51.13.9", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069CpeReqPassword},
	OIDAttr{Name: "k518_tr069_inform_timeout", OID: "1800.51.13.10", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Tr069InformTimeout},
	OIDAttr{Name: "k518_tr069_manufacturer", OID: "1800.51.13.11", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069Manufacturer},
	OIDAttr{Name: "k518_tr069_oui", OID: "1800.51.13.12", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069Oui},
	OIDAttr{Name: "k518_tr069_product_class", OID: "1800.51.13.13", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069ProductClass},
	OIDAttr{Name: "k518_tr069_serial_namber", OID: "1800.51.13.14", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069SerialNamber},
	OIDAttr{Name: "k518_tr069_hardware_version", OID: "1800.51.13.15", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069HardwareVersion},
	OIDAttr{Name: "k518_tr069_software_version", OID: "1800.51.13.16", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069SoftwareVersion},
	OIDAttr{Name: "k518_tr069_conreq_port", OID: "1800.51.13.17", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Tr069ConreqPort},
	OIDAttr{Name: "k518_tr069_conreg_uri", OID: "1800.51.13.18", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069ConregUri},
	OIDAttr{Name: "k518_tr069_conreq_url", OID: "1800.51.13.19", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069ConreqUrl},
	OIDAttr{Name: "k518_tr069_dn0_number", OID: "1800.51.13.20", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069Dn0Number},
	OIDAttr{Name: "k518_tr069_dn0_name", OID: "1800.51.13.21", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518Tr069Dn0Name},
	OIDAttr{Name: "k518_tr069_auto_provision", OID: "1800.51.13.22", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518Tr069AutoProvision},
	OIDAttr{Name: "k518_control_out1_put_type", OID: "1800.51.14.1", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut1PutType},
	OIDAttr{Name: "k518_control_out1_connect_lev", OID: "1800.51.14.2", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut1ConnectLev},
	OIDAttr{Name: "k518_control_out1_connect_sec", OID: "1800.51.14.3", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut1ConnectSec},
	OIDAttr{Name: "k518_control_out1_trigger_type", OID: "1800.51.14.4", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut1TriggerType},
	OIDAttr{Name: "k518_control_out1_door_lev", OID: "1800.51.14.5", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut1DoorLev},
	OIDAttr{Name: "k518_control_out1_door_sec", OID: "1800.51.14.6", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut1DoorSec},
	OIDAttr{Name: "k518_control_out1_door_num", OID: "1800.51.14.7", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518ControlOut1DoorNum},
	OIDAttr{Name: "k518_control_out1_door_psw", OID: "1800.51.14.8", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518ControlOut1DoorPsw},
	OIDAttr{Name: "k518_control_out1_door_max_talk_len", OID: "1800.51.14.9", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut1DoorMaxTalkLen},
	OIDAttr{Name: "k518_control_out2_put_type", OID: "1800.51.14.10", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut2PutType},
	OIDAttr{Name: "k518_control_out2_connect_lev", OID: "1800.51.14.11", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut2ConnectLev},
	OIDAttr{Name: "k518_control_out2_connect_sec", OID: "1800.51.14.12", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut2ConnectSec},
	OIDAttr{Name: "k518_control_out2_trigger_type", OID: "1800.51.14.13", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut2TriggerType},
	OIDAttr{Name: "k518_control_out2_door_lev", OID: "1800.51.14.14", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut2DoorLev},
	OIDAttr{Name: "k518_control_out2_door_sec", OID: "1800.51.14.15", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut2DoorSec},
	OIDAttr{Name: "k518_control_out2_door_num", OID: "1800.51.14.16", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518ControlOut2DoorNum},
	OIDAttr{Name: "k518_control_out2_door_psw", OID: "1800.51.14.17", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518ControlOut2DoorPsw},
	OIDAttr{Name: "k518_control_out2_door_max_talk_len", OID: "1800.51.14.18", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ControlOut2DoorMaxTalkLen},
	OIDAttr{Name: "k518_pnm_switch", OID: "1800.51.14.19", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518PnmSwitch},
	OIDAttr{Name: "k518_pnm_server_ip", OID: "1800.51.14.20", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518PnmServerIp},
	OIDAttr{Name: "k518_pnm_server_port", OID: "1800.51.14.21", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518PnmServerPort},
	OIDAttr{Name: "k518_pnm_area_name", OID: "1800.51.14.22", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518PnmAreaName},
	OIDAttr{Name: "k518_pnm_device_name", OID: "1800.51.14.23", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518PnmDeviceName},
	OIDAttr{Name: "k518_pnm_self_detect_timer", OID: "1800.51.14.24", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518PnmSelfDetectTimer},
	OIDAttr{Name: "k518_pnm_self_pin_control_list", OID: "1800.51.14.25", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518PnmSelfPinControlList},
	OIDAttr{Name: "k518_pnm_self_pin_name_list", OID: "1800.51.14.26", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518PnmSelfPinNameList},
	OIDAttr{Name: "k518_union_switch", OID: "1800.51.14.27", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518UnionSwitch},
	OIDAttr{Name: "k518_dtmf_relay_mode", OID: "1800.51.15.1", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518DtmfRelayMode},
	OIDAttr{Name: "k518_codec_type_list", OID: "1800.51.15.2", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518CodecTypeList},
	OIDAttr{Name: "k518_hotline_enable", OID: "1800.51.15.3", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518HotlineEnable},
	OIDAttr{Name: "k518_hotline_number", OID: "1800.51.15.4", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518HotlineNumber},
	OIDAttr{Name: "k518_hotline_account", OID: "1800.51.15.5", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518HotlineAccount},
	OIDAttr{Name: "k518_quick_attrib", OID: "1800.51.15.6", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518QuickAttrib},
	OIDAttr{Name: "k518_quick_number_list", OID: "1800.51.15.7", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518QuickNumberList},
	OIDAttr{Name: "k518_quick_auto_list", OID: "1800.51.15.8", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518QuickAutoList},
	OIDAttr{Name: "k518_quick_account_list", OID: "1800.51.15.9", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518QuickAccountList},
	OIDAttr{Name: "k518_boardcast_ip_list", OID: "1800.51.15.10", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518BoardcastIpList},
	OIDAttr{Name: "k518_boardcast_port_list", OID: "1800.51.15.11", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518BoardcastPortList},
	OIDAttr{Name: "k518_recv_boardcast_voice_priority", OID: "1800.51.15.12", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518RecvBoardcastVoicePriority},
	OIDAttr{Name: "k518_recv_boardcast_enable_list", OID: "1800.51.15.13", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518RecvBoardcastEnableList},
	OIDAttr{Name: "k518_recv_boardcast_priority_list", OID: "1800.51.15.14", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518RecvBoardcastPriorityList},
	OIDAttr{Name: "k518_recv_boardcast_ip_list", OID: "1800.51.15.15", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518RecvBoardcastIpList},
	OIDAttr{Name: "k518_recv_boardcast_port_list", OID: "1800.51.15.16", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518RecvBoardcastPortList},
	OIDAttr{Name: "k518_microphone_volume", OID: "1800.51.15.17", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518MicrophoneVolume},
	OIDAttr{Name: "k518_speaker_volume", OID: "1800.51.15.18", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SpeakerVolume},
	OIDAttr{Name: "k518_hookon_wait_time", OID: "1800.51.15.19", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518HookonWaitTime},
	OIDAttr{Name: "k518_ring_style", OID: "1800.51.15.20", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518RingStyle},
	OIDAttr{Name: "k518_ring_volume", OID: "1800.51.15.21", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518RingVolume},
	OIDAttr{Name: "k518_forward_enable", OID: "1800.51.15.22", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518ForwardEnable},
	OIDAttr{Name: "k518_forward_number", OID: "1800.51.15.23", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518ForwardNumber},
	OIDAttr{Name: "k518_sip_account_enable", OID: "1800.51.16.1", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipAccountEnable},
	OIDAttr{Name: "k518_sip_register_switch", OID: "1800.51.16.2", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipRegisterSwitch},
	OIDAttr{Name: "k518_sip_unregister_switch", OID: "1800.51.16.3", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipUnregisterSwitch},
	OIDAttr{Name: "k518_sip_register_expire", OID: "1800.51.16.4", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipRegisterExpire},
	OIDAttr{Name: "k518_sip_display_name", OID: "1800.51.16.5", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipDisplayName},
	OIDAttr{Name: "k518_sip_account_name", OID: "1800.51.16.6", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipAccountName},
	OIDAttr{Name: "k518_sip_auth_name", OID: "1800.51.16.7", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipAuthName},
	OIDAttr{Name: "k518_sip_account_password", OID: "1800.51.16.8", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipAccountPassword},
	OIDAttr{Name: "k518_sip_reg_server0", OID: "1800.51.16.9", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipRegServer0},
	OIDAttr{Name: "k518_sip_reg_port0", OID: "1800.51.16.10", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipRegPort0},
	OIDAttr{Name: "k518_sip_reg_domain0", OID: "1800.51.16.11", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipRegDomain0},
	OIDAttr{Name: "k518_sip_reg_server1", OID: "1800.51.16.12", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipRegServer1},
	OIDAttr{Name: "k518_sip_reg_port1", OID: "1800.51.16.13", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipRegPort1},
	OIDAttr{Name: "k518_sip_reg_domain1", OID: "1800.51.16.14", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipRegDomain1},
	OIDAttr{Name: "k518_sip_reg_server2", OID: "1800.51.16.15", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipRegServer2},
	OIDAttr{Name: "k518_sip_reg_port2", OID: "1800.51.16.16", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipRegPort2},
	OIDAttr{Name: "k518_sip_reg_domain2", OID: "1800.51.16.17", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipRegDomain2},
	OIDAttr{Name: "k518_sip_user_agent", OID: "1800.51.16.18", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518SipUserAgent},
	OIDAttr{Name: "k518_sip_heart_beat_enable", OID: "1800.51.16.19", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipHeartBeatEnable},
	OIDAttr{Name: "k518_sip_heart_beat_time", OID: "1800.51.16.20", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipHeartBeatTime},
	OIDAttr{Name: "k518_sip_auto_answer_enable", OID: "1800.51.16.21", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipAutoAnswerEnable},
	OIDAttr{Name: "k518_sip_auto_answer_time", OID: "1800.51.16.22", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipAutoAnswerTime},
	OIDAttr{Name: "k518_sip_user_param_switch", OID: "1800.51.16.23", Type: gosnmp.Integer, ReadOnly: false, ValidHander: k518SipUserParamSwitch},
	OIDAttr{Name: "k518_digit_rule_1_5", OID: "1800.51.17.1", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule15},
	OIDAttr{Name: "k518_digit_rule_6_10", OID: "1800.51.17.2", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule610},
	OIDAttr{Name: "k518_digit_rule_11_15", OID: "1800.51.17.3", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule1115},
	OIDAttr{Name: "k518_digit_rule_16_20", OID: "1800.51.17.4", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule1620},
	OIDAttr{Name: "k518_digit_rule_21_25", OID: "1800.51.17.5", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule2125},
	OIDAttr{Name: "k518_digit_rule_26_30", OID: "1800.51.17.6", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule2630},
	OIDAttr{Name: "k518_digit_rule_31_35", OID: "1800.51.17.7", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule3135},
	OIDAttr{Name: "k518_digit_rule_36_40", OID: "1800.51.17.8", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule3640},
	OIDAttr{Name: "k518_digit_rule_41_45", OID: "1800.51.17.9", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule4145},
	OIDAttr{Name: "k518_digit_rule_46_50", OID: "1800.51.17.10", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: k518DigitRule4650},
}

func init() {
	for _, attr := range oid518 {
		cAttr := attr
		cAttr.OID = oidPrefix + cAttr.OID

		jsonKeyAttr[cAttr.Name] = &cAttr
		oidKeyAttr[cAttr.OID] = &cAttr
	}
}
