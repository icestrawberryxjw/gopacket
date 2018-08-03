// Copyright 2017 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"testing"

	"github.com/icestrawberrxjw/gopacket"
)

// First packet is a REGISTER Request
//
// REGISTER sip:sip.provider.com SIP/2.0
// Via:SIP/2.0/UDP 172.16.254.66:5060;branch=z9hG4bK3e5380d454981e88702eb2269669462;rport
// From:"Bob" <sip:bob@sip.provider.com>;tag=3718850509
// To:"Alice" <sip:alice@sip.provider.com>
// Call-ID:306366781@172_16_254_66
// CSeq:3 REGISTER
// Max-Forwards:70
// Allow:INVITE,ACK,CANCEL,BYE,OPTIONS,INFO,SUBSCRIBE,NOTIFY,REFER,UPDATE
// Contact: <sip:bob@172.16.254.66:5060>
// Expires:1800
// User-Agent:C530 IP/42.245.00.000.000
// Content-Length:0
//
var testPacketSIPRequest = []byte{
	0x00, 0x07, 0x7d, 0x41, 0x2e, 0x40, 0x00, 0xd0, 0x03, 0x75, 0xe0, 0x00, 0x08, 0x00, 0x45, 0x00,
	0x01, 0xf4, 0x73, 0x74, 0x00, 0x00, 0x75, 0x11, 0xca, 0x7f, 0x01, 0x01, 0x01, 0x01, 0x02, 0x02,
	0x02, 0x02, 0x13, 0xc4, 0x13, 0xc4, 0x01, 0xe0, 0x86, 0xa0, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x45, 0x52, 0x20, 0x73, 0x69, 0x70, 0x3a, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x20, 0x53, 0x49, 0x50, 0x2f, 0x32, 0x2e, 0x30, 0x0d,
	0x0a, 0x56, 0x69, 0x61, 0x3a, 0x53, 0x49, 0x50, 0x2f, 0x32, 0x2e, 0x30, 0x2f, 0x55, 0x44, 0x50,
	0x20, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x36, 0x2e, 0x32, 0x35, 0x34, 0x2e, 0x36, 0x36, 0x3a, 0x35,
	0x30, 0x36, 0x30, 0x3b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x3d, 0x7a, 0x39, 0x68, 0x47, 0x34,
	0x62, 0x4b, 0x33, 0x65, 0x35, 0x33, 0x38, 0x30, 0x64, 0x34, 0x35, 0x34, 0x39, 0x38, 0x31, 0x65,
	0x38, 0x38, 0x37, 0x30, 0x32, 0x65, 0x62, 0x32, 0x32, 0x36, 0x39, 0x36, 0x36, 0x39, 0x34, 0x36,
	0x32, 0x3b, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x0d, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x3a, 0x22, 0x42,
	0x6f, 0x62, 0x22, 0x20, 0x3c, 0x73, 0x69, 0x70, 0x3a, 0x62, 0x6f, 0x62, 0x40, 0x73, 0x69, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x3e, 0x3b, 0x74,
	0x61, 0x67, 0x3d, 0x33, 0x37, 0x31, 0x38, 0x38, 0x35, 0x30, 0x35, 0x30, 0x39, 0x0d, 0x0a, 0x54,
	0x6f, 0x3a, 0x22, 0x41, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x20, 0x3c, 0x73, 0x69, 0x70, 0x3a, 0x61,
	0x6c, 0x69, 0x63, 0x65, 0x40, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x3e, 0x0d, 0x0a, 0x43, 0x61, 0x6c, 0x6c, 0x2d, 0x49, 0x44, 0x3a,
	0x33, 0x30, 0x36, 0x33, 0x36, 0x36, 0x37, 0x38, 0x31, 0x40, 0x31, 0x37, 0x32, 0x5f, 0x31, 0x36,
	0x5f, 0x32, 0x35, 0x34, 0x5f, 0x36, 0x36, 0x0d, 0x0a, 0x43, 0x53, 0x65, 0x71, 0x3a, 0x33, 0x20,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x0d, 0x0a, 0x4d, 0x61, 0x78, 0x2d, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x73, 0x3a, 0x37, 0x30, 0x0d, 0x0a, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x3a, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x2c, 0x41, 0x43, 0x4b, 0x2c, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x2c, 0x42, 0x59, 0x45, 0x2c, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x2c, 0x49,
	0x4e, 0x46, 0x4f, 0x2c, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x2c, 0x4e, 0x4f,
	0x54, 0x49, 0x46, 0x59, 0x2c, 0x52, 0x45, 0x46, 0x45, 0x52, 0x2c, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x3a, 0x20, 0x3c, 0x73, 0x69, 0x70,
	0x3a, 0x62, 0x6f, 0x62, 0x40, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x36, 0x2e, 0x32, 0x35, 0x34, 0x2e,
	0x36, 0x36, 0x3a, 0x35, 0x30, 0x36, 0x30, 0x3e, 0x0d, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x3a, 0x31, 0x38, 0x30, 0x30, 0x0d, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2d, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x3a, 0x43, 0x35, 0x33, 0x30, 0x20, 0x49, 0x50, 0x2f, 0x34, 0x32, 0x2e, 0x32, 0x34,
	0x35, 0x2e, 0x30, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x0d, 0x0a, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x30, 0x0d, 0x0a,
	0x0d, 0x0a,
}

