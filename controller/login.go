package controller

import (
	"fmt"
	"net/http"

	"github.com/Dunkelheit/feedbackapp/database"
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/Dunkelheit/feedbackapp/util"

	"gopkg.in/gin-gonic/gin.v1"
	ldap "gopkg.in/ldap.v2"
)

const (
	ldapHost               = "10.41.100.152"
	ldapPort               = 389
	ldapAttrMail           = "mail"
	ldapAttrName           = "name"
	ldapAttrGivenName      = "givenName"
	ldapAttrSn             = "sn"
	ldapAttrTitle          = "title"
	ldapAttrDepartment     = "department"
	ldapAttrCompany        = "company"
	ldapAttrSAMAccountName = "sAMAccountName"
	ldapBaseDN             = "ou=Icemobile,dc=brandloyaltyint,dc=corp"
	ldapFilterAllUsers     = "(&(objectClass=user))"
	ldapFilterSingleUsers  = "(&(objectClass=user)(sAMAccountName=%s))"
	ldapUsername           = "%s@brandloyaltyint.corp"
)

type login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var ldapAttributes = []string{ldapAttrMail, ldapAttrName, ldapAttrGivenName, ldapAttrSn, ldapAttrTitle, ldapAttrDepartment, ldapAttrCompany, ldapAttrSAMAccountName}

func buildSearchRequest(filter string) *ldap.SearchRequest {
	return ldap.NewSearchRequest(ldapBaseDN, ldap.ScopeWholeSubtree, ldap.DerefAlways,
		0, 0, false, filter, ldapAttributes, nil)
}

func filterEntriesWithoutMail(entries []*ldap.Entry) []*ldap.Entry {
	result := make([]*ldap.Entry, 0)
	for _, entry := range entries {
		if mail := entry.GetRawAttributeValue(ldapAttrMail); len(mail) > 0 {
			result = append(result, entry)
		}
	}
	return result
}

func getLDAPUsers(l *ldap.Conn) ([]*ldap.Entry, error) {
	searchRequest := buildSearchRequest(ldapFilterAllUsers)

	sr, err := l.Search(searchRequest)

	if err != nil {
		return nil, err
	}

	return filterEntriesWithoutMail(sr.Entries), nil
}

func entryToUser(user *ldap.Entry) *model.User {
	return &model.User{
		Username:   user.GetAttributeValue(ldapAttrSAMAccountName),
		FirstName:  user.GetAttributeValue(ldapAttrGivenName),
		Surname:    user.GetAttributeValue(ldapAttrSn),
		FullName:   user.GetAttributeValue(ldapAttrName),
		JobTitle:   user.GetAttributeValue(ldapAttrTitle),
		Department: user.GetAttributeValue(ldapAttrDepartment),
		Company:    user.GetAttributeValue(ldapAttrCompany),
		Email:      user.GetAttributeValue(ldapAttrMail),
	}
}

func preloadLDAPUsers(l *ldap.Conn) (int, error) {
	users, err := getLDAPUsers(l)
	if err != nil {
		return 0, err
	}
	for _, user := range users {
		database.DB.Create(entryToUser(user))
	}
	return len(users), nil
}

// Login using LDAP
func Login(c *gin.Context) {
	in := &login{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	err = l.Bind(fmt.Sprintf(ldapUsername, in.Username), in.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	var currentUsers int64
	database.DB.Model(&model.User{}).Count(&currentUsers)

	if currentUsers == 0 {
		go func() {
			defer l.Close()
			newUsers, err := preloadLDAPUsers(l)
			fmt.Println(fmt.Sprintf("New users: %d", newUsers))
			if err != nil {
				fmt.Println("Error!")
				fmt.Println(err)
			}
		}()
	} else {
		defer l.Close()
	}

	searchResult, err := l.Search(buildSearchRequest(fmt.Sprintf(ldapFilterSingleUsers, in.Username)))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if foundUsers := len(searchResult.Entries); foundUsers != 1 {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Unexpected amount of users found (%d)", foundUsers))
		return
	}

	user := entryToUser(searchResult.Entries[0])

	tokenString, err := util.EncodeToken(user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error setting the authentication header")
		return
	}

	c.Writer.Header().Set("x-auth-token", tokenString)
	c.JSON(http.StatusOK, user)
}
