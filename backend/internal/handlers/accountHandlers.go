package handlers

import "github.com/gin-gonic/gin"

func GetAllAccounts(c *gin.Context)           {}
func GetAccountPublicDetails(c *gin.Context)  {}
func GetAccountPrivateDetails(c *gin.Context) {}

func CreateAccount(c *gin.Context) {}
func UpdateAccountById(c *gin.Context)   {}

func GenerateAccountToken(c *gin.Context)  {}
func RestoreAccountById(c *gin.Context) {}

func DeleteAccountById(c *gin.Context)     {}
func SoftDeleteAccountById(c *gin.Context) {}
