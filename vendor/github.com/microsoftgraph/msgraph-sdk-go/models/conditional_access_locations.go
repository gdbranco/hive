package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessLocations 
type ConditionalAccessLocations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Location IDs excluded from scope of policy.
    excludeLocations []string
    // Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
    includeLocations []string
    // The OdataType property
    odataType *string
}
// NewConditionalAccessLocations instantiates a new conditionalAccessLocations and sets the default values.
func NewConditionalAccessLocations()(*ConditionalAccessLocations) {
    m := &ConditionalAccessLocations{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessLocationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessLocationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessLocations(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessLocations) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExcludeLocations gets the excludeLocations property value. Location IDs excluded from scope of policy.
func (m *ConditionalAccessLocations) GetExcludeLocations()([]string) {
    return m.excludeLocations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessLocations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeLocations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetExcludeLocations)
    res["includeLocations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetIncludeLocations)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIncludeLocations gets the includeLocations property value. Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
func (m *ConditionalAccessLocations) GetIncludeLocations()([]string) {
    return m.includeLocations
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessLocations) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ConditionalAccessLocations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeLocations() != nil {
        err := writer.WriteCollectionOfStringValues("excludeLocations", m.GetExcludeLocations())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeLocations() != nil {
        err := writer.WriteCollectionOfStringValues("includeLocations", m.GetIncludeLocations())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessLocations) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExcludeLocations sets the excludeLocations property value. Location IDs excluded from scope of policy.
func (m *ConditionalAccessLocations) SetExcludeLocations(value []string)() {
    m.excludeLocations = value
}
// SetIncludeLocations sets the includeLocations property value. Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
func (m *ConditionalAccessLocations) SetIncludeLocations(value []string)() {
    m.includeLocations = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessLocations) SetOdataType(value *string)() {
    m.odataType = value
}
