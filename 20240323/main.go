package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {
	a := app.New()

	desk, ok := a.(desktop.App)
	if !ok {
		return
	}

	var menuItems []*fyne.MenuItem
	menuItems = append(
		menuItems,
		fyne.NewMenuItem("YES", nil),
		fyne.NewMenuItem("NO", nil),
		fyne.NewMenuItemSeparator(),
	)
	menuItems = append(menuItems, MakeUtilsMenuItems("JsonUtils", []string{"CMD+SHIFT+E", "SHIFT+CTRL+E"})...)
	menuItems = append(menuItems, MakeUtilsMenuItems("自动生成Markdown超链接", []string{"CMD+SHIFT+K", "SHIFT+CTRL+K"})...)
	menuItems = append(menuItems, MakeUtilsMenuItems("根据Markdown表格生成Java实体类", []string{"CMD+SHIFT+CTRL+J"})...)
	menuItems = append(menuItems, MakeUtilsMenuItems("根据Java实体类生成Markdown表格", []string{"CMD+SHIFT+CTRL+G"})...)
	menuItems = append(menuItems, MakeUtilsMenuItems("生成Json可用的转义字符串", []string{"CMD+SHIFT+CTRL+E"})...)
	scriptsMenu := fyne.NewMenuItem("Scripts", nil)
	scriptsMenu.ChildMenu = fyne.NewMenu("scripts",
		fyne.NewMenuItem("File1.gogogo", nil),
		fyne.NewMenuItem("File2.gogogo", nil),
		fyne.NewMenuItem("File3.gogogo", nil),
	)
	menuItems = append(menuItems, scriptsMenu)
	menuItems = append(menuItems, MakeUtilsMenuItems("CustomWild", []string{"CMD+SHIFT+X"})...)

	m := fyne.NewMenu("ClipboardUtils-20240323",
		menuItems...,
	)
	desk.SetSystemTrayMenu(m)

	a.Run()
}

func MakeUtilsMenuItems(toolName string, shortKeyDefList []string) []*fyne.MenuItem {
	var res = make([]*fyne.MenuItem, len(shortKeyDefList)+2)
	for i, shortKeyDef := range shortKeyDefList {
		res[i] = fyne.NewMenuItem(shortKeyDef, nil)
	}
	res[len(res)-2] = fyne.NewMenuItem(fmt.Sprintf("为%s自定义快捷键", toolName), nil)
	res[len(res)-1] = fyne.NewMenuItemSeparator()

	return res
}
