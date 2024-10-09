package ldapservice

import (
	"context"
	"os"
	"strings"

	"github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/sets"
	"github.com/go-kit/log/level"
	goldap "github.com/go-ldap/ldap/v3"

	"github.com/MicroOps-cn/fuck-web/pkg/client/ldap"
)

const (
	ClassIdasCore         = "fuckWebCore"
	ClassIdasApp          = "fuckWebApp"
	ClassExtensibleObject = "extensibleObject"
)

func NewUserAndAppService(ctx context.Context, name string, client *ldap.Client) *UserService {
	classes, err := ldap.GetAvailableObjectClass(client.Session(ctx))
	if err != nil {
		return nil
	}
	var hasIDASClass bool
	if classes.HasAll(ClassIdasCore, ClassIdasApp) {
		hasIDASClass = true
	} else if !classes.Has(ClassExtensibleObject) {
		level.Error(log.GetContextLogger(ctx)).Log("msg", "If you want to run the platform normally, you need `extensibleObject` or fuck-web related classes")
		os.Exit(1)
	}
	return &UserService{name: name, Client: client, hasIDASClass: hasIDASClass}
}

type UserService struct {
	*ldap.Client
	name         string
	hasIDASClass bool
	memberAttr   string
}

func (s UserService) GetUserClass() sets.Set[string] {
	if s.hasIDASClass {
		return sets.New[string](ClassIdasCore)
	}
	return sets.New[string](ClassExtensibleObject)
}

func (s UserService) GetMemberAttr() string {
	return s.memberAttr
}

func (s UserService) Name() string {
	return s.name
}

func (s UserService) AutoCreateOrganizationalUnit(ctx context.Context, name string) error {
	session := s.Session(ctx)
	_, err := session.Search(goldap.NewSearchRequest(
		name,
		goldap.ScopeBaseObject, goldap.NeverDerefAliases, 1, 0, false,
		"(objectClass=*)",
		nil,
		nil))
	if ldap.IsLdapError(err, 32) {
		if _, suffix, found := strings.Cut(name, ","); found && len(suffix) > 0 {
			if err = s.AutoCreateOrganizationalUnit(ctx, suffix); err != nil {
				return err
			}
		}
		addReq := goldap.NewAddRequest(name, nil)
		addReq.Attributes = append(addReq.Attributes, goldap.Attribute{
			Type: "objectClass", Vals: []string{"top", "organizationalUnit"},
		})
		err = session.Add(addReq)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s UserService) AutoMigrate(ctx context.Context) error {
	return nil
}

type ldapUpdateColumn struct {
	columnName     string
	ldapColumnName string
	val            []string
	oriVals        []string
}
