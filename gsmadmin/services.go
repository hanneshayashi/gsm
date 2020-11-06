/*
Copyright © 2020 Hannes Hayashi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package gsmadmin

import (
	"context"
	"log"
	"net/http"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

var (
	client                    *http.Client
	adminService              *admin.Service
	customersService          *admin.CustomersService
	usersService              *admin.UsersService
	groupsService             *admin.GroupsService
	membersService            *admin.MembersService
	orgunitsService           *admin.OrgunitsService
	rolesService              *admin.RolesService
	rolesAssignmentsService   *admin.RoleAssignmentsService
	privilegesService         *admin.PrivilegesService
	verificationCodesService  *admin.VerificationCodesService
	usersAliasesService       *admin.UsersAliasesService
	groupsAliasesService      *admin.GroupsAliasesService
	tokensService             *admin.TokensService
	aspsService               *admin.AspsService
	domainsService            *admin.DomainsService
	domainAliasesService      *admin.DomainAliasesService
	mobiledevicesService      *admin.MobiledevicesService
	chromeosdevicesService    *admin.ChromeosdevicesService
	resourcesBuildingsService *admin.ResourcesBuildingsService
	resourcesCalendarsService *admin.ResourcesCalendarsService
	resourcesFeaturesService  *admin.ResourcesFeaturesService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getAdminService() *admin.Service {
	if client == nil {
		log.Fatalf("gsmadmin.client is not set. Set with gsmadmin.SetClient(client)")
	}
	if adminService == nil {
		var err error
		adminService, err = admin.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating admin service: %v", err)
		}
	}
	return adminService
}

func getCustomersService() *admin.CustomersService {
	if customersService == nil {
		customersService = admin.NewCustomersService(getAdminService())
	}
	return customersService
}

func getUsersService() *admin.UsersService {
	if usersService == nil {
		usersService = admin.NewUsersService(getAdminService())
	}
	return usersService
}

func getGroupsService() *admin.GroupsService {
	if groupsService == nil {
		groupsService = admin.NewGroupsService(getAdminService())
	}
	return groupsService
}

func getMembersService() *admin.MembersService {
	if membersService == nil {
		membersService = admin.NewMembersService(getAdminService())
	}
	return membersService
}

func getOrgunitsService() *admin.OrgunitsService {
	if orgunitsService == nil {
		orgunitsService = admin.NewOrgunitsService(getAdminService())
	}
	return orgunitsService
}

func getRolesService() *admin.RolesService {
	if rolesService == nil {
		rolesService = admin.NewRolesService(getAdminService())
	}
	return rolesService
}

func getRoleAssignmentsService() *admin.RoleAssignmentsService {
	if rolesAssignmentsService == nil {
		rolesAssignmentsService = admin.NewRoleAssignmentsService(getAdminService())
	}
	return rolesAssignmentsService
}

func getPrivilegesService() *admin.PrivilegesService {
	if privilegesService == nil {
		privilegesService = admin.NewPrivilegesService(getAdminService())
	}
	return privilegesService
}

func getVerificationCodesService() *admin.VerificationCodesService {
	if verificationCodesService == nil {
		verificationCodesService = admin.NewVerificationCodesService(getAdminService())
	}
	return verificationCodesService
}

func getUsersAliasesService() (usersAliasesService *admin.UsersAliasesService) {
	if usersAliasesService == nil {
		usersAliasesService = admin.NewUsersAliasesService(getAdminService())
	}
	return usersAliasesService
}

func getGroupsAliasesService() *admin.GroupsAliasesService {
	if groupsAliasesService == nil {
		groupsAliasesService = admin.NewGroupsAliasesService(getAdminService())
	}
	return groupsAliasesService
}

func getTokensService() *admin.TokensService {
	if tokensService == nil {
		tokensService = admin.NewTokensService(getAdminService())
	}
	return tokensService
}

func getAspsService() *admin.AspsService {
	if aspsService == nil {
		aspsService = admin.NewAspsService(getAdminService())
	}
	return aspsService
}

func getDomainsService() *admin.DomainsService {
	if domainsService == nil {
		domainsService = admin.NewDomainsService(getAdminService())
	}
	return domainsService
}

func getDomainAliasesService() *admin.DomainAliasesService {
	if domainAliasesService == nil {
		domainAliasesService = admin.NewDomainAliasesService(getAdminService())
	}
	return domainAliasesService
}

func getMobiledevicesService() *admin.MobiledevicesService {
	if mobiledevicesService == nil {
		mobiledevicesService = admin.NewMobiledevicesService(getAdminService())
	}
	return mobiledevicesService
}

func getChromeosdevicesService() *admin.ChromeosdevicesService {
	if chromeosdevicesService == nil {
		chromeosdevicesService = admin.NewChromeosdevicesService(getAdminService())
	}
	return chromeosdevicesService
}

func getResourcesBuildingsService() *admin.ResourcesBuildingsService {
	if resourcesBuildingsService == nil {
		resourcesBuildingsService = admin.NewResourcesBuildingsService(getAdminService())
	}
	return resourcesBuildingsService
}

func getResourcesCalendarsService() *admin.ResourcesCalendarsService {
	if resourcesCalendarsService == nil {
		resourcesCalendarsService = admin.NewResourcesCalendarsService(getAdminService())
	}
	return resourcesCalendarsService
}

func getResourcesFeaturesService() *admin.ResourcesFeaturesService {
	if resourcesFeaturesService == nil {
		resourcesFeaturesService = admin.NewResourcesFeaturesService(getAdminService())
	}
	return resourcesFeaturesService
}
