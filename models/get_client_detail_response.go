// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetClientDetailResponse get client detail response
// swagger:model GetClientDetailResponse
type GetClientDetailResponse struct {

	// response
	Response *GetClientDetailResponseResponse `json:"response,omitempty"`
}

// Validate validates this get client detail response
func (m *GetClientDetailResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClientDetailResponse) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponse) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientDetailResponseResponse get client detail response response
// swagger:model GetClientDetailResponseResponse
type GetClientDetailResponseResponse struct {

	// connection info
	ConnectionInfo *GetClientDetailResponseResponseConnectionInfo `json:"connectionInfo,omitempty"`

	// detail
	Detail *GetClientDetailResponseResponseDetail `json:"detail,omitempty"`

	// topology
	Topology *GetClientDetailResponseResponseTopology `json:"topology,omitempty"`
}

// Validate validates this get client detail response response
func (m *GetClientDetailResponseResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopology(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClientDetailResponseResponse) validateConnectionInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionInfo) { // not required
		return nil
	}

	if m.ConnectionInfo != nil {
		if err := m.ConnectionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response" + "." + "connectionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *GetClientDetailResponseResponse) validateDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.Detail) { // not required
		return nil
	}

	if m.Detail != nil {
		if err := m.Detail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response" + "." + "detail")
			}
			return err
		}
	}

	return nil
}