// Second packet is a REGISTER Response
//
// SIP/2.0 200 OK
// Via:SIP/2.0/UDP 172.16.254.66:5060;received=8.8.8.8;rport=5060;branch=z9hG4bK3e5380d454981e88702eb2269669462
// From:"Bob" <sip:bob@sip.provider.com>;tag=3718850509
// To:"Alice" <sip:alice@sip.provider.com>;tag=02-32748-1417c4ac-24835dbf3
// Call-ID:306366781@172_16_254_66
// CSeq:3 REGISTER
// Contact: <sip:bob@172.16.254.66:5060>;expires=1800
// P-Associated-URI: <sip:bob@sip.provider.com>
// Content-Length:0
//
var testPacketSIPResponse = []byte{
	0x00, 0xd0, 0x00, 0x4a, 0x2c, 0x00, 0x00, 0x07, 0x7d, 0x41, 0x2e, 0x40, 0x08, 0x00, 0x45, 0x00,
	0x01, 0xc1, 0x00, 0x00, 0x40, 0x00, 0x3f, 0x11, 0x34, 0x27, 0x02, 0x02, 0x02, 0x02, 0x01, 0x01,
	0x01, 0x01, 0x13, 0xc4, 0x13, 0xc4, 0x01, 0xad, 0x60, 0x36, 0x53, 0x49, 0x50, 0x2f, 0x32, 0x2e,
	0x30, 0x20, 0x32, 0x30, 0x30, 0x20, 0x4f, 0x4b, 0x0d, 0x0a, 0x56, 0x69, 0x61, 0x3a, 0x53, 0x49,
	0x50, 0x2f, 0x32, 0x2e, 0x30, 0x2f, 0x55, 0x44, 0x50, 0x20, 0x31, 0x37, 0x32, 0x2e, 0x31, 0x36,
	0x2e, 0x32, 0x35, 0x34, 0x2e, 0x36, 0x36, 0x3a, 0x35, 0x30, 0x36, 0x30, 0x3b, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x3d, 0x38, 0x2e, 0x38, 0x2e, 0x38, 0x2e, 0x38, 0x3b, 0x72, 0x70,
	0x6f, 0x72, 0x74, 0x3d, 0x35, 0x30, 0x36, 0x30, 0x3b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x3d,
	0x7a, 0x39, 0x68, 0x47, 0x34, 0x62, 0x4b, 0x33, 0x65, 0x35, 0x33, 0x38, 0x30, 0x64, 0x34, 0x35,
	0x34, 0x39, 0x38, 0x31, 0x65, 0x38, 0x38, 0x37, 0x30, 0x32, 0x65, 0x62, 0x32, 0x32, 0x36, 0x39,
	0x36, 0x36, 0x39, 0x34, 0x36, 0x32, 0x0d, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x3a, 0x22, 0x42, 0x6f,
	0x62, 0x22, 0x20, 0x3c, 0x73, 0x69, 0x70, 0x3a, 0x62, 0x6f, 0x62, 0x40, 0x73, 0x69, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x3e, 0x3b, 0x74, 0x61,
	0x67, 0x3d, 0x33, 0x37, 0x31, 0x38, 0x38, 0x35, 0x30, 0x35, 0x30, 0x39, 0x0d, 0x0a, 0x54, 0x6f,
	0x3a, 0x22, 0x41, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x20, 0x3c, 0x73, 0x69, 0x70, 0x3a, 0x61, 0x6c,
	0x69, 0x63, 0x65, 0x40, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x6d, 0x3e, 0x3b, 0x74, 0x61, 0x67, 0x3d, 0x30, 0x32, 0x2d, 0x33, 0x32, 0x37,
	0x34, 0x38, 0x2d, 0x31, 0x34, 0x31, 0x37, 0x63, 0x34, 0x61, 0x63, 0x2d, 0x32, 0x34, 0x38, 0x33,
	0x35, 0x64, 0x62, 0x66, 0x33, 0x0d, 0x0a, 0x43, 0x61, 0x6c, 0x6c, 0x2d, 0x49, 0x44, 0x3a, 0x33,
	0x30, 0x36, 0x33, 0x36, 0x36, 0x37, 0x38, 0x31, 0x40, 0x31, 0x37, 0x32, 0x5f, 0x31, 0x36, 0x5f,
	0x32, 0x35, 0x34, 0x5f, 0x36, 0x36, 0x0d, 0x0a, 0x43, 0x53, 0x65, 0x71, 0x3a, 0x33, 0x20, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x3a, 0x20, 0x3c, 0x73, 0x69, 0x70, 0x3a, 0x62, 0x6f, 0x62, 0x40, 0x31, 0x37, 0x32, 0x2e, 0x31,
	0x36, 0x2e, 0x32, 0x35, 0x34, 0x2e, 0x36, 0x36, 0x3a, 0x35, 0x30, 0x36, 0x30, 0x3e, 0x3b, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x3d, 0x31, 0x38, 0x30, 0x30, 0x0d, 0x0a, 0x50, 0x2d, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x55, 0x52, 0x49, 0x3a, 0x20, 0x3c,
	0x73, 0x69, 0x70, 0x3a, 0x62, 0x6f, 0x62, 0x40, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x3e, 0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2d, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x30, 0x0d, 0x0a, 0x0d, 0x0a,
}

