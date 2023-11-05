// Code generated by ent, DO NOT EDIT.

package generated

import (
	"time"

	"github.com/Pyakz/buildbox-api/ent/generated/account"
	"github.com/Pyakz/buildbox-api/ent/generated/issue"
	"github.com/Pyakz/buildbox-api/ent/generated/milestone"
	"github.com/Pyakz/buildbox-api/ent/generated/plan"
	"github.com/Pyakz/buildbox-api/ent/generated/project"
	"github.com/Pyakz/buildbox-api/ent/generated/role"
	"github.com/Pyakz/buildbox-api/ent/generated/serviceprovider"
	"github.com/Pyakz/buildbox-api/ent/generated/serviceproviderprojects"
	"github.com/Pyakz/buildbox-api/ent/generated/subscription"
	"github.com/Pyakz/buildbox-api/ent/generated/task"
	"github.com/Pyakz/buildbox-api/ent/generated/team"
	"github.com/Pyakz/buildbox-api/ent/generated/user"
	"github.com/Pyakz/buildbox-api/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescName is the schema descriptor for name field.
	accountDescName := accountFields[0].Descriptor()
	// account.NameValidator is a validator for the "name" field. It is called by the builders before save.
	account.NameValidator = accountDescName.Validators[0].(func(string) error)
	// accountDescEmail is the schema descriptor for email field.
	accountDescEmail := accountFields[1].Descriptor()
	// account.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	account.EmailValidator = accountDescEmail.Validators[0].(func(string) error)
	// accountDescPhoneNumber is the schema descriptor for phone_number field.
	accountDescPhoneNumber := accountFields[2].Descriptor()
	// account.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	account.PhoneNumberValidator = accountDescPhoneNumber.Validators[0].(func(string) error)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[3].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[4].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUUID is the schema descriptor for uuid field.
	accountDescUUID := accountFields[5].Descriptor()
	// account.DefaultUUID holds the default value on creation for the uuid field.
	account.DefaultUUID = accountDescUUID.Default.(func() uuid.UUID)
	issueFields := schema.Issue{}.Fields()
	_ = issueFields
	// issueDescTitle is the schema descriptor for title field.
	issueDescTitle := issueFields[3].Descriptor()
	// issue.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	issue.TitleValidator = issueDescTitle.Validators[0].(func(string) error)
	// issueDescDescription is the schema descriptor for description field.
	issueDescDescription := issueFields[4].Descriptor()
	// issue.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	issue.DescriptionValidator = issueDescDescription.Validators[0].(func(string) error)
	// issueDescUpdatedAt is the schema descriptor for updated_at field.
	issueDescUpdatedAt := issueFields[5].Descriptor()
	// issue.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	issue.DefaultUpdatedAt = issueDescUpdatedAt.Default.(func() time.Time)
	// issueDescCreatedAt is the schema descriptor for created_at field.
	issueDescCreatedAt := issueFields[6].Descriptor()
	// issue.DefaultCreatedAt holds the default value on creation for the created_at field.
	issue.DefaultCreatedAt = issueDescCreatedAt.Default.(func() time.Time)
	// issueDescDeleted is the schema descriptor for deleted field.
	issueDescDeleted := issueFields[7].Descriptor()
	// issue.DefaultDeleted holds the default value on creation for the deleted field.
	issue.DefaultDeleted = issueDescDeleted.Default.(bool)
	// issueDescUUID is the schema descriptor for uuid field.
	issueDescUUID := issueFields[8].Descriptor()
	// issue.DefaultUUID holds the default value on creation for the uuid field.
	issue.DefaultUUID = issueDescUUID.Default.(func() uuid.UUID)
	milestoneFields := schema.Milestone{}.Fields()
	_ = milestoneFields
	// milestoneDescTitle is the schema descriptor for title field.
	milestoneDescTitle := milestoneFields[3].Descriptor()
	// milestone.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	milestone.TitleValidator = milestoneDescTitle.Validators[0].(func(string) error)
	// milestoneDescDescription is the schema descriptor for description field.
	milestoneDescDescription := milestoneFields[4].Descriptor()
	// milestone.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	milestone.DescriptionValidator = milestoneDescDescription.Validators[0].(func(string) error)
	// milestoneDescUpdatedAt is the schema descriptor for updated_at field.
	milestoneDescUpdatedAt := milestoneFields[6].Descriptor()
	// milestone.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	milestone.DefaultUpdatedAt = milestoneDescUpdatedAt.Default.(func() time.Time)
	// milestoneDescCreatedAt is the schema descriptor for created_at field.
	milestoneDescCreatedAt := milestoneFields[7].Descriptor()
	// milestone.DefaultCreatedAt holds the default value on creation for the created_at field.
	milestone.DefaultCreatedAt = milestoneDescCreatedAt.Default.(func() time.Time)
	// milestoneDescDeleted is the schema descriptor for deleted field.
	milestoneDescDeleted := milestoneFields[8].Descriptor()
	// milestone.DefaultDeleted holds the default value on creation for the deleted field.
	milestone.DefaultDeleted = milestoneDescDeleted.Default.(bool)
	// milestoneDescUUID is the schema descriptor for uuid field.
	milestoneDescUUID := milestoneFields[9].Descriptor()
	// milestone.DefaultUUID holds the default value on creation for the uuid field.
	milestone.DefaultUUID = milestoneDescUUID.Default.(func() uuid.UUID)
	planFields := schema.Plan{}.Fields()
	_ = planFields
	// planDescName is the schema descriptor for name field.
	planDescName := planFields[0].Descriptor()
	// plan.NameValidator is a validator for the "name" field. It is called by the builders before save.
	plan.NameValidator = planDescName.Validators[0].(func(string) error)
	// planDescPrice is the schema descriptor for price field.
	planDescPrice := planFields[2].Descriptor()
	// plan.DefaultPrice holds the default value on creation for the price field.
	plan.DefaultPrice = planDescPrice.Default.(float64)
	// planDescUpdatedAt is the schema descriptor for updated_at field.
	planDescUpdatedAt := planFields[3].Descriptor()
	// plan.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	plan.DefaultUpdatedAt = planDescUpdatedAt.Default.(func() time.Time)
	// planDescCreatedAt is the schema descriptor for created_at field.
	planDescCreatedAt := planFields[4].Descriptor()
	// plan.DefaultCreatedAt holds the default value on creation for the created_at field.
	plan.DefaultCreatedAt = planDescCreatedAt.Default.(func() time.Time)
	// planDescUUID is the schema descriptor for uuid field.
	planDescUUID := planFields[5].Descriptor()
	// plan.DefaultUUID holds the default value on creation for the uuid field.
	plan.DefaultUUID = planDescUUID.Default.(func() uuid.UUID)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescName is the schema descriptor for name field.
	projectDescName := projectFields[4].Descriptor()
	// project.NameValidator is a validator for the "name" field. It is called by the builders before save.
	project.NameValidator = projectDescName.Validators[0].(func(string) error)
	// projectDescBudget is the schema descriptor for budget field.
	projectDescBudget := projectFields[7].Descriptor()
	// project.BudgetValidator is a validator for the "budget" field. It is called by the builders before save.
	project.BudgetValidator = projectDescBudget.Validators[0].(func(float64) error)
	// projectDescStartDate is the schema descriptor for start_date field.
	projectDescStartDate := projectFields[10].Descriptor()
	// project.DefaultStartDate holds the default value on creation for the start_date field.
	project.DefaultStartDate = projectDescStartDate.Default.(func() time.Time)
	// projectDescUUID is the schema descriptor for uuid field.
	projectDescUUID := projectFields[12].Descriptor()
	// project.DefaultUUID holds the default value on creation for the uuid field.
	project.DefaultUUID = projectDescUUID.Default.(func() uuid.UUID)
	// projectDescDeleted is the schema descriptor for deleted field.
	projectDescDeleted := projectFields[13].Descriptor()
	// project.DefaultDeleted holds the default value on creation for the deleted field.
	project.DefaultDeleted = projectDescDeleted.Default.(bool)
	// projectDescUpdatedAt is the schema descriptor for updated_at field.
	projectDescUpdatedAt := projectFields[14].Descriptor()
	// project.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	project.DefaultUpdatedAt = projectDescUpdatedAt.Default.(func() time.Time)
	// projectDescCreatedAt is the schema descriptor for created_at field.
	projectDescCreatedAt := projectFields[15].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the created_at field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleFields[4].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[5].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUUID is the schema descriptor for uuid field.
	roleDescUUID := roleFields[7].Descriptor()
	// role.DefaultUUID holds the default value on creation for the uuid field.
	role.DefaultUUID = roleDescUUID.Default.(func() uuid.UUID)
	serviceproviderFields := schema.ServiceProvider{}.Fields()
	_ = serviceproviderFields
	// serviceproviderDescName is the schema descriptor for name field.
	serviceproviderDescName := serviceproviderFields[2].Descriptor()
	// serviceprovider.NameValidator is a validator for the "name" field. It is called by the builders before save.
	serviceprovider.NameValidator = serviceproviderDescName.Validators[0].(func(string) error)
	// serviceproviderDescUpdatedAt is the schema descriptor for updated_at field.
	serviceproviderDescUpdatedAt := serviceproviderFields[7].Descriptor()
	// serviceprovider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	serviceprovider.DefaultUpdatedAt = serviceproviderDescUpdatedAt.Default.(func() time.Time)
	// serviceproviderDescCreatedAt is the schema descriptor for created_at field.
	serviceproviderDescCreatedAt := serviceproviderFields[8].Descriptor()
	// serviceprovider.DefaultCreatedAt holds the default value on creation for the created_at field.
	serviceprovider.DefaultCreatedAt = serviceproviderDescCreatedAt.Default.(func() time.Time)
	// serviceproviderDescUUID is the schema descriptor for uuid field.
	serviceproviderDescUUID := serviceproviderFields[9].Descriptor()
	// serviceprovider.DefaultUUID holds the default value on creation for the uuid field.
	serviceprovider.DefaultUUID = serviceproviderDescUUID.Default.(func() uuid.UUID)
	serviceproviderprojectsFields := schema.ServiceProviderProjects{}.Fields()
	_ = serviceproviderprojectsFields
	// serviceproviderprojectsDescUpdatedAt is the schema descriptor for updated_at field.
	serviceproviderprojectsDescUpdatedAt := serviceproviderprojectsFields[3].Descriptor()
	// serviceproviderprojects.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	serviceproviderprojects.DefaultUpdatedAt = serviceproviderprojectsDescUpdatedAt.Default.(func() time.Time)
	// serviceproviderprojectsDescCreatedAt is the schema descriptor for created_at field.
	serviceproviderprojectsDescCreatedAt := serviceproviderprojectsFields[4].Descriptor()
	// serviceproviderprojects.DefaultCreatedAt holds the default value on creation for the created_at field.
	serviceproviderprojects.DefaultCreatedAt = serviceproviderprojectsDescCreatedAt.Default.(func() time.Time)
	// serviceproviderprojectsDescUUID is the schema descriptor for uuid field.
	serviceproviderprojectsDescUUID := serviceproviderprojectsFields[5].Descriptor()
	// serviceproviderprojects.DefaultUUID holds the default value on creation for the uuid field.
	serviceproviderprojects.DefaultUUID = serviceproviderprojectsDescUUID.Default.(func() uuid.UUID)
	subscriptionFields := schema.Subscription{}.Fields()
	_ = subscriptionFields
	// subscriptionDescStartDate is the schema descriptor for start_date field.
	subscriptionDescStartDate := subscriptionFields[2].Descriptor()
	// subscription.DefaultStartDate holds the default value on creation for the start_date field.
	subscription.DefaultStartDate = subscriptionDescStartDate.Default.(func() time.Time)
	// subscriptionDescUpdatedAt is the schema descriptor for updated_at field.
	subscriptionDescUpdatedAt := subscriptionFields[7].Descriptor()
	// subscription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	subscription.DefaultUpdatedAt = subscriptionDescUpdatedAt.Default.(func() time.Time)
	// subscriptionDescCreatedAt is the schema descriptor for created_at field.
	subscriptionDescCreatedAt := subscriptionFields[8].Descriptor()
	// subscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	subscription.DefaultCreatedAt = subscriptionDescCreatedAt.Default.(func() time.Time)
	// subscriptionDescUUID is the schema descriptor for uuid field.
	subscriptionDescUUID := subscriptionFields[9].Descriptor()
	// subscription.DefaultUUID holds the default value on creation for the uuid field.
	subscription.DefaultUUID = subscriptionDescUUID.Default.(func() uuid.UUID)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescTitle is the schema descriptor for title field.
	taskDescTitle := taskFields[5].Descriptor()
	// task.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	task.TitleValidator = taskDescTitle.Validators[0].(func(string) error)
	// taskDescDescription is the schema descriptor for description field.
	taskDescDescription := taskFields[6].Descriptor()
	// task.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	task.DescriptionValidator = taskDescDescription.Validators[0].(func(string) error)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskFields[7].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[8].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescDeleted is the schema descriptor for deleted field.
	taskDescDeleted := taskFields[9].Descriptor()
	// task.DefaultDeleted holds the default value on creation for the deleted field.
	task.DefaultDeleted = taskDescDeleted.Default.(bool)
	// taskDescUUID is the schema descriptor for uuid field.
	taskDescUUID := taskFields[10].Descriptor()
	// task.DefaultUUID holds the default value on creation for the uuid field.
	task.DefaultUUID = taskDescUUID.Default.(func() uuid.UUID)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescName is the schema descriptor for name field.
	teamDescName := teamFields[2].Descriptor()
	// team.NameValidator is a validator for the "name" field. It is called by the builders before save.
	team.NameValidator = teamDescName.Validators[0].(func(string) error)
	// teamDescUpdatedAt is the schema descriptor for updated_at field.
	teamDescUpdatedAt := teamFields[5].Descriptor()
	// team.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	team.DefaultUpdatedAt = teamDescUpdatedAt.Default.(func() time.Time)
	// teamDescCreatedAt is the schema descriptor for created_at field.
	teamDescCreatedAt := teamFields[6].Descriptor()
	// team.DefaultCreatedAt holds the default value on creation for the created_at field.
	team.DefaultCreatedAt = teamDescCreatedAt.Default.(func() time.Time)
	// teamDescUUID is the schema descriptor for uuid field.
	teamDescUUID := teamFields[7].Descriptor()
	// team.DefaultUUID holds the default value on creation for the uuid field.
	team.DefaultUUID = teamDescUUID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[9].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[10].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[11].Descriptor()
	// user.DefaultUUID holds the default value on creation for the uuid field.
	user.DefaultUUID = userDescUUID.Default.(func() uuid.UUID)
}
