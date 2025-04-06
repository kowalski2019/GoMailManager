package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kowaslki2019/mailmanager/cmd"
	"github.com/kowaslki2019/mailmanager/internal/logger"
	"github.com/kowaslki2019/mailmanager/internal/models"
	"github.com/kowaslki2019/mailmanager/internal/util"
)

func Health(c *gin.Context) {
	res, err := cmd.ExecuteCommand("ls", "-la")
	if err != nil {
		logger.Error("Error by executeCommand: ", err.Error())
	}
	logger.Debug("Result: ", res)
	util.SendOkResponse(c, "Everything is fine")
}

func HandleAutoRelply(c *gin.Context) {
	var req models.AutoReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendBadRequestResponse(c, "Invalid request body"+err.Error())
		return
	}

	deploymentMode := c.GetHeader("Deployment-Mode")
	logger.Debug("Deployment Mode: ", deploymentMode)

	out, err := cmd.ExecuteCommand("mailmanager_handle_autoreply_request", req.Name, req.Email, req.Subject, req.Message, strconv.FormatBool(req.Enabled), req.StartDate, req.EndDate, deploymentMode)
	if err != nil {
		logger.Error("Error by executeCommand: ", err.Error())
	}
	logger.Debug("Execute command result: ", out)
	util.SendOkResponse(c, out)
}

func HandleRules(c *gin.Context) {
	var req models.EmailRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendBadRequestResponse(c, "Invalid request body"+err.Error())
		return
	}
	out, err := cmd.ExecuteCommand("mailmanager_handle_rules_request", strconv.FormatInt(req.Id, 10), req.UserEmail, req.Name, req.Condition, req.ConditionValue, req.Action, req.ActionValue, req.Description)
	if err != nil {
		logger.Error("Error by executeCommand: ", err.Error())
	}
	logger.Debug("Execute command result: ", out)
	util.SendOkResponse(c, out)
}

func GetMailboxes(c *gin.Context) {
	username := c.Param("username")
	unameFull := username + "@industryinmotion.de"
	mailboxes, err := cmd.ExecuteCommand("doveadm", "mailbox", "list", "-u", unameFull)
	if err != nil {
		logger.Error("Error by executeCommand: ", err.Error())
	}
	responseObject := gin.H{
		"mailboxes": mailboxes,
	}
	util.SendJsonResponse(c, 200, responseObject)
}

func DeleteRule(c *gin.Context) {
	ruleId := c.Param("rule_id")
	username := c.Param("username")
	email := username + "@industryinmotion.de"
	out, err := cmd.ExecuteCommand("mailmanager_delete_rule", ruleId, email)
	if err != nil {
		logger.Error("Error by executeCommand: ", err.Error())
	}
	logger.Debug("Execute command result: ", out)
	util.SendOkResponse(c, out)
}
