package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditActor a class containing the properties for Audit Actor.
type AuditActor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the Application.
    applicationDisplayName *string
    // AAD Application Id.
    applicationId *string
    // Actor Type.
    auditActorType *string
    // IPAddress.
    ipAddress *string
    // The OdataType property
    odataType *string
    // Service Principal Name (SPN).
    servicePrincipalName *string
    // User Id.
    userId *string
    // List of user permissions when the audit was performed.
    userPermissions []string
    // User Principal Name (UPN).
    userPrincipalName *string
}
// NewAuditActor instantiates a new auditActor and sets the default values.
func NewAuditActor()(*AuditActor) {
    m := &AuditActor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuditActorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditActorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditActor(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditActor) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplicationDisplayName gets the applicationDisplayName property value. Name of the Application.
func (m *AuditActor) GetApplicationDisplayName()(*string) {
    return m.applicationDisplayName
}
// GetApplicationId gets the applicationId property value. AAD Application Id.
func (m *AuditActor) GetApplicationId()(*string) {
    return m.applicationId
}
// GetAuditActorType gets the auditActorType property value. Actor Type.
func (m *AuditActor) GetAuditActorType()(*string) {
    return m.auditActorType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditActor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplicationDisplayName)
    res["applicationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplicationId)
    res["auditActorType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAuditActorType)
    res["ipAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIpAddress)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["servicePrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetServicePrincipalName)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userPermissions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetUserPermissions)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetIpAddress gets the ipAddress property value. IPAddress.
func (m *AuditActor) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuditActor) GetOdataType()(*string) {
    return m.odataType
}
// GetServicePrincipalName gets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *AuditActor) GetServicePrincipalName()(*string) {
    return m.servicePrincipalName
}
// GetUserId gets the userId property value. User Id.
func (m *AuditActor) GetUserId()(*string) {
    return m.userId
}
// GetUserPermissions gets the userPermissions property value. List of user permissions when the audit was performed.
func (m *AuditActor) GetUserPermissions()([]string) {
    return m.userPermissions
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name (UPN).
func (m *AuditActor) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *AuditActor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationDisplayName", m.GetApplicationDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("auditActorType", m.GetAuditActorType())
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    if m.GetUserPermissions() != nil {
        err := writer.WriteCollectionOfStringValues("userPermissions", m.GetUserPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *AuditActor) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplicationDisplayName sets the applicationDisplayName property value. Name of the Application.
func (m *AuditActor) SetApplicationDisplayName(value *string)() {
    m.applicationDisplayName = value
}
// SetApplicationId sets the applicationId property value. AAD Application Id.
func (m *AuditActor) SetApplicationId(value *string)() {
    m.applicationId = value
}
// SetAuditActorType sets the auditActorType property value. Actor Type.
func (m *AuditActor) SetAuditActorType(value *string)() {
    m.auditActorType = value
}
// SetIpAddress sets the ipAddress property value. IPAddress.
func (m *AuditActor) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuditActor) SetOdataType(value *string)() {
    m.odataType = value
}
// SetServicePrincipalName sets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *AuditActor) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
// SetUserId sets the userId property value. User Id.
func (m *AuditActor) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPermissions sets the userPermissions property value. List of user permissions when the audit was performed.
func (m *AuditActor) SetUserPermissions(value []string)() {
    m.userPermissions = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name (UPN).
func (m *AuditActor) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
