package ui

import "github.com/rivo/tview"

var _messagePanel = BuildMessagePanel()
var _inputPanel = BuildInputPanel()
var _messageLayout = BuildMessageLayout(_messagePanel, _inputPanel)
var _channelFrame, _channelPanel = BuildChannelPanel()
var _userFrame, _userPanel = BuildUserPanel()
var _sidebarLayout = BuildSideBarLayout(_channelFrame, _userFrame)
var _appLayout = BuildAppLayout(_messageLayout, _sidebarLayout)
var _app = tview.NewApplication()

func GetMessagePanel() *tview.List {
	return _messagePanel
}

func GetInputPanel() *tview.InputField {
	return _inputPanel
}

func GetMessageLayout() *tview.Flex {
	return _messageLayout
}

func GetChannelFrame() *tview.Frame {
	return _channelFrame
}

func GetChannelPanel() *tview.List {
	return _channelPanel
}

func GetUserFrame() *tview.Frame {
	return _userFrame
}

func GetUserPanel() *tview.List {
	return _userPanel
}

func GetSideBarLayout() *tview.Flex {
	return _sidebarLayout
}

func GetAppLayout() *tview.Flex {
	return _appLayout
}

func GetApplication() *tview.Application {
	return _app;
}