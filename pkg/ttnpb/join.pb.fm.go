// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	go_thethings_network_lorawan_stack_pkg_types "go.thethings.network/lorawan-stack/pkg/types"
	time "time"
)

var JoinRequestFieldPathsNested = []string{
	"Payload",
	"cf_list",
	"cf_list.ch_masks",
	"cf_list.freq",
	"cf_list.type",
	"correlation_ids",
	"downlink_settings",
	"downlink_settings.opt_neg",
	"downlink_settings.rx1_dr_offset",
	"downlink_settings.rx2_dr",
	"end_device_ids",
	"end_device_ids.application_ids",
	"end_device_ids.application_ids.application_id",
	"end_device_ids.dev_addr",
	"end_device_ids.dev_eui",
	"end_device_ids.device_id",
	"end_device_ids.join_eui",
	"net_id",
	"payload",
	"payload.Payload.join_accept_payload",
	"payload.Payload.join_accept_payload.cf_list",
	"payload.Payload.join_accept_payload.cf_list.ch_masks",
	"payload.Payload.join_accept_payload.cf_list.freq",
	"payload.Payload.join_accept_payload.cf_list.type",
	"payload.Payload.join_accept_payload.dev_addr",
	"payload.Payload.join_accept_payload.dl_settings",
	"payload.Payload.join_accept_payload.dl_settings.opt_neg",
	"payload.Payload.join_accept_payload.dl_settings.rx1_dr_offset",
	"payload.Payload.join_accept_payload.dl_settings.rx2_dr",
	"payload.Payload.join_accept_payload.encrypted",
	"payload.Payload.join_accept_payload.join_nonce",
	"payload.Payload.join_accept_payload.net_id",
	"payload.Payload.join_accept_payload.rx_delay",
	"payload.Payload.join_request_payload",
	"payload.Payload.join_request_payload.dev_eui",
	"payload.Payload.join_request_payload.dev_nonce",
	"payload.Payload.join_request_payload.join_eui",
	"payload.Payload.mac_payload",
	"payload.Payload.mac_payload.decoded_payload",
	"payload.Payload.mac_payload.f_hdr",
	"payload.Payload.mac_payload.f_hdr.dev_addr",
	"payload.Payload.mac_payload.f_hdr.f_cnt",
	"payload.Payload.mac_payload.f_hdr.f_ctrl",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.ack",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.adr",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.adr_ack_req",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.class_b",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.f_pending",
	"payload.Payload.mac_payload.f_hdr.f_opts",
	"payload.Payload.mac_payload.f_port",
	"payload.Payload.mac_payload.frm_payload",
	"payload.Payload.rejoin_request_payload",
	"payload.Payload.rejoin_request_payload.dev_eui",
	"payload.Payload.rejoin_request_payload.join_eui",
	"payload.Payload.rejoin_request_payload.net_id",
	"payload.Payload.rejoin_request_payload.rejoin_cnt",
	"payload.Payload.rejoin_request_payload.rejoin_type",
	"payload.m_hdr",
	"payload.m_hdr.m_type",
	"payload.m_hdr.major",
	"payload.mic",
	"raw_payload",
	"rx_delay",
	"selected_mac_version",
}

var JoinRequestFieldPathsTopLevel = []string{
	"Payload",
	"cf_list",
	"correlation_ids",
	"downlink_settings",
	"end_device_ids",
	"net_id",
	"payload",
	"raw_payload",
	"rx_delay",
	"selected_mac_version",
}

