package constants

const INVALID_FORMAT_ID = "Invalid or Malformed ID format"

var DEFAULT_OWNER_PERMISSIONS = []string{
	"account:read",
	"account:update",
	"account:delete",
	"account:settings",
	"account:billing",
	"account:subscription",
	"users:read",
	"users:create",
	"users:update",
	"users:update_others",
	"users:delete",
	"issues:read",
	"issues:create",
	"issues:update",
	"issues:delete",
	"plans:read",
	"plans:create",
	"plans:update",
	"plans:delete",
	"projects:read",
	"projects:create",
	"projects:update",
	"projects:delete",
	"milestones:read",
	"milestones:create",
	"milestones:update",
	"milestones:delete",
	"roles:read",
	"roles:create",
	"roles:update",
	"roles:delete",
	"service_providers:read",
	"service_providers:create",
	"service_providers:update",
	"service_providers:delete",
	"tasks:read",
	"tasks:create",
	"tasks:update",
	"tasks:delete",
	"teams:read",
	"teams:create",
	"teams:update",
	"teams:delete",
}