func (m *GetClientDetailResponseResponse) validateTopology(formats strfmt.Registry) error {

	if swag.IsZero(m.Topology) { // not required
		return nil
	}

	if m.Topology != nil {
		if err := m.Topology.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response" + "." + "topology")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponseResponse) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientDetailResponseResponseConnectionInfo get client detail response response connection info
// swagger:model GetClientDetailResponseResponseConnectionInfo
type GetClientDetailResponseResponseConnectionInfo struct {

	// band
	Band string `json:"band,omitempty"`

	// channel
	Channel string `json:"channel,omitempty"`

	// channel width
	ChannelWidth string `json:"channelWidth,omitempty"`

	// host type
	HostType string `json:"hostType,omitempty"`

	// nw device mac
	NwDeviceMac string `json:"nwDeviceMac,omitempty"`

	// nw device name
	NwDeviceName string `json:"nwDeviceName,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// spatial stream
	SpatialStream string `json:"spatialStream,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`

	// uapsd
	Uapsd string `json:"uapsd,omitempty"`

	// wmm
	Wmm string `json:"wmm,omitempty"`
}

// Validate validates this get client detail response response connection info
func (m *GetClientDetailResponseResponseConnectionInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponseResponseConnectionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponseResponseConnectionInfo) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponseResponseConnectionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientDetailResponseResponseDetail get client detail response response detail
// swagger:model GetClientDetailResponseResponseDetail
type GetClientDetailResponseResponseDetail struct {

	// ap group
	ApGroup string `json:"apGroup,omitempty"`

	// auth type
	AuthType string `json:"authType,omitempty"`

	// avg rssi
	AvgRssi string `json:"avgRssi,omitempty"`

	// avg snr
	AvgSnr string `json:"avgSnr,omitempty"`

	// channel
	Channel string `json:"channel,omitempty"`

	// client connection
	ClientConnection string `json:"clientConnection,omitempty"`

	// connected device
	ConnectedDevice []string `json:"connectedDevice"`

	// connection status
	ConnectionStatus string `json:"connectionStatus,omitempty"`

	// data rate
	DataRate string `json:"dataRate,omitempty"`

	// dns failure
	DNSFailure string `json:"dnsFailure,omitempty"`

	// dns success
	DNSSuccess string `json:"dnsSuccess,omitempty"`

	// frequency
	Frequency string `json:"frequency,omitempty"`

	// health score
	HealthScore []*GetClientDetailResponseResponseDetailHealthScoreItems0 `json:"healthScore"`

	// host Ip v4
	HostIPV4 string `json:"hostIpV4,omitempty"`

	// host Ip v6
	HostIPV6 []string `json:"hostIpV6"`

	// host mac
	HostMac string `json:"hostMac,omitempty"`

	// host name
	HostName string `json:"hostName,omitempty"`

	// host os
	HostOs string `json:"hostOs,omitempty"`

	// host type
	HostType string `json:"hostType,omitempty"`

	// host version
	HostVersion string `json:"hostVersion,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// issue count
	IssueCount string `json:"issueCount,omitempty"`

	// last updated
	LastUpdated string `json:"lastUpdated,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// onboarding
	Onboarding *GetClientDetailResponseResponseDetailOnboarding `json:"onboarding,omitempty"`

	// onboarding time
	OnboardingTime string `json:"onboardingTime,omitempty"`

	// port
	Port string `json:"port,omitempty"`

	// rssi
	Rssi string `json:"rssi,omitempty"`

	// rx bytes
	RxBytes string `json:"rxBytes,omitempty"`

	// snr
	Snr string `json:"snr,omitempty"`

	// ssid
	Ssid string `json:"ssid,omitempty"`

	// sub type
	SubType string `json:"subType,omitempty"`

	// tx bytes
	TxBytes string `json:"txBytes,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// vlan Id
	VlanID string `json:"vlanId,omitempty"`
}

// Validate validates this get client detail response response detail
func (m *GetClientDetailResponseResponseDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnboarding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClientDetailResponseResponseDetail) validateHealthScore(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthScore) { // not required
		return nil
	}

	for i := 0; i < len(m.HealthScore); i++ {
		if swag.IsZero(m.HealthScore[i]) { // not required
			continue
		}

		if m.HealthScore[i] != nil {
			if err := m.HealthScore[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("response" + "." + "detail" + "." + "healthScore" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetClientDetailResponseResponseDetail) validateOnboarding(formats strfmt.Registry) error {

	if swag.IsZero(m.Onboarding) { // not required
		return nil
	}

	if m.Onboarding != nil {
		if err := m.Onboarding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response" + "." + "detail" + "." + "onboarding")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponseResponseDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponseResponseDetail) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponseResponseDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientDetailResponseResponseDetailHealthScoreItems0 get client detail response response detail health score items0
// swagger:model GetClientDetailResponseResponseDetailHealthScoreItems0
type GetClientDetailResponseResponseDetailHealthScoreItems0 struct {

	// health type
	HealthType string `json:"healthType,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// score
	Score string `json:"score,omitempty"`
}

// Validate validates this get client detail response response detail health score items0
func (m *GetClientDetailResponseResponseDetailHealthScoreItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponseResponseDetailHealthScoreItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponseResponseDetailHealthScoreItems0) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponseResponseDetailHealthScoreItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientDetailResponseResponseDetailOnboarding get client detail response response detail onboarding
// swagger:model GetClientDetailResponseResponseDetailOnboarding
type GetClientDetailResponseResponseDetailOnboarding struct {

	// aaa server Ip
	AaaServerIP string `json:"aaaServerIp,omitempty"`

	// assoc done time
	AssocDoneTime string `json:"assocDoneTime,omitempty"`

	// auth done time
	AuthDoneTime string `json:"authDoneTime,omitempty"`

	// average assoc duration
	AverageAssocDuration string `json:"averageAssocDuration,omitempty"`

	// average auth duration
	AverageAuthDuration string `json:"averageAuthDuration,omitempty"`

	// average dhcp duration
	AverageDhcpDuration string `json:"averageDhcpDuration,omitempty"`

	// average run duration
	AverageRunDuration string `json:"averageRunDuration,omitempty"`

	// dhcp done time
	DhcpDoneTime string `json:"dhcpDoneTime,omitempty"`

	// dhcp server Ip
	DhcpServerIP string `json:"dhcpServerIp,omitempty"`

	// max assoc duration
	MaxAssocDuration string `json:"maxAssocDuration,omitempty"`

	// max auth duration
	MaxAuthDuration string `json:"maxAuthDuration,omitempty"`

	// max dhcp duration
	MaxDhcpDuration string `json:"maxDhcpDuration,omitempty"`

	// max run duration
	MaxRunDuration string `json:"maxRunDuration,omitempty"`
}

// Validate validates this get client detail response response detail onboarding
func (m *GetClientDetailResponseResponseDetailOnboarding) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponseResponseDetailOnboarding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponseResponseDetailOnboarding) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponseResponseDetailOnboarding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientDetailResponseResponseTopology get client detail response response topology
// swagger:model GetClientDetailResponseResponseTopology
type GetClientDetailResponseResponseTopology struct {

	// links
	Links []*GetClientDetailResponseResponseTopologyLinksItems0 `json:"links"`

	// nodes
	Nodes []*GetClientDetailResponseResponseTopologyNodesItems0 `json:"nodes"`
}

// Validate validates this get client detail response response topology
func (m *GetClientDetailResponseResponseTopology) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClientDetailResponseResponseTopology) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("response" + "." + "topology" + "." + "links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetClientDetailResponseResponseTopology) validateNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("response" + "." + "topology" + "." + "nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponseResponseTopology) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponseResponseTopology) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponseResponseTopology
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientDetailResponseResponseTopologyLinksItems0 get client detail response response topology links items0
// swagger:model GetClientDetailResponseResponseTopologyLinksItems0
type GetClientDetailResponseResponseTopologyLinksItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// label
	Label []string `json:"label"`

	// link status
	LinkStatus string `json:"linkStatus,omitempty"`

	// port utilization
	PortUtilization string `json:"portUtilization,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// target
	Target string `json:"target,omitempty"`
}

// Validate validates this get client detail response response topology links items0
func (m *GetClientDetailResponseResponseTopologyLinksItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponseResponseTopologyLinksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponseResponseTopologyLinksItems0) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponseResponseTopologyLinksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientDetailResponseResponseTopologyNodesItems0 get client detail response response topology nodes items0
// swagger:model GetClientDetailResponseResponseTopologyNodesItems0
type GetClientDetailResponseResponseTopologyNodesItems0 struct {

	// clients
	Clients int `json:"clients,omitempty"`

	// connected device
	ConnectedDevice string `json:"connectedDevice,omitempty"`

	// count
	Count int `json:"count,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// device type
	DeviceType string `json:"deviceType,omitempty"`

	// fabric group
	FabricGroup string `json:"fabricGroup,omitempty"`

	// family
	Family string `json:"family,omitempty"`

	// health score
	HealthScore int `json:"healthScore,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ip
	IP string `json:"ip,omitempty"`

	// level
	Level int `json:"level,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// node type
	NodeType string `json:"nodeType,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// radio frequency
	RadioFrequency string `json:"radioFrequency,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// software version
	SoftwareVersion string `json:"softwareVersion,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this get client detail response response topology nodes items0
func (m *GetClientDetailResponseResponseTopologyNodesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetClientDetailResponseResponseTopologyNodesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientDetailResponseResponseTopologyNodesItems0) UnmarshalBinary(b []byte) error {
	var res GetClientDetailResponseResponseTopologyNodesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