func (dst *JoinRequest) SetFields(src *JoinRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "raw_payload":
			if len(subs) > 0 {
				return fmt.Errorf("'raw_payload' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RawPayload = src.RawPayload
			} else {
				var zero []byte
				dst.RawPayload = zero
			}
		case "payload":
			if len(subs) > 0 {
				newDst := dst.Payload
				if newDst == nil {
					newDst = &Message{}
					dst.Payload = newDst
				}
				var newSrc *Message
				if src != nil {
					newSrc = src.Payload
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Payload = src.Payload
				} else {
					dst.Payload = nil
				}
			}
		case "end_device_ids":
			if len(subs) > 0 {
				newDst := &dst.EndDeviceIdentifiers
				var newSrc *EndDeviceIdentifiers
				if src != nil {
					newSrc = &src.EndDeviceIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIdentifiers = src.EndDeviceIdentifiers
				} else {
					var zero EndDeviceIdentifiers
					dst.EndDeviceIdentifiers = zero
				}
			}
		case "selected_mac_version":
			if len(subs) > 0 {
				return fmt.Errorf("'selected_mac_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SelectedMACVersion = src.SelectedMACVersion
			} else {
				var zero MACVersion
				dst.SelectedMACVersion = zero
			}
		case "net_id":
			if len(subs) > 0 {
				return fmt.Errorf("'net_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NetID = src.NetID
			} else {
				var zero go_thethings_network_lorawan_stack_pkg_types.NetID
				dst.NetID = zero
			}
		case "downlink_settings":
			if len(subs) > 0 {
				newDst := &dst.DownlinkSettings
				var newSrc *DLSettings
				if src != nil {
					newSrc = &src.DownlinkSettings
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkSettings = src.DownlinkSettings
				} else {
					var zero DLSettings
					dst.DownlinkSettings = zero
				}
			}
		case "rx_delay":
			if len(subs) > 0 {
				return fmt.Errorf("'rx_delay' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RxDelay = src.RxDelay
			} else {
				var zero RxDelay
				dst.RxDelay = zero
			}
		case "cf_list":
			if len(subs) > 0 {
				newDst := dst.CFList
				if newDst == nil {
					newDst = &CFList{}
					dst.CFList = newDst
				}
				var newSrc *CFList
				if src != nil {
					newSrc = src.CFList
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.CFList = src.CFList
				} else {
					dst.CFList = nil
				}
			}
		case "correlation_ids":
			if len(subs) > 0 {
				return fmt.Errorf("'correlation_ids' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CorrelationIDs = src.CorrelationIDs
			} else {
				dst.CorrelationIDs = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var JoinResponseFieldPathsNested = []string{
	"correlation_ids",
	"lifetime",
	"raw_payload",
	"session_keys",
	"session_keys.app_s_key",
	"session_keys.app_s_key.kek_label",
	"session_keys.app_s_key.key",
	"session_keys.f_nwk_s_int_key",
	"session_keys.f_nwk_s_int_key.kek_label",
	"session_keys.f_nwk_s_int_key.key",
	"session_keys.nwk_s_enc_key",
	"session_keys.nwk_s_enc_key.kek_label",
	"session_keys.nwk_s_enc_key.key",
	"session_keys.s_nwk_s_int_key",
	"session_keys.s_nwk_s_int_key.kek_label",
	"session_keys.s_nwk_s_int_key.key",
	"session_keys.session_key_id",
}

var JoinResponseFieldPathsTopLevel = []string{
	"correlation_ids",
	"lifetime",
	"raw_payload",
	"session_keys",
}

func (dst *JoinResponse) SetFields(src *JoinResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "raw_payload":
			if len(subs) > 0 {
				return fmt.Errorf("'raw_payload' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.RawPayload = src.RawPayload
			} else {
				var zero []byte
				dst.RawPayload = zero
			}
		case "session_keys":
			if len(subs) > 0 {
				newDst := &dst.SessionKeys
				var newSrc *SessionKeys
				if src != nil {
					newSrc = &src.SessionKeys
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.SessionKeys = src.SessionKeys
				} else {
					var zero SessionKeys
					dst.SessionKeys = zero
				}
			}
		case "lifetime":
			if len(subs) > 0 {
				return fmt.Errorf("'lifetime' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Lifetime = src.Lifetime
			} else {
				var zero time.Duration
				dst.Lifetime = zero
			}
		case "correlation_ids":
			if len(subs) > 0 {
				return fmt.Errorf("'correlation_ids' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CorrelationIDs = src.CorrelationIDs
			} else {
				dst.CorrelationIDs = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}