func TestSIPMain(t *testing.T) {

	expectedHeaders := map[string]string{"Call-ID": "306366781@172_16_254_66", "Contact": "<sip:bob@172.16.254.66:5060>"}
	_TestPacketSIP(t, testPacketSIPRequest, SIPMethodRegister, false, 3, expectedHeaders)

	expectedHeaders = map[string]string{"Call-ID": "306366781@172_16_254_66", "Contact": "<sip:bob@172.16.254.66:5060>;expires=1800"}
	_TestPacketSIP(t, testPacketSIPResponse, SIPMethodRegister, true, 3, expectedHeaders)

}

func _TestPacketSIP(t *testing.T, packetData []byte, methodWanted SIPMethod, isResponse bool, wantedCseq int64, expectedHeaders map[string]string) {

	p := gopacket.NewPacket(packetData, LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}

	if got, ok := p.Layer(LayerTypeSIP).(*SIP); ok {

		// Check method
		if got.Method != methodWanted {
			t.Errorf("SIP Packet should be a %s method, got : %s", methodWanted, got.Method)
		}

		// Check if it's right packet type
		if got.IsResponse != isResponse {
			t.Errorf("SIP packet type is not the same as expected")
		}

		// Check headers
		for headerName, headerValue := range expectedHeaders {
			if got.GetFirstHeader(headerName) != headerValue {
				t.Errorf("Header %s shoud be %s, got : %s", headerName, headerValue, got.GetFirstHeader(headerName))
			}
		}

		// Check CSeq
		if got.GetCSeq() != wantedCseq {
			t.Errorf("SIP Packet should be %d. Got : %d", wantedCseq, got.GetCSeq())
		}
	}
}
