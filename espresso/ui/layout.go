package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildAppLayout(messageLayout *tview.Flex, sidebarLayout *tview.Flex) *tview.Flex {
	appLayout := tview.NewFlex()
	appLayout.AddItem(messageLayout, 0, 1, false)
	appLayout.AddItem(sidebarLayout, 20, 1, false)
	return appLayout
}

func BuildMessageLayout(messagePanel *tview.List, inputPanel *tview.InputField) *tview.Flex {
	messageLayout := tview.NewFlex()
	messageLayout.SetDirection(tview.FlexRow)
	messageLayout.AddItem(messagePanel, 0, 1, false)
	messageLayout.AddItem(inputPanel, 2, 1, true)
	return messageLayout
}

func BuildMessagePanel() *tview.List {
	messagePanel := tview.NewList()
	messagePanel.ShowSecondaryText(false)
	messagePanel.SetBorderPadding(1, 1, 0, 0)
	messagePanel.SetWrapAround(false)
	messagePanel.SetSelectedFocusOnly(true)
	return messagePanel
}

func BuildInputPanel() *tview.InputField {
	inputPanel := tview.NewInputField()
	inputPanel.SetPlaceholder("Type message here...")
	inputPanel.SetPlaceholderStyle(inputPanel.GetPlaceholderStyle().Background(tcell.ColorDarkSeaGreen))
	inputPanel.SetPlaceholderTextColor(tcell.ColorWhite)
	inputPanel.SetFieldBackgroundColor(tcell.ColorDarkSeaGreen)
	return inputPanel
}

func BuildSideBarLayout(channelPanel *tview.Frame, userPanel *tview.Frame) *tview.Flex {
	sidebarLayout := tview.NewFlex()
	sidebarLayout.SetDirection(tview.FlexRow)
	sidebarLayout.AddItem(channelPanel, 0, 1, false)
	sidebarLayout.AddItem(userPanel, 0, 1, false)
	return sidebarLayout
}

func BuildUserPanel() (*tview.Frame, *tview.List) {
	userPanel := tview.NewList()
	userPanel.ShowSecondaryText(false)
	userPanel.SetBorderPadding(0, 0, 2, 0)
	frame := tview.NewFrame(userPanel)
	frame.AddText("[::b]Online Users", true, tview.AlignLeft, tcell.ColorDarkSeaGreen)
	return frame, userPanel
}

func BuildChannelPanel() (*tview.Frame, *tview.List) {
	channelPanel := tview.NewList()
	channelPanel.ShowSecondaryText(false)
	channelPanel.SetBorderPadding(0, 0, 2, 0)
	frame := tview.NewFrame(channelPanel)
	frame.AddText("[::b]Channels", true, tview.AlignLeft, tcell.ColorDarkSeaGreen)
	return frame, channelPanel
}
