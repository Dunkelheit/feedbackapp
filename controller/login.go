package controller

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	ldap "gopkg.in/ldap.v2"
)

// Login using LDAP
func Login(c *gin.Context) {
	type Login struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	in := &Login{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "10.41.100.152", 389))
	defer l.Close()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	err = l.Bind(in.Username+"@brandloyaltyint.corp", in.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	attributes := []string{"mail", "name", "givenName", "sn", "title", "department", "company", "sAMAccountName"}
	searchRequest := ldap.NewSearchRequest("ou=Icemobile,dc=brandloyaltyint,dc=corp", ldap.ScopeWholeSubtree, ldap.DerefAlways,
		0, 0, false, "(&(objectClass=user)(sAMAccountName="+in.Username+"))", attributes, nil)
	sr, err := l.Search(searchRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	for _, elem := range sr.Entries {
		for _, at := range attributes {
			fmt.Println(at + ": " + elem.GetAttributeValue(at))
		}
	}

	c.JSON(http.StatusOK, sr)
}
