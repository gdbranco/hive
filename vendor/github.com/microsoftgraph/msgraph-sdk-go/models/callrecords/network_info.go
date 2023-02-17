package callrecords

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkInfo 
type NetworkInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
    bandwidthLowEventRatio *float32
    // The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
    basicServiceSetIdentifier *string
    // The connectionType property
    connectionType *NetworkConnectionType
    // Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
    delayEventRatio *float32
    // DNS suffix associated with the network adapter of the media endpoint.
    dnsSuffix *string
    // IP address of the media endpoint.
    ipAddress *string
    // Link speed in bits per second reported by the network adapter used by the media endpoint.
    linkSpeed *int64
    // The media access control (MAC) address of the media endpoint's network device.
    macAddress *string
    // The networkTransportProtocol property
    networkTransportProtocol *NetworkTransportProtocol
    // The OdataType property
    odataType *string
    // Network port number used by media endpoint.
    port *int32
    // Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
    receivedQualityEventRatio *float32
    // IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
    reflexiveIPAddress *string
    // IP address of the media relay server allocated by the media endpoint.
    relayIPAddress *string
    // Network port number allocated on the media relay server by the media endpoint.
    relayPort *int32
    // Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
    sentQualityEventRatio *float32
    // Subnet used for media stream by the media endpoint.
    subnet *string
    // List of network trace route hops collected for this media stream.*
    traceRouteHops []TraceRouteHopable
    // The wifiBand property
    wifiBand *WifiBand
    // Estimated remaining battery charge in percentage reported by the media endpoint.
    wifiBatteryCharge *int32
    // WiFi channel used by the media endpoint.
    wifiChannel *int32
    // Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
    wifiMicrosoftDriver *string
    // Version of the Microsoft WiFi driver used by the media endpoint.
    wifiMicrosoftDriverVersion *string
    // The wifiRadioType property
    wifiRadioType *WifiRadioType
    // WiFi signal strength in percentage reported by the media endpoint.
    wifiSignalStrength *int32
    // Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
    wifiVendorDriver *string
    // Version of the WiFi driver used by the media endpoint.
    wifiVendorDriverVersion *string
}
// NewNetworkInfo instantiates a new networkInfo and sets the default values.
func NewNetworkInfo()(*NetworkInfo) {
    m := &NetworkInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateNetworkInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNetworkInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBandwidthLowEventRatio gets the bandwidthLowEventRatio property value. Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
func (m *NetworkInfo) GetBandwidthLowEventRatio()(*float32) {
    return m.bandwidthLowEventRatio
}
// GetBasicServiceSetIdentifier gets the basicServiceSetIdentifier property value. The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
func (m *NetworkInfo) GetBasicServiceSetIdentifier()(*string) {
    return m.basicServiceSetIdentifier
}
// GetConnectionType gets the connectionType property value. The connectionType property
func (m *NetworkInfo) GetConnectionType()(*NetworkConnectionType) {
    return m.connectionType
}
// GetDelayEventRatio gets the delayEventRatio property value. Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
func (m *NetworkInfo) GetDelayEventRatio()(*float32) {
    return m.delayEventRatio
}
// GetDnsSuffix gets the dnsSuffix property value. DNS suffix associated with the network adapter of the media endpoint.
func (m *NetworkInfo) GetDnsSuffix()(*string) {
    return m.dnsSuffix
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NetworkInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bandwidthLowEventRatio"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat32Value(m.SetBandwidthLowEventRatio)
    res["basicServiceSetIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBasicServiceSetIdentifier)
    res["connectionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseNetworkConnectionType , m.SetConnectionType)
    res["delayEventRatio"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat32Value(m.SetDelayEventRatio)
    res["dnsSuffix"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDnsSuffix)
    res["ipAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIpAddress)
    res["linkSpeed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetLinkSpeed)
    res["macAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMacAddress)
    res["networkTransportProtocol"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseNetworkTransportProtocol , m.SetNetworkTransportProtocol)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["port"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPort)
    res["receivedQualityEventRatio"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat32Value(m.SetReceivedQualityEventRatio)
    res["reflexiveIPAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetReflexiveIPAddress)
    res["relayIPAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRelayIPAddress)
    res["relayPort"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRelayPort)
    res["sentQualityEventRatio"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat32Value(m.SetSentQualityEventRatio)
    res["subnet"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubnet)
    res["traceRouteHops"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTraceRouteHopFromDiscriminatorValue , m.SetTraceRouteHops)
    res["wifiBand"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWifiBand , m.SetWifiBand)
    res["wifiBatteryCharge"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWifiBatteryCharge)
    res["wifiChannel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWifiChannel)
    res["wifiMicrosoftDriver"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWifiMicrosoftDriver)
    res["wifiMicrosoftDriverVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWifiMicrosoftDriverVersion)
    res["wifiRadioType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWifiRadioType , m.SetWifiRadioType)
    res["wifiSignalStrength"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWifiSignalStrength)
    res["wifiVendorDriver"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWifiVendorDriver)
    res["wifiVendorDriverVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWifiVendorDriverVersion)
    return res
}
// GetIpAddress gets the ipAddress property value. IP address of the media endpoint.
func (m *NetworkInfo) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetLinkSpeed gets the linkSpeed property value. Link speed in bits per second reported by the network adapter used by the media endpoint.
func (m *NetworkInfo) GetLinkSpeed()(*int64) {
    return m.linkSpeed
}
// GetMacAddress gets the macAddress property value. The media access control (MAC) address of the media endpoint's network device.
func (m *NetworkInfo) GetMacAddress()(*string) {
    return m.macAddress
}
// GetNetworkTransportProtocol gets the networkTransportProtocol property value. The networkTransportProtocol property
func (m *NetworkInfo) GetNetworkTransportProtocol()(*NetworkTransportProtocol) {
    return m.networkTransportProtocol
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NetworkInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetPort gets the port property value. Network port number used by media endpoint.
func (m *NetworkInfo) GetPort()(*int32) {
    return m.port
}
// GetReceivedQualityEventRatio gets the receivedQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
func (m *NetworkInfo) GetReceivedQualityEventRatio()(*float32) {
    return m.receivedQualityEventRatio
}
// GetReflexiveIPAddress gets the reflexiveIPAddress property value. IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
func (m *NetworkInfo) GetReflexiveIPAddress()(*string) {
    return m.reflexiveIPAddress
}
// GetRelayIPAddress gets the relayIPAddress property value. IP address of the media relay server allocated by the media endpoint.
func (m *NetworkInfo) GetRelayIPAddress()(*string) {
    return m.relayIPAddress
}
// GetRelayPort gets the relayPort property value. Network port number allocated on the media relay server by the media endpoint.
func (m *NetworkInfo) GetRelayPort()(*int32) {
    return m.relayPort
}
// GetSentQualityEventRatio gets the sentQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
func (m *NetworkInfo) GetSentQualityEventRatio()(*float32) {
    return m.sentQualityEventRatio
}
// GetSubnet gets the subnet property value. Subnet used for media stream by the media endpoint.
func (m *NetworkInfo) GetSubnet()(*string) {
    return m.subnet
}
// GetTraceRouteHops gets the traceRouteHops property value. List of network trace route hops collected for this media stream.*
func (m *NetworkInfo) GetTraceRouteHops()([]TraceRouteHopable) {
    return m.traceRouteHops
}
// GetWifiBand gets the wifiBand property value. The wifiBand property
func (m *NetworkInfo) GetWifiBand()(*WifiBand) {
    return m.wifiBand
}
// GetWifiBatteryCharge gets the wifiBatteryCharge property value. Estimated remaining battery charge in percentage reported by the media endpoint.
func (m *NetworkInfo) GetWifiBatteryCharge()(*int32) {
    return m.wifiBatteryCharge
}
// GetWifiChannel gets the wifiChannel property value. WiFi channel used by the media endpoint.
func (m *NetworkInfo) GetWifiChannel()(*int32) {
    return m.wifiChannel
}
// GetWifiMicrosoftDriver gets the wifiMicrosoftDriver property value. Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) GetWifiMicrosoftDriver()(*string) {
    return m.wifiMicrosoftDriver
}
// GetWifiMicrosoftDriverVersion gets the wifiMicrosoftDriverVersion property value. Version of the Microsoft WiFi driver used by the media endpoint.
func (m *NetworkInfo) GetWifiMicrosoftDriverVersion()(*string) {
    return m.wifiMicrosoftDriverVersion
}
// GetWifiRadioType gets the wifiRadioType property value. The wifiRadioType property
func (m *NetworkInfo) GetWifiRadioType()(*WifiRadioType) {
    return m.wifiRadioType
}
// GetWifiSignalStrength gets the wifiSignalStrength property value. WiFi signal strength in percentage reported by the media endpoint.
func (m *NetworkInfo) GetWifiSignalStrength()(*int32) {
    return m.wifiSignalStrength
}
// GetWifiVendorDriver gets the wifiVendorDriver property value. Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) GetWifiVendorDriver()(*string) {
    return m.wifiVendorDriver
}
// GetWifiVendorDriverVersion gets the wifiVendorDriverVersion property value. Version of the WiFi driver used by the media endpoint.
func (m *NetworkInfo) GetWifiVendorDriverVersion()(*string) {
    return m.wifiVendorDriverVersion
}
// Serialize serializes information the current object
func (m *NetworkInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat32Value("bandwidthLowEventRatio", m.GetBandwidthLowEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("basicServiceSetIdentifier", m.GetBasicServiceSetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err := writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("delayEventRatio", m.GetDelayEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dnsSuffix", m.GetDnsSuffix())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("linkSpeed", m.GetLinkSpeed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkTransportProtocol() != nil {
        cast := (*m.GetNetworkTransportProtocol()).String()
        err := writer.WriteStringValue("networkTransportProtocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("receivedQualityEventRatio", m.GetReceivedQualityEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reflexiveIPAddress", m.GetReflexiveIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("relayIPAddress", m.GetRelayIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("relayPort", m.GetRelayPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("sentQualityEventRatio", m.GetSentQualityEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnet", m.GetSubnet())
        if err != nil {
            return err
        }
    }
    if m.GetTraceRouteHops() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTraceRouteHops())
        err := writer.WriteCollectionOfObjectValues("traceRouteHops", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWifiBand() != nil {
        cast := (*m.GetWifiBand()).String()
        err := writer.WriteStringValue("wifiBand", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiBatteryCharge", m.GetWifiBatteryCharge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiChannel", m.GetWifiChannel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMicrosoftDriver", m.GetWifiMicrosoftDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMicrosoftDriverVersion", m.GetWifiMicrosoftDriverVersion())
        if err != nil {
            return err
        }
    }
    if m.GetWifiRadioType() != nil {
        cast := (*m.GetWifiRadioType()).String()
        err := writer.WriteStringValue("wifiRadioType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiSignalStrength", m.GetWifiSignalStrength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiVendorDriver", m.GetWifiVendorDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiVendorDriverVersion", m.GetWifiVendorDriverVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBandwidthLowEventRatio sets the bandwidthLowEventRatio property value. Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
func (m *NetworkInfo) SetBandwidthLowEventRatio(value *float32)() {
    m.bandwidthLowEventRatio = value
}
// SetBasicServiceSetIdentifier sets the basicServiceSetIdentifier property value. The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
func (m *NetworkInfo) SetBasicServiceSetIdentifier(value *string)() {
    m.basicServiceSetIdentifier = value
}
// SetConnectionType sets the connectionType property value. The connectionType property
func (m *NetworkInfo) SetConnectionType(value *NetworkConnectionType)() {
    m.connectionType = value
}
// SetDelayEventRatio sets the delayEventRatio property value. Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
func (m *NetworkInfo) SetDelayEventRatio(value *float32)() {
    m.delayEventRatio = value
}
// SetDnsSuffix sets the dnsSuffix property value. DNS suffix associated with the network adapter of the media endpoint.
func (m *NetworkInfo) SetDnsSuffix(value *string)() {
    m.dnsSuffix = value
}
// SetIpAddress sets the ipAddress property value. IP address of the media endpoint.
func (m *NetworkInfo) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetLinkSpeed sets the linkSpeed property value. Link speed in bits per second reported by the network adapter used by the media endpoint.
func (m *NetworkInfo) SetLinkSpeed(value *int64)() {
    m.linkSpeed = value
}
// SetMacAddress sets the macAddress property value. The media access control (MAC) address of the media endpoint's network device.
func (m *NetworkInfo) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetNetworkTransportProtocol sets the networkTransportProtocol property value. The networkTransportProtocol property
func (m *NetworkInfo) SetNetworkTransportProtocol(value *NetworkTransportProtocol)() {
    m.networkTransportProtocol = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NetworkInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPort sets the port property value. Network port number used by media endpoint.
func (m *NetworkInfo) SetPort(value *int32)() {
    m.port = value
}
// SetReceivedQualityEventRatio sets the receivedQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
func (m *NetworkInfo) SetReceivedQualityEventRatio(value *float32)() {
    m.receivedQualityEventRatio = value
}
// SetReflexiveIPAddress sets the reflexiveIPAddress property value. IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
func (m *NetworkInfo) SetReflexiveIPAddress(value *string)() {
    m.reflexiveIPAddress = value
}
// SetRelayIPAddress sets the relayIPAddress property value. IP address of the media relay server allocated by the media endpoint.
func (m *NetworkInfo) SetRelayIPAddress(value *string)() {
    m.relayIPAddress = value
}
// SetRelayPort sets the relayPort property value. Network port number allocated on the media relay server by the media endpoint.
func (m *NetworkInfo) SetRelayPort(value *int32)() {
    m.relayPort = value
}
// SetSentQualityEventRatio sets the sentQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
func (m *NetworkInfo) SetSentQualityEventRatio(value *float32)() {
    m.sentQualityEventRatio = value
}
// SetSubnet sets the subnet property value. Subnet used for media stream by the media endpoint.
func (m *NetworkInfo) SetSubnet(value *string)() {
    m.subnet = value
}
// SetTraceRouteHops sets the traceRouteHops property value. List of network trace route hops collected for this media stream.*
func (m *NetworkInfo) SetTraceRouteHops(value []TraceRouteHopable)() {
    m.traceRouteHops = value
}
// SetWifiBand sets the wifiBand property value. The wifiBand property
func (m *NetworkInfo) SetWifiBand(value *WifiBand)() {
    m.wifiBand = value
}
// SetWifiBatteryCharge sets the wifiBatteryCharge property value. Estimated remaining battery charge in percentage reported by the media endpoint.
func (m *NetworkInfo) SetWifiBatteryCharge(value *int32)() {
    m.wifiBatteryCharge = value
}
// SetWifiChannel sets the wifiChannel property value. WiFi channel used by the media endpoint.
func (m *NetworkInfo) SetWifiChannel(value *int32)() {
    m.wifiChannel = value
}
// SetWifiMicrosoftDriver sets the wifiMicrosoftDriver property value. Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) SetWifiMicrosoftDriver(value *string)() {
    m.wifiMicrosoftDriver = value
}
// SetWifiMicrosoftDriverVersion sets the wifiMicrosoftDriverVersion property value. Version of the Microsoft WiFi driver used by the media endpoint.
func (m *NetworkInfo) SetWifiMicrosoftDriverVersion(value *string)() {
    m.wifiMicrosoftDriverVersion = value
}
// SetWifiRadioType sets the wifiRadioType property value. The wifiRadioType property
func (m *NetworkInfo) SetWifiRadioType(value *WifiRadioType)() {
    m.wifiRadioType = value
}
// SetWifiSignalStrength sets the wifiSignalStrength property value. WiFi signal strength in percentage reported by the media endpoint.
func (m *NetworkInfo) SetWifiSignalStrength(value *int32)() {
    m.wifiSignalStrength = value
}
// SetWifiVendorDriver sets the wifiVendorDriver property value. Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) SetWifiVendorDriver(value *string)() {
    m.wifiVendorDriver = value
}
// SetWifiVendorDriverVersion sets the wifiVendorDriverVersion property value. Version of the WiFi driver used by the media endpoint.
func (m *NetworkInfo) SetWifiVendorDriverVersion(value *string)() {
    m.wifiVendorDriverVersion = value
}
