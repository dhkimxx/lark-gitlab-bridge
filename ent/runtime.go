// Code generated by ent, DO NOT EDIT.

package ent

import (
	"lark-gitlab-bridge/ent/schema"
	"lark-gitlab-bridge/ent/user"
	"lark-gitlab-bridge/ent/webhook"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	webhookFields := schema.Webhook{}.Fields()
	_ = webhookFields
	// webhookDescName is the schema descriptor for name field.
	webhookDescName := webhookFields[1].Descriptor()
	// webhook.NameValidator is a validator for the "name" field. It is called by the builders before save.
	webhook.NameValidator = webhookDescName.Validators[0].(func(string) error)
	// webhookDescURL is the schema descriptor for url field.
	webhookDescURL := webhookFields[2].Descriptor()
	// webhook.URLValidator is a validator for the "url" field. It is called by the builders before save.
	webhook.URLValidator = webhookDescURL.Validators[0].(func(string) error)
	// webhookDescRequiredVerification is the schema descriptor for required_verification field.
	webhookDescRequiredVerification := webhookFields[3].Descriptor()
	// webhook.DefaultRequiredVerification holds the default value on creation for the required_verification field.
	webhook.DefaultRequiredVerification = webhookDescRequiredVerification.Default.(bool)
	// webhookDescCreatedAt is the schema descriptor for created_at field.
	webhookDescCreatedAt := webhookFields[5].Descriptor()
	// webhook.DefaultCreatedAt holds the default value on creation for the created_at field.
	webhook.DefaultCreatedAt = webhookDescCreatedAt.Default.(func() time.Time)
	// webhookDescUpdatedAt is the schema descriptor for updated_at field.
	webhookDescUpdatedAt := webhookFields[6].Descriptor()
	// webhook.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	webhook.DefaultUpdatedAt = webhookDescUpdatedAt.Default.(func() time.Time)
	// webhook.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	webhook.UpdateDefaultUpdatedAt = webhookDescUpdatedAt.UpdateDefault.(func() time.Time)
	// webhookDescID is the schema descriptor for id field.
	webhookDescID := webhookFields[0].Descriptor()
	// webhook.DefaultID holds the default value on creation for the id field.
	webhook.DefaultID = webhookDescID.Default.(func() uuid.UUID)
